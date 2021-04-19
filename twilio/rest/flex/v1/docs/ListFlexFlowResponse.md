# ListFlexFlowResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**FlexFlows** | Pointer to [**[]FlexV1FlexFlow**](FlexV1FlexFlow.md) |  | [optional] 
**Meta** | Pointer to [**ListChannelResponseMeta**](ListChannelResponse_meta.md) |  | [optional] 

## Methods

### NewListFlexFlowResponse

`func NewListFlexFlowResponse() *ListFlexFlowResponse`

NewListFlexFlowResponse instantiates a new ListFlexFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlexFlowResponseWithDefaults

`func NewListFlexFlowResponseWithDefaults() *ListFlexFlowResponse`

NewListFlexFlowResponseWithDefaults instantiates a new ListFlexFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlexFlows

`func (o *ListFlexFlowResponse) GetFlexFlows() []FlexV1FlexFlow`

GetFlexFlows returns the FlexFlows field if non-nil, zero value otherwise.

### GetFlexFlowsOk

`func (o *ListFlexFlowResponse) GetFlexFlowsOk() (*[]FlexV1FlexFlow, bool)`

GetFlexFlowsOk returns a tuple with the FlexFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlows

`func (o *ListFlexFlowResponse) SetFlexFlows(v []FlexV1FlexFlow)`

SetFlexFlows sets FlexFlows field to given value.

### HasFlexFlows

`func (o *ListFlexFlowResponse) HasFlexFlows() bool`

HasFlexFlows returns a boolean if a field has been set.

### GetMeta

`func (o *ListFlexFlowResponse) GetMeta() ListChannelResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListFlexFlowResponse) GetMetaOk() (*ListChannelResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListFlexFlowResponse) SetMeta(v ListChannelResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListFlexFlowResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


