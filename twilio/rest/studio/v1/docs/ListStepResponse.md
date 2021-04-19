# ListStepResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 
**Steps** | Pointer to [**[]StudioV1FlowEngagementStep**](StudioV1FlowEngagementStep.md) |  | [optional] 

## Methods

### NewListStepResponse

`func NewListStepResponse() *ListStepResponse`

NewListStepResponse instantiates a new ListStepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStepResponseWithDefaults

`func NewListStepResponseWithDefaults() *ListStepResponse`

NewListStepResponseWithDefaults instantiates a new ListStepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListStepResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStepResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStepResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStepResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSteps

`func (o *ListStepResponse) GetSteps() []StudioV1FlowEngagementStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ListStepResponse) GetStepsOk() (*[]StudioV1FlowEngagementStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ListStepResponse) SetSteps(v []StudioV1FlowEngagementStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ListStepResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


