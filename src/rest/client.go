package twilio

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 30,
}

type Request struct {
	AccountSid string
	AuthToken  string

	APIKeySid    string
	APIKeySecret string
}

func (request *Request) BasicAuth() (string, string) {
	if request.APIKeySid != "" {
		return request.APIKeySid, request.APIKeySecret
	}

	return request.AccountSid, request.AuthToken
}

func (request *Request) post(payload url.Values, url string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return client.Do(req)
}

func (request *Request) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())

	return client.Do(req)
}

func (request *Request) delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(request.BasicAuth())

	return client.Do(req)
}
