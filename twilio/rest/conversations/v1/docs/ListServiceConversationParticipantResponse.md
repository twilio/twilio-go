# ListServiceConversationParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListConversationResponseMeta**](ListConversationResponse_meta.md) |  | [optional] 
**Participants** | Pointer to [**[]ConversationsV1ServiceServiceConversationServiceConversationParticipant**](ConversationsV1ServiceServiceConversationServiceConversationParticipant.md) |  | [optional] 

## Methods

### NewListServiceConversationParticipantResponse

`func NewListServiceConversationParticipantResponse() *ListServiceConversationParticipantResponse`

NewListServiceConversationParticipantResponse instantiates a new ListServiceConversationParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceConversationParticipantResponseWithDefaults

`func NewListServiceConversationParticipantResponseWithDefaults() *ListServiceConversationParticipantResponse`

NewListServiceConversationParticipantResponseWithDefaults instantiates a new ListServiceConversationParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListServiceConversationParticipantResponse) GetMeta() ListConversationResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListServiceConversationParticipantResponse) GetMetaOk() (*ListConversationResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListServiceConversationParticipantResponse) SetMeta(v ListConversationResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListServiceConversationParticipantResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParticipants

`func (o *ListServiceConversationParticipantResponse) GetParticipants() []ConversationsV1ServiceServiceConversationServiceConversationParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ListServiceConversationParticipantResponse) GetParticipantsOk() (*[]ConversationsV1ServiceServiceConversationServiceConversationParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ListServiceConversationParticipantResponse) SetParticipants(v []ConversationsV1ServiceServiceConversationServiceConversationParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ListServiceConversationParticipantResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


