# ListExecutionStepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 
**Steps** | Pointer to [**[]StudioV2FlowExecutionExecutionStep**](StudioV2FlowExecutionExecutionStep.md) |  | [optional] 

## Methods

### NewListExecutionStepResponse

`func NewListExecutionStepResponse() *ListExecutionStepResponse`

NewListExecutionStepResponse instantiates a new ListExecutionStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExecutionStepResponseWithDefaults

`func NewListExecutionStepResponseWithDefaults() *ListExecutionStepResponse`

NewListExecutionStepResponseWithDefaults instantiates a new ListExecutionStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListExecutionStepResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListExecutionStepResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListExecutionStepResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListExecutionStepResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSteps

`func (o *ListExecutionStepResponse) GetSteps() []StudioV2FlowExecutionExecutionStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ListExecutionStepResponse) GetStepsOk() (*[]StudioV2FlowExecutionExecutionStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ListExecutionStepResponse) SetSteps(v []StudioV2FlowExecutionExecutionStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ListExecutionStepResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


