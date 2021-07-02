/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific Flow.
func (c *ApiService) DeleteFlow(Sid string) error {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Retrieve a specific Flow.
func (c *ApiService) FetchFlow(Sid string) (*StudioV1Flow, error) {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Flows.
func (c *ApiService) ListFlow(params *ListFlowParams) (*ListFlowResponse, error) {
	path := "/v1/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Flow records from the API. Request is executed immediately.
func (c *ApiService) FlowPage(params *ListFlowParams, pageToken string, pageNumber string) (*client.Page, error) {
	path := "/v1/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	return client.NewPage(c.baseURL, response), nil
}

//Streams Flow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) FlowStream(params *ListFlowParams, limit int) (chan map[string]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.FlowPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.Stream(page, limit, 0), nil
}

//Lists Flow records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) FlowList(params *ListFlowParams, limit int) ([]interface{}, error) {
	params.SetPageSize(c.requestHandler.ReadLimits(params.PageSize, limit))
	page, err := c.FlowPage(params, "", "")
	if err != nil {
		return nil, err
	}
	return c.requestHandler.List(page, limit, 0), nil
}
