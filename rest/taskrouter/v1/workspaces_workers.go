/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.18.0
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

// Optional parameters for the method 'CreateWorker'
type CreateWorkerParams struct {
	// The SID of a valid Activity that will describe the new Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker's initial state is the `default_activity_sid` configured on the Workspace.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// A valid JSON string that describes the new Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateWorkerParams) SetActivitySid(ActivitySid string) *CreateWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *CreateWorkerParams) SetAttributes(Attributes string) *CreateWorkerParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateWorkerParams) SetFriendlyName(FriendlyName string) *CreateWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) CreateWorker(WorkspaceSid string, params *CreateWorkerParams) (*TaskrouterV1WorkspaceWorker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWorker(WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
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

func (c *ApiService) FetchWorker(WorkspaceSid string, Sid string) (*TaskrouterV1WorkspaceWorker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorker'
type ListWorkerParams struct {
	// The `activity_name` of the Worker resources to read.
	ActivityName *string `json:"ActivityName,omitempty"`
	// The `activity_sid` of the Worker resources to read.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// Whether to return only Worker resources that are available or unavailable. Can be `true`, `1`, or `yes` to return Worker resources that are available, and `false`, or any value returns the Worker resources that are not available.
	Available *string `json:"Available,omitempty"`
	// The `friendly_name` of the Worker resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue.
	TargetWorkersExpression *string `json:"TargetWorkersExpression,omitempty"`
	// The `friendly_name` of the TaskQueue that the Workers to read are eligible for.
	TaskQueueName *string `json:"TaskQueueName,omitempty"`
	// The SID of the TaskQueue that the Workers to read are eligible for.
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListWorkerParams) SetActivityName(ActivityName string) *ListWorkerParams {
	params.ActivityName = &ActivityName
	return params
}
func (params *ListWorkerParams) SetActivitySid(ActivitySid string) *ListWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *ListWorkerParams) SetAvailable(Available string) *ListWorkerParams {
	params.Available = &Available
	return params
}
func (params *ListWorkerParams) SetFriendlyName(FriendlyName string) *ListWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkerParams) SetTargetWorkersExpression(TargetWorkersExpression string) *ListWorkerParams {
	params.TargetWorkersExpression = &TargetWorkersExpression
	return params
}
func (params *ListWorkerParams) SetTaskQueueName(TaskQueueName string) *ListWorkerParams {
	params.TaskQueueName = &TaskQueueName
	return params
}
func (params *ListWorkerParams) SetTaskQueueSid(TaskQueueSid string) *ListWorkerParams {
	params.TaskQueueSid = &TaskQueueSid
	return params
}
func (params *ListWorkerParams) SetPageSize(PageSize int) *ListWorkerParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListWorker(WorkspaceSid string, params *ListWorkerParams) (*ListWorkerResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivityName != nil {
		data.Set("ActivityName", *params.ActivityName)
	}
	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Available != nil {
		data.Set("Available", *params.Available)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TargetWorkersExpression != nil {
		data.Set("TargetWorkersExpression", *params.TargetWorkersExpression)
	}
	if params != nil && params.TaskQueueName != nil {
		data.Set("TaskQueueName", *params.TaskQueueName)
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkerResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//Retrieve a single page of Worker records from the API. Request is executed immediately.
func (c *ApiService) WorkerPage(WorkspaceSid string, params *ListWorkerParams, pageToken string, pageNumber string, pageSize string) *client.Page {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivityName != nil {
		data.Set("ActivityName", *params.ActivityName)
	}
	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Available != nil {
		data.Set("Available", *params.Available)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TargetWorkersExpression != nil {
		data.Set("TargetWorkersExpression", *params.TargetWorkersExpression)
	}
	if params != nil && params.TaskQueueName != nil {
		data.Set("TaskQueueName", *params.TaskQueueName)
	}
	if params != nil && params.TaskQueueSid != nil {
		data.Set("TaskQueueSid", *params.TaskQueueSid)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	data.Set("PageToken", pageToken)
	data.Set("PageNumber", pageNumber)
	data.Set("PageSize", pageSize)

	response, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil
	}

	return client.NewPage(c.baseURL, response)
}

//Streams Worker records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) WorkerStream(WorkspaceSid string, params *ListWorkerParams, meta client.PaginationData) chan map[string]interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.WorkerPage(WorkspaceSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.Stream(page, limits.Limit, limits.PageLimit)
}

//Lists Worker records from the API as a list. Unlike stream, this operation is eager and will loads 'limit' records into memory before returning.
func (c *ApiService) WorkerList(WorkspaceSid string, params *ListWorkerParams, meta client.PaginationData) []interface{} {
	limits := c.requestHandler.ReadLimits(meta)
	page := c.WorkerPage(WorkspaceSid, params, "", "", fmt.Sprint(limits.PageSize))
	return c.requestHandler.List(page, limits.Limit, limits.PageLimit)
}

// Optional parameters for the method 'UpdateWorker'
type UpdateWorkerParams struct {
	// The SID of a valid Activity that will describe the Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information.
	ActivitySid *string `json:"ActivitySid,omitempty"`
	// The JSON string that describes the Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}.
	Attributes *string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the Worker. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to reject pending reservations.
	RejectPendingReservations *bool `json:"RejectPendingReservations,omitempty"`
}

func (params *UpdateWorkerParams) SetActivitySid(ActivitySid string) *UpdateWorkerParams {
	params.ActivitySid = &ActivitySid
	return params
}
func (params *UpdateWorkerParams) SetAttributes(Attributes string) *UpdateWorkerParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateWorkerParams) SetFriendlyName(FriendlyName string) *UpdateWorkerParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkerParams) SetRejectPendingReservations(RejectPendingReservations bool) *UpdateWorkerParams {
	params.RejectPendingReservations = &RejectPendingReservations
	return params
}

func (c *ApiService) UpdateWorker(WorkspaceSid string, Sid string, params *UpdateWorkerParams) (*TaskrouterV1WorkspaceWorker, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workers/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActivitySid != nil {
		data.Set("ActivitySid", *params.ActivitySid)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.RejectPendingReservations != nil {
		data.Set("RejectPendingReservations", fmt.Sprint(*params.RejectPendingReservations))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorker{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
