# VideoV1RoomRoomRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Codec** | Pointer to **NullableString** | The codec used for the recording | [optional] 
**ContainerFormat** | Pointer to **NullableString** | The file format for the recording | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**Duration** | Pointer to **NullableInt32** | The duration of the recording in seconds | [optional] 
**GroupingSids** | Pointer to **map[string]interface{}** | A list of SIDs related to the Recording | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Offset** | Pointer to **NullableInt32** | The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started | [optional] 
**RoomSid** | Pointer to **NullableString** | The SID of the Room resource the recording is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Size** | Pointer to **NullableInt32** | The size of the recorded track in bytes | [optional] 
**SourceSid** | Pointer to **NullableString** | The SID of the recording source | [optional] 
**Status** | Pointer to **NullableString** | The status of the recording | [optional] 
**TrackName** | Pointer to **NullableString** | The name that was given to the source track of the recording | [optional] 
**Type** | Pointer to **NullableString** | The recording&#39;s media type | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1RoomRoomRecording

`func NewVideoV1RoomRoomRecording() *VideoV1RoomRoomRecording`

NewVideoV1RoomRoomRecording instantiates a new VideoV1RoomRoomRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RoomRoomRecordingWithDefaults

`func NewVideoV1RoomRoomRecordingWithDefaults() *VideoV1RoomRoomRecording`

NewVideoV1RoomRoomRecordingWithDefaults instantiates a new VideoV1RoomRoomRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1RoomRoomRecording) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1RoomRoomRecording) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1RoomRoomRecording) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1RoomRoomRecording) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1RoomRoomRecording) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1RoomRoomRecording) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCodec

`func (o *VideoV1RoomRoomRecording) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *VideoV1RoomRoomRecording) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *VideoV1RoomRoomRecording) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *VideoV1RoomRoomRecording) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *VideoV1RoomRoomRecording) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *VideoV1RoomRoomRecording) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetContainerFormat

`func (o *VideoV1RoomRoomRecording) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *VideoV1RoomRoomRecording) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *VideoV1RoomRoomRecording) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *VideoV1RoomRoomRecording) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *VideoV1RoomRoomRecording) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *VideoV1RoomRoomRecording) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetDateCreated

`func (o *VideoV1RoomRoomRecording) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1RoomRoomRecording) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1RoomRoomRecording) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1RoomRoomRecording) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1RoomRoomRecording) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1RoomRoomRecording) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDuration

`func (o *VideoV1RoomRoomRecording) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoV1RoomRoomRecording) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoV1RoomRoomRecording) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoV1RoomRoomRecording) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *VideoV1RoomRoomRecording) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *VideoV1RoomRoomRecording) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetGroupingSids

`func (o *VideoV1RoomRoomRecording) GetGroupingSids() map[string]interface{}`

GetGroupingSids returns the GroupingSids field if non-nil, zero value otherwise.

### GetGroupingSidsOk

`func (o *VideoV1RoomRoomRecording) GetGroupingSidsOk() (*map[string]interface{}, bool)`

GetGroupingSidsOk returns a tuple with the GroupingSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingSids

`func (o *VideoV1RoomRoomRecording) SetGroupingSids(v map[string]interface{})`

SetGroupingSids sets GroupingSids field to given value.

### HasGroupingSids

`func (o *VideoV1RoomRoomRecording) HasGroupingSids() bool`

HasGroupingSids returns a boolean if a field has been set.

### SetGroupingSidsNil

`func (o *VideoV1RoomRoomRecording) SetGroupingSidsNil(b bool)`

 SetGroupingSidsNil sets the value for GroupingSids to be an explicit nil

### UnsetGroupingSids
`func (o *VideoV1RoomRoomRecording) UnsetGroupingSids()`

UnsetGroupingSids ensures that no value is present for GroupingSids, not even an explicit nil
### GetLinks

`func (o *VideoV1RoomRoomRecording) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoV1RoomRoomRecording) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoV1RoomRoomRecording) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoV1RoomRoomRecording) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VideoV1RoomRoomRecording) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VideoV1RoomRoomRecording) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetOffset

`func (o *VideoV1RoomRoomRecording) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VideoV1RoomRoomRecording) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VideoV1RoomRoomRecording) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VideoV1RoomRoomRecording) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *VideoV1RoomRoomRecording) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *VideoV1RoomRoomRecording) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetRoomSid

`func (o *VideoV1RoomRoomRecording) GetRoomSid() string`

GetRoomSid returns the RoomSid field if non-nil, zero value otherwise.

### GetRoomSidOk

`func (o *VideoV1RoomRoomRecording) GetRoomSidOk() (*string, bool)`

GetRoomSidOk returns a tuple with the RoomSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomSid

`func (o *VideoV1RoomRoomRecording) SetRoomSid(v string)`

SetRoomSid sets RoomSid field to given value.

### HasRoomSid

`func (o *VideoV1RoomRoomRecording) HasRoomSid() bool`

HasRoomSid returns a boolean if a field has been set.

### SetRoomSidNil

`func (o *VideoV1RoomRoomRecording) SetRoomSidNil(b bool)`

 SetRoomSidNil sets the value for RoomSid to be an explicit nil

### UnsetRoomSid
`func (o *VideoV1RoomRoomRecording) UnsetRoomSid()`

UnsetRoomSid ensures that no value is present for RoomSid, not even an explicit nil
### GetSid

`func (o *VideoV1RoomRoomRecording) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1RoomRoomRecording) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1RoomRoomRecording) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1RoomRoomRecording) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1RoomRoomRecording) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1RoomRoomRecording) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSize

`func (o *VideoV1RoomRoomRecording) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoV1RoomRoomRecording) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoV1RoomRoomRecording) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoV1RoomRoomRecording) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VideoV1RoomRoomRecording) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VideoV1RoomRoomRecording) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSourceSid

`func (o *VideoV1RoomRoomRecording) GetSourceSid() string`

GetSourceSid returns the SourceSid field if non-nil, zero value otherwise.

### GetSourceSidOk

`func (o *VideoV1RoomRoomRecording) GetSourceSidOk() (*string, bool)`

GetSourceSidOk returns a tuple with the SourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSid

`func (o *VideoV1RoomRoomRecording) SetSourceSid(v string)`

SetSourceSid sets SourceSid field to given value.

### HasSourceSid

`func (o *VideoV1RoomRoomRecording) HasSourceSid() bool`

HasSourceSid returns a boolean if a field has been set.

### SetSourceSidNil

`func (o *VideoV1RoomRoomRecording) SetSourceSidNil(b bool)`

 SetSourceSidNil sets the value for SourceSid to be an explicit nil

### UnsetSourceSid
`func (o *VideoV1RoomRoomRecording) UnsetSourceSid()`

UnsetSourceSid ensures that no value is present for SourceSid, not even an explicit nil
### GetStatus

`func (o *VideoV1RoomRoomRecording) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoV1RoomRoomRecording) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoV1RoomRoomRecording) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoV1RoomRoomRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VideoV1RoomRoomRecording) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VideoV1RoomRoomRecording) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTrackName

`func (o *VideoV1RoomRoomRecording) GetTrackName() string`

GetTrackName returns the TrackName field if non-nil, zero value otherwise.

### GetTrackNameOk

`func (o *VideoV1RoomRoomRecording) GetTrackNameOk() (*string, bool)`

GetTrackNameOk returns a tuple with the TrackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackName

`func (o *VideoV1RoomRoomRecording) SetTrackName(v string)`

SetTrackName sets TrackName field to given value.

### HasTrackName

`func (o *VideoV1RoomRoomRecording) HasTrackName() bool`

HasTrackName returns a boolean if a field has been set.

### SetTrackNameNil

`func (o *VideoV1RoomRoomRecording) SetTrackNameNil(b bool)`

 SetTrackNameNil sets the value for TrackName to be an explicit nil

### UnsetTrackName
`func (o *VideoV1RoomRoomRecording) UnsetTrackName()`

UnsetTrackName ensures that no value is present for TrackName, not even an explicit nil
### GetType

`func (o *VideoV1RoomRoomRecording) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoV1RoomRoomRecording) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoV1RoomRoomRecording) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoV1RoomRoomRecording) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VideoV1RoomRoomRecording) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VideoV1RoomRoomRecording) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *VideoV1RoomRoomRecording) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1RoomRoomRecording) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1RoomRoomRecording) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1RoomRoomRecording) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1RoomRoomRecording) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1RoomRoomRecording) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


