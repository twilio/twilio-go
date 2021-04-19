# ListCompositionHookResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**CompositionHooks** | Pointer to [**[]VideoV1CompositionHook**](VideoV1CompositionHook.md) |  | [optional] 
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 

## Methods

### NewListCompositionHookResponse

`func NewListCompositionHookResponse() *ListCompositionHookResponse`

NewListCompositionHookResponse instantiates a new ListCompositionHookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompositionHookResponseWithDefaults

`func NewListCompositionHookResponseWithDefaults() *ListCompositionHookResponse`

NewListCompositionHookResponseWithDefaults instantiates a new ListCompositionHookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompositionHooks

`func (o *ListCompositionHookResponse) GetCompositionHooks() []VideoV1CompositionHook`

GetCompositionHooks returns the CompositionHooks field if non-nil, zero value otherwise.

### GetCompositionHooksOk

`func (o *ListCompositionHookResponse) GetCompositionHooksOk() (*[]VideoV1CompositionHook, bool)`

GetCompositionHooksOk returns a tuple with the CompositionHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionHooks

`func (o *ListCompositionHookResponse) SetCompositionHooks(v []VideoV1CompositionHook)`

SetCompositionHooks sets CompositionHooks field to given value.

### HasCompositionHooks

`func (o *ListCompositionHookResponse) HasCompositionHooks() bool`

HasCompositionHooks returns a boolean if a field has been set.

### GetMeta

`func (o *ListCompositionHookResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCompositionHookResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCompositionHookResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCompositionHookResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


