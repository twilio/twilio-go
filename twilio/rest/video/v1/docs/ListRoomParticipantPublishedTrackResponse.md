# ListRoomParticipantPublishedTrackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCompositionHookResponseMeta**](ListCompositionHookResponse_meta.md) |  | [optional] 
**PublishedTracks** | Pointer to [**[]VideoV1RoomRoomParticipantRoomParticipantPublishedTrack**](VideoV1RoomRoomParticipantRoomParticipantPublishedTrack.md) |  | [optional] 

## Methods

### NewListRoomParticipantPublishedTrackResponse

`func NewListRoomParticipantPublishedTrackResponse() *ListRoomParticipantPublishedTrackResponse`

NewListRoomParticipantPublishedTrackResponse instantiates a new ListRoomParticipantPublishedTrackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoomParticipantPublishedTrackResponseWithDefaults

`func NewListRoomParticipantPublishedTrackResponseWithDefaults() *ListRoomParticipantPublishedTrackResponse`

NewListRoomParticipantPublishedTrackResponseWithDefaults instantiates a new ListRoomParticipantPublishedTrackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListRoomParticipantPublishedTrackResponse) GetMeta() ListCompositionHookResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoomParticipantPublishedTrackResponse) GetMetaOk() (*ListCompositionHookResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoomParticipantPublishedTrackResponse) SetMeta(v ListCompositionHookResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoomParticipantPublishedTrackResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPublishedTracks

`func (o *ListRoomParticipantPublishedTrackResponse) GetPublishedTracks() []VideoV1RoomRoomParticipantRoomParticipantPublishedTrack`

GetPublishedTracks returns the PublishedTracks field if non-nil, zero value otherwise.

### GetPublishedTracksOk

`func (o *ListRoomParticipantPublishedTrackResponse) GetPublishedTracksOk() (*[]VideoV1RoomRoomParticipantRoomParticipantPublishedTrack, bool)`

GetPublishedTracksOk returns a tuple with the PublishedTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedTracks

`func (o *ListRoomParticipantPublishedTrackResponse) SetPublishedTracks(v []VideoV1RoomRoomParticipantRoomParticipantPublishedTrack)`

SetPublishedTracks sets PublishedTracks field to given value.

### HasPublishedTracks

`func (o *ListRoomParticipantPublishedTrackResponse) HasPublishedTracks() bool`

HasPublishedTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


