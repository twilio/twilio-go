// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"encoding/json"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	twilioError "github.com/twilio/twilio-go/framework/error"
)

// Credentials store user authentication credentials.
type Credentials struct {
	AccountSID string
	AuthToken  string
}

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	*Credentials
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient initializes a new Client with the given credentials
func NewClient(accountSid string, authToken string) *Client {
	c := &Client{}
	creds := &Credentials{AccountSID: accountSid, AuthToken: authToken}
	c.Credentials = creds
	return c
}

func (c *Client) basicAuth() (string, string) {
	return c.Credentials.AccountSID, c.Credentials.AuthToken
}

const errorStatusCode = 400
const (
	keepZeros = true
	delimiter = '.'
	escapee   = '\\'
)

func doWithErr(req *http.Request, client *http.Client) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= errorStatusCode {
		err = &twilioError.TwilioRestError{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c Client) SendRequest(method string, rawURL string, queryParams interface{}, formData url.Values) (*http.Response, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	valueReader := &strings.Reader{}

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

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return doWithErr(req, c.HTTPClient)
}

// Post performs a POST request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Post(path string, bodyData url.Values, headers interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodPost, path, nil, bodyData)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Get(path string, queryData interface{}, headers interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodGet, path, queryData, nil)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Delete(path string, nothing interface{}, headers interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodDelete, path, nil, nil)
}
