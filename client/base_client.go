package client

import (
	"net/http"
	"net/url"
)

// BaseClient is the interface that wraps the Get/Post/Delete calls.
type BaseClient interface {
	Post(path string, bodyData url.Values, headers interface{}) (*http.Response, error)
	Get(path string, queryData interface{}, headers interface{}) (*http.Response, error)
	Delete(path string, nothing interface{}, headers interface{}) (*http.Response, error)
}
