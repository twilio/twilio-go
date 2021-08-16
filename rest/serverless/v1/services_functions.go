/*
 * Twilio - Serverless
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

// Optional parameters for the method 'CreateFunction'
type CreateFunctionParams struct {
	// A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateFunctionParams) SetFriendlyName(FriendlyName string) *CreateFunctionParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new Function resource.
func (c *ApiService) CreateFunction(ServiceSid string, params *CreateFunctionParams) (*ServerlessV1Function, error) {
	path := "/v1/Services/{ServiceSid}/Functions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Function{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Function resource.
func (c *ApiService) DeleteFunction(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Functions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

// Retrieve a specific Function resource.
func (c *ApiService) FetchFunction(ServiceSid string, Sid string) (*ServerlessV1Function, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Function{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFunction'
type ListFunctionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFunctionParams) SetPageSize(PageSize int) *ListFunctionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFunctionParams) SetLimit(Limit int) *ListFunctionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Function records from the API. Request is executed immediately.
func (c *ApiService) PageFunction(ServiceSid string, params *ListFunctionParams, pageToken string, pageNumber string) (*ListFunctionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Functions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListFunctionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Function records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFunction(ServiceSid string, params *ListFunctionParams) ([]ServerlessV1Function, error) {
	if params == nil {
		params = &ListFunctionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFunction(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ServerlessV1Function

	for response != nil {
		records = append(records, response.Functions...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFunctionResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFunctionResponse)
	}

	return records, err
}

// Streams Function records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFunction(ServiceSid string, params *ListFunctionParams) (chan ServerlessV1Function, error) {
	if params == nil {
		params = &ListFunctionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFunction(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ServerlessV1Function, 1)

	go func() {
		for response != nil {
			for item := range response.Functions {
				channel <- response.Functions[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFunctionResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFunctionResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFunctionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFunctionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFunction'
type UpdateFunctionParams struct {
	// A descriptive string that you create to describe the Function resource. It can be a maximum of 255 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateFunctionParams) SetFriendlyName(FriendlyName string) *UpdateFunctionParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Update a specific Function resource.
func (c *ApiService) UpdateFunction(ServiceSid string, Sid string, params *UpdateFunctionParams) (*ServerlessV1Function, error) {
	path := "/v1/Services/{ServiceSid}/Functions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ServerlessV1Function{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
