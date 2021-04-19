# ListServiceConversationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | Pointer to [**[]ConversationsV1ServiceServiceConversation**](ConversationsV1ServiceServiceConversation.md) |  | [optional] 
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 

## Methods

### NewListServiceConversationResponse

`func NewListServiceConversationResponse() *ListServiceConversationResponse`

NewListServiceConversationResponse instantiates a new ListServiceConversationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceConversationResponseWithDefaults

`func NewListServiceConversationResponseWithDefaults() *ListServiceConversationResponse`

NewListServiceConversationResponseWithDefaults instantiates a new ListServiceConversationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversations

`func (o *ListServiceConversationResponse) GetConversations() []ConversationsV1ServiceServiceConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *ListServiceConversationResponse) GetConversationsOk() (*[]ConversationsV1ServiceServiceConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *ListServiceConversationResponse) SetConversations(v []ConversationsV1ServiceServiceConversation)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *ListServiceConversationResponse) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetMeta

`func (o *ListServiceConversationResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceConversationResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceConversationResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceConversationResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


