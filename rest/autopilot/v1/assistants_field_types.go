/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateFieldType'
type CreateFieldTypeParams struct {
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. The first 64 characters must be unique.
	UniqueName *string `json:"UniqueName,omitempty"`
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateFieldTypeParams) SetUniqueName(UniqueName string) *CreateFieldTypeParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateFieldTypeParams) SetFriendlyName(FriendlyName string) *CreateFieldTypeParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateFieldType(AssistantSid string, params *CreateFieldTypeParams) (*AutopilotV1FieldType, error) {
	return c.CreateFieldTypeWithCtx(context.TODO(), AssistantSid, params)
}

func (c *ApiService) CreateFieldTypeWithCtx(ctx context.Context, AssistantSid string, params *CreateFieldTypeParams) (*AutopilotV1FieldType, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1FieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteFieldType(AssistantSid string, Sid string) error {
	return c.DeleteFieldTypeWithCtx(context.TODO(), AssistantSid, Sid)
}

func (c *ApiService) DeleteFieldTypeWithCtx(ctx context.Context, AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchFieldType(AssistantSid string, Sid string) (*AutopilotV1FieldType, error) {
	return c.FetchFieldTypeWithCtx(context.TODO(), AssistantSid, Sid)
}

func (c *ApiService) FetchFieldTypeWithCtx(ctx context.Context, AssistantSid string, Sid string) (*AutopilotV1FieldType, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1FieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFieldType'
type ListFieldTypeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFieldTypeParams) SetPageSize(PageSize int) *ListFieldTypeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFieldTypeParams) SetLimit(Limit int) *ListFieldTypeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FieldType records from the API. Request is executed immediately.
func (c *ApiService) PageFieldType(AssistantSid string, params *ListFieldTypeParams, pageToken, pageNumber string) (*ListFieldTypeResponse, error) {
	return c.PageFieldTypeWithCtx(context.TODO(), AssistantSid, params, pageToken, pageNumber)
}

// Retrieve a single page of FieldType records from the API. Request is executed immediately.
func (c *ApiService) PageFieldTypeWithCtx(ctx context.Context, AssistantSid string, params *ListFieldTypeParams, pageToken, pageNumber string) (*ListFieldTypeResponse, error) {
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
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
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
func (c *ApiService) ListFieldType(AssistantSid string, params *ListFieldTypeParams) ([]AutopilotV1FieldType, error) {
	return c.ListFieldTypeWithCtx(context.TODO(), AssistantSid, params)
}

// Lists FieldType records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFieldTypeWithCtx(ctx context.Context, AssistantSid string, params *ListFieldTypeParams) ([]AutopilotV1FieldType, error) {
	response, errors := c.StreamFieldTypeWithCtx(ctx, AssistantSid, params)

	records := make([]AutopilotV1FieldType, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams FieldType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFieldType(AssistantSid string, params *ListFieldTypeParams) (chan AutopilotV1FieldType, chan error) {
	return c.StreamFieldTypeWithCtx(context.TODO(), AssistantSid, params)
}

// Streams FieldType records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFieldTypeWithCtx(ctx context.Context, AssistantSid string, params *ListFieldTypeParams) (chan AutopilotV1FieldType, chan error) {
	if params == nil {
		params = &ListFieldTypeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AutopilotV1FieldType, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFieldTypeWithCtx(ctx, AssistantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFieldType(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFieldType(ctx context.Context, response *ListFieldTypeResponse, params *ListFieldTypeParams, recordChannel chan AutopilotV1FieldType, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.FieldTypes
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListFieldTypeResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFieldTypeResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFieldTypeResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
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

func (c *ApiService) UpdateFieldType(AssistantSid string, Sid string, params *UpdateFieldTypeParams) (*AutopilotV1FieldType, error) {
	return c.UpdateFieldTypeWithCtx(context.TODO(), AssistantSid, Sid, params)
}

func (c *ApiService) UpdateFieldTypeWithCtx(ctx context.Context, AssistantSid string, Sid string, params *UpdateFieldTypeParams) (*AutopilotV1FieldType, error) {
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

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1FieldType{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
