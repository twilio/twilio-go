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
	AccountSid string
	AuthToken  string
}

// Request provides a standard HTTP backend.
type Request struct {
	Credentials
	Client  *http.Client
	BaseURL string
}

// Error provides information about an unsuccessful request.
type Error struct {
	Code     int    `json:"code"`
	Detail   string `json:"detail"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("Status: %d - Error %d: %s (%s) More info: %s",
		err.Status, err.Code, err.Message, err.Detail, err.MoreInfo)
}

func (request *Request) basicAuth() (string, string) {
	return request.Credentials.AccountSid, request.Credentials.AuthToken
}

func doWithErr(req *http.Request, client *http.Client) (*http.Response, error) {
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return nil, httpErr
	}

	if res.StatusCode >= errorStatusCode {
		apiErr := &Error{}
		if decodeErr := json.NewDecoder(res.Body).Decode(apiErr); decodeErr != nil {
			return nil, decodeErr
		}

		return nil, apiErr
	}

	return res, nil
}

// MakeRequest makes HTTP request.
func (request Request) MakeRequest(method string, fullyQualifiedURI string, data interface{}) (*http.Response, error) {
	valueReader := &strings.Reader{}

	if data != nil {
		v, _ := query.Values(data)
		qs := v.Encode()
		replacer := strings.NewReplacer("%5B", ".", "%5D", "")
		dotNotationQs := replacer.Replace(qs)
		valueReader = strings.NewReader(dotNotationQs)
	}

	req, err := http.NewRequest(method, fullyQualifiedURI, valueReader)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(request.basicAuth())

	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return doWithErr(req, request.Client)
}

// Post performs a POST request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (request Request) Post(uri string, data interface{}) (*http.Response, error) {
	fullyQualifiedURI := request.BaseURL + uri
	return request.MakeRequest(http.MethodPost, fullyQualifiedURI, data)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (request Request) Get(uri string) (*http.Response, error) {
	fullyQualifiedURI := request.BaseURL + uri
	return request.MakeRequest(http.MethodGet, fullyQualifiedURI, nil)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL
// with the provided data as parameters.
func (request Request) Delete(uri string) (*http.Response, error) {
	fullyQualifiedURI := request.BaseURL + uri
	return request.MakeRequest(http.MethodDelete, fullyQualifiedURI, nil)
}
