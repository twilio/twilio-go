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

//Retrieve a single page of Day records from the API. Request is executed immediately.
func (c *ApiService) PageDay(ResourceType string, params *ListDayParams, pageToken string, pageNumber string) (*ListDayResponse, error) {
	path := "/v1/Exports/{ResourceType}/Days"

	path = strings.Replace(path, "{"+"ResourceType"+"}", ResourceType, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
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

//Lists Day records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDay(ResourceType string, params *ListDayParams, limit *int) ([]*ListDayResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDay(ResourceType, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []*ListDayResponse

	for response != nil {
		records = append(records, response)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDayResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListDayResponse)
	}

	return records, err
}

//Streams Day records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDay(ResourceType string, params *ListDayParams, limit *int) (chan *ListDayResponse, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageDay(ResourceType, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan *ListDayResponse, 1)

	go func() {
		for response != nil {
			channel <- response

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListDayResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListDayResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListDayResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDayResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
