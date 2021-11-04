/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.0
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

// Retrieve a Step.
func (c *ApiService) FetchExecutionStep(FlowSid string, ExecutionSid string, Sid string) (*StudioV2ExecutionStep, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV2ExecutionStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListExecutionStepParams) SetPageSize(PageSize int) *ListExecutionStepParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListExecutionStepParams) SetLimit(Limit int) *ListExecutionStepParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ExecutionStep records from the API. Request is executed immediately.
func (c *ApiService) PageExecutionStep(FlowSid string, ExecutionSid string, params *ListExecutionStepParams, pageToken, pageNumber string) (*ListExecutionStepResponse, error) {
	path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ExecutionStep records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListExecutionStep(FlowSid string, ExecutionSid string, params *ListExecutionStepParams) ([]StudioV2ExecutionStep, error) {
	if params == nil {
		params = &ListExecutionStepParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageExecutionStep(FlowSid, ExecutionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []StudioV2ExecutionStep

	for response != nil {
		records = append(records, response.Steps...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListExecutionStepResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListExecutionStepResponse)
	}

	return records, err
}

// Streams ExecutionStep records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamExecutionStep(FlowSid string, ExecutionSid string, params *ListExecutionStepParams) (chan StudioV2ExecutionStep, error) {
	if params == nil {
		params = &ListExecutionStepParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageExecutionStep(FlowSid, ExecutionSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan StudioV2ExecutionStep, 1)

	go func() {
		for response != nil {
			for item := range response.Steps {
				channel <- response.Steps[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListExecutionStepResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListExecutionStepResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListExecutionStepResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
