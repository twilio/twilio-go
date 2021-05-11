/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
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

// Optional parameters for the method 'CreateEngagement'
type CreateEngagementParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow Engagement. Available as variable `{{flow.channel.address}}`
	From *string `json:"From,omitempty"`
	// A JSON string we will add to your flow's context and that you can access as variables inside your flow. For example, if you pass in `Parameters={'name':'Zeke'}` then inside a widget you can reference the variable `{{flow.data.name}}` which will return the string 'Zeke'. Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode your JSON string.
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Engagement, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
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

// Triggers a new Engagement for the Flow
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

// Optional parameters for the method 'CreateExecution'
type CreateExecutionParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`.
	From *string `json:"From,omitempty"`
	// JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
	Parameters *map[string]interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
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

// Triggers a new Execution for the Flow
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

// Delete this Engagement and all Steps relating to it.
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

// Delete the Execution and all Steps relating to it.
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

// Delete a specific Flow.
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

// Retrieve an Engagement
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

// Retrieve the most recent context for an Engagement.
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

// Retrieve an Execution
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

// Retrieve the most recent context for an Execution.
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

// Retrieve a Step.
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

// Retrieve the context for an Execution Step.
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

// Retrieve a specific Flow.
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

// Retrieve a Step.
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

// Retrieve the context for an Engagement Step.
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

// Optional parameters for the method 'ListEngagement'
type ListEngagementParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListEngagementParams) SetPageSize(PageSize int32) *ListEngagementParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Engagements for the Flow.
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

// Optional parameters for the method 'ListExecution'
type ListExecutionParams struct {
	// Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedFrom *time.Time `json:"DateCreatedFrom,omitempty"`
	// Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedTo *time.Time `json:"DateCreatedTo,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
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

// Retrieve a list of all Executions for the Flow.
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

// Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListExecutionStepParams) SetPageSize(PageSize int32) *ListExecutionStepParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Steps for an Execution.
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

// Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListFlowParams) SetPageSize(PageSize int32) *ListFlowParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Flows.
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

// Optional parameters for the method 'ListStep'
type ListStepParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListStepParams) SetPageSize(PageSize int32) *ListStepParams {
	params.PageSize = &PageSize
	return params
}

// Retrieve a list of all Steps for an Engagement.
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

// Optional parameters for the method 'UpdateExecution'
type UpdateExecutionParams struct {
	// The status of the Execution. Can only be `ended`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateExecutionParams) SetStatus(Status string) *UpdateExecutionParams {
	params.Status = &Status
	return params
}

// Update the status of an Execution to &#x60;ended&#x60;.
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
