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

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	*Credentials
	HTTPClient *http.Client
	BaseURL    string
	Edge       string
	Region     string
	AccountSid string
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

// BuildHost builds the target host string taking into account region and edge configurations.
func (c *Client) BuildHost(rawHost string) string {
	var (
		edge    = ""
		region  = ""
		pieces  = strings.Split(rawHost, ".")
		product = pieces[0]
		result  = []string{}
	)
	suffix := ""

	if len(pieces) >= 3 {
		suffix = strings.Join(pieces[len(pieces)-2:], ".")
	} else {
		return rawHost
	}

	if len(pieces) == 4 {
		// product.region.twilio.com
		region = pieces[1]
	} else if len(pieces) == 5 {
		// product.edge.region.twilio.com
		edge = pieces[1]
		region = pieces[2]
	}

	if c.Edge != "" {
		edge = c.Edge
	}

	if c.Region != "" {
		region = c.Region
	} else if region == "" && edge != "" {
		region = "us1"
	}

	if c.BaseURL != "" {
		suffix = c.BaseURL
	}

	for _, item := range []string{product, edge, region, suffix} {
		if item != "" {
			result = append(result, item)
		}
	}

	return strings.Join(result, ".")
}

// SetTimeout sets the Timeout for HTTP requests.
func (c *Client) SetTimeout(timeout time.Duration) {
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
func (c Client) SendRequest(method string, rawURL string, queryParams interface{}, formData url.Values,
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

	u.Host = c.BuildHost(u.Host)

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

	return c.doWithErr(req)
}

// Returns the Account SID.
func (c Client) GetAccountSid() string {
	return c.AccountSid
}

// Post performs a POST request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodPost, path, nil, bodyData, headers)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Get(path string, queryData interface{}, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodGet, path, queryData, nil, headers)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Delete(path string, nothing interface{}, headers map[string]interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodDelete, path, nil, nil, headers)
}
