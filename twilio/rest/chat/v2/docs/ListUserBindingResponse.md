# ListUserBindingResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Bindings** | Pointer to [**[]ChatV2ServiceUserUserBinding**](ChatV2ServiceUserUserBinding.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListUserBindingResponse

`func NewListUserBindingResponse() *ListUserBindingResponse`

NewListUserBindingResponse instantiates a new ListUserBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserBindingResponseWithDefaults

`func NewListUserBindingResponseWithDefaults() *ListUserBindingResponse`

NewListUserBindingResponseWithDefaults instantiates a new ListUserBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindings

`func (o *ListUserBindingResponse) GetBindings() []ChatV2ServiceUserUserBinding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *ListUserBindingResponse) GetBindingsOk() (*[]ChatV2ServiceUserUserBinding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *ListUserBindingResponse) SetBindings(v []ChatV2ServiceUserUserBinding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *ListUserBindingResponse) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetMeta

`func (o *ListUserBindingResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUserBindingResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUserBindingResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUserBindingResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


