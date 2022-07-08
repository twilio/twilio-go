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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateTask'
type CreateTaskParams struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task. It is optional and not unique.
	Actions *interface{} `json:"Actions,omitempty"`
	// The URL from which the Assistant can fetch actions.
	ActionsUrl *string `json:"ActionsUrl,omitempty"`
	// A descriptive string that you create to describe the new resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the new resource. It can be used as an alternative to the `sid` in the URL path to address the resource. This value must be unique and 64 characters or less in length.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateTaskParams) SetActions(Actions interface{}) *CreateTaskParams {
	params.Actions = &Actions
	return params
}
func (params *CreateTaskParams) SetActionsUrl(ActionsUrl string) *CreateTaskParams {
	params.ActionsUrl = &ActionsUrl
	return params
}
func (params *CreateTaskParams) SetFriendlyName(FriendlyName string) *CreateTaskParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTaskParams) SetUniqueName(UniqueName string) *CreateTaskParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateTask(AssistantSid string, params *CreateTaskParams) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Actions != nil {
		v, err := json.Marshal(params.Actions)

		if err != nil {
			return nil, err
		}

		data.Set("Actions", string(v))
	}
	if params != nil && params.ActionsUrl != nil {
		data.Set("ActionsUrl", *params.ActionsUrl)
	}
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

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTask(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
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

func (c *ApiService) FetchTask(AssistantSid string, Sid string) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTask'
type ListTaskParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTaskParams) SetPageSize(PageSize int) *ListTaskParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTaskParams) SetLimit(Limit int) *ListTaskParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Task records from the API. Request is executed immediately.
func (c *ApiService) PageTask(AssistantSid string, params *ListTaskParams, pageToken, pageNumber string) (*ListTaskResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks"

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

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Task records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTask(AssistantSid string, params *ListTaskParams) ([]AutopilotV1Task, error) {
	response, errors := c.StreamTask(AssistantSid, params)

	records := make([]AutopilotV1Task, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Task records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTask(AssistantSid string, params *ListTaskParams) (chan AutopilotV1Task, chan error) {
	if params == nil {
		params = &ListTaskParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AutopilotV1Task, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTask(AssistantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTask(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTask(response *ListTaskResponse, params *ListTaskParams, recordChannel chan AutopilotV1Task, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Tasks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListTaskResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTaskResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTaskResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTask'
type UpdateTaskParams struct {
	// The JSON string that specifies the [actions](https://www.twilio.com/docs/autopilot/actions) that instruct the Assistant on how to perform the task.
	Actions *interface{} `json:"Actions,omitempty"`
	// The URL from which the Assistant can fetch actions.
	ActionsUrl *string `json:"ActionsUrl,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 64 characters or less in length and be unique. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateTaskParams) SetActions(Actions interface{}) *UpdateTaskParams {
	params.Actions = &Actions
	return params
}
func (params *UpdateTaskParams) SetActionsUrl(ActionsUrl string) *UpdateTaskParams {
	params.ActionsUrl = &ActionsUrl
	return params
}
func (params *UpdateTaskParams) SetFriendlyName(FriendlyName string) *UpdateTaskParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTaskParams) SetUniqueName(UniqueName string) *UpdateTaskParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) UpdateTask(AssistantSid string, Sid string, params *UpdateTaskParams) (*AutopilotV1Task, error) {
	path := "/v1/Assistants/{AssistantSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Actions != nil {
		v, err := json.Marshal(params.Actions)

		if err != nil {
			return nil, err
		}

		data.Set("Actions", string(v))
	}
	if params != nil && params.ActionsUrl != nil {
		data.Set("ActionsUrl", *params.ActionsUrl)
	}
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

	ps := &AutopilotV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
