package backend

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

// Credentials store user authentication credentials.
type Credentials struct {
	AccountSid string
	AuthToken  string
}

// Request provides a standard HTTP backend.
type Request struct {
	Credentials
	BaseURL string
}

func (request *Request) basicAuth() (string, string) {
	return request.Credentials.AccountSid, request.Credentials.AuthToken
}

// Post performs a POST request on the object at the provided URI in the context of the Request's BaseURL with the provided data as parameters.
func (request *Request) Post(uri string, data interface{}) (*http.Response, error) {
	v, _ := query.Values(data)
	valueReader := strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", request.BaseURL+uri, valueReader)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.basicAuth())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return httpClient.Do(req)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL with the provided data as parameters.
func (request *Request) Get(uri string) (*http.Response, error) {
	req, err := http.NewRequest("GET", request.BaseURL+uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.basicAuth())

	return httpClient.Do(req)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL with the provided data as parameters.
func (request *Request) Delete(uri string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", request.BaseURL+uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.basicAuth())

	return httpClient.Do(req)
}
