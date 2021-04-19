# ListServiceConversationMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]ConversationsV1ServiceServiceConversationServiceConversationMessage**](ConversationsV1ServiceServiceConversationServiceConversationMessage.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListServiceConversationMessageResponse

`func NewListServiceConversationMessageResponse() *ListServiceConversationMessageResponse`

NewListServiceConversationMessageResponse instantiates a new ListServiceConversationMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceConversationMessageResponseWithDefaults

`func NewListServiceConversationMessageResponseWithDefaults() *ListServiceConversationMessageResponse`

NewListServiceConversationMessageResponseWithDefaults instantiates a new ListServiceConversationMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListServiceConversationMessageResponse) GetMessages() []ConversationsV1ServiceServiceConversationServiceConversationMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListServiceConversationMessageResponse) GetMessagesOk() (*[]ConversationsV1ServiceServiceConversationServiceConversationMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListServiceConversationMessageResponse) SetMessages(v []ConversationsV1ServiceServiceConversationServiceConversationMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListServiceConversationMessageResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMeta

`func (o *ListServiceConversationMessageResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceConversationMessageResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceConversationMessageResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceConversationMessageResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


