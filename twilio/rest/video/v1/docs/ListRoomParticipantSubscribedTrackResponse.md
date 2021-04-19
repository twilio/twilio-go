# ListRoomParticipantSubscribedTrackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**SubscribedTracks** | Pointer to [**[]VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack**](VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack.md) |  | [optional] 

## Methods

### NewListRoomParticipantSubscribedTrackResponse

`func NewListRoomParticipantSubscribedTrackResponse() *ListRoomParticipantSubscribedTrackResponse`

NewListRoomParticipantSubscribedTrackResponse instantiates a new ListRoomParticipantSubscribedTrackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomParticipantSubscribedTrackResponseWithDefaults

`func NewListRoomParticipantSubscribedTrackResponseWithDefaults() *ListRoomParticipantSubscribedTrackResponse`

NewListRoomParticipantSubscribedTrackResponseWithDefaults instantiates a new ListRoomParticipantSubscribedTrackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoomParticipantSubscribedTrackResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomParticipantSubscribedTrackResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomParticipantSubscribedTrackResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomParticipantSubscribedTrackResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSubscribedTracks

`func (o *ListRoomParticipantSubscribedTrackResponse) GetSubscribedTracks() []VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack`

GetSubscribedTracks returns the SubscribedTracks field if non-nil, zero value otherwise.

### GetSubscribedTracksOk

`func (o *ListRoomParticipantSubscribedTrackResponse) GetSubscribedTracksOk() (*[]VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack, bool)`

GetSubscribedTracksOk returns a tuple with the SubscribedTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedTracks

`func (o *ListRoomParticipantSubscribedTrackResponse) SetSubscribedTracks(v []VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack)`

SetSubscribedTracks sets SubscribedTracks field to given value.

### HasSubscribedTracks

`func (o *ListRoomParticipantSubscribedTrackResponse) HasSubscribedTracks() bool`

HasSubscribedTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


