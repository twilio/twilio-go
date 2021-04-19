# ListRoomParticipantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**Participants** | Pointer to [**[]VideoV1RoomRoomParticipant**](VideoV1RoomRoomParticipant.md) |  | [optional] 

## Methods

### NewListRoomParticipantResponse

`func NewListRoomParticipantResponse() *ListRoomParticipantResponse`

NewListRoomParticipantResponse instantiates a new ListRoomParticipantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomParticipantResponseWithDefaults

`func NewListRoomParticipantResponseWithDefaults() *ListRoomParticipantResponse`

NewListRoomParticipantResponseWithDefaults instantiates a new ListRoomParticipantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoomParticipantResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomParticipantResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomParticipantResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomParticipantResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetParticipants

`func (o *ListRoomParticipantResponse) GetParticipants() []VideoV1RoomRoomParticipant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *ListRoomParticipantResponse) GetParticipantsOk() (*[]VideoV1RoomRoomParticipant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *ListRoomParticipantResponse) SetParticipants(v []VideoV1RoomRoomParticipant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *ListRoomParticipantResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


