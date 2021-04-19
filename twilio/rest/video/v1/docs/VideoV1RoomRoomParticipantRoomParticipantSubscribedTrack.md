# VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the track is enabled | [optional] 
**Kind** | Pointer to **NullableString** | The track type | [optional] 
**Name** | Pointer to **NullableString** | The track name | [optional] 
**ParticipantSid** | Pointer to **NullableString** | The SID of the participant that subscribes to the track | [optional] 
**PublisherSid** | Pointer to **NullableString** | The SID of the participant that publishes the track | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the room where the track is published | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrack

`func NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrack() *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack`

NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrack instantiates a new VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrackWithDefaults

`func NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrackWithDefaults() *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack`

NewVideoV1RoomRoomParticipantRoomParticipantSubscribedTrackWithDefaults instantiates a new VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnabled

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetKind

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetName

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetPublisherSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetPublisherSid() string`

GetPublisherSid returns the PublisherSid field if non-nil, zero value otherwise.

### GetPublisherSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetPublisherSidOk() (*string, bool)`

GetPublisherSidOk returns a tuple with the PublisherSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetPublisherSid(v string)`

SetPublisherSid sets PublisherSid field to given value.

### HasPublisherSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasPublisherSid() bool`

HasPublisherSid returns a boolean if a field has been set.

### SetPublisherSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetPublisherSidNil(b bool)`

 SetPublisherSidNil sets the value for PublisherSid to be an explicit nil

### UnsetPublisherSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetPublisherSid()`

UnsetPublisherSid ensures that no value is present for PublisherSid, not even an explicit nil
### GetRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


