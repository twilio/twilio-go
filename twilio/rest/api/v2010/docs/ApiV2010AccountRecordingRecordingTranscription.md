# ApiV2010AccountRecordingRecordingTranscription

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to create the transcription | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Duration** | Pointer to **NullableString** | The duration of the transcribed audio in seconds. | [optional] 
**Price** | Pointer to **NullableFloat32** | The charge for the transcription | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which price is measured | [optional] 
**RecordingSid** | Pointer to **NullableString** | The SID that identifies the transcription&#39;s recording | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the transcription | [optional] 
**TranscriptionText** | Pointer to **NullableString** | The text content of the transcription. | [optional] 
**Type** | Pointer to **NullableString** | The transcription type | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountRecordingRecordingTranscription

`func NewApiV2010AccountRecordingRecordingTranscription() *ApiV2010AccountRecordingRecordingTranscription`

NewApiV2010AccountRecordingRecordingTranscription instantiates a new ApiV2010AccountRecordingRecordingTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountRecordingRecordingTranscriptionWithDefaults

`func NewApiV2010AccountRecordingRecordingTranscriptionWithDefaults() *ApiV2010AccountRecordingRecordingTranscription`

NewApiV2010AccountRecordingRecordingTranscriptionWithDefaults instantiates a new ApiV2010AccountRecordingRecordingTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDuration

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetRecordingSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetRecordingSid() string`

GetRecordingSid returns the RecordingSid field if non-nil, zero value otherwise.

### GetRecordingSidOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetRecordingSidOk() (*string, bool)`

GetRecordingSidOk returns a tuple with the RecordingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetRecordingSid(v string)`

SetRecordingSid sets RecordingSid field to given value.

### HasRecordingSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasRecordingSid() bool`

HasRecordingSid returns a boolean if a field has been set.

### SetRecordingSidNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetRecordingSidNil(b bool)`

 SetRecordingSidNil sets the value for RecordingSid to be an explicit nil

### UnsetRecordingSid
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetRecordingSid()`

UnsetRecordingSid ensures that no value is present for RecordingSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTranscriptionText

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetTranscriptionText() string`

GetTranscriptionText returns the TranscriptionText field if non-nil, zero value otherwise.

### GetTranscriptionTextOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetTranscriptionTextOk() (*string, bool)`

GetTranscriptionTextOk returns a tuple with the TranscriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionText

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetTranscriptionText(v string)`

SetTranscriptionText sets TranscriptionText field to given value.

### HasTranscriptionText

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasTranscriptionText() bool`

HasTranscriptionText returns a boolean if a field has been set.

### SetTranscriptionTextNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetTranscriptionTextNil(b bool)`

 SetTranscriptionTextNil sets the value for TranscriptionText to be an explicit nil

### UnsetTranscriptionText
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetTranscriptionText()`

UnsetTranscriptionText ensures that no value is present for TranscriptionText, not even an explicit nil
### GetType

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountRecordingRecordingTranscription) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountRecordingRecordingTranscription) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountRecordingRecordingTranscription) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountRecordingRecordingTranscription) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


