// Package twilio provides internal utilities for the twilio-go client library.
package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Error provides information about an unsuccessful request.
type Error struct {
	Code     int    `json:"code"`
	Detail   string `json:"detail"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (err Error) Error() string {
	return fmt.Sprintf("Status: %d - Error %d: %s (%s) More info: %s",
		err.Status, err.Code, err.Message, err.Detail, err.MoreInfo)
}

// Credentials store user authentication credentials.
type Credentials struct {
	AccountSid string
	AuthToken  string
}

// Client encapsulates a standard HTTP backend with authorization.
type Client struct {
	Credentials
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) basicAuth() (string, string) {
	return c.Credentials.AccountSid, c.Credentials.AuthToken
}

const errorStatusCode = 400

func doWithErr(req *http.Request, client *http.Client) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= errorStatusCode {
		err = &Error{}
		if decodeErr := json.NewDecoder(res.Body).Decode(err); decodeErr != nil {
			err = errors.Wrap(decodeErr, "error decoding the response for an HTTP error code: "+strconv.Itoa(res.StatusCode))
			return nil, err
		}

		return nil, err
	}

	return res, nil
}

// SendRequest verifies, constructs, and authorizes an HTTP request.
func (c Client) SendRequest(method string, url string, data interface{}) (*http.Response, error) {
	valueReader := &strings.Reader{}

	if data != nil {
		v, _ := query.Values(data)
		qs := v.Encode()
		// Convert "[" and "]" (%5B and %5D) to "." and "" to conform to Twilio API specs.
		replacer := strings.NewReplacer("%5B", ".", "%5D", "")
		dotNotationQs := replacer.Replace(qs)
		valueReader = strings.NewReader(dotNotationQs)
	}

	req, err := http.NewRequest(method, url, valueReader)
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
func (c Client) Post(path string, data interface{}) (*http.Response, error) {
	return c.SendRequest(http.MethodPost, path, data)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Get(path string) (*http.Response, error) {
	return c.SendRequest(http.MethodGet, path, nil)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (c Client) Delete(path string) (*http.Response, error) {
	return c.SendRequest(http.MethodDelete, path, nil)
}
