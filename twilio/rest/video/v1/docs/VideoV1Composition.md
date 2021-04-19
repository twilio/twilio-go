# VideoV1Composition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AudioSources** | Pointer to **[]string** | The array of track names to include in the composition | [optional] 
**AudioSourcesExcluded** | Pointer to **[]string** | The array of track names to exclude from the composition | [optional] 
**Bitrate** | Pointer to **NullableInt32** | The average bit rate of the composition&#39;s media | [optional] 
**DateCompleted** | Pointer to **NullableTime** | Date when the media processing task finished | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateDeleted** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the composition generated media was deleted | [optional] 
**Duration** | Pointer to **NullableInt32** | The duration of the composition&#39;s media file in seconds | [optional] 
**Format** | Pointer to **NullableString** | The container format of the composition&#39;s media files as specified in the POST request that created the Composition resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URL of the media file associated with the composition | [optional] 
**Resolution** | Pointer to **NullableString** | The dimensions of the video image in pixels expressed as columns (width) and rows (height) | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the Group Room that generated the audio and video tracks used in the composition | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Size** | Pointer to **NullableInt32** | The size of the composed media file in bytes | [optional] 
**Status** | Pointer to **NullableString** | The status of the composition | [optional] 
**Trim** | Pointer to **NullableBool** | Whether to remove intervals with no media | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VideoLayout** | Pointer to **map[string]interface{}** | An object that describes the video layout of the composition | [optional] 

## Methods

### NewVideoV1Composition

`func NewVideoV1Composition() *VideoV1Composition`

NewVideoV1Composition instantiates a new VideoV1Composition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1CompositionWithDefaults

`func NewVideoV1CompositionWithDefaults() *VideoV1Composition`

NewVideoV1CompositionWithDefaults instantiates a new VideoV1Composition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1Composition) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1Composition) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1Composition) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1Composition) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1Composition) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1Composition) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAudioSources

`func (o *VideoV1Composition) GetAudioSources() []string`

GetAudioSources returns the AudioSources field if non-nil, zero value otherwise.

### GetAudioSourcesOk

`func (o *VideoV1Composition) GetAudioSourcesOk() (*[]string, bool)`

GetAudioSourcesOk returns a tuple with the AudioSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSources

`func (o *VideoV1Composition) SetAudioSources(v []string)`

SetAudioSources sets AudioSources field to given value.

### HasAudioSources

`func (o *VideoV1Composition) HasAudioSources() bool`

HasAudioSources returns a boolean if a field has been set.

### SetAudioSourcesNil

`func (o *VideoV1Composition) SetAudioSourcesNil(b bool)`

 SetAudioSourcesNil sets the value for AudioSources to be an explicit nil

### UnsetAudioSources
`func (o *VideoV1Composition) UnsetAudioSources()`

UnsetAudioSources ensures that no value is present for AudioSources, not even an explicit nil
### GetAudioSourcesExcluded

`func (o *VideoV1Composition) GetAudioSourcesExcluded() []string`

GetAudioSourcesExcluded returns the AudioSourcesExcluded field if non-nil, zero value otherwise.

### GetAudioSourcesExcludedOk

`func (o *VideoV1Composition) GetAudioSourcesExcludedOk() (*[]string, bool)`

GetAudioSourcesExcludedOk returns a tuple with the AudioSourcesExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSourcesExcluded

`func (o *VideoV1Composition) SetAudioSourcesExcluded(v []string)`

SetAudioSourcesExcluded sets AudioSourcesExcluded field to given value.

### HasAudioSourcesExcluded

`func (o *VideoV1Composition) HasAudioSourcesExcluded() bool`

HasAudioSourcesExcluded returns a boolean if a field has been set.

### SetAudioSourcesExcludedNil

`func (o *VideoV1Composition) SetAudioSourcesExcludedNil(b bool)`

 SetAudioSourcesExcludedNil sets the value for AudioSourcesExcluded to be an explicit nil

### UnsetAudioSourcesExcluded
`func (o *VideoV1Composition) UnsetAudioSourcesExcluded()`

UnsetAudioSourcesExcluded ensures that no value is present for AudioSourcesExcluded, not even an explicit nil
### GetBitrate

`func (o *VideoV1Composition) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *VideoV1Composition) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *VideoV1Composition) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *VideoV1Composition) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *VideoV1Composition) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *VideoV1Composition) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetDateCompleted

`func (o *VideoV1Composition) GetDateCompleted() time.Time`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *VideoV1Composition) GetDateCompletedOk() (*time.Time, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *VideoV1Composition) SetDateCompleted(v time.Time)`

SetDateCompleted sets DateCompleted field to given value.

### HasDateCompleted

`func (o *VideoV1Composition) HasDateCompleted() bool`

HasDateCompleted returns a boolean if a field has been set.

### SetDateCompletedNil

`func (o *VideoV1Composition) SetDateCompletedNil(b bool)`

 SetDateCompletedNil sets the value for DateCompleted to be an explicit nil

### UnsetDateCompleted
`func (o *VideoV1Composition) UnsetDateCompleted()`

UnsetDateCompleted ensures that no value is present for DateCompleted, not even an explicit nil
### GetDateCreated

`func (o *VideoV1Composition) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1Composition) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1Composition) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1Composition) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1Composition) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1Composition) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateDeleted

`func (o *VideoV1Composition) GetDateDeleted() time.Time`

GetDateDeleted returns the DateDeleted field if non-nil, zero value otherwise.

### GetDateDeletedOk

`func (o *VideoV1Composition) GetDateDeletedOk() (*time.Time, bool)`

GetDateDeletedOk returns a tuple with the DateDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateDeleted

`func (o *VideoV1Composition) SetDateDeleted(v time.Time)`

SetDateDeleted sets DateDeleted field to given value.

### HasDateDeleted

`func (o *VideoV1Composition) HasDateDeleted() bool`

HasDateDeleted returns a boolean if a field has been set.

### SetDateDeletedNil

`func (o *VideoV1Composition) SetDateDeletedNil(b bool)`

 SetDateDeletedNil sets the value for DateDeleted to be an explicit nil

### UnsetDateDeleted
`func (o *VideoV1Composition) UnsetDateDeleted()`

UnsetDateDeleted ensures that no value is present for DateDeleted, not even an explicit nil
### GetDuration

`func (o *VideoV1Composition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoV1Composition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoV1Composition) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoV1Composition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *VideoV1Composition) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *VideoV1Composition) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetFormat

`func (o *VideoV1Composition) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *VideoV1Composition) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *VideoV1Composition) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *VideoV1Composition) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *VideoV1Composition) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *VideoV1Composition) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetLinks

`func (o *VideoV1Composition) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoV1Composition) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoV1Composition) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoV1Composition) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VideoV1Composition) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VideoV1Composition) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetResolution

`func (o *VideoV1Composition) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VideoV1Composition) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VideoV1Composition) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VideoV1Composition) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *VideoV1Composition) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *VideoV1Composition) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetRoomSid

`func (o *VideoV1Composition) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1Composition) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1Composition) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1Composition) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1Composition) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1Composition) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetSid

`func (o *VideoV1Composition) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1Composition) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1Composition) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1Composition) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1Composition) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1Composition) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSize

`func (o *VideoV1Composition) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoV1Composition) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoV1Composition) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoV1Composition) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VideoV1Composition) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VideoV1Composition) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *VideoV1Composition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoV1Composition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoV1Composition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoV1Composition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VideoV1Composition) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VideoV1Composition) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTrim

`func (o *VideoV1Composition) GetTrim() bool`

GetTrim returns the Trim field if non-nil, zero value otherwise.

### GetTrimOk

`func (o *VideoV1Composition) GetTrimOk() (*bool, bool)`

GetTrimOk returns a tuple with the Trim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrim

`func (o *VideoV1Composition) SetTrim(v bool)`

SetTrim sets Trim field to given value.

### HasTrim

`func (o *VideoV1Composition) HasTrim() bool`

HasTrim returns a boolean if a field has been set.

### SetTrimNil

`func (o *VideoV1Composition) SetTrimNil(b bool)`

 SetTrimNil sets the value for Trim to be an explicit nil

### UnsetTrim
`func (o *VideoV1Composition) UnsetTrim()`

UnsetTrim ensures that no value is present for Trim, not even an explicit nil
### GetUrl

`func (o *VideoV1Composition) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1Composition) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1Composition) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1Composition) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1Composition) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1Composition) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVideoLayout

`func (o *VideoV1Composition) GetVideoLayout() map[string]interface{}`

GetVideoLayout returns the VideoLayout field if non-nil, zero value otherwise.

### GetVideoLayoutOk

`func (o *VideoV1Composition) GetVideoLayoutOk() (*map[string]interface{}, bool)`

GetVideoLayoutOk returns a tuple with the VideoLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLayout

`func (o *VideoV1Composition) SetVideoLayout(v map[string]interface{})`

SetVideoLayout sets VideoLayout field to given value.

### HasVideoLayout

`func (o *VideoV1Composition) HasVideoLayout() bool`

HasVideoLayout returns a boolean if a field has been set.

### SetVideoLayoutNil

`func (o *VideoV1Composition) SetVideoLayoutNil(b bool)`

 SetVideoLayoutNil sets the value for VideoLayout to be an explicit nil

### UnsetVideoLayout
`func (o *VideoV1Composition) UnsetVideoLayout()`

UnsetVideoLayout ensures that no value is present for VideoLayout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


