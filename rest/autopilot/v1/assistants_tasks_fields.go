/*
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateField'
type CreateFieldParams struct {
	// The Field Type of the new field. Can be: a [Built-in Field Type](https://www.twilio.com/docs/autopilot/built-in-field-types), the `unique_name`, or the `sid` of a custom Field Type.
	FieldType *string `json:"FieldType,omitempty"`
	// An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateFieldParams) SetFieldType(FieldType string) *CreateFieldParams {
	params.FieldType = &FieldType
	return params
}
func (params *CreateFieldParams) SetUniqueName(UniqueName string) *CreateFieldParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateField(AssistantSid string, TaskSid string, params *CreateFieldParams) (*AutopilotV1Field, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FieldType != nil {
		data.Set("FieldType", *params.FieldType)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Field{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteField(AssistantSid string, TaskSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
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

func (c *ApiService) FetchField(AssistantSid string, TaskSid string, Sid string) (*AutopilotV1Field, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Field{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListField'
type ListFieldParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFieldParams) SetPageSize(PageSize int) *ListFieldParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFieldParams) SetLimit(Limit int) *ListFieldParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Field records from the API. Request is executed immediately.
func (c *ApiService) PageField(AssistantSid string, TaskSid string, params *ListFieldParams, pageToken, pageNumber string) (*ListFieldResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{TaskSid}/Fields"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"TaskSid"+"}", TaskSid, -1)

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

	ps := &ListFieldResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Field records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListField(AssistantSid string, TaskSid string, params *ListFieldParams) ([]AutopilotV1Field, error) {
	if params == nil {
		params = &ListFieldParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageField(AssistantSid, TaskSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1Field

	for response != nil {
		records = append(records, response.Fields...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFieldResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFieldResponse)
	}

	return records, err
}

// Streams Field records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamField(AssistantSid string, TaskSid string, params *ListFieldParams) (chan AutopilotV1Field, error) {
	if params == nil {
		params = &ListFieldParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageField(AssistantSid, TaskSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1Field, 1)

	go func() {
		for response != nil {
			for item := range response.Fields {
				channel <- response.Fields[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFieldResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFieldResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFieldResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFieldResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
