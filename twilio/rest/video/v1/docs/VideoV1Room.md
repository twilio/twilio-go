# VideoV1Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Duration** | Pointer to **NullableInt32** | The duration of the room in seconds | [optional] 
**EnableTurn** | Pointer to **NullableBool** | Enable Twilio&#39;s Network Traversal TURN service | [optional] 
**EndTime** | Pointer to **NullableTime** | The UTC end time of the room in UTC ISO 8601 format | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**MaxConcurrentPublishedTracks** | Pointer to **NullableInt32** | The maximum number of published tracks allowed in the room at the same time | [optional] 
**MaxParticipants** | Pointer to **NullableInt32** | The maximum number of concurrent Participants allowed in the room | [optional] 
**MediaRegion** | Pointer to **NullableString** | The region for the media server in Group Rooms | [optional] 
**RecordParticipantsOnConnect** | Pointer to **NullableBool** | Whether to start recording when Participants connect | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the room | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call status_callback | [optional] 
**Type** | Pointer to **NullableString** | The type of room | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VideoCodecs** | Pointer to **[]string** | An array of the video codecs that are supported when publishing a track in the room | [optional] 

## Methods

### NewVideoV1Room

`func NewVideoV1Room() *VideoV1Room`

NewVideoV1Room instantiates a new VideoV1Room object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomWithDefaults

`func NewVideoV1RoomWithDefaults() *VideoV1Room`

NewVideoV1RoomWithDefaults instantiates a new VideoV1Room object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1Room) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1Room) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1Room) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1Room) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1Room) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1Room) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *VideoV1Room) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1Room) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1Room) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1Room) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1Room) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1Room) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VideoV1Room) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VideoV1Room) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VideoV1Room) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VideoV1Room) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VideoV1Room) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VideoV1Room) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDuration

`func (o *VideoV1Room) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoV1Room) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoV1Room) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoV1Room) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *VideoV1Room) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *VideoV1Room) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEnableTurn

`func (o *VideoV1Room) GetEnableTurn() bool`

GetEnableTurn returns the EnableTurn field if non-nil, zero value otherwise.

### GetEnableTurnOk

`func (o *VideoV1Room) GetEnableTurnOk() (*bool, bool)`

GetEnableTurnOk returns a tuple with the EnableTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTurn

`func (o *VideoV1Room) SetEnableTurn(v bool)`

SetEnableTurn sets EnableTurn field to given value.

### HasEnableTurn

`func (o *VideoV1Room) HasEnableTurn() bool`

HasEnableTurn returns a boolean if a field has been set.

### SetEnableTurnNil

`func (o *VideoV1Room) SetEnableTurnNil(b bool)`

 SetEnableTurnNil sets the value for EnableTurn to be an explicit nil

### UnsetEnableTurn
`func (o *VideoV1Room) UnsetEnableTurn()`

UnsetEnableTurn ensures that no value is present for EnableTurn, not even an explicit nil
### GetEndTime

`func (o *VideoV1Room) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VideoV1Room) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VideoV1Room) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VideoV1Room) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *VideoV1Room) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *VideoV1Room) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetLinks

`func (o *VideoV1Room) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoV1Room) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoV1Room) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoV1Room) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VideoV1Room) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VideoV1Room) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMaxConcurrentPublishedTracks

`func (o *VideoV1Room) GetMaxConcurrentPublishedTracks() int32`

GetMaxConcurrentPublishedTracks returns the MaxConcurrentPublishedTracks field if non-nil, zero value otherwise.

### GetMaxConcurrentPublishedTracksOk

`func (o *VideoV1Room) GetMaxConcurrentPublishedTracksOk() (*int32, bool)`

GetMaxConcurrentPublishedTracksOk returns a tuple with the MaxConcurrentPublishedTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentPublishedTracks

`func (o *VideoV1Room) SetMaxConcurrentPublishedTracks(v int32)`

SetMaxConcurrentPublishedTracks sets MaxConcurrentPublishedTracks field to given value.

### HasMaxConcurrentPublishedTracks

`func (o *VideoV1Room) HasMaxConcurrentPublishedTracks() bool`

HasMaxConcurrentPublishedTracks returns a boolean if a field has been set.

### SetMaxConcurrentPublishedTracksNil

`func (o *VideoV1Room) SetMaxConcurrentPublishedTracksNil(b bool)`

 SetMaxConcurrentPublishedTracksNil sets the value for MaxConcurrentPublishedTracks to be an explicit nil

### UnsetMaxConcurrentPublishedTracks
`func (o *VideoV1Room) UnsetMaxConcurrentPublishedTracks()`

UnsetMaxConcurrentPublishedTracks ensures that no value is present for MaxConcurrentPublishedTracks, not even an explicit nil
### GetMaxParticipants

`func (o *VideoV1Room) GetMaxParticipants() int32`

GetMaxParticipants returns the MaxParticipants field if non-nil, zero value otherwise.

### GetMaxParticipantsOk

`func (o *VideoV1Room) GetMaxParticipantsOk() (*int32, bool)`

GetMaxParticipantsOk returns a tuple with the MaxParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParticipants

`func (o *VideoV1Room) SetMaxParticipants(v int32)`

SetMaxParticipants sets MaxParticipants field to given value.

### HasMaxParticipants

`func (o *VideoV1Room) HasMaxParticipants() bool`

HasMaxParticipants returns a boolean if a field has been set.

### SetMaxParticipantsNil

`func (o *VideoV1Room) SetMaxParticipantsNil(b bool)`

 SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil

### UnsetMaxParticipants
`func (o *VideoV1Room) UnsetMaxParticipants()`

UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
### GetMediaRegion

`func (o *VideoV1Room) GetMediaRegion() string`

GetMediaRegion returns the MediaRegion field if non-nil, zero value otherwise.

### GetMediaRegionOk

`func (o *VideoV1Room) GetMediaRegionOk() (*string, bool)`

GetMediaRegionOk returns a tuple with the MediaRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaRegion

`func (o *VideoV1Room) SetMediaRegion(v string)`

SetMediaRegion sets MediaRegion field to given value.

### HasMediaRegion

`func (o *VideoV1Room) HasMediaRegion() bool`

HasMediaRegion returns a boolean if a field has been set.

### SetMediaRegionNil

`func (o *VideoV1Room) SetMediaRegionNil(b bool)`

 SetMediaRegionNil sets the value for MediaRegion to be an explicit nil

### UnsetMediaRegion
`func (o *VideoV1Room) UnsetMediaRegion()`

UnsetMediaRegion ensures that no value is present for MediaRegion, not even an explicit nil
### GetRecordParticipantsOnConnect

`func (o *VideoV1Room) GetRecordParticipantsOnConnect() bool`

GetRecordParticipantsOnConnect returns the RecordParticipantsOnConnect field if non-nil, zero value otherwise.

### GetRecordParticipantsOnConnectOk

`func (o *VideoV1Room) GetRecordParticipantsOnConnectOk() (*bool, bool)`

GetRecordParticipantsOnConnectOk returns a tuple with the RecordParticipantsOnConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordParticipantsOnConnect

`func (o *VideoV1Room) SetRecordParticipantsOnConnect(v bool)`

SetRecordParticipantsOnConnect sets RecordParticipantsOnConnect field to given value.

### HasRecordParticipantsOnConnect

`func (o *VideoV1Room) HasRecordParticipantsOnConnect() bool`

HasRecordParticipantsOnConnect returns a boolean if a field has been set.

### SetRecordParticipantsOnConnectNil

`func (o *VideoV1Room) SetRecordParticipantsOnConnectNil(b bool)`

 SetRecordParticipantsOnConnectNil sets the value for RecordParticipantsOnConnect to be an explicit nil

### UnsetRecordParticipantsOnConnect
`func (o *VideoV1Room) UnsetRecordParticipantsOnConnect()`

UnsetRecordParticipantsOnConnect ensures that no value is present for RecordParticipantsOnConnect, not even an explicit nil
### GetSid

`func (o *VideoV1Room) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1Room) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1Room) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1Room) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1Room) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1Room) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *VideoV1Room) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoV1Room) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoV1Room) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoV1Room) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VideoV1Room) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VideoV1Room) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *VideoV1Room) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *VideoV1Room) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *VideoV1Room) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *VideoV1Room) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *VideoV1Room) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *VideoV1Room) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *VideoV1Room) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *VideoV1Room) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *VideoV1Room) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *VideoV1Room) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *VideoV1Room) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *VideoV1Room) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetType

`func (o *VideoV1Room) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoV1Room) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoV1Room) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoV1Room) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VideoV1Room) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VideoV1Room) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUniqueName

`func (o *VideoV1Room) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *VideoV1Room) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *VideoV1Room) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *VideoV1Room) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *VideoV1Room) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *VideoV1Room) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *VideoV1Room) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1Room) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1Room) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1Room) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1Room) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1Room) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVideoCodecs

`func (o *VideoV1Room) GetVideoCodecs() []string`

GetVideoCodecs returns the VideoCodecs field if non-nil, zero value otherwise.

### GetVideoCodecsOk

`func (o *VideoV1Room) GetVideoCodecsOk() (*[]string, bool)`

GetVideoCodecsOk returns a tuple with the VideoCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodecs

`func (o *VideoV1Room) SetVideoCodecs(v []string)`

SetVideoCodecs sets VideoCodecs field to given value.

### HasVideoCodecs

`func (o *VideoV1Room) HasVideoCodecs() bool`

HasVideoCodecs returns a boolean if a field has been set.

### SetVideoCodecsNil

`func (o *VideoV1Room) SetVideoCodecsNil(b bool)`

 SetVideoCodecsNil sets the value for VideoCodecs to be an explicit nil

### UnsetVideoCodecs
`func (o *VideoV1Room) UnsetVideoCodecs()`

UnsetVideoCodecs ensures that no value is present for VideoCodecs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


