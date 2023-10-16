// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"net/url"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"github.com/pkg/errors"
	"github.com/twilio/twilio-go/client/form"
)

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

func extractContentTypeHeader(headers map[string]interface{}) (cType string){
		headerType, ok := headers["Content-Type"]
			if !ok {
				return urlEncodedContentType
			  } 
			return headerType.(string)
}

func decodeJsonPayload(jsonBody string)(s []byte){
	prefixTrimmedString := jsonBody[2:]
	suffixTrimmedString := prefixTrimmedString[:len(prefixTrimmedString)-2]
	finalJsonPayload := jsonPrefix + suffixTrimmedString + jsonSuffix
	finalJsonPayloadArray := []byte(finalJsonPayload)
	return finalJsonPayloadArray
}

const (
	urlEncodedContentType = "application/x-www-form-urlencoded"
	jsonContentType = "application/json"
	jsonPrefix = "{"
	jsonSuffix = "}"
	keepZeros = true
	delimiter = '.'
	escapee   = '\\'
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

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c *Client) SendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	modRawURL := "https://preview.messaging.twilio.com" + rawURL

	contentType := extractContentTypeHeader(headers)

	u, err := url.Parse(modRawURL)
	if err != nil {
		return nil, err
	}
	
	valueReader := &strings.Reader{}
	goVersion := runtime.Version()
	var req *http.Request

	if method == http.MethodGet {
		if data != nil {
			v, _ := form.EncodeToStringWith(data, delimiter, escapee, keepZeros)
			regex := regexp.MustCompile(`\.\d+`)
			s := regex.ReplaceAllString(v, "")

			u.RawQuery = s
		}
	}

	if contentType == jsonContentType {
		var jsonData []byte
		var err error
		var encodedString string
		for _, value := range data {
			jsonData, err = json.Marshal(value[0])
			encodedString, err = url.QueryUnescape(string(jsonData))
			if err != nil {
				return nil, err
			}
		}
		jsonBody := decodeJsonPayload(encodedString)
		req, err = http.NewRequest(method, u.String(), bytes.NewBuffer(jsonBody))
		if err != nil {
			return nil, err
		}		
	} else {
		if method == http.MethodPost {
			valueReader = strings.NewReader(data.Encode())
		}
		req, err = http.NewRequest(method, u.String(), valueReader)
		if err != nil {
			return nil, err
		}

	}

	if contentType == urlEncodedContentType{
		req.Header.Add("Content-Type", urlEncodedContentType)
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