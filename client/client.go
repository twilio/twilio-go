// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/twilio/twilio-go/client/form"
)

var alphanumericRegex *regexp.Regexp
var delimitingRegex *regexp.Regexp

func init() {
	alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	delimitingRegex = regexp.MustCompile(`\.\d+`)
}

// Credentials store user authentication credentials.
type Credentials struct {
	Username string
	Password string
}

func NewCredentials(username string, password string) *Credentials {
	return &Credentials{Username: username, Password: password}
}

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	*Credentials
	HTTPClient          *http.Client
	accountSid          string
	UserAgentExtensions []string
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
	urlEncodedContentType = "application/x-www-form-urlencoded"
	jsonContentType       = "application/json"
	keepZeros             = true
	delimiter             = '.'
	escapee               = '\\'
)

func (c *Client) doWithErr(req *http.Request) (*http.Response, error) {
	client := c.HTTPClient

	if client == nil {
		client = defaultHTTPClient()
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Note that 3XX response codes are allowed for fetches
	if res.StatusCode < 200 || res.StatusCode >= 400 {
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
		//Here the HTTP POST methods which is not having json content type are processed
		//All the values will be added in data and encoded (all body, query, path parameters)
		if method == http.MethodPost || method == http.MethodPut {
			valueReader = strings.NewReader(data.Encode())
		}
		req, err = http.NewRequest(method, u.String(), valueReader)
		if err != nil {
			return nil, err
		}

	}

	credErr := c.validateCredentials()
	if credErr != nil {
		return nil, credErr
	}
	req.SetBasicAuth(c.basicAuth())

	// E.g. "User-Agent": "twilio-go/1.0.0 (darwin amd64) go/go1.17.8"
	userAgent := fmt.Sprintf("twilio-go/%s (%s %s) go/%s", LibraryVersion, runtime.GOOS, runtime.GOARCH, goVersion)

	if len(c.UserAgentExtensions) > 0 {
		userAgent += " " + strings.Join(c.UserAgentExtensions, " ")
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

// Returns the Account SID.
func (c *Client) AccountSid() string {
	return c.accountSid
}
