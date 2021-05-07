package client

import (
	"net/http"
	"net/url"
)

type BaseClient interface {
	GetAccountSid() string
	Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error)
	Get(path string, queryData interface{}, headers map[string]interface{}) (*http.Response, error)
	Delete(path string, nothing interface{}, headers map[string]interface{}) (*http.Response, error)
}
