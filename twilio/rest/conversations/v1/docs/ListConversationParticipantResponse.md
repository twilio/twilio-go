# ListConversationParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Participants** | Pointer to [**[]ConversationsV1ConversationConversationParticipant**](ConversationsV1ConversationConversationParticipant.md) |  | [optional] 

## Methods

### NewListConversationParticipantResponse

`func NewListConversationParticipantResponse() *ListConversationParticipantResponse`

NewListConversationParticipantResponse instantiates a new ListConversationParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversationParticipantResponseWithDefaults

`func NewListConversationParticipantResponseWithDefaults() *ListConversationParticipantResponse`

NewListConversationParticipantResponseWithDefaults instantiates a new ListConversationParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListConversationParticipantResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListConversationParticipantResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListConversationParticipantResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListConversationParticipantResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParticipants

`func (o *ListConversationParticipantResponse) GetParticipants() []ConversationsV1ConversationConversationParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ListConversationParticipantResponse) GetParticipantsOk() (*[]ConversationsV1ConversationConversationParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ListConversationParticipantResponse) SetParticipants(v []ConversationsV1ConversationConversationParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ListConversationParticipantResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


