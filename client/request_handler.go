// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

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
	return c.Client.SendRequest(method, c.BuildUrl(rawURL), data, headers)
}

// BuildUrl builds the target host string taking into account region and edge configurations.
func (c *RequestHandler) BuildUrl(rawURL string) string {
	u, _ := url.Parse(rawURL)

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
		return u.String()
	}

	if len(pieces) == 4 {
		// product.region.twilio.com
		region = pieces[1]
	} else if len(pieces) == 5 {
		// product.edge.region.twilio.com
		edge = pieces[1]
		region = pieces[2]
	}

	if c.Edge != "" {
		edge = c.Edge
	}

	if c.Region != "" {
		region = c.Region
	} else if region == "" && edge != "" {
		region = "us1"
	}

	for _, item := range []string{product, edge, region, suffix} {
		if item != "" {
			result = append(result, item)
		}
	}

	u.Host = strings.Join(result, ".")
	return u.String()
}

func (c *RequestHandler) Post(path string, bodyData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodPost, path, bodyData, headers)
}

func (c *RequestHandler) Get(path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodGet, path, queryData, headers)
}

func (c *RequestHandler) Delete(path string, nothing url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodDelete, path, nil, headers)
}
