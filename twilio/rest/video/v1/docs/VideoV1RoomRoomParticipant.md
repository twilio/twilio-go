# VideoV1RoomRoomParticipant

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Duration** | Pointer to **NullableInt32** | Duration of time in seconds the participant was connected | [optional] 
**EndTime** | Pointer to **NullableTime** | The time when the participant disconnected from the room in ISO 8601 format | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the participant&#39;s room | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**StartTime** | Pointer to **NullableTime** | The time of participant connected to the room in ISO 8601 format | [optional] 
**Status** | Pointer to **NullableString** | The status of the Participant | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1RoomRoomParticipant

`func NewVideoV1RoomRoomParticipant() *VideoV1RoomRoomParticipant`

NewVideoV1RoomRoomParticipant instantiates a new VideoV1RoomRoomParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomRoomParticipantWithDefaults

`func NewVideoV1RoomRoomParticipantWithDefaults() *VideoV1RoomRoomParticipant`

NewVideoV1RoomRoomParticipantWithDefaults instantiates a new VideoV1RoomRoomParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1RoomRoomParticipant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1RoomRoomParticipant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1RoomRoomParticipant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1RoomRoomParticipant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1RoomRoomParticipant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1RoomRoomParticipant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *VideoV1RoomRoomParticipant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1RoomRoomParticipant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1RoomRoomParticipant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1RoomRoomParticipant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1RoomRoomParticipant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1RoomRoomParticipant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1RoomRoomParticipant) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1RoomRoomParticipant) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1RoomRoomParticipant) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1RoomRoomParticipant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1RoomRoomParticipant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1RoomRoomParticipant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDuration

`func (o *VideoV1RoomRoomParticipant) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoV1RoomRoomParticipant) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoV1RoomRoomParticipant) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoV1RoomRoomParticipant) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *VideoV1RoomRoomParticipant) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *VideoV1RoomRoomParticipant) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEndTime

`func (o *VideoV1RoomRoomParticipant) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VideoV1RoomRoomParticipant) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VideoV1RoomRoomParticipant) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VideoV1RoomRoomParticipant) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *VideoV1RoomRoomParticipant) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *VideoV1RoomRoomParticipant) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetIdentity

`func (o *VideoV1RoomRoomParticipant) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *VideoV1RoomRoomParticipant) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *VideoV1RoomRoomParticipant) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *VideoV1RoomRoomParticipant) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *VideoV1RoomRoomParticipant) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *VideoV1RoomRoomParticipant) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLinks

`func (o *VideoV1RoomRoomParticipant) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoV1RoomRoomParticipant) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoV1RoomRoomParticipant) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoV1RoomRoomParticipant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VideoV1RoomRoomParticipant) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VideoV1RoomRoomParticipant) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRoomSid

`func (o *VideoV1RoomRoomParticipant) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1RoomRoomParticipant) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1RoomRoomParticipant) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1RoomRoomParticipant) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1RoomRoomParticipant) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1RoomRoomParticipant) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetSid

`func (o *VideoV1RoomRoomParticipant) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1RoomRoomParticipant) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1RoomRoomParticipant) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1RoomRoomParticipant) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1RoomRoomParticipant) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1RoomRoomParticipant) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStartTime

`func (o *VideoV1RoomRoomParticipant) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VideoV1RoomRoomParticipant) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VideoV1RoomRoomParticipant) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VideoV1RoomRoomParticipant) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *VideoV1RoomRoomParticipant) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *VideoV1RoomRoomParticipant) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStatus

`func (o *VideoV1RoomRoomParticipant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoV1RoomRoomParticipant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoV1RoomRoomParticipant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoV1RoomRoomParticipant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VideoV1RoomRoomParticipant) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VideoV1RoomRoomParticipant) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *VideoV1RoomRoomParticipant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1RoomRoomParticipant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1RoomRoomParticipant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1RoomRoomParticipant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1RoomRoomParticipant) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1RoomRoomParticipant) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


