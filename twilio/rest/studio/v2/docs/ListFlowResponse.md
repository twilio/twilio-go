# ListFlowResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Flows** | Pointer to [**[]StudioV2Flow**](StudioV2Flow.md) |  | [optional] 
**Meta** | Pointer to [**ListFlowResponseMeta**](ListFlowResponse_meta.md) |  | [optional] 

## Methods

### NewListFlowResponse

`func NewListFlowResponse() *ListFlowResponse`

NewListFlowResponse instantiates a new ListFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlowResponseWithDefaults

`func NewListFlowResponseWithDefaults() *ListFlowResponse`

NewListFlowResponseWithDefaults instantiates a new ListFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlows

`func (o *ListFlowResponse) GetFlows() []StudioV2Flow`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ListFlowResponse) GetFlowsOk() (*[]StudioV2Flow, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ListFlowResponse) SetFlows(v []StudioV2Flow)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *ListFlowResponse) HasFlows() bool`

HasFlows returns a boolean if a field has been set.

### GetMeta

`func (o *ListFlowResponse) GetMeta() ListFlowResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFlowResponse) GetMetaOk() (*ListFlowResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFlowResponse) SetMeta(v ListFlowResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFlowResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


