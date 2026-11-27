// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/twilio/twilio-go/client/form"
)

var alphanumericRegex *regexp.Regexp
var delimitingRegex *regexp.Regexp

func init() {
	alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	delimitingRegex = regexp.MustCompile(`\.\d+`)
	rand.Seed(time.Now().UnixNano())
}

// Credentials store user authentication credentials.
type Credentials struct {
	Username string
	Password string
}

func NewCredentials(username string, password string) *Credentials {
	return &Credentials{Username: username, Password: password}
}

type OAuth interface {
	GetAccessToken(context.Context) (string, error)
}

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	*Credentials
	HTTPClient          *http.Client
	accountSid          string
	UserAgentExtensions []string
	oAuth               OAuth
	// AutoRetry enables automatic retry with exponential backoff on 429 responses
	AutoRetry bool
	// MaxRetryDelay is the maximum retry delay in milliseconds for 429 responses
	MaxRetryDelay int
	// MaxRetries is the maximum number of request retries for 429 responses
	MaxRetries int
}

// default http Client should not follow redirects and return the most recent response.
func defaultHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 10,
	}
}

func (c *Client) basicAuth() (string, string) {
	return c.Credentials.Username, c.Credentials.Password
}

// SetTimeout sets the Timeout for HTTP requests.
func (c *Client) SetTimeout(timeout time.Duration) {
	if c.HTTPClient == nil {
		c.HTTPClient = defaultHTTPClient()
	}
	c.HTTPClient.Timeout = timeout
}

func extractContentTypeHeader(headers map[string]interface{}) (cType string) {
	headerType, ok := headers["Content-Type"]
	if !ok {
		return urlEncodedContentType
	}
	return headerType.(string)
}

const (
	urlEncodedContentType            = "application/x-www-form-urlencoded"
	jsonContentType                  = "application/json"
	keepZeros                        = true
	delimiter                        = '.'
	escapee                          = '\\'
	defaultInitialRetryIntervalMillis = 100
	defaultMaxRetryDelay              = 3000
	defaultMaxRetries                 = 3
)

func (c *Client) doWithErr(req *http.Request) (*http.Response, error) {
	return c.doWithErrAndRetry(req, 0)
}

func (c *Client) doWithErrAndRetry(req *http.Request, retryCount int) (*http.Response, error) {
	client := c.HTTPClient

	if client == nil {
		client = defaultHTTPClient()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Handle 429 responses with retry if AutoRetry is enabled
	if res.StatusCode == http.StatusTooManyRequests && c.AutoRetry {
		maxRetries := c.MaxRetries
		if maxRetries == 0 {
			maxRetries = defaultMaxRetries
		}

		if retryCount < maxRetries {
			res.Body.Close()

			// Calculate exponential backoff delay with full jitter
			maxRetryDelay := c.MaxRetryDelay
			if maxRetryDelay == 0 {
				maxRetryDelay = defaultMaxRetryDelay
			}

			baseDelay := math.Min(
				float64(maxRetryDelay),
				float64(defaultInitialRetryIntervalMillis)*math.Pow(2, float64(retryCount+1)),
			)
			// Full jitter: random delay between 0 and baseDelay
			delay := time.Duration(rand.Float64()*baseDelay) * time.Millisecond

			time.Sleep(delay)

			// Retry the request
			return c.doWithErrAndRetry(req, retryCount+1)
		}
	}

	// Note that 3XX response codes are allowed for fetches
	if res.StatusCode < 200 || res.StatusCode >= 400 {
		// First try to decode as a RestErrorV1 (Twilio API Standards V1.0)
		apiVersion := req.Header.Get("X-Twilio-ApiVersion")
		if apiVersion == "v1.0" {
			v1Err := &RestErrorV1{}
			if decodeErr := json.NewDecoder(res.Body).Decode(v1Err); decodeErr == nil {
				// Successfully decoded as V1 error format
				return nil, v1Err
			}
			// Reset the response body reader for the next attempt
			res.Body.Close()
			res, err = client.Do(req)
			if err != nil {
				return nil, err
			}
		}

		// Fall back to legacy TwilioRestError format
		err = &TwilioRestError{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}
	return res, nil
}

// throws error if username and password contains special characters
func (c *Client) validateCredentials() error {
	username, password := c.basicAuth()
	if !alphanumericRegex.MatchString(username) {
		return &TwilioRestError{
			Status:   400,
			Code:     21222,
			Message:  "Invalid Username. Illegal chars",
			MoreInfo: "https://www.twilio.com/docs/errors/21222"}
	}
	if !alphanumericRegex.MatchString(password) {
		return &TwilioRestError{
			Status:   400,
			Code:     21224,
			Message:  "Invalid Password. Illegal chars",
			MoreInfo: "https://www.twilio.com/docs/errors/21224"}
	}
	return nil
}

var baseUserAgent string
var userAgentOnce sync.Once

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c *Client) SendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}, body ...byte) (*http.Response, error) {

	contentType := extractContentTypeHeader(headers)

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	valueReader := &strings.Reader{}
	goVersion := runtime.Version()
	var req *http.Request

	//For HTTP GET Method there are no body parameters. All other parameters like query, path etc
	// are added as information in the url itself. Also while Content-Type is json, we are sending
	// json body. In that case, data variable contains all other parameters than body, which is the
	//same case as GET method. In that case as well all parameters will be added to url
	if method == http.MethodGet || method == http.MethodDelete || contentType == jsonContentType {
		if data != nil {
			v, _ := form.EncodeToStringWith(data, delimiter, escapee, keepZeros)
			s := delimitingRegex.ReplaceAllString(v, "")

			u.RawQuery = s
		}
	}

	//data is already processed and information will be added to u(the url) in the
	//previous step. Now body will solely contain json payload
	if contentType == jsonContentType {
		req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	} else {
		// Here the HTTP POST methods which do not have json content type are processed
		// All the values will be added in data and encoded (all body, query, path parameters)
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
			valueReader = strings.NewReader(data.Encode())
		}
		req, err = http.NewRequestWithContext(context.Background(), method, u.String(), valueReader)
		if err != nil {
			return nil, err
		}

	}

	credErr := c.validateCredentials()
	if credErr != nil {
		return nil, credErr
	}
	if c.OAuth() == nil && c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.basicAuth())
	}

	// E.g. "User-Agent": "twilio-go/1.0.0 (darwin amd64) go/go1.17.8"
	userAgentOnce.Do(func() {
		baseUserAgent = fmt.Sprintf("twilio-go/%s (%s %s) go/%s", LibraryVersion, runtime.GOOS, runtime.GOARCH, goVersion)
	})
	userAgent := baseUserAgent

	if len(c.UserAgentExtensions) > 0 {
		userAgent += " " + strings.Join(c.UserAgentExtensions, " ")
	}
	if c.OAuth() != nil {
		oauth := c.OAuth()
		token, _ := c.OAuth().GetAccessToken(context.TODO())
		if token != "" {
			req.Header.Add("Authorization", "Bearer "+token)
		}
		c.SetOauth(oauth) // Set the OAuth token in the client which gets nullified after the token fetch
	} else if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.basicAuth())
	}

	req.Header.Add("User-Agent", userAgent)

	for k, v := range headers {
		req.Header.Add(k, fmt.Sprint(v))
	}
	return c.doWithErr(req)
}

// SetAccountSid sets the Client's accountSid field
func (c *Client) SetAccountSid(sid string) {
	c.accountSid = sid
}

// AccountSid returns the Account SID.
func (c *Client) AccountSid() string {
	return c.accountSid
}

func (c *Client) SetOauth(oauth OAuth) {
	c.oAuth = oauth
}

func (c *Client) OAuth() OAuth {
	return c.oAuth
}

// SetAutoRetry enables or disables automatic retry with exponential backoff on 429 responses.
func (c *Client) SetAutoRetry(autoRetry bool) {
	c.AutoRetry = autoRetry
}

// SetMaxRetryDelay sets the maximum retry delay in milliseconds for 429 responses.
// Defaults to 3000ms if not set.
func (c *Client) SetMaxRetryDelay(maxRetryDelay int) {
	c.MaxRetryDelay = maxRetryDelay
}

// SetMaxRetries sets the maximum number of request retries for 429 responses.
// Defaults to 3 if not set.
func (c *Client) SetMaxRetries(maxRetries int) {
	c.MaxRetries = maxRetries
}
