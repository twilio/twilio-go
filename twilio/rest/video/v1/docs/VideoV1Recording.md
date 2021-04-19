# VideoV1Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Codec** | Pointer to **NullableString** | The codec used to encode the track | [optional] 
**ContainerFormat** | Pointer to **NullableString** | The file format for the recording | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**Duration** | Pointer to **NullableInt32** | The duration of the recording in seconds | [optional] 
**GroupingSids** | Pointer to **map[string]interface{}** | A list of SIDs related to the recording | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Offset** | Pointer to **NullableInt32** | The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Size** | Pointer to **NullableInt32** | The size of the recorded track, in bytes | [optional] 
**SourceSid** | Pointer to **NullableString** | The SID of the recording source | [optional] 
**Status** | Pointer to **NullableString** | The status of the recording | [optional] 
**TrackName** | Pointer to **NullableString** | The name that was given to the source track of the recording | [optional] 
**Type** | Pointer to **NullableString** | The recording&#39;s media type | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVideoV1Recording

`func NewVideoV1Recording() *VideoV1Recording`

NewVideoV1Recording instantiates a new VideoV1Recording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoV1RecordingWithDefaults

`func NewVideoV1RecordingWithDefaults() *VideoV1Recording`

NewVideoV1RecordingWithDefaults instantiates a new VideoV1Recording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VideoV1Recording) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VideoV1Recording) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VideoV1Recording) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VideoV1Recording) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VideoV1Recording) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VideoV1Recording) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCodec

`func (o *VideoV1Recording) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *VideoV1Recording) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *VideoV1Recording) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *VideoV1Recording) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *VideoV1Recording) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *VideoV1Recording) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetContainerFormat

`func (o *VideoV1Recording) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *VideoV1Recording) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *VideoV1Recording) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *VideoV1Recording) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *VideoV1Recording) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *VideoV1Recording) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetDateCreated

`func (o *VideoV1Recording) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VideoV1Recording) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VideoV1Recording) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VideoV1Recording) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VideoV1Recording) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VideoV1Recording) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDuration

`func (o *VideoV1Recording) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoV1Recording) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoV1Recording) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoV1Recording) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *VideoV1Recording) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *VideoV1Recording) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetGroupingSids

`func (o *VideoV1Recording) GetGroupingSids() map[string]interface{}`

GetGroupingSids returns the GroupingSids field if non-nil, zero value otherwise.

### GetGroupingSidsOk

`func (o *VideoV1Recording) GetGroupingSidsOk() (*map[string]interface{}, bool)`

GetGroupingSidsOk returns a tuple with the GroupingSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingSids

`func (o *VideoV1Recording) SetGroupingSids(v map[string]interface{})`

SetGroupingSids sets GroupingSids field to given value.

### HasGroupingSids

`func (o *VideoV1Recording) HasGroupingSids() bool`

HasGroupingSids returns a boolean if a field has been set.

### SetGroupingSidsNil

`func (o *VideoV1Recording) SetGroupingSidsNil(b bool)`

 SetGroupingSidsNil sets the value for GroupingSids to be an explicit nil

### UnsetGroupingSids
`func (o *VideoV1Recording) UnsetGroupingSids()`

UnsetGroupingSids ensures that no value is present for GroupingSids, not even an explicit nil
### GetLinks

`func (o *VideoV1Recording) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VideoV1Recording) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VideoV1Recording) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VideoV1Recording) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VideoV1Recording) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VideoV1Recording) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetOffset

`func (o *VideoV1Recording) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *VideoV1Recording) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *VideoV1Recording) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *VideoV1Recording) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *VideoV1Recording) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *VideoV1Recording) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetSid

`func (o *VideoV1Recording) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VideoV1Recording) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VideoV1Recording) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VideoV1Recording) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VideoV1Recording) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VideoV1Recording) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSize

`func (o *VideoV1Recording) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoV1Recording) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoV1Recording) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoV1Recording) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VideoV1Recording) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VideoV1Recording) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSourceSid

`func (o *VideoV1Recording) GetSourceSid() string`

GetSourceSid returns the SourceSid field if non-nil, zero value otherwise.

### GetSourceSidOk

`func (o *VideoV1Recording) GetSourceSidOk() (*string, bool)`

GetSourceSidOk returns a tuple with the SourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSid

`func (o *VideoV1Recording) SetSourceSid(v string)`

SetSourceSid sets SourceSid field to given value.

### HasSourceSid

`func (o *VideoV1Recording) HasSourceSid() bool`

HasSourceSid returns a boolean if a field has been set.

### SetSourceSidNil

`func (o *VideoV1Recording) SetSourceSidNil(b bool)`

 SetSourceSidNil sets the value for SourceSid to be an explicit nil

### UnsetSourceSid
`func (o *VideoV1Recording) UnsetSourceSid()`

UnsetSourceSid ensures that no value is present for SourceSid, not even an explicit nil
### GetStatus

`func (o *VideoV1Recording) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoV1Recording) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoV1Recording) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoV1Recording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VideoV1Recording) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VideoV1Recording) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTrackName

`func (o *VideoV1Recording) GetTrackName() string`

GetTrackName returns the TrackName field if non-nil, zero value otherwise.

### GetTrackNameOk

`func (o *VideoV1Recording) GetTrackNameOk() (*string, bool)`

GetTrackNameOk returns a tuple with the TrackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackName

`func (o *VideoV1Recording) SetTrackName(v string)`

SetTrackName sets TrackName field to given value.

### HasTrackName

`func (o *VideoV1Recording) HasTrackName() bool`

HasTrackName returns a boolean if a field has been set.

### SetTrackNameNil

`func (o *VideoV1Recording) SetTrackNameNil(b bool)`

 SetTrackNameNil sets the value for TrackName to be an explicit nil

### UnsetTrackName
`func (o *VideoV1Recording) UnsetTrackName()`

UnsetTrackName ensures that no value is present for TrackName, not even an explicit nil
### GetType

`func (o *VideoV1Recording) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoV1Recording) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoV1Recording) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoV1Recording) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VideoV1Recording) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VideoV1Recording) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *VideoV1Recording) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoV1Recording) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoV1Recording) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoV1Recording) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VideoV1Recording) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VideoV1Recording) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


