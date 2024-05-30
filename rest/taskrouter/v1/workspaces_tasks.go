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
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ghostmonitor/twilio-go/client"
)

// Optional parameters for the method 'CreateTask'
type CreateTaskParams struct {
	// The amount of time in seconds the new task can live before being assigned. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the `task.canceled` event will fire with description `Task TTL Exceeded`.
	Timeout *int `json:"Timeout,omitempty"`
	// The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647).
	Priority *int `json:"Priority,omitempty"`
	// When MultiTasking is enabled, specify the TaskChannel by passing either its `unique_name` or `sid`. Default value is `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
	// The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional.
	WorkflowSid *string `json:"WorkflowSid,omitempty"`
	// A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow's `assignment_callback_url` when the Task is assigned to a Worker. For example: `{ \\\"task_type\\\": \\\"call\\\", \\\"twilio_call_sid\\\": \\\"CAxxx\\\", \\\"customer_ticket_number\\\": \\\"12345\\\" }`.
	Attributes *string `json:"Attributes,omitempty"`
	// The virtual start time to assign the new task and override the default. When supplied, the new task will have this virtual start time. When not supplied, the new task will have the virtual start time equal to `date_created`. Value can't be in the future.
	VirtualStartTime *time.Time `json:"VirtualStartTime,omitempty"`
	// A SID of a Worker, Queue, or Workflow to route a Task to
	RoutingTarget *string `json:"RoutingTarget,omitempty"`
	// A boolean indicating if a new task should respect a worker's capacity during assignment
	IgnoreCapacity *string `json:"IgnoreCapacity,omitempty"`
	// The SID of the TaskQueue in which the Task belongs
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
}

func (params *CreateTaskParams) SetTimeout(Timeout int) *CreateTaskParams {
	params.Timeout = &Timeout
	return params
}
func (params *CreateTaskParams) SetPriority(Priority int) *CreateTaskParams {
	params.Priority = &Priority
	return params
}
func (params *CreateTaskParams) SetTaskChannel(TaskChannel string) *CreateTaskParams {
	params.TaskChannel = &TaskChannel
	return params
}
func (params *CreateTaskParams) SetWorkflowSid(WorkflowSid string) *CreateTaskParams {
	params.WorkflowSid = &WorkflowSid
	return params
}
func (params *CreateTaskParams) SetAttributes(Attributes string) *CreateTaskParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateTaskParams) SetVirtualStartTime(VirtualStartTime time.Time) *CreateTaskParams {
	params.VirtualStartTime = &VirtualStartTime
	return params
}
func (params *CreateTaskParams) SetRoutingTarget(RoutingTarget string) *CreateTaskParams {
	params.RoutingTarget = &RoutingTarget
	return params
}
func (params *CreateTaskParams) SetIgnoreCapacity(IgnoreCapacity string) *CreateTaskParams {
	params.IgnoreCapacity = &IgnoreCapacity
	return params
}
func (params *CreateTaskParams) SetTaskQueueSid(TaskQueueSid string) *CreateTaskParams {
	params.TaskQueueSid = &TaskQueueSid
	return params
}

//
func (c *ApiService) CreateTask(WorkspaceSid string, params *CreateTaskParams) (*TaskrouterV1Task, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Timeout != nil {
		data.Set("Timeout", fmt.Sprint(*params.Timeout))
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}
	if params != nil && params.WorkflowSid != nil {
		data.Set("WorkflowSid", *params.WorkflowSid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.VirtualStartTime != nil {
		data.Set("VirtualStartTime", fmt.Sprint((*params.VirtualStartTime).Format(time.RFC3339)))
	}
	if params != nil && params.RoutingTarget != nil {
		data.Set("RoutingTarget", *params.RoutingTarget)
	}
	if params != nil && params.IgnoreCapacity != nil {
		data.Set("IgnoreCapacity", *params.IgnoreCapacity)
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteTask'
type DeleteTaskParams struct {
	// If provided, deletes this Task if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
}

func (params *DeleteTaskParams) SetIfMatch(IfMatch string) *DeleteTaskParams {
	params.IfMatch = &IfMatch
	return params
}

//
func (c *ApiService) DeleteTask(WorkspaceSid string, Sid string, params *DeleteTaskParams) error {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}
	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

//
func (c *ApiService) FetchTask(WorkspaceSid string, Sid string) (*TaskrouterV1Task, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListTask'
type ListTaskParams struct {
	// The priority value of the Tasks to read. Returns the list of all Tasks in the Workspace with the specified priority.
	Priority *int `json:"Priority,omitempty"`
	// The `assignment_status` of the Tasks you want to read. Can be: `pending`, `reserved`, `assigned`, `canceled`, `wrapping`, or `completed`. Returns all Tasks in the Workspace with the specified `assignment_status`.
	AssignmentStatus *[]string `json:"AssignmentStatus,omitempty"`
	// The SID of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this SID.
	WorkflowSid *string `json:"WorkflowSid,omitempty"`
	// The friendly name of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this friendly name.
	WorkflowName *string `json:"WorkflowName,omitempty"`
	// The SID of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this SID.
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// The `friendly_name` of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this friendly name.
	TaskQueueName *string `json:"TaskQueueName,omitempty"`
	// The attributes of the Tasks to read. Returns the Tasks that match the attributes specified in this parameter.
	EvaluateTaskAttributes *string `json:"EvaluateTaskAttributes,omitempty"`
	// A SID of a Worker, Queue, or Workflow to route a Task to
	RoutingTarget *string `json:"RoutingTarget,omitempty"`
	// How to order the returned Task resources. By default, Tasks are sorted by ascending DateCreated. This value is specified as: `Attribute:Order`, where `Attribute` can be either `DateCreated`, `Priority`, or `VirtualStartTime` and `Order` can be either `asc` or `desc`. For example, `Priority:desc` returns Tasks ordered in descending order of their Priority. Pairings of sort orders can be specified in a comma-separated list such as `Priority:desc,DateCreated:asc`, which returns the Tasks in descending Priority order and ascending DateCreated Order. The only ordering pairing not allowed is DateCreated and VirtualStartTime.
	Ordering *string `json:"Ordering,omitempty"`
	// Whether to read Tasks with Add-ons. If `true`, returns only Tasks with Add-ons. If `false`, returns only Tasks without Add-ons.
	HasAddons *bool `json:"HasAddons,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListTaskParams) SetPriority(Priority int) *ListTaskParams {
	params.Priority = &Priority
	return params
}
func (params *ListTaskParams) SetAssignmentStatus(AssignmentStatus []string) *ListTaskParams {
	params.AssignmentStatus = &AssignmentStatus
	return params
}
func (params *ListTaskParams) SetWorkflowSid(WorkflowSid string) *ListTaskParams {
	params.WorkflowSid = &WorkflowSid
	return params
}
func (params *ListTaskParams) SetWorkflowName(WorkflowName string) *ListTaskParams {
	params.WorkflowName = &WorkflowName
	return params
}
func (params *ListTaskParams) SetTaskQueueSid(TaskQueueSid string) *ListTaskParams {
	params.TaskQueueSid = &TaskQueueSid
	return params
}
func (params *ListTaskParams) SetTaskQueueName(TaskQueueName string) *ListTaskParams {
	params.TaskQueueName = &TaskQueueName
	return params
}
func (params *ListTaskParams) SetEvaluateTaskAttributes(EvaluateTaskAttributes string) *ListTaskParams {
	params.EvaluateTaskAttributes = &EvaluateTaskAttributes
	return params
}
func (params *ListTaskParams) SetRoutingTarget(RoutingTarget string) *ListTaskParams {
	params.RoutingTarget = &RoutingTarget
	return params
}
func (params *ListTaskParams) SetOrdering(Ordering string) *ListTaskParams {
	params.Ordering = &Ordering
	return params
}
func (params *ListTaskParams) SetHasAddons(HasAddons bool) *ListTaskParams {
	params.HasAddons = &HasAddons
	return params
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
func (c *ApiService) PageTask(
	WorkspaceSid string,
	params *ListTaskParams,
	pageToken, pageNumber string,
) (*ListTaskResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.AssignmentStatus != nil {
		for _, item := range *params.AssignmentStatus {
			data.Add("AssignmentStatus", item)
		}
	}
	if params != nil && params.WorkflowSid != nil {
		data.Set("WorkflowSid", *params.WorkflowSid)
	}
	if params != nil && params.WorkflowName != nil {
		data.Set("WorkflowName", *params.WorkflowName)
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
	}
	if params != nil && params.TaskQueueName != nil {
		data.Set("TaskQueueName", *params.TaskQueueName)
	}
	if params != nil && params.EvaluateTaskAttributes != nil {
		data.Set("EvaluateTaskAttributes", *params.EvaluateTaskAttributes)
	}
	if params != nil && params.RoutingTarget != nil {
		data.Set("RoutingTarget", *params.RoutingTarget)
	}
	if params != nil && params.Ordering != nil {
		data.Set("Ordering", *params.Ordering)
	}
	if params != nil && params.HasAddons != nil {
		data.Set("HasAddons", fmt.Sprint(*params.HasAddons))
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
func (c *ApiService) ListTask(WorkspaceSid string, params *ListTaskParams) ([]TaskrouterV1Task, error) {
	response, errors := c.StreamTask(WorkspaceSid, params)

	records := make([]TaskrouterV1Task, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Task records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamTask(WorkspaceSid string, params *ListTaskParams) (chan TaskrouterV1Task, chan error) {
	if params == nil {
		params = &ListTaskParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1Task, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageTask(WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamTask(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamTask(
	response *ListTaskResponse,
	params *ListTaskParams,
	recordChannel chan TaskrouterV1Task,
	errorChannel chan error,
) {
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
	// If provided, applies this mutation if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
	// The JSON string that describes the custom attributes of the task.
	Attributes *string `json:"Attributes,omitempty"`
	//
	AssignmentStatus *string `json:"AssignmentStatus,omitempty"`
	// The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason.
	Reason *string `json:"Reason,omitempty"`
	// The Task's new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647).
	Priority *int `json:"Priority,omitempty"`
	// When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
	TaskChannel *string `json:"TaskChannel,omitempty"`
	// The task's new virtual start time value. When supplied, the Task takes on the specified virtual start time. Value can't be in the future.
	VirtualStartTime *time.Time `json:"VirtualStartTime,omitempty"`
}

func (params *UpdateTaskParams) SetIfMatch(IfMatch string) *UpdateTaskParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateTaskParams) SetAttributes(Attributes string) *UpdateTaskParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateTaskParams) SetAssignmentStatus(AssignmentStatus string) *UpdateTaskParams {
	params.AssignmentStatus = &AssignmentStatus
	return params
}
func (params *UpdateTaskParams) SetReason(Reason string) *UpdateTaskParams {
	params.Reason = &Reason
	return params
}
func (params *UpdateTaskParams) SetPriority(Priority int) *UpdateTaskParams {
	params.Priority = &Priority
	return params
}
func (params *UpdateTaskParams) SetTaskChannel(TaskChannel string) *UpdateTaskParams {
	params.TaskChannel = &TaskChannel
	return params
}
func (params *UpdateTaskParams) SetVirtualStartTime(VirtualStartTime time.Time) *UpdateTaskParams {
	params.VirtualStartTime = &VirtualStartTime
	return params
}

//
func (c *ApiService) UpdateTask(WorkspaceSid string, Sid string, params *UpdateTaskParams) (*TaskrouterV1Task, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Tasks/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.AssignmentStatus != nil {
		data.Set("AssignmentStatus", *params.AssignmentStatus)
	}
	if params != nil && params.Reason != nil {
		data.Set("Reason", *params.Reason)
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", fmt.Sprint(*params.Priority))
	}
	if params != nil && params.TaskChannel != nil {
		data.Set("TaskChannel", *params.TaskChannel)
	}
	if params != nil && params.VirtualStartTime != nil {
		data.Set("VirtualStartTime", fmt.Sprint((*params.VirtualStartTime).Format(time.RFC3339)))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Task{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
