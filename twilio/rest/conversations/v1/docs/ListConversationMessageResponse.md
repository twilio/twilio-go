# ListConversationMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]ConversationsV1ConversationConversationMessage**](ConversationsV1ConversationConversationMessage.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListConversationMessageResponse

`func NewListConversationMessageResponse() *ListConversationMessageResponse`

NewListConversationMessageResponse instantiates a new ListConversationMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationMessageResponseWithDefaults

`func NewListConversationMessageResponseWithDefaults() *ListConversationMessageResponse`

NewListConversationMessageResponseWithDefaults instantiates a new ListConversationMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListConversationMessageResponse) GetMessages() []ConversationsV1ConversationConversationMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListConversationMessageResponse) GetMessagesOk() (*[]ConversationsV1ConversationConversationMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListConversationMessageResponse) SetMessages(v []ConversationsV1ConversationConversationMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListConversationMessageResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMeta

`func (o *ListConversationMessageResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConversationMessageResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConversationMessageResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConversationMessageResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


