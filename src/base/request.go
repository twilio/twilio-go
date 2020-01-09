package twilio

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 30,
}

type Request struct {
	AccountSid string
	AuthToken  string

	APIKeySid    string
	APIKeySecret string

	BaseURL string
}

func (request *Request) BasicAuth() (string, string) {
	if request.APIKeySid != "" {
		return request.APIKeySid, request.APIKeySecret
	}

	return request.AccountSid, request.AuthToken
}

func (request *Request) Post(uri string, data url.Values) (*http.Response, error) {
	req, err := http.NewRequest("POST", request.BaseURL+uri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return httpClient.Do(req)
}

func (request *Request) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())

	return httpClient.Do(req)
}

func (request *Request) Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())

	return httpClient.Do(req)
}
