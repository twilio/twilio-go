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

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateWorkflow'
type CreateWorkflowParams struct {
	// A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
	Configuration *string `json:"Configuration,omitempty"`
	// The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
	AssignmentCallbackUrl *string `json:"AssignmentCallbackUrl,omitempty"`
	// The URL that we should call when a call to the `assignment_callback_url` fails.
	FallbackAssignmentCallbackUrl *string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.
	TaskReservationTimeout *int `json:"TaskReservationTimeout,omitempty"`
}

func (params *CreateWorkflowParams) SetFriendlyName(FriendlyName string) *CreateWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateWorkflowParams) SetConfiguration(Configuration string) *CreateWorkflowParams {
	params.Configuration = &Configuration
	return params
}
func (params *CreateWorkflowParams) SetAssignmentCallbackUrl(AssignmentCallbackUrl string) *CreateWorkflowParams {
	params.AssignmentCallbackUrl = &AssignmentCallbackUrl
	return params
}
func (params *CreateWorkflowParams) SetFallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl string) *CreateWorkflowParams {
	params.FallbackAssignmentCallbackUrl = &FallbackAssignmentCallbackUrl
	return params
}
func (params *CreateWorkflowParams) SetTaskReservationTimeout(TaskReservationTimeout int) *CreateWorkflowParams {
	params.TaskReservationTimeout = &TaskReservationTimeout
	return params
}

func (c *ApiService) CreateWorkflow(WorkspaceSid string, params *CreateWorkflowParams) (*TaskrouterV1Workflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Configuration != nil {
		data.Set("Configuration", *params.Configuration)
	}
	if params != nil && params.AssignmentCallbackUrl != nil {
		data.Set("AssignmentCallbackUrl", *params.AssignmentCallbackUrl)
	}
	if params != nil && params.FallbackAssignmentCallbackUrl != nil {
		data.Set("FallbackAssignmentCallbackUrl", *params.FallbackAssignmentCallbackUrl)
	}
	if params != nil && params.TaskReservationTimeout != nil {
		data.Set("TaskReservationTimeout", fmt.Sprint(*params.TaskReservationTimeout))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workflow{}
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
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchWorkflow(WorkspaceSid string, Sid string) (*TaskrouterV1Workflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workflow{}
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
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWorkflowParams) SetFriendlyName(FriendlyName string) *ListWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkflowParams) SetPageSize(PageSize int) *ListWorkflowParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWorkflowParams) SetLimit(Limit int) *ListWorkflowParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Workflow records from the API. Request is executed immediately.
func (c *ApiService) PageWorkflow(WorkspaceSid string, params *ListWorkflowParams, pageToken, pageNumber string) (*ListWorkflowResponse, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows"

	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListWorkflowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Workflow records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWorkflow(WorkspaceSid string, params *ListWorkflowParams) ([]TaskrouterV1Workflow, error) {
	response, errors := c.StreamWorkflow(WorkspaceSid, params)

	records := make([]TaskrouterV1Workflow, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Workflow records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWorkflow(WorkspaceSid string, params *ListWorkflowParams) (chan TaskrouterV1Workflow, chan error) {
	if params == nil {
		params = &ListWorkflowParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1Workflow, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageWorkflow(WorkspaceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamWorkflow(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamWorkflow(response *ListWorkflowResponse, params *ListWorkflowParams, recordChannel chan TaskrouterV1Workflow, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Workflows
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListWorkflowResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListWorkflowResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListWorkflowResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkflowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWorkflow'
type UpdateWorkflowParams struct {
	// A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
	AssignmentCallbackUrl *string `json:"AssignmentCallbackUrl,omitempty"`
	// The URL that we should call when a call to the `assignment_callback_url` fails.
	FallbackAssignmentCallbackUrl *string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	// A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
	Configuration *string `json:"Configuration,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.
	TaskReservationTimeout *int `json:"TaskReservationTimeout,omitempty"`
	// Whether or not to re-evaluate Tasks. The default is `false`, which means Tasks in the Workflow will not be processed through the assignment loop again.
	ReEvaluateTasks *string `json:"ReEvaluateTasks,omitempty"`
}

func (params *UpdateWorkflowParams) SetFriendlyName(FriendlyName string) *UpdateWorkflowParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkflowParams) SetAssignmentCallbackUrl(AssignmentCallbackUrl string) *UpdateWorkflowParams {
	params.AssignmentCallbackUrl = &AssignmentCallbackUrl
	return params
}
func (params *UpdateWorkflowParams) SetFallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl string) *UpdateWorkflowParams {
	params.FallbackAssignmentCallbackUrl = &FallbackAssignmentCallbackUrl
	return params
}
func (params *UpdateWorkflowParams) SetConfiguration(Configuration string) *UpdateWorkflowParams {
	params.Configuration = &Configuration
	return params
}
func (params *UpdateWorkflowParams) SetTaskReservationTimeout(TaskReservationTimeout int) *UpdateWorkflowParams {
	params.TaskReservationTimeout = &TaskReservationTimeout
	return params
}
func (params *UpdateWorkflowParams) SetReEvaluateTasks(ReEvaluateTasks string) *UpdateWorkflowParams {
	params.ReEvaluateTasks = &ReEvaluateTasks
	return params
}

func (c *ApiService) UpdateWorkflow(WorkspaceSid string, Sid string, params *UpdateWorkflowParams) (*TaskrouterV1Workflow, error) {
	path := "/v1/Workspaces/{WorkspaceSid}/Workflows/{Sid}"
	path = strings.Replace(path, "{"+"WorkspaceSid"+"}", WorkspaceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.AssignmentCallbackUrl != nil {
		data.Set("AssignmentCallbackUrl", *params.AssignmentCallbackUrl)
	}
	if params != nil && params.FallbackAssignmentCallbackUrl != nil {
		data.Set("FallbackAssignmentCallbackUrl", *params.FallbackAssignmentCallbackUrl)
	}
	if params != nil && params.Configuration != nil {
		data.Set("Configuration", *params.Configuration)
	}
	if params != nil && params.TaskReservationTimeout != nil {
		data.Set("TaskReservationTimeout", fmt.Sprint(*params.TaskReservationTimeout))
	}
	if params != nil && params.ReEvaluateTasks != nil {
		data.Set("ReEvaluateTasks", *params.ReEvaluateTasks)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workflow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
