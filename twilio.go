package twilio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const errorStatusCode = 400

// Credentials store user authentication credentials.
type Credentials struct {
	AccountSID string
	AuthToken  string
}

// Client provides a standard HTTP backend.
type Client struct {
	Credentials
	Client  *http.Client
	BaseURL string
}

const (
	defaultBaseURL = "https://api.twilio.com"
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

func (c *Client) basicAuth() (string, string) {
	return c.Credentials.AccountSID, c.Credentials.AuthToken
}

func doWithErr(req *http.Request, client *http.Client) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}

	res, httpErr := client.Do(req)
	if httpErr != nil {
		return nil, httpErr
	}

	if res.StatusCode >= errorStatusCode {
		apiErr := &Error{}
		if decodeErr := json.NewDecoder(res.Body).Decode(apiErr); decodeErr != nil {
			apiErr.Code = res.StatusCode
			apiErr.Message = "Error decoding response: " + decodeErr.Error()

			return nil, apiErr
		}

		return nil, apiErr
	}

	return res, nil
}

// SendRequest makes HTTP request.
func (c Client) SendRequest(method string, path string, data interface{}) (*http.Response, error) {
	baseURL := c.BaseURL

	if len(strings.TrimSpace(baseURL)) == 0 {
		baseURL = defaultBaseURL
	}

	fullyQualifiedURI := baseURL + path

	valueReader := &strings.Reader{}

	if data != nil {
		v, _ := query.Values(data)
		valueReader = strings.NewReader(v.Encode())
	}

	req, err := http.NewRequest(method, fullyQualifiedURI, valueReader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.basicAuth())

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return doWithErr(req, c.Client)
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
