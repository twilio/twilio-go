# ListExecutionResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Executions** | Pointer to [**[]StudioV2FlowExecution**](StudioV2FlowExecution.md) |  | [optional] 
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 

## Methods

### NewListExecutionResponse

`func NewListExecutionResponse() *ListExecutionResponse`

NewListExecutionResponse instantiates a new ListExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExecutionResponseWithDefaults

`func NewListExecutionResponseWithDefaults() *ListExecutionResponse`

NewListExecutionResponseWithDefaults instantiates a new ListExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *ListExecutionResponse) GetExecutions() []StudioV2FlowExecution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ListExecutionResponse) GetExecutionsOk() (*[]StudioV2FlowExecution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ListExecutionResponse) SetExecutions(v []StudioV2FlowExecution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *ListExecutionResponse) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetMeta

`func (o *ListExecutionResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListExecutionResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListExecutionResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListExecutionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


