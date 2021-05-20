package client

import (
	"net/http"
	"net/url"
	"time"
)

type BaseClient interface {
	AccountSid() string
	SetTimeout(timeout time.Duration)
	SendRequest(method string, rawURL string, data url.Values,
		headers map[string]interface{}) (*http.Response, error)
}
