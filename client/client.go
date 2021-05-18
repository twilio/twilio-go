// Package client provides internal utilities for the twilio-go client library.
package client

import (
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
	"github.com/twilio/twilio-go/config"
	twilioError "github.com/twilio/twilio-go/framework/error"
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
	HTTPClient *http.Client
	accountSid string
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

const (
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
		err = &twilioError.TwilioRestError{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}
	return res, nil
}

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c *Client) SendRequest(method string, rawURL string, queryParams interface{}, formData url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	valueReader := &strings.Reader{}
	goVersion := runtime.Version()

	if queryParams != nil {
		v, _ := EncodeToStringWith(queryParams, delimiter, escapee, keepZeros)
		regex := regexp.MustCompile(`\.\d+`)
		s := regex.ReplaceAllString(v, "")

		u.RawQuery = s
	}

	if formData != nil {
		valueReader = strings.NewReader(formData.Encode())
	}

	req, err := http.NewRequest(method, u.String(), valueReader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.basicAuth())

	// E.g. "User-Agent": "twilio-go/1.0.0 (go1.16)"
	userAgent := fmt.Sprint("twilio-go/", config.LibraryVersion, " (", goVersion, ")")
	req.Header.Add("User-Agent", userAgent)

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	for k, v := range headers {
		req.Header.Add(k, fmt.Sprint(v))
	}

	c.PreProcessRequest(req)

	response, err := c.doWithErr(req)

	c.PreProcessResponse(response, err)

	return response, err
}

func (c *Client) PreProcessRequest(req *http.Request) {
	// No-op
}

func (c *Client) PreProcessResponse(response *http.Response, err error) {
	// No-op
}

func (c *Client) SetAccountSid(sid string) {
	c.accountSid = sid
}

// Returns the Account SID.
func (c *Client) AccountSid() string {
	return c.accountSid
}

// Post performs a POST request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c *Client) Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodPost, path, nil, bodyData, headers)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c *Client) Get(path string, queryData interface{}, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodGet, path, queryData, nil, headers)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c *Client) Delete(path string, nothing interface{}, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodDelete, path, nil, nil, headers)
}
