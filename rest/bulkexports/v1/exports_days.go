/*
 * Twilio - Bulkexports
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

// Fetch a specific Day.
func (c *ApiService) FetchDay(ResourceType string, Day string) error {
	path := "/v1/Exports/{ResourceType}/Days/{Day}"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)
	path = strings.Replace(path, "{"+"Day"+"}", Day, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'ListDay'
type ListDayParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListDayParams) SetPageSize(PageSize int) *ListDayParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Days for a resource.
func (c *ApiService) ListDay(ResourceType string, params *ListDayParams) (*ListDayResponse, error) {
	path := "/v1/Exports/{ResourceType}/Days"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

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

	ps := &ListDayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Day records from the API. Request is executed immediately.
func (c *ApiService) DayPage(ResourceType string, params *ListDayParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Exports/{ResourceType}/Days"
	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Day records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) DayStream(ResourceType string, params *ListDayParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.DayPage(ResourceType, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Day records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) DayList(ResourceType string, params *ListDayParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.DayPage(ResourceType, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}
