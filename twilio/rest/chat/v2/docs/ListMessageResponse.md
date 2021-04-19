# ListMessageResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Messages** | Pointer to [**[]ChatV2ServiceChannelMessage**](ChatV2ServiceChannelMessage.md) |  | [optional] 
**Meta** | Pointer to [**ListCredentialResponseMeta**](ListCredentialResponse_meta.md) |  | [optional] 

## Methods

### NewListMessageResponse

`func NewListMessageResponse() *ListMessageResponse`

NewListMessageResponse instantiates a new ListMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMessageResponseWithDefaults

`func NewListMessageResponseWithDefaults() *ListMessageResponse`

NewListMessageResponseWithDefaults instantiates a new ListMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListMessageResponse) GetMessages() []ChatV2ServiceChannelMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListMessageResponse) GetMessagesOk() (*[]ChatV2ServiceChannelMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListMessageResponse) SetMessages(v []ChatV2ServiceChannelMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListMessageResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMeta

`func (o *ListMessageResponse) GetMeta() ListCredentialResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListMessageResponse) GetMetaOk() (*ListCredentialResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListMessageResponse) SetMeta(v ListCredentialResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListMessageResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


