/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"fmt"
    twilio "github.com/twilio/twilio-go/client"
    "net/url"
    "strings"
    ""
)

type DefaultApiService struct {
    baseURL string
    client  *twilio.Client
}

func NewDefaultApiService(client *twilio.Client) *DefaultApiService {
    return &DefaultApiService{
        client: client,
        baseURL: fmt.Sprintf("https://studio.%s", client.BaseURL),
    }
}
// CreateExecutionParams Optional parameters for the method 'CreateExecution'
type CreateExecutionParams struct {
    From *string `json:"From,omitempty"`
    Parameters *map[string]interface{} `json:"Parameters,omitempty"`
    To *string `json:"To,omitempty"`
}

/*
CreateExecution Method for CreateExecution
Triggers a new Execution for the Flow
 * @param flowSid The SID of the Excecution's Flow.
 * @param optional nil or *CreateExecutionOpts - Optional Parameters:
 * @param "From" (string) - The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`.
 * @param "Parameters" (map[string]interface{}) - JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
 * @param "To" (string) - The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
@return StudioV2FlowExecution
*/
func (c *DefaultApiService) CreateExecution(flowSid string, params *CreateExecutionParams) (*StudioV2FlowExecution, error) {
    path := "/v2/Flows/{FlowSid}/Executions"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)

    data := url.Values{}
    headers := 0

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

    ps := &StudioV2FlowExecution{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// CreateFlowParams Optional parameters for the method 'CreateFlow'
type CreateFlowParams struct {
    CommitMessage *string `json:"CommitMessage,omitempty"`
    Definition *map[string]interface{} `json:"Definition,omitempty"`
    FriendlyName *string `json:"FriendlyName,omitempty"`
    Status *string `json:"Status,omitempty"`
}

/*
CreateFlow Method for CreateFlow
Create a Flow.
 * @param optional nil or *CreateFlowOpts - Optional Parameters:
 * @param "CommitMessage" (string) - Description on change made in the revision.
 * @param "Definition" (map[string]interface{}) - JSON representation of flow definition.
 * @param "FriendlyName" (string) - The string that you assigned to describe the Flow.
 * @param "Status" (string) - The status of the Flow. Can be: `draft` or `published`.
@return StudioV2Flow
*/
func (c *DefaultApiService) CreateFlow(params *CreateFlowParams) (*StudioV2Flow, error) {
    path := "/v2/Flows"

    data := url.Values{}
    headers := 0

    if params != nil && params.CommitMessage != nil {
        data.Set("CommitMessage", *params.CommitMessage)
    }
    if params != nil && params.Definition != nil {
        v, err := json.Marshal(params.Definition)

        if err != nil {
            return nil, err
        }

        data.Set("Definition", string(v))
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }
    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2Flow{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
DeleteExecution Method for DeleteExecution
Delete the Execution and all Steps relating to it.
 * @param flowSid The SID of the Flow with the Execution resources to delete.
 * @param sid The SID of the Execution resource to delete.
*/
func (c *DefaultApiService) DeleteExecution(flowSid string, sid string) (error) {
    path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
DeleteFlow Method for DeleteFlow
Delete a specific Flow.
 * @param sid The SID of the Flow resource to delete.
*/
func (c *DefaultApiService) DeleteFlow(sid string) (error) {
    path := "/v2/Flows/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Delete(c.baseURL+path, data, headers)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    return nil
}

/*
FetchExecution Method for FetchExecution
Retrieve an Execution
 * @param flowSid The SID of the Flow with the Execution resource to fetch
 * @param sid The SID of the Execution resource to fetch.
@return StudioV2FlowExecution
*/
func (c *DefaultApiService) FetchExecution(flowSid string, sid string) (*StudioV2FlowExecution, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowExecution{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchExecutionContext Method for FetchExecutionContext
Retrieve the most recent context for an Execution.
 * @param flowSid The SID of the Flow with the Execution context to fetch.
 * @param executionSid The SID of the Execution context to fetch.
@return StudioV2FlowExecutionExecutionContext
*/
func (c *DefaultApiService) FetchExecutionContext(flowSid string, executionSid string) (*StudioV2FlowExecutionExecutionContext, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Context"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"ExecutionSid"+"}", executionSid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowExecutionExecutionContext{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchExecutionStep Method for FetchExecutionStep
Retrieve a Step.
 * @param flowSid The SID of the Flow with the Step to fetch.
 * @param executionSid The SID of the Execution resource with the Step to fetch.
 * @param sid The SID of the ExecutionStep resource to fetch.
@return StudioV2FlowExecutionExecutionStep
*/
func (c *DefaultApiService) FetchExecutionStep(flowSid string, executionSid string, sid string) (*StudioV2FlowExecutionExecutionStep, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{Sid}"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"ExecutionSid"+"}", executionSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowExecutionExecutionStep{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchExecutionStepContext Method for FetchExecutionStepContext
Retrieve the context for an Execution Step.
 * @param flowSid The SID of the Flow with the Step to fetch.
 * @param executionSid The SID of the Execution resource with the Step to fetch.
 * @param stepSid The SID of the Step to fetch.
@return StudioV2FlowExecutionExecutionStepExecutionStepContext
*/
func (c *DefaultApiService) FetchExecutionStepContext(flowSid string, executionSid string, stepSid string) (*StudioV2FlowExecutionExecutionStepExecutionStepContext, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps/{StepSid}/Context"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"ExecutionSid"+"}", executionSid, -1)
    path = strings.Replace(path, "{"+"StepSid"+"}", stepSid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowExecutionExecutionStepExecutionStepContext{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchFlow Method for FetchFlow
Retrieve a specific Flow.
 * @param sid The SID of the Flow resource to fetch.
@return StudioV2Flow
*/
func (c *DefaultApiService) FetchFlow(sid string) (*StudioV2Flow, error) {
    path := "/v2/Flows/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2Flow{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchFlowRevision Method for FetchFlowRevision
Retrieve a specific Flow revision.
 * @param sid The SID of the Flow resource to fetch.
 * @param revision Specific Revision number or can be `LatestPublished` and `LatestRevision`.
@return StudioV2FlowFlowRevision
*/
func (c *DefaultApiService) FetchFlowRevision(sid string, revision string) (*StudioV2FlowFlowRevision, error) {
    path := "/v2/Flows/{Sid}/Revisions/{Revision}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)
    path = strings.Replace(path, "{"+"Revision"+"}", revision, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowFlowRevision{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}

/*
FetchTestUser Method for FetchTestUser
 * @param sid
@return StudioV2FlowTestUser
*/
func (c *DefaultApiService) FetchTestUser(sid string) (*StudioV2FlowTestUser, error) {
    path := "/v2/Flows/{Sid}/TestUsers"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := 0
    headers := 0



    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowTestUser{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListExecutionParams Optional parameters for the method 'ListExecution'
type ListExecutionParams struct {
    DateCreatedFrom *time.Time `json:"DateCreatedFrom,omitempty"`
    DateCreatedTo *time.Time `json:"DateCreatedTo,omitempty"`
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListExecution Method for ListExecution
Retrieve a list of all Executions for the Flow.
 * @param flowSid The SID of the Flow with the Execution resources to read.
 * @param optional nil or *ListExecutionOpts - Optional Parameters:
 * @param "DateCreatedFrom" (time.Time) - Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
 * @param "DateCreatedTo" (time.Time) - Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return InlineResponse2001
*/
func (c *DefaultApiService) ListExecution(flowSid string, params *ListExecutionParams) (*InlineResponse2001, error) {
    path := "/v2/Flows/{FlowSid}/Executions"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.DateCreatedFrom != nil {
        data.Set("DateCreatedFrom", string(*params.DateCreatedFrom))
    }
    if params != nil && params.DateCreatedTo != nil {
        data.Set("DateCreatedTo", string(*params.DateCreatedTo))
    }
    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InlineResponse2001{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListExecutionStepParams Optional parameters for the method 'ListExecutionStep'
type ListExecutionStepParams struct {
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListExecutionStep Method for ListExecutionStep
Retrieve a list of all Steps for an Execution.
 * @param flowSid The SID of the Flow with the Steps to read.
 * @param executionSid The SID of the Execution with the Steps to read.
 * @param optional nil or *ListExecutionStepOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return InlineResponse2002
*/
func (c *DefaultApiService) ListExecutionStep(flowSid string, executionSid string, params *ListExecutionStepParams) (*InlineResponse2002, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{ExecutionSid}/Steps"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"ExecutionSid"+"}", executionSid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InlineResponse2002{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListFlowParams Optional parameters for the method 'ListFlow'
type ListFlowParams struct {
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListFlow Method for ListFlow
Retrieve a list of all Flows.
 * @param optional nil or *ListFlowOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return InlineResponse200
*/
func (c *DefaultApiService) ListFlow(params *ListFlowParams) (*InlineResponse200, error) {
    path := "/v2/Flows"

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InlineResponse200{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// ListFlowRevisionParams Optional parameters for the method 'ListFlowRevision'
type ListFlowRevisionParams struct {
    PageSize *int32 `json:"PageSize,omitempty"`
}

/*
ListFlowRevision Method for ListFlowRevision
Retrieve a list of all Flows revisions.
 * @param sid The SID of the Flow resource to fetch.
 * @param optional nil or *ListFlowRevisionOpts - Optional Parameters:
 * @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
@return InlineResponse2003
*/
func (c *DefaultApiService) ListFlowRevision(sid string, params *ListFlowRevisionParams) (*InlineResponse2003, error) {
    path := "/v2/Flows/{Sid}/Revisions"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.PageSize != nil {
        data.Set("PageSize", string(*params.PageSize))
    }


    resp, err := c.client.Get(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &InlineResponse2003{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateExecutionParams Optional parameters for the method 'UpdateExecution'
type UpdateExecutionParams struct {
    Status *string `json:"Status,omitempty"`
}

/*
UpdateExecution Method for UpdateExecution
Update the status of an Execution to &#x60;ended&#x60;.
 * @param flowSid The SID of the Flow with the Execution resources to update.
 * @param sid The SID of the Execution resource to update.
 * @param optional nil or *UpdateExecutionOpts - Optional Parameters:
 * @param "Status" (string) - The status of the Execution. Can only be `ended`.
@return StudioV2FlowExecution
*/
func (c *DefaultApiService) UpdateExecution(flowSid string, sid string, params *UpdateExecutionParams) (*StudioV2FlowExecution, error) {
    path := "/v2/Flows/{FlowSid}/Executions/{Sid}"
    path = strings.Replace(path, "{"+"FlowSid"+"}", flowSid, -1)
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowExecution{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateFlowParams Optional parameters for the method 'UpdateFlow'
type UpdateFlowParams struct {
    CommitMessage *string `json:"CommitMessage,omitempty"`
    Definition *map[string]interface{} `json:"Definition,omitempty"`
    FriendlyName *string `json:"FriendlyName,omitempty"`
    Status *string `json:"Status,omitempty"`
}

/*
UpdateFlow Method for UpdateFlow
Update a Flow.
 * @param sid The SID of the Flow resource to fetch.
 * @param optional nil or *UpdateFlowOpts - Optional Parameters:
 * @param "CommitMessage" (string) - Description on change made in the revision.
 * @param "Definition" (map[string]interface{}) - JSON representation of flow definition.
 * @param "FriendlyName" (string) - The string that you assigned to describe the Flow.
 * @param "Status" (string) - The status of the Flow. Can be: `draft` or `published`.
@return StudioV2Flow
*/
func (c *DefaultApiService) UpdateFlow(sid string, params *UpdateFlowParams) (*StudioV2Flow, error) {
    path := "/v2/Flows/{Sid}"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.CommitMessage != nil {
        data.Set("CommitMessage", *params.CommitMessage)
    }
    if params != nil && params.Definition != nil {
        v, err := json.Marshal(params.Definition)

        if err != nil {
            return nil, err
        }

        data.Set("Definition", string(v))
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }
    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2Flow{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateFlowValidateParams Optional parameters for the method 'UpdateFlowValidate'
type UpdateFlowValidateParams struct {
    CommitMessage *string `json:"CommitMessage,omitempty"`
    Definition *map[string]interface{} `json:"Definition,omitempty"`
    FriendlyName *string `json:"FriendlyName,omitempty"`
    Status *string `json:"Status,omitempty"`
}

/*
UpdateFlowValidate Method for UpdateFlowValidate
 * @param optional nil or *UpdateFlowValidateOpts - Optional Parameters:
 * @param "CommitMessage" (string) - 
 * @param "Definition" (map[string]interface{}) - 
 * @param "FriendlyName" (string) - 
 * @param "Status" (string) - 
@return StudioV2FlowValidate
*/
func (c *DefaultApiService) UpdateFlowValidate(params *UpdateFlowValidateParams) (*StudioV2FlowValidate, error) {
    path := "/v2/Flows/Validate"

    data := url.Values{}
    headers := 0

    if params != nil && params.CommitMessage != nil {
        data.Set("CommitMessage", *params.CommitMessage)
    }
    if params != nil && params.Definition != nil {
        v, err := json.Marshal(params.Definition)

        if err != nil {
            return nil, err
        }

        data.Set("Definition", string(v))
    }
    if params != nil && params.FriendlyName != nil {
        data.Set("FriendlyName", *params.FriendlyName)
    }
    if params != nil && params.Status != nil {
        data.Set("Status", *params.Status)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowValidate{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
// UpdateTestUserParams Optional parameters for the method 'UpdateTestUser'
type UpdateTestUserParams struct {
    TestUsers *[]string `json:"TestUsers,omitempty"`
}

/*
UpdateTestUser Method for UpdateTestUser
 * @param sid
 * @param optional nil or *UpdateTestUserOpts - Optional Parameters:
 * @param "TestUsers" ([]string) - 
@return StudioV2FlowTestUser
*/
func (c *DefaultApiService) UpdateTestUser(sid string, params *UpdateTestUserParams) (*StudioV2FlowTestUser, error) {
    path := "/v2/Flows/{Sid}/TestUsers"
    path = strings.Replace(path, "{"+"Sid"+"}", sid, -1)

    data := url.Values{}
    headers := 0

    if params != nil && params.TestUsers != nil {
        data.Set("TestUsers", *params.TestUsers)
    }


    resp, err := c.client.Post(c.baseURL+path, data, headers)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    ps := &StudioV2FlowTestUser{}
    if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
        return nil, err
    }

    return ps, err
}
