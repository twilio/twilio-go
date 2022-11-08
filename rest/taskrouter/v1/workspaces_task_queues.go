/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateTaskQueue'
type CreateTaskQueueParams struct {
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, `'\\\"language\\\" == \\\"spanish\\\"'`. The default value is `1==1`. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers).
	TargetWorkers *string `json:"TargetWorkers,omitempty"`
	// The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
	MaxReservedWorkers *int `json:"MaxReservedWorkers,omitempty"`
	//
	TaskOrder *string `json:"TaskOrder,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid *string `json:"ReservationActivitySid,omitempty"`
	// The SID of the Activity to assign Workers when a task is assigned to them.
	AssignmentActivitySid *string `json:"AssignmentActivitySid,omitempty"`
}

func (params *CreateTaskQueueParams) SetFriendlyName(FriendlyName string) *CreateTaskQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateTaskQueueParams) SetTargetWorkers(TargetWorkers string) *CreateTaskQueueParams {
	params.TargetWorkers = &TargetWorkers
	return params
}
func (params *CreateTaskQueueParams) SetMaxReservedWorkers(MaxReservedWorkers int) *CreateTaskQueueParams {
	params.MaxReservedWorkers = &MaxReservedWorkers
	return params
}
func (params *CreateTaskQueueParams) SetTaskOrder(TaskOrder string) *CreateTaskQueueParams {
	params.TaskOrder = &TaskOrder
	return params
}
func (params *CreateTaskQueueParams) SetReservationActivitySid(ReservationActivitySid string) *CreateTaskQueueParams {
	params.ReservationActivitySid = &ReservationActivitySid
	return params
}
func (params *CreateTaskQueueParams) SetAssignmentActivitySid(AssignmentActivitySid string) *CreateTaskQueueParams {
	params.AssignmentActivitySid = &AssignmentActivitySid
	return params
}

func (c *ApiService) CreateTaskQueue(WorkspaceSid string, params *CreateTaskQueueParams) (*TaskrouterV1TaskQueue, error) {
	return c.CreateTaskQueueWithCtx(context.TODO(), WorkspaceSid, params)
}

func (c *ApiService) CreateTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, params *CreateTaskQueueParams) (*TaskrouterV1TaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TargetWorkers != nil {
		data.Set("TargetWorkers", *params.TargetWorkers)
	}
	if params != nil && params.MaxReservedWorkers != nil {
		data.Set("MaxReservedWorkers", fmt.Sprint(*params.MaxReservedWorkers))
	}
	if params != nil && params.TaskOrder != nil {
		data.Set("TaskOrder", *params.TaskOrder)
	}
	if params != nil && params.ReservationActivitySid != nil {
		data.Set("ReservationActivitySid", *params.ReservationActivitySid)
	}
	if params != nil && params.AssignmentActivitySid != nil {
		data.Set("AssignmentActivitySid", *params.AssignmentActivitySid)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteTaskQueue(WorkspaceSid string, Sid string) error {
	return c.DeleteTaskQueueWithCtx(context.TODO(), WorkspaceSid, Sid)
}

func (c *ApiService) DeleteTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
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

func (c *ApiService) FetchTaskQueue(WorkspaceSid string, Sid string) (*TaskrouterV1TaskQueue, error) {
	return c.FetchTaskQueueWithCtx(context.TODO(), WorkspaceSid, Sid)
}

func (c *ApiService) FetchTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, Sid string) (*TaskrouterV1TaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskQueue{}
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
	// Sorting parameter for TaskQueues
	Ordering *string `json:"Ordering,omitempty"`
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
func (params *ListTaskQueueParams) SetOrdering(Ordering string) *ListTaskQueueParams {
	params.Ordering = &Ordering
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
func (c *ApiService) PageTaskQueue(WorkspaceSid string, params *ListTaskQueueParams, pageToken, pageNumber string) (*ListTaskQueueResponse, error) {
	return c.PageTaskQueueWithCtx(context.TODO(), WorkspaceSid, params, pageToken, pageNumber)
}

// Retrieve a single page of TaskQueue records from the API. Request is executed immediately.
func (c *ApiService) PageTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, params *ListTaskQueueParams, pageToken, pageNumber string) (*ListTaskQueueResponse, error) {
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
	if params != nil && params.Ordering != nil {
		data.Set("Ordering", *params.Ordering)
	}
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

	ps := &ListTaskQueueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists TaskQueue records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTaskQueue(WorkspaceSid string, params *ListTaskQueueParams) ([]TaskrouterV1TaskQueue, error) {
	return c.ListTaskQueueWithCtx(context.TODO(), WorkspaceSid, params)
}

// Lists TaskQueue records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, params *ListTaskQueueParams) ([]TaskrouterV1TaskQueue, error) {
	response, errors := c.StreamTaskQueueWithCtx(ctx, WorkspaceSid, params)

	records := make([]TaskrouterV1TaskQueue, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams TaskQueue records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTaskQueue(WorkspaceSid string, params *ListTaskQueueParams) (chan TaskrouterV1TaskQueue, chan error) {
	return c.StreamTaskQueueWithCtx(context.TODO(), WorkspaceSid, params)
}

// Streams TaskQueue records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, params *ListTaskQueueParams) (chan TaskrouterV1TaskQueue, chan error) {
	if params == nil {
		params = &ListTaskQueueParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1TaskQueue, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTaskQueueWithCtx(ctx, WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTaskQueue(ctx, response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTaskQueue(ctx context.Context, response *ListTaskQueueResponse, params *ListTaskQueueParams, recordChannel chan TaskrouterV1TaskQueue, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.TaskQueues
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNextWithCtx(ctx, c.baseURL, response, c.getNextListTaskQueueResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListTaskQueueResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListTaskQueueResponse(ctx context.Context, nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(ctx, nextPageUrl, nil, nil)
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
	// A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example '\\\"language\\\" == \\\"spanish\\\"' If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below.
	TargetWorkers *string `json:"TargetWorkers,omitempty"`
	// The SID of the Activity to assign Workers when a task is reserved for them.
	ReservationActivitySid *string `json:"ReservationActivitySid,omitempty"`
	// The SID of the Activity to assign Workers when a task is assigned for them.
	AssignmentActivitySid *string `json:"AssignmentActivitySid,omitempty"`
	// The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50.
	MaxReservedWorkers *int `json:"MaxReservedWorkers,omitempty"`
	//
	TaskOrder *string `json:"TaskOrder,omitempty"`
}

func (params *UpdateTaskQueueParams) SetFriendlyName(FriendlyName string) *UpdateTaskQueueParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateTaskQueueParams) SetTargetWorkers(TargetWorkers string) *UpdateTaskQueueParams {
	params.TargetWorkers = &TargetWorkers
	return params
}
func (params *UpdateTaskQueueParams) SetReservationActivitySid(ReservationActivitySid string) *UpdateTaskQueueParams {
	params.ReservationActivitySid = &ReservationActivitySid
	return params
}
func (params *UpdateTaskQueueParams) SetAssignmentActivitySid(AssignmentActivitySid string) *UpdateTaskQueueParams {
	params.AssignmentActivitySid = &AssignmentActivitySid
	return params
}
func (params *UpdateTaskQueueParams) SetMaxReservedWorkers(MaxReservedWorkers int) *UpdateTaskQueueParams {
	params.MaxReservedWorkers = &MaxReservedWorkers
	return params
}
func (params *UpdateTaskQueueParams) SetTaskOrder(TaskOrder string) *UpdateTaskQueueParams {
	params.TaskOrder = &TaskOrder
	return params
}

func (c *ApiService) UpdateTaskQueue(WorkspaceSid string, Sid string, params *UpdateTaskQueueParams) (*TaskrouterV1TaskQueue, error) {
	return c.UpdateTaskQueueWithCtx(context.TODO(), WorkspaceSid, Sid, params)
}

func (c *ApiService) UpdateTaskQueueWithCtx(ctx context.Context, WorkspaceSid string, Sid string, params *UpdateTaskQueueParams) (*TaskrouterV1TaskQueue, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TargetWorkers != nil {
		data.Set("TargetWorkers", *params.TargetWorkers)
	}
	if params != nil && params.ReservationActivitySid != nil {
		data.Set("ReservationActivitySid", *params.ReservationActivitySid)
	}
	if params != nil && params.AssignmentActivitySid != nil {
		data.Set("AssignmentActivitySid", *params.AssignmentActivitySid)
	}
	if params != nil && params.MaxReservedWorkers != nil {
		data.Set("MaxReservedWorkers", fmt.Sprint(*params.MaxReservedWorkers))
	}
	if params != nil && params.TaskOrder != nil {
		data.Set("TaskOrder", *params.TaskOrder)
	}

	resp, err := c.requestHandler.Post(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1TaskQueue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
