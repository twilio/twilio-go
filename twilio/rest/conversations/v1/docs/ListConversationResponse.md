# ListConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | Pointer to [**[]ConversationsV1Conversation**](ConversationsV1Conversation.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListConversationResponse

`func NewListConversationResponse() *ListConversationResponse`

NewListConversationResponse instantiates a new ListConversationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationResponseWithDefaults

`func NewListConversationResponseWithDefaults() *ListConversationResponse`

NewListConversationResponseWithDefaults instantiates a new ListConversationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversations

`func (o *ListConversationResponse) GetConversations() []ConversationsV1Conversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *ListConversationResponse) GetConversationsOk() (*[]ConversationsV1Conversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *ListConversationResponse) SetConversations(v []ConversationsV1Conversation)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *ListConversationResponse) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetMeta

`func (o *ListConversationResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConversationResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConversationResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConversationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


