# ListServiceBindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bindings** | Pointer to [**[]ConversationsV1ServiceServiceBinding**](ConversationsV1ServiceServiceBinding.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListServiceBindingResponse

`func NewListServiceBindingResponse() *ListServiceBindingResponse`

NewListServiceBindingResponse instantiates a new ListServiceBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceBindingResponseWithDefaults

`func NewListServiceBindingResponseWithDefaults() *ListServiceBindingResponse`

NewListServiceBindingResponseWithDefaults instantiates a new ListServiceBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *ListServiceBindingResponse) GetBindings() []ConversationsV1ServiceServiceBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *ListServiceBindingResponse) GetBindingsOk() (*[]ConversationsV1ServiceServiceBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *ListServiceBindingResponse) SetBindings(v []ConversationsV1ServiceServiceBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *ListServiceBindingResponse) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetMeta

`func (o *ListServiceBindingResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceBindingResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceBindingResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceBindingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


