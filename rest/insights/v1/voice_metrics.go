/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'ListMetric'
type ListMetricParams struct {
	//
	Edge *string `json:"Edge,omitempty"`
	//
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListMetricParams) SetEdge(Edge string) *ListMetricParams {
	params.Edge = &Edge
	return params
}
func (params *ListMetricParams) SetDirection(Direction string) *ListMetricParams {
	params.Direction = &Direction
	return params
}
func (params *ListMetricParams) SetPageSize(PageSize int) *ListMetricParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListMetric(CallSid string, params *ListMetricParams) (*ListMetricResponse, error) {
	path := "/v1/Voice/{CallSid}/Metrics"
	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMetricResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
