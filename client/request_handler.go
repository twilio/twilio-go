// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func buildUrlInternal(overrideEdge, overrideRegion, rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	var (
		edge    = ""
		region  = ""
		pieces  = strings.Split(u.Host, ".")
		product = pieces[0]
		result  []string
	)
	suffix := ""

	if len(pieces) >= 3 {
		suffix = strings.Join(pieces[len(pieces)-2:], ".")
	} else {
		return u.String(), nil
	}

	if len(pieces) == 4 {
		// product.region.twilio.com
		region = pieces[1]
	} else if len(pieces) == 5 {
		// product.edge.region.twilio.com
		edge = pieces[1]
		region = pieces[2]
	}

	if overrideEdge != "" {
		edge = overrideEdge
	}

	if overrideRegion != "" {
		region = overrideRegion
	} else if region == "" && edge != "" {
		region = "us1"
	}

	for _, item := range []string{product, edge, region, suffix} {
		if item != "" {
			result = append(result, item)
		}
	}

	u.Host = strings.Join(result, ".")
	return u.String(), nil
}

type RequestHandler struct {
	Client BaseClient
	Edge   string
	Region string
}

func NewRequestHandler(client BaseClient) *RequestHandler {
	return &RequestHandler{
		Client: client,
		Edge:   os.Getenv("TWILIO_EDGE"),
		Region: os.Getenv("TWILIO_REGION"),
	}
}

func (c *RequestHandler) sendRequest(method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	parsedURL, err := c.BuildUrl(rawURL)
	if err != nil {
		return nil, err
	}

	return c.Client.SendRequest(method, parsedURL, data, headers)
}

// BuildUrl builds the target host string taking into account region and edge configurations.
func (c *RequestHandler) BuildUrl(rawURL string) (string, error) {
	return buildUrlInternal(c.Edge, c.Region, rawURL)
}

// deprecated
func (c *RequestHandler) Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodPost, path, bodyData, headers)
}

// deprecated
func (c *RequestHandler) Get(path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodGet, path, queryData, headers)
}

// deprecated
func (c *RequestHandler) Delete(path string, nothing url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodDelete, path, nil, headers)
}

func UpgradeRequestHandler(h *RequestHandler) *RequestHandlerWithCtx {
	return &RequestHandlerWithCtx{
		// wrapped client will supply context.TODO() to all API calls
		Client: wrapBaseClientWithNoopCtx(h.Client),
		Edge:   h.Edge,
		Region: h.Region,
	}
}

type RequestHandlerWithCtx struct {
	Client BaseClientWithCtx
	Edge   string
	Region string
}

func (c *RequestHandlerWithCtx) sendRequest(ctx context.Context, method string, rawURL string, data url.Values,
	headers map[string]interface{}) (*http.Response, error) {
	parsedURL, err := c.BuildUrl(rawURL)
	if err != nil {
		return nil, err
	}

	return c.Client.SendRequestWithCtx(ctx, method, parsedURL, data, headers)
}

func NewRequestHandlerWithCtx(client BaseClientWithCtx) *RequestHandlerWithCtx {
	return &RequestHandlerWithCtx{
		Client: client,
		Edge:   os.Getenv("TWILIO_EDGE"),
		Region: os.Getenv("TWILIO_REGION"),
	}
}

// BuildUrl builds the target host string taking into account region and edge configurations.
func (c *RequestHandlerWithCtx) BuildUrl(rawURL string) (string, error) {
	return buildUrlInternal(c.Edge, c.Region, rawURL)
}

func (c *RequestHandlerWithCtx) Post(ctx context.Context, path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(ctx, http.MethodPost, path, bodyData, headers)
}

func (c *RequestHandlerWithCtx) Get(ctx context.Context, path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(ctx, http.MethodGet, path, queryData, headers)
}

func (c *RequestHandlerWithCtx) Delete(ctx context.Context, path string, nothing url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(ctx, http.MethodDelete, path, nil, headers)
}
