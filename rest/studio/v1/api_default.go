/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://studio.twilio.com",
	}
}

// CreateEngagementParams Optional parameters for the method 'CreateEngagement'
type CreateEngagementParams struct {
	From       *string                 `json:"From,omitempty"`
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	To         *string                 `json:"To,omitempty"`
}

func (params *CreateEngagementParams) SetFrom(From string) *CreateEngagementParams {
	params.From = &From
	return params
}
func (params *CreateEngagementParams) SetParameters(Parameters map[string]interface{}) *CreateEngagementParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateEngagementParams) SetTo(To string) *CreateEngagementParams {
	params.To = &To
	return params
}

// CreateEngagement Method for CreateEngagement
//
// Triggers a new Engagement for the Flow
//
// param: FlowSid The SID of the Flow.
//
// param: optional nil or *CreateEngagementParams - Optional Parameters:
//
// param: "From" (string) - The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
//
// param: "Parameters" (map[string]interface{}) - A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
//
// param: "To" (string) - The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
//
// return: StudioV1FlowEngagement
func (c *DefaultApiService) CreateEngagement(FlowSid string, params *CreateEngagementParams) (*StudioV1FlowEngagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// CreateExecutionParams Optional parameters for the method 'CreateExecution'
type CreateExecutionParams struct {
	From       *string                 `json:"From,omitempty"`
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	To         *string                 `json:"To,omitempty"`
}

func (params *CreateExecutionParams) SetFrom(From string) *CreateExecutionParams {
	params.From = &From
	return params
}
func (params *CreateExecutionParams) SetParameters(Parameters map[string]interface{}) *CreateExecutionParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateExecutionParams) SetTo(To string) *CreateExecutionParams {
	params.To = &To
	return params
}

// CreateExecution Method for CreateExecution
//
// Triggers a new Execution for the Flow
//
// param: FlowSid The SID of the Excecution's Flow.
//
// param: optional nil or *CreateExecutionParams - Optional Parameters:
//
// param: "From" (string) - The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. For SMS, this can also be a Messaging Service SID.
//
// param: "Parameters" (map[string]interface{}) - JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
//
// param: "To" (string) - The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
//
// return: StudioV1FlowExecution
func (c *DefaultApiService) CreateExecution(FlowSid string, params *CreateExecutionParams) (*StudioV1FlowExecution, error) {
	path := "/v1/Flows/{FlowSid}/Executions"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// DeleteEngagement Method for DeleteEngagement
//
// Delete this Engagement and all Steps relating to it.
//
// param: FlowSid The SID of the Flow to delete Engagements from.
//
// param: Sid The SID of the Engagement resource to delete.
//
func (c *DefaultApiService) DeleteEngagement(FlowSid string, Sid string) error {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// DeleteExecution Method for DeleteExecution
//
// Delete the Execution and all Steps relating to it.
//
// param: FlowSid The SID of the Flow with the Execution resources to delete.
//
// param: Sid The SID of the Execution resource to delete.
//
func (c *DefaultApiService) DeleteExecution(FlowSid string, Sid string) error {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// DeleteFlow Method for DeleteFlow
//
// Delete a specific Flow.
//
// param: Sid The SID of the Flow resource to delete.
//
func (c *DefaultApiService) DeleteFlow(Sid string) error {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// FetchEngagement Method for FetchEngagement
//
// Retrieve an Engagement
//
// param: FlowSid The SID of the Flow.
//
// param: Sid The SID of the Engagement resource to fetch.
//
// return: StudioV1FlowEngagement
func (c *DefaultApiService) FetchEngagement(FlowSid string, Sid string) (*StudioV1FlowEngagement, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagement{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchEngagementContext Method for FetchEngagementContext
//
// Retrieve the most recent context for an Engagement.
//
// param: FlowSid The SID of the Flow.
//
// param: EngagementSid The SID of the Engagement.
//
// return: StudioV1FlowEngagementEngagementContext
func (c *DefaultApiService) FetchEngagementContext(FlowSid string, EngagementSid string) (*StudioV1FlowEngagementEngagementContext, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagementEngagementContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchExecution Method for FetchExecution
//
// Retrieve an Execution
//
// param: FlowSid The SID of the Flow with the Execution resource to fetch
//
// param: Sid The SID of the Execution resource to fetch.
//
// return: StudioV1FlowExecution
func (c *DefaultApiService) FetchExecution(FlowSid string, Sid string) (*StudioV1FlowExecution, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchExecutionContext Method for FetchExecutionContext
//
// Retrieve the most recent context for an Execution.
//
// param: FlowSid The SID of the Flow with the Execution context to fetch.
//
// param: ExecutionSid The SID of the Execution context to fetch.
//
// return: StudioV1FlowExecutionExecutionContext
func (c *DefaultApiService) FetchExecutionContext(FlowSid string, ExecutionSid string) (*StudioV1FlowExecutionExecutionContext, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecutionExecutionContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchExecutionStep Method for FetchExecutionStep
//
// Retrieve a Step.
//
// param: FlowSid The SID of the Flow with the Step to fetch.
//
// param: ExecutionSid The SID of the Execution resource with the Step to fetch.
//
// param: Sid The SID of the ExecutionStep resource to fetch.
//
// return: StudioV1FlowExecutionExecutionStep
func (c *DefaultApiService) FetchExecutionStep(FlowSid string, ExecutionSid string, Sid string) (*StudioV1FlowExecutionExecutionStep, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecutionExecutionStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchExecutionStepContext Method for FetchExecutionStepContext
//
// Retrieve the context for an Execution Step.
//
// param: FlowSid The SID of the Flow with the Step to fetch.
//
// param: ExecutionSid The SID of the Execution resource with the Step to fetch.
//
// param: StepSid The SID of the Step to fetch.
//
// return: StudioV1FlowExecutionExecutionStepExecutionStepContext
func (c *DefaultApiService) FetchExecutionStepContext(FlowSid string, ExecutionSid string, StepSid string) (*StudioV1FlowExecutionExecutionStepExecutionStepContext, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)
	path = strings.Replace(path, "{"+"StepSid"+"}", StepSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecutionExecutionStepExecutionStepContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchFlow Method for FetchFlow
//
// Retrieve a specific Flow.
//
// param: Sid The SID of the Flow resource to fetch.
//
// return: StudioV1Flow
func (c *DefaultApiService) FetchFlow(Sid string) (*StudioV1Flow, error) {
	path := "/v1/Flows/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Flow{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchStep Method for FetchStep
//
// Retrieve a Step.
//
// param: FlowSid The SID of the Flow with the Step to fetch.
//
// param: EngagementSid The SID of the Engagement with the Step to fetch.
//
// param: Sid The SID of the Step resource to fetch.
//
// return: StudioV1FlowEngagementStep
func (c *DefaultApiService) FetchStep(FlowSid string, EngagementSid string, Sid string) (*StudioV1FlowEngagementStep, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagementStep{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchStepContext Method for FetchStepContext
//
// Retrieve the context for an Engagement Step.
//
// param: FlowSid The SID of the Flow with the Step to fetch.
//
// param: EngagementSid The SID of the Engagement with the Step to fetch.
//
// param: StepSid The SID of the Step to fetch
//
// return: StudioV1FlowEngagementStepStepContext
func (c *DefaultApiService) FetchStepContext(FlowSid string, EngagementSid string, StepSid string) (*StudioV1FlowEngagementStepStepContext, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps/{StepSid}/Context"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)
	path = strings.Replace(path, "{"+"StepSid"+"}", StepSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowEngagementStepStepContext{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListEngagementParams Optional parameters for the method 'ListEngagement'
type ListEngagementParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListEngagementParams) SetPageSize(PageSize int32) *ListEngagementParams {
	params.PageSize = &PageSize
	return params
}

// ListEngagement Method for ListEngagement
//
// Retrieve a list of all Engagements for the Flow.
//
// param: FlowSid The SID of the Flow to read Engagements from.
//
// param: optional nil or *ListEngagementParams - Optional Parameters:
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListEngagementResponse
func (c *DefaultApiService) ListEngagement(FlowSid string, params *ListEngagementParams) (*ListEngagementResponse, error) {
	path := "/v1/Flows/{FlowSid}/Engagements"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEngagementResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListExecutionParams Optional parameters for the method 'ListExecution'
type ListExecutionParams struct {
	DateCreatedFrom *time.Time `json:"DateCreatedFrom,omitempty"`
	DateCreatedTo   *time.Time `json:"DateCreatedTo,omitempty"`
	PageSize        *int32     `json:"PageSize,omitempty"`
}

func (params *ListExecutionParams) SetDateCreatedFrom(DateCreatedFrom time.Time) *ListExecutionParams {
	params.DateCreatedFrom = &DateCreatedFrom
	return params
}
func (params *ListExecutionParams) SetDateCreatedTo(DateCreatedTo time.Time) *ListExecutionParams {
	params.DateCreatedTo = &DateCreatedTo
	return params
}
func (params *ListExecutionParams) SetPageSize(PageSize int32) *ListExecutionParams {
	params.PageSize = &PageSize
	return params
}

// ListExecution Method for ListExecution
//
// Retrieve a list of all Executions for the Flow.
//
// param: FlowSid The SID of the Flow with the Execution resources to read.
//
// param: optional nil or *ListExecutionParams - Optional Parameters:
//
// param: "DateCreatedFrom" (time.Time) - Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
//
// param: "DateCreatedTo" (time.Time) - Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListExecutionResponse
func (c *DefaultApiService) ListExecution(FlowSid string, params *ListExecutionParams) (*ListExecutionResponse, error) {
	path := "/v1/Flows/{FlowSid}/Executions"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreatedFrom != nil {
		data.Set("DateCreatedFrom", fmt.Sprint((*params.DateCreatedFrom).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedTo != nil {
		data.Set("DateCreatedTo", fmt.Sprint((*params.DateCreatedTo).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListExecutionStepParams Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListExecutionStepParams) SetPageSize(PageSize int32) *ListExecutionStepParams {
	params.PageSize = &PageSize
	return params
}

// ListExecutionStep Method for ListExecutionStep
//
// Retrieve a list of all Steps for an Execution.
//
// param: FlowSid The SID of the Flow with the Steps to read.
//
// param: ExecutionSid The SID of the Execution with the Steps to read.
//
// param: optional nil or *ListExecutionStepParams - Optional Parameters:
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListExecutionStepResponse
func (c *DefaultApiService) ListExecutionStep(FlowSid string, ExecutionSid string, params *ListExecutionStepParams) (*ListExecutionStepResponse, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"ExecutionSid"+"}", ExecutionSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListFlowParams Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int32) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}

// ListFlow Method for ListFlow
//
// Retrieve a list of all Flows.
//
// param: optional nil or *ListFlowParams - Optional Parameters:
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListFlowResponse
func (c *DefaultApiService) ListFlow(params *ListFlowParams) (*ListFlowResponse, error) {
	path := "/v1/Flows"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFlowResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListStepParams Optional parameters for the method 'ListStep'
type ListStepParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListStepParams) SetPageSize(PageSize int32) *ListStepParams {
	params.PageSize = &PageSize
	return params
}

// ListStep Method for ListStep
//
// Retrieve a list of all Steps for an Engagement.
//
// param: FlowSid The SID of the Flow with the Step to read.
//
// param: EngagementSid The SID of the Engagement with the Step to read.
//
// param: optional nil or *ListStepParams - Optional Parameters:
//
// param: "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
//
// return: ListStepResponse
func (c *DefaultApiService) ListStep(FlowSid string, EngagementSid string, params *ListStepParams) (*ListStepResponse, error) {
	path := "/v1/Flows/{FlowSid}/Engagements/{EngagementSid}/Steps"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"EngagementSid"+"}", EngagementSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListStepResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// UpdateExecutionParams Optional parameters for the method 'UpdateExecution'
type UpdateExecutionParams struct {
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateExecutionParams) SetStatus(Status string) *UpdateExecutionParams {
	params.Status = &Status
	return params
}

// UpdateExecution Method for UpdateExecution
//
// Update the status of an Execution to &#x60;ended&#x60;.
//
// param: FlowSid The SID of the Flow with the Execution resources to update.
//
// param: Sid The SID of the Execution resource to update.
//
// param: optional nil or *UpdateExecutionParams - Optional Parameters:
//
// param: "Status" (string) - The status of the Execution. Can only be `ended`.
//
// return: StudioV1FlowExecution
func (c *DefaultApiService) UpdateExecution(FlowSid string, Sid string, params *UpdateExecutionParams) (*StudioV1FlowExecution, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.client.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1FlowExecution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
