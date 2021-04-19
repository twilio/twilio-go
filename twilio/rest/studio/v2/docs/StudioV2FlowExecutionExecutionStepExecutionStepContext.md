# StudioV2FlowExecutionExecutionStepExecutionStepContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Context** | Pointer to **map[string]interface{}** | The current state of the flow | [optional] 
**ExecutionSid** | Pointer to **NullableString** | The SID of the Execution | [optional] 
**FlowSid** | Pointer to **NullableString** | The SID of the Flow | [optional] 
**StepSid** | Pointer to **NullableString** | Step SID | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewStudioV2FlowExecutionExecutionStepExecutionStepContext

`func NewStudioV2FlowExecutionExecutionStepExecutionStepContext() *StudioV2FlowExecutionExecutionStepExecutionStepContext`

NewStudioV2FlowExecutionExecutionStepExecutionStepContext instantiates a new StudioV2FlowExecutionExecutionStepExecutionStepContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV2FlowExecutionExecutionStepExecutionStepContextWithDefaults

`func NewStudioV2FlowExecutionExecutionStepExecutionStepContextWithDefaults() *StudioV2FlowExecutionExecutionStepExecutionStepContext`

NewStudioV2FlowExecutionExecutionStepExecutionStepContextWithDefaults instantiates a new StudioV2FlowExecutionExecutionStepExecutionStepContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContext

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetExecutionSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetExecutionSid() string`

GetExecutionSid returns the ExecutionSid field if non-nil, zero value otherwise.

### GetExecutionSidOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetExecutionSidOk() (*string, bool)`

GetExecutionSidOk returns a tuple with the ExecutionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetExecutionSid(v string)`

SetExecutionSid sets ExecutionSid field to given value.

### HasExecutionSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasExecutionSid() bool`

HasExecutionSid returns a boolean if a field has been set.

### SetExecutionSidNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetExecutionSidNil(b bool)`

 SetExecutionSidNil sets the value for ExecutionSid to be an explicit nil

### UnsetExecutionSid
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetExecutionSid()`

UnsetExecutionSid ensures that no value is present for ExecutionSid, not even an explicit nil
### GetFlowSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetFlowSid() string`

GetFlowSid returns the FlowSid field if non-nil, zero value otherwise.

### GetFlowSidOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetFlowSidOk() (*string, bool)`

GetFlowSidOk returns a tuple with the FlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetFlowSid(v string)`

SetFlowSid sets FlowSid field to given value.

### HasFlowSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasFlowSid() bool`

HasFlowSid returns a boolean if a field has been set.

### SetFlowSidNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetFlowSidNil(b bool)`

 SetFlowSidNil sets the value for FlowSid to be an explicit nil

### UnsetFlowSid
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetFlowSid()`

UnsetFlowSid ensures that no value is present for FlowSid, not even an explicit nil
### GetStepSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetStepSid() string`

GetStepSid returns the StepSid field if non-nil, zero value otherwise.

### GetStepSidOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetStepSidOk() (*string, bool)`

GetStepSidOk returns a tuple with the StepSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetStepSid(v string)`

SetStepSid sets StepSid field to given value.

### HasStepSid

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasStepSid() bool`

HasStepSid returns a boolean if a field has been set.

### SetStepSidNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetStepSidNil(b bool)`

 SetStepSidNil sets the value for StepSid to be an explicit nil

### UnsetStepSid
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetStepSid()`

UnsetStepSid ensures that no value is present for StepSid, not even an explicit nil
### GetUrl

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV2FlowExecutionExecutionStepExecutionStepContext) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


