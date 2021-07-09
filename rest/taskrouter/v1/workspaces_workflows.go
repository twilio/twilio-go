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
)

// Optional parameters for the method 'CreateWorkflow'
type CreateWorkflowParams struct {
	// The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
	AssignmentCallbackUrl *string `json:"AssignmentCallbackUrl,omitempty"`
	// A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
	Configuration *string `json:"Configuration,omitempty"`
	// The URL that we should call when a call to the `assignment_callback_url` fails.
	FallbackAssignmentCallbackUrl *string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	// A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.
	TaskReservationTimeout *int `json:"TaskReservationTimeout,omitempty"`
}

func (params *CreateWorkflowParams) SetAssignmentCallbackUrl(AssignmentCallbackUrl string) *CreateWorkflowParams {
	params.AssignmentCallbackUrl = &AssignmentCallbackUrl
	return params
}
func (params *CreateWorkflowParams) SetConfiguration(Configuration string) *CreateWorkflowParams {
	params.Configuration = &Configuration
	return params
}
func (params *CreateWorkflowParams) SetFallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl string) *CreateWorkflowParams {
	params.FallbackAssignmentCallbackUrl = &FallbackAssignmentCallbackUrl
	return params
}
func (params *CreateWorkflowParams) SetFriendlyName(FriendlyName string) *CreateWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateWorkflowParams) SetTaskReservationTimeout(TaskReservationTimeout int) *CreateWorkflowParams {
	params.TaskReservationTimeout = &TaskReservationTimeout
	return params
}

func (c *ApiService) CreateWorkflow(WorkspaceSid string, params *CreateWorkflowParams) (*TaskrouterV1WorkspaceWorkflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	if params != nil && params.AssignmentCallbackUrl != nil {
		data.Set("AssignmentCallbackUrl", *params.AssignmentCallbackUrl)
	}
	if params != nil && params.Configuration != nil {
		data.Set("Configuration", *params.Configuration)
	}
	if params != nil && params.FallbackAssignmentCallbackUrl != nil {
		data.Set("FallbackAssignmentCallbackUrl", *params.FallbackAssignmentCallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.TaskReservationTimeout != nil {
		data.Set("TaskReservationTimeout", fmt.Sprint(*params.TaskReservationTimeout))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorkflow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWorkflow(WorkspaceSid string, Sid string) error {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows/{Sid}"
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

func (c *ApiService) FetchWorkflow(WorkspaceSid string, Sid string) (*TaskrouterV1WorkspaceWorkflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorkflow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorkflow'
type ListWorkflowParams struct {
	// The `friendly_name` of the Workflow resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListWorkflowParams) SetFriendlyName(FriendlyName string) *ListWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkflowParams) SetPageSize(PageSize int) *ListWorkflowParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListWorkflow(WorkspaceSid string, params *ListWorkflowParams) (*ListWorkflowResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkflowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateWorkflow'
type UpdateWorkflowParams struct {
	// The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
	AssignmentCallbackUrl *string `json:"AssignmentCallbackUrl,omitempty"`
	// A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
	Configuration *string `json:"Configuration,omitempty"`
	// The URL that we should call when a call to the `assignment_callback_url` fails.
	FallbackAssignmentCallbackUrl *string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	// A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether or not to re-evaluate Tasks. The default is `false`, which means Tasks in the Workflow will not be processed through the assignment loop again.
	ReEvaluateTasks *string `json:"ReEvaluateTasks,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.
	TaskReservationTimeout *int `json:"TaskReservationTimeout,omitempty"`
}

func (params *UpdateWorkflowParams) SetAssignmentCallbackUrl(AssignmentCallbackUrl string) *UpdateWorkflowParams {
	params.AssignmentCallbackUrl = &AssignmentCallbackUrl
	return params
}
func (params *UpdateWorkflowParams) SetConfiguration(Configuration string) *UpdateWorkflowParams {
	params.Configuration = &Configuration
	return params
}
func (params *UpdateWorkflowParams) SetFallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl string) *UpdateWorkflowParams {
	params.FallbackAssignmentCallbackUrl = &FallbackAssignmentCallbackUrl
	return params
}
func (params *UpdateWorkflowParams) SetFriendlyName(FriendlyName string) *UpdateWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkflowParams) SetReEvaluateTasks(ReEvaluateTasks string) *UpdateWorkflowParams {
	params.ReEvaluateTasks = &ReEvaluateTasks
	return params
}
func (params *UpdateWorkflowParams) SetTaskReservationTimeout(TaskReservationTimeout int) *UpdateWorkflowParams {
	params.TaskReservationTimeout = &TaskReservationTimeout
	return params
}

func (c *ApiService) UpdateWorkflow(WorkspaceSid string, Sid string, params *UpdateWorkflowParams) (*TaskrouterV1WorkspaceWorkflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.AssignmentCallbackUrl != nil {
		data.Set("AssignmentCallbackUrl", *params.AssignmentCallbackUrl)
	}
	if params != nil && params.Configuration != nil {
		data.Set("Configuration", *params.Configuration)
	}
	if params != nil && params.FallbackAssignmentCallbackUrl != nil {
		data.Set("FallbackAssignmentCallbackUrl", *params.FallbackAssignmentCallbackUrl)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ReEvaluateTasks != nil {
		data.Set("ReEvaluateTasks", *params.ReEvaluateTasks)
	}
	if params != nil && params.TaskReservationTimeout != nil {
		data.Set("TaskReservationTimeout", fmt.Sprint(*params.TaskReservationTimeout))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1WorkspaceWorkflow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
