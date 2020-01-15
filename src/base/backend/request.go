package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
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

// Error provides information about an unsuccesful request.
type Error struct {
	Code     int    `json:"code"`
	Detail   string `json:"detail"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("Status: %d - Error %d: %s (%s) More info: %s", err.Status, err.Code, err.Message, err.Detail, err.MoreInfo)
}

func (request *Request) basicAuth() (string, string) {
	return request.Credentials.AccountSid, request.Credentials.AuthToken
}

func doWithErr(req *http.Request) (*http.Response, error) {
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode >= 400 {
		err = &Error{}
		json.NewDecoder(res.Body).Decode(err)
	}

	return res, err
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

	return doWithErr(req)
}

// Get performs a GET request on the object at the provided URI in the context of the Request's BaseURL with the provided data as parameters.
func (request *Request) Get(uri string) (*http.Response, error) {
	req, err := http.NewRequest("GET", request.BaseURL+uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.basicAuth())

	return doWithErr(req)
}

// Delete performs a DELETE request on the object at the provided URI in the context of the Request's BaseURL with the provided data as parameters.
func (request *Request) Delete(uri string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", request.BaseURL+uri, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.basicAuth())

	return doWithErr(req)
}
