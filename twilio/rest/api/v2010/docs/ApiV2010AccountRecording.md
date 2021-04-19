# ApiV2010AccountRecording

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used during the recording. | [optional] 
**CallSid** | Pointer to **NullableString** | The SID of the Call the resource is associated with | [optional] 
**Channels** | Pointer to **NullableInt32** | The number of channels in the final recording file as an integer. | [optional] 
**ConferenceSid** | Pointer to **NullableString** | The unique ID for the conference associated with the recording. | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Duration** | Pointer to **NullableString** | The length of the recording in seconds. | [optional] 
**EncryptionDetails** | Pointer to **map[string]interface{}** | How to decrypt the recording. | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | More information about why the recording is missing, if status is &#x60;absent&#x60;. | [optional] 
**Price** | Pointer to **NullableString** | The one-time cost of creating the recording. | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency used in the price property. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Source** | Pointer to **NullableString** | How the recording was created | [optional] 
**StartTime** | Pointer to **NullableString** | The start time of the recording, given in RFC 2822 format | [optional] 
**Status** | Pointer to **NullableString** | The status of the recording. | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountRecording

`func NewApiV2010AccountRecording() *ApiV2010AccountRecording`

NewApiV2010AccountRecording instantiates a new ApiV2010AccountRecording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountRecordingWithDefaults

`func NewApiV2010AccountRecordingWithDefaults() *ApiV2010AccountRecording`

NewApiV2010AccountRecordingWithDefaults instantiates a new ApiV2010AccountRecording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountRecording) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountRecording) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountRecording) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountRecording) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountRecording) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountRecording) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountRecording) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountRecording) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountRecording) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountRecording) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountRecording) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountRecording) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetCallSid

`func (o *ApiV2010AccountRecording) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *ApiV2010AccountRecording) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *ApiV2010AccountRecording) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *ApiV2010AccountRecording) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### SetCallSidNil

`func (o *ApiV2010AccountRecording) SetCallSidNil(b bool)`

 SetCallSidNil sets the value for CallSid to be an explicit nil

### UnsetCallSid
`func (o *ApiV2010AccountRecording) UnsetCallSid()`

UnsetCallSid ensures that no value is present for CallSid, not even an explicit nil
### GetChannels

`func (o *ApiV2010AccountRecording) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ApiV2010AccountRecording) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ApiV2010AccountRecording) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ApiV2010AccountRecording) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *ApiV2010AccountRecording) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *ApiV2010AccountRecording) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetConferenceSid

`func (o *ApiV2010AccountRecording) GetConferenceSid() string`

GetConferenceSid returns the ConferenceSid field if non-nil, zero value otherwise.

### GetConferenceSidOk

`func (o *ApiV2010AccountRecording) GetConferenceSidOk() (*string, bool)`

GetConferenceSidOk returns a tuple with the ConferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceSid

`func (o *ApiV2010AccountRecording) SetConferenceSid(v string)`

SetConferenceSid sets ConferenceSid field to given value.

### HasConferenceSid

`func (o *ApiV2010AccountRecording) HasConferenceSid() bool`

HasConferenceSid returns a boolean if a field has been set.

### SetConferenceSidNil

`func (o *ApiV2010AccountRecording) SetConferenceSidNil(b bool)`

 SetConferenceSidNil sets the value for ConferenceSid to be an explicit nil

### UnsetConferenceSid
`func (o *ApiV2010AccountRecording) UnsetConferenceSid()`

UnsetConferenceSid ensures that no value is present for ConferenceSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountRecording) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountRecording) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountRecording) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountRecording) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountRecording) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountRecording) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountRecording) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountRecording) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountRecording) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountRecording) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountRecording) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountRecording) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDuration

`func (o *ApiV2010AccountRecording) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2010AccountRecording) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2010AccountRecording) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2010AccountRecording) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2010AccountRecording) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2010AccountRecording) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetEncryptionDetails

`func (o *ApiV2010AccountRecording) GetEncryptionDetails() map[string]interface{}`

GetEncryptionDetails returns the EncryptionDetails field if non-nil, zero value otherwise.

### GetEncryptionDetailsOk

`func (o *ApiV2010AccountRecording) GetEncryptionDetailsOk() (*map[string]interface{}, bool)`

GetEncryptionDetailsOk returns a tuple with the EncryptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionDetails

`func (o *ApiV2010AccountRecording) SetEncryptionDetails(v map[string]interface{})`

SetEncryptionDetails sets EncryptionDetails field to given value.

### HasEncryptionDetails

`func (o *ApiV2010AccountRecording) HasEncryptionDetails() bool`

HasEncryptionDetails returns a boolean if a field has been set.

### SetEncryptionDetailsNil

`func (o *ApiV2010AccountRecording) SetEncryptionDetailsNil(b bool)`

 SetEncryptionDetailsNil sets the value for EncryptionDetails to be an explicit nil

### UnsetEncryptionDetails
`func (o *ApiV2010AccountRecording) UnsetEncryptionDetails()`

UnsetEncryptionDetails ensures that no value is present for EncryptionDetails, not even an explicit nil
### GetErrorCode

`func (o *ApiV2010AccountRecording) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApiV2010AccountRecording) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApiV2010AccountRecording) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApiV2010AccountRecording) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ApiV2010AccountRecording) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ApiV2010AccountRecording) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountRecording) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountRecording) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountRecording) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountRecording) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountRecording) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountRecording) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountRecording) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountRecording) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountRecording) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountRecording) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountRecording) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountRecording) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountRecording) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountRecording) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountRecording) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountRecording) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountRecording) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountRecording) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSource

`func (o *ApiV2010AccountRecording) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiV2010AccountRecording) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiV2010AccountRecording) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiV2010AccountRecording) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ApiV2010AccountRecording) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ApiV2010AccountRecording) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStartTime

`func (o *ApiV2010AccountRecording) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApiV2010AccountRecording) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApiV2010AccountRecording) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApiV2010AccountRecording) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *ApiV2010AccountRecording) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *ApiV2010AccountRecording) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountRecording) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountRecording) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountRecording) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountRecording) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountRecording) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountRecording) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountRecording) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountRecording) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountRecording) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountRecording) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountRecording) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountRecording) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountRecording) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountRecording) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountRecording) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountRecording) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountRecording) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountRecording) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


