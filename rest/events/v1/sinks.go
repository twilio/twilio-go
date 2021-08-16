/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
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

// Optional parameters for the method 'CreateSink'
type CreateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
	// The information required for Twilio to connect to the provided Sink encoded as JSON.
	SinkConfiguration *map[string]interface{} `json:"SinkConfiguration,omitempty"`
	// The Sink type. Can only be \\\"kinesis\\\" or \\\"webhook\\\" currently.
	SinkType *string `json:"SinkType,omitempty"`
}

func (params *CreateSinkParams) SetDescription(Description string) *CreateSinkParams {
	params.Description = &Description
	return params
}
func (params *CreateSinkParams) SetSinkConfiguration(SinkConfiguration map[string]interface{}) *CreateSinkParams {
	params.SinkConfiguration = &SinkConfiguration
	return params
}
func (params *CreateSinkParams) SetSinkType(SinkType string) *CreateSinkParams {
	params.SinkType = &SinkType
	return params
}

// Create a new Sink
func (c *ApiService) CreateSink(params *CreateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.SinkConfiguration != nil {
		v, err := json.Marshal(params.SinkConfiguration)

		if err != nil {
			return nil, err
		}

		data.Set("SinkConfiguration", string(v))
	}
	if params != nil && params.SinkType != nil {
		data.Set("SinkType", *params.SinkType)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Sink.
func (c *ApiService) DeleteSink(Sid string) error {
	path := "/v1/Sinks/{Sid}"
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

// Fetch a specific Sink.
func (c *ApiService) FetchSink(Sid string) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSink'
type ListSinkParams struct {
	// A boolean query parameter filtering the results to return sinks used/not used by a subscription.
	InUse *bool `json:"InUse,omitempty"`
	// A String query parameter filtering the results by status `initialized`, `validating`, `active` or `failed`.
	Status *string `json:"Status,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSinkParams) SetInUse(InUse bool) *ListSinkParams {
	params.InUse = &InUse
	return params
}
func (params *ListSinkParams) SetStatus(Status string) *ListSinkParams {
	params.Status = &Status
	return params
}
func (params *ListSinkParams) SetPageSize(PageSize int) *ListSinkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSinkParams) SetLimit(Limit int) *ListSinkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Sink records from the API. Request is executed immediately.
func (c *ApiService) PageSink(params *ListSinkParams, pageToken string, pageNumber string) (*ListSinkResponse, error) {
	path := "/v1/Sinks"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.InUse != nil {
		data.Set("InUse", fmt.Sprint(*params.InUse))
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
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

	ps := &ListSinkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Sink records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSink(params *ListSinkParams) ([]EventsV1Sink, error) {
	if params == nil {
		params = &ListSinkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSink(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []EventsV1Sink

	for response != nil {
		records = append(records, response.Sinks...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSinkResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSinkResponse)
	}

	return records, err
}

// Streams Sink records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSink(params *ListSinkParams) (chan EventsV1Sink, error) {
	if params == nil {
		params = &ListSinkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSink(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan EventsV1Sink, 1)

	go func() {
		for response != nil {
			for item := range response.Sinks {
				channel <- response.Sinks[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSinkResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSinkResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSinkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSinkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSink'
type UpdateSinkParams struct {
	// A human readable description for the Sink **This value should not contain PII.**
	Description *string `json:"Description,omitempty"`
}

func (params *UpdateSinkParams) SetDescription(Description string) *UpdateSinkParams {
	params.Description = &Description
	return params
}

// Update a specific Sink
func (c *ApiService) UpdateSink(Sid string, params *UpdateSinkParams) (*EventsV1Sink, error) {
	path := "/v1/Sinks/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &EventsV1Sink{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
