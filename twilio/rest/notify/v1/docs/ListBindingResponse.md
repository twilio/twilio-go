# ListBindingResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Bindings** | Pointer to [**[]NotifyV1ServiceBinding**](NotifyV1ServiceBinding.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListBindingResponse

`func NewListBindingResponse() *ListBindingResponse`

NewListBindingResponse instantiates a new ListBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBindingResponseWithDefaults

`func NewListBindingResponseWithDefaults() *ListBindingResponse`

NewListBindingResponseWithDefaults instantiates a new ListBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *ListBindingResponse) GetBindings() []NotifyV1ServiceBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *ListBindingResponse) GetBindingsOk() (*[]NotifyV1ServiceBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *ListBindingResponse) SetBindings(v []NotifyV1ServiceBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *ListBindingResponse) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetMeta

`func (o *ListBindingResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBindingResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBindingResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBindingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


