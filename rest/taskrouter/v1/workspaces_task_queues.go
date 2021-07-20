/*
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateTaskQueue'
type CreateTaskQueueParams struct {
	// The SID of the Activity to assign Workers when a task is assigned to them.
	AssignmentActivitySid *string `json:"AssignmentActivitySid,omitempty"`
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
	MaxReservedWorkers *int `json:"MaxReservedWorkers,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid *string `json:"ReservationActivitySid,omitempty"`
	// A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, `'\\\"language\\\" == \\\"spanish\\\"'`. The default value is `1==1`. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers).
	TargetWorkers *string `json:"TargetWorkers,omitempty"`
	// How Tasks will be assigned to Workers. Set this parameter to `LIFO` to assign most recently created Task first or FIFO to assign the oldest Task first. Default is `FIFO`. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more.
	TaskOrder *string `json:"TaskOrder,omitempty"`
}

func (params *CreateTaskQueueParams) SetAssignmentActivitySid(AssignmentActivitySid string) *CreateTaskQueueParams {
	params.AssignmentActivitySid = &AssignmentActivitySid
	return params
}
func (params *CreateTaskQueueParams) SetFriendlyName(FriendlyName string) *CreateTaskQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTaskQueueParams) SetMaxReservedWorkers(MaxReservedWorkers int) *CreateTaskQueueParams {
	params.MaxReservedWorkers = &MaxReservedWorkers
	return params
}
func (params *CreateTaskQueueParams) SetReservationActivitySid(ReservationActivitySid string) *CreateTaskQueueParams {
	params.ReservationActivitySid = &ReservationActivitySid
	return params
}
func (params *CreateTaskQueueParams) SetTargetWorkers(TargetWorkers string) *CreateTaskQueueParams {
	params.TargetWorkers = &TargetWorkers
	return params
}
func (params *CreateTaskQueueParams) SetTaskOrder(TaskOrder string) *CreateTaskQueueParams {
	params.TaskOrder = &TaskOrder
	return params
}

func (c *ApiService) CreateTaskQueue(WorkspaceSid string, params *CreateTaskQueueParams) (*TaskrouterV1WorkspaceTaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AssignmentActivitySid != nil {
		data.Set("AssignmentActivitySid", *params.AssignmentActivitySid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MaxReservedWorkers != nil {
		data.Set("MaxReservedWorkers", fmt.Sprint(*params.MaxReservedWorkers))
	}
	if params != nil && params.ReservationActivitySid != nil {
		data.Set("ReservationActivitySid", *params.ReservationActivitySid)
	}
	if params != nil && params.TargetWorkers != nil {
		data.Set("TargetWorkers", *params.TargetWorkers)
	}
	if params != nil && params.TaskOrder != nil {
		data.Set("TaskOrder", *params.TaskOrder)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTaskQueue(WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
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

func (c *ApiService) FetchTaskQueue(WorkspaceSid string, Sid string) (*TaskrouterV1WorkspaceTaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTaskQueue'
type ListTaskQueueParams struct {
	// The `friendly_name` of the TaskQueue resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The attributes of the Workers to read. Returns the TaskQueues with Workers that match the attributes specified in this parameter.
	EvaluateWorkerAttributes *string `json:"EvaluateWorkerAttributes,omitempty"`
	// The SID of the Worker with the TaskQueue resources to read.
	WorkerSid *string `json:"WorkerSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTaskQueueParams) SetFriendlyName(FriendlyName string) *ListTaskQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListTaskQueueParams) SetEvaluateWorkerAttributes(EvaluateWorkerAttributes string) *ListTaskQueueParams {
	params.EvaluateWorkerAttributes = &EvaluateWorkerAttributes
	return params
}
func (params *ListTaskQueueParams) SetWorkerSid(WorkerSid string) *ListTaskQueueParams {
	params.WorkerSid = &WorkerSid
	return params
}
func (params *ListTaskQueueParams) SetPageSize(PageSize int) *ListTaskQueueParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListTaskQueueParams) SetLimit(Limit int) *ListTaskQueueParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of TaskQueue records from the API. Request is executed immediately.
func (c *ApiService) PageTaskQueue(WorkspaceSid string, params *ListTaskQueueParams, pageToken string, pageNumber string) (*ListTaskQueueResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.EvaluateWorkerAttributes != nil {
		data.Set("EvaluateWorkerAttributes", *params.EvaluateWorkerAttributes)
	}
	if params != nil && params.WorkerSid != nil {
		data.Set("WorkerSid", *params.WorkerSid)
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

	ps := &ListTaskQueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TaskQueue records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTaskQueue(WorkspaceSid string, params *ListTaskQueueParams) ([]TaskrouterV1WorkspaceTaskQueue, error) {
	if params == nil {
		params = &ListTaskQueueParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTaskQueue(WorkspaceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []TaskrouterV1WorkspaceTaskQueue

	for response != nil {
		records = append(records, response.TaskQueues...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListTaskQueueResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListTaskQueueResponse)
	}

	return records, err
}

// Streams TaskQueue records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTaskQueue(WorkspaceSid string, params *ListTaskQueueParams) (chan TaskrouterV1WorkspaceTaskQueue, error) {
	if params == nil {
		params = &ListTaskQueueParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageTaskQueue(WorkspaceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan TaskrouterV1WorkspaceTaskQueue, 1)

	go func() {
		for response != nil {
			for item := range response.TaskQueues {
				channel <- response.TaskQueues[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListTaskQueueResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListTaskQueueResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListTaskQueueResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListTaskQueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateTaskQueue'
type UpdateTaskQueueParams struct {
	// The SID of the Activity to assign Workers when a task is assigned for them.
	AssignmentActivitySid *string `json:"AssignmentActivitySid,omitempty"`
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50.
	MaxReservedWorkers *int `json:"MaxReservedWorkers,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid *string `json:"ReservationActivitySid,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example '\\\"language\\\" == \\\"spanish\\\"' If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below.
	TargetWorkers *string `json:"TargetWorkers,omitempty"`
	// How Tasks will be assigned to Workers. Can be: `FIFO` or `LIFO` and the default is `FIFO`. Use `FIFO` to assign the oldest task first and `LIFO` to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
	TaskOrder *string `json:"TaskOrder,omitempty"`
}

func (params *UpdateTaskQueueParams) SetAssignmentActivitySid(AssignmentActivitySid string) *UpdateTaskQueueParams {
	params.AssignmentActivitySid = &AssignmentActivitySid
	return params
}
func (params *UpdateTaskQueueParams) SetFriendlyName(FriendlyName string) *UpdateTaskQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTaskQueueParams) SetMaxReservedWorkers(MaxReservedWorkers int) *UpdateTaskQueueParams {
	params.MaxReservedWorkers = &MaxReservedWorkers
	return params
}
func (params *UpdateTaskQueueParams) SetReservationActivitySid(ReservationActivitySid string) *UpdateTaskQueueParams {
	params.ReservationActivitySid = &ReservationActivitySid
	return params
}
func (params *UpdateTaskQueueParams) SetTargetWorkers(TargetWorkers string) *UpdateTaskQueueParams {
	params.TargetWorkers = &TargetWorkers
	return params
}
func (params *UpdateTaskQueueParams) SetTaskOrder(TaskOrder string) *UpdateTaskQueueParams {
	params.TaskOrder = &TaskOrder
	return params
}

func (c *ApiService) UpdateTaskQueue(WorkspaceSid string, Sid string, params *UpdateTaskQueueParams) (*TaskrouterV1WorkspaceTaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.AssignmentActivitySid != nil {
		data.Set("AssignmentActivitySid", *params.AssignmentActivitySid)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MaxReservedWorkers != nil {
		data.Set("MaxReservedWorkers", fmt.Sprint(*params.MaxReservedWorkers))
	}
	if params != nil && params.ReservationActivitySid != nil {
		data.Set("ReservationActivitySid", *params.ReservationActivitySid)
	}
	if params != nil && params.TargetWorkers != nil {
		data.Set("TargetWorkers", *params.TargetWorkers)
	}
	if params != nil && params.TaskOrder != nil {
		data.Set("TaskOrder", *params.TaskOrder)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceTaskQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
