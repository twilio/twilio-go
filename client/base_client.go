package client

import (
	"net/http"
	"time"
)

type BaseClient interface {
	AccountSid() string
	SetTimeout(timeout time.Duration)
	SendRequest(method string, rawURL string, data interface{},
		headers map[string]interface{}) (*http.Response, error)
}
