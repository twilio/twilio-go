// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
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
		err = &TwilioRestError{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}
	return res, nil
}

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
	formContentType   = "application/x-www-form-urlencoded"
)

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c *Client) SendRequest(method string, rawURL string, data interface{},
	headers map[string]interface{}) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	goVersion := runtime.Version()

	if method == http.MethodGet {
		v, _ := EncodeToStringWith(data, delimiter, escapee, keepZeros)
		regex := regexp.MustCompile(`\.\d+`)
		s := regex.ReplaceAllString(v, "")

		u.RawQuery = s
	}

	var valueReader io.Reader
	if method == http.MethodPost {
		if headers == nil || headers[contentTypeHeader] == nil {
			return nil, fmt.Errorf("the '%s' header must be set on a POST request", contentTypeHeader)
		}

		requestBody, err := requestBodyToReader(headers[contentTypeHeader].(string), data)
		if err != nil {
			return nil, err
		}
		valueReader = requestBody
	}

	req, err := http.NewRequest(method, u.String(), valueReader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.basicAuth())

	// E.g. "User-Agent": "twilio-go/1.0.0 (go1.16)"
	userAgent := fmt.Sprint("twilio-go/", LibraryVersion, " (", goVersion, ")")
	req.Header.Add("User-Agent", userAgent)

	for k, v := range headers {
		req.Header.Add(k, fmt.Sprint(v))
	}

	return c.doWithErr(req)
}

func requestBodyToReader(contentTypeHeaderValue string, data interface{}) (io.Reader, error) {
	kind := reflect.ValueOf(data).Kind()

	if contentTypeHeaderValue == formContentType {
		if v, ok := data.(url.Values); ok {
			return strings.NewReader(v.Encode()), nil
		}
		return nil, fmt.Errorf("expected data to be of type url.Values for '%s' but got %s", formContentType, kind)
	}

	if contentTypeHeaderValue == jsonContentType {
		if kind == reflect.Struct || kind == reflect.Map || kind == reflect.Slice {
			body, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			return bytes.NewBuffer(body), nil
		}
		return nil, fmt.Errorf("expected data to be either a struct, map or slice for '%s' but got %s", jsonContentType, kind)
	}

	return nil, fmt.Errorf("%s is not a supported media type", contentTypeHeaderValue)
}

// SetAccountSid sets the Client's accountSid field
func (c *Client) SetAccountSid(sid string) {
	c.accountSid = sid
}

// Returns the Account SID.
func (c *Client) AccountSid() string {
	return c.accountSid
}
