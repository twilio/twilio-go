// Package client provides internal utilities for the twilio-go client library.
package client

import (
	"log"
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
	headers map[string]interface{}, body ...byte) (*http.Response, error) {
	if (c.Edge == "" && c.Region != "") || (c.Edge != "" && c.Region == "") {
		log.Println("For regional processing, DNS is of format product.city.region.twilio.com; otherwise use product.twilio.com.")
	}
	if c.Edge == "" && c.Region != "" {
		log.Println("Setting default `Edge` for the provided `region`")
		var regionToEdge = map[string]string{
			"au1": "sydney",
			"br1": "sao-paulo",
			"de1": "frankfurt",
			"ie1": "dublin",
			"jp1": "tokyo",
			"jp2": "osaka",
			"sg1": "singapore",
			"us1": "ashburn",
			"us2": "umatilla",
		}
		if edge, ok := regionToEdge[c.Region]; ok {
			c.Edge = edge
		}
	}
	parsedURL, err := c.BuildUrl(rawURL)
	if err != nil {
		return nil, err
	}
	return c.Client.SendRequest(method, parsedURL, data, headers, body...)
}

// BuildUrl builds the target host string taking into account region and edge configurations.
func (c *RequestHandler) BuildUrl(rawURL string) (string, error) {
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
	return u.String(), nil
}

func (c *RequestHandler) Post(path string, bodyData url.Values, headers map[string]interface{}, body ...byte) (*http.Response, error) {
	return c.sendRequest(http.MethodPost, path, bodyData, headers, body...)
}

func (c *RequestHandler) Put(path string, bodyData url.Values, headers map[string]interface{}, body ...byte) (*http.Response, error) {
	return c.sendRequest(http.MethodPut, path, bodyData, headers, body...)
}

func (c *RequestHandler) Patch(path string, bodyData url.Values, headers map[string]interface{}, body ...byte) (*http.Response, error) {
	return c.sendRequest(http.MethodPatch, path, bodyData, headers, body...)
}

func (c *RequestHandler) Get(path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodGet, path, queryData, headers)
}

func (c *RequestHandler) Delete(path string, queryData url.Values, headers map[string]interface{}) (*http.Response, error) {
	return c.sendRequest(http.MethodDelete, path, queryData, headers)
}
