package client

import (
	"context"
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

// BaseClientWithCtx is an extension of BaseClient with the ability to associate a contex with
// the request
type BaseClientWithCtx interface {
	BaseClient
	SendRequestWithCtx(ctx context.Context, method string, rawURL string, data url.Values,
		headers map[string]interface{}) (*http.Response, error)
}

type wrapperClient struct {
	// embed the BaseClient so the functions remain accessible
	BaseClient
}

func (w wrapperClient) SendRequestWithCtx(ctx context.Context, method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	return w.SendRequest(method, rawURL, data, headers)
}

// wrapBaseClientWithNoopCtx "upgrades" a BaseClient to BaseClientWithCtx so that requests can be
// send with a request context.
func wrapBaseClientWithNoopCtx(c BaseClient) BaseClientWithCtx {
	// the default library client has SendRequestWithCtx, use it if available.
	if typedClient, ok := c.(BaseClientWithCtx); ok {
		return typedClient
	}
	return wrapperClient{BaseClient: c}
}
