/*
 * Twilio - Autopilot
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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateFieldType'
type CreateFieldTypeParams struct {
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateFieldTypeParams) SetFriendlyName(FriendlyName string) *CreateFieldTypeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateFieldTypeParams) SetUniqueName(UniqueName string) *CreateFieldTypeParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateFieldType(AssistantSid string, params *CreateFieldTypeParams) (*AutopilotV1AssistantFieldType, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantFieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteFieldType(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
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

func (c *ApiService) FetchFieldType(AssistantSid string, Sid string) (*AutopilotV1AssistantFieldType, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantFieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFieldType'
type ListFieldTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListFieldTypeParams) SetPageSize(PageSize int) *ListFieldTypeParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a single page of FieldType records from the API. Request is executed immediately.
func (c *ApiService) PageFieldType(AssistantSid string, params *ListFieldTypeParams, pageToken string, pageNumber string) (*ListFieldTypeResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

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

	ps := &ListFieldTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FieldType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFieldType(AssistantSid string, params *ListFieldTypeParams, limit int) ([]AutopilotV1AssistantFieldType, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageFieldType(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1AssistantFieldType

	for response != nil {
		records = append(records, response.FieldTypes...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, limit, c.getNextListFieldTypeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFieldTypeResponse)
	}

	return records, err
}

// Streams FieldType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFieldType(AssistantSid string, params *ListFieldTypeParams, limit int) (chan AutopilotV1AssistantFieldType, error) {
	params.SetPageSize(client.ReadLimits(params.PageSize, limit))

	response, err := c.PageFieldType(AssistantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1AssistantFieldType, 1)

	go func() {
		for response != nil {
			for item := range response.FieldTypes {
				channel <- response.FieldTypes[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, limit, c.getNextListFieldTypeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFieldTypeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFieldTypeResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFieldTypeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateFieldType'
type UpdateFieldTypeParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateFieldTypeParams) SetFriendlyName(FriendlyName string) *UpdateFieldTypeParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateFieldTypeParams) SetUniqueName(UniqueName string) *UpdateFieldTypeParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateFieldType(AssistantSid string, Sid string, params *UpdateFieldTypeParams) (*AutopilotV1AssistantFieldType, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1AssistantFieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
