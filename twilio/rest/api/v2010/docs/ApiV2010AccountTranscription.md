# ApiV2010AccountTranscription

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

### NewApiV2010AccountTranscription

`func NewApiV2010AccountTranscription() *ApiV2010AccountTranscription`

NewApiV2010AccountTranscription instantiates a new ApiV2010AccountTranscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountTranscriptionWithDefaults

`func NewApiV2010AccountTranscriptionWithDefaults() *ApiV2010AccountTranscription`

NewApiV2010AccountTranscriptionWithDefaults instantiates a new ApiV2010AccountTranscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountTranscription) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountTranscription) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountTranscription) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountTranscription) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountTranscription) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountTranscription) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountTranscription) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountTranscription) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountTranscription) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountTranscription) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountTranscription) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountTranscription) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountTranscription) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountTranscription) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountTranscription) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountTranscription) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountTranscription) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountTranscription) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountTranscription) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountTranscription) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountTranscription) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountTranscription) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountTranscription) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountTranscription) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDuration

`func (o *ApiV2010AccountTranscription) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ApiV2010AccountTranscription) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ApiV2010AccountTranscription) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ApiV2010AccountTranscription) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ApiV2010AccountTranscription) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ApiV2010AccountTranscription) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountTranscription) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountTranscription) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountTranscription) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountTranscription) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountTranscription) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountTranscription) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountTranscription) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountTranscription) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountTranscription) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountTranscription) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountTranscription) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountTranscription) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetRecordingSid

`func (o *ApiV2010AccountTranscription) GetRecordingSid() string`

GetRecordingSid returns the RecordingSid field if non-nil, zero value otherwise.

### GetRecordingSidOk

`func (o *ApiV2010AccountTranscription) GetRecordingSidOk() (*string, bool)`

GetRecordingSidOk returns a tuple with the RecordingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingSid

`func (o *ApiV2010AccountTranscription) SetRecordingSid(v string)`

SetRecordingSid sets RecordingSid field to given value.

### HasRecordingSid

`func (o *ApiV2010AccountTranscription) HasRecordingSid() bool`

HasRecordingSid returns a boolean if a field has been set.

### SetRecordingSidNil

`func (o *ApiV2010AccountTranscription) SetRecordingSidNil(b bool)`

 SetRecordingSidNil sets the value for RecordingSid to be an explicit nil

### UnsetRecordingSid
`func (o *ApiV2010AccountTranscription) UnsetRecordingSid()`

UnsetRecordingSid ensures that no value is present for RecordingSid, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountTranscription) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountTranscription) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountTranscription) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountTranscription) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountTranscription) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountTranscription) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountTranscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountTranscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountTranscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountTranscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountTranscription) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountTranscription) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTranscriptionText

`func (o *ApiV2010AccountTranscription) GetTranscriptionText() string`

GetTranscriptionText returns the TranscriptionText field if non-nil, zero value otherwise.

### GetTranscriptionTextOk

`func (o *ApiV2010AccountTranscription) GetTranscriptionTextOk() (*string, bool)`

GetTranscriptionTextOk returns a tuple with the TranscriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptionText

`func (o *ApiV2010AccountTranscription) SetTranscriptionText(v string)`

SetTranscriptionText sets TranscriptionText field to given value.

### HasTranscriptionText

`func (o *ApiV2010AccountTranscription) HasTranscriptionText() bool`

HasTranscriptionText returns a boolean if a field has been set.

### SetTranscriptionTextNil

`func (o *ApiV2010AccountTranscription) SetTranscriptionTextNil(b bool)`

 SetTranscriptionTextNil sets the value for TranscriptionText to be an explicit nil

### UnsetTranscriptionText
`func (o *ApiV2010AccountTranscription) UnsetTranscriptionText()`

UnsetTranscriptionText ensures that no value is present for TranscriptionText, not even an explicit nil
### GetType

`func (o *ApiV2010AccountTranscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV2010AccountTranscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV2010AccountTranscription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV2010AccountTranscription) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ApiV2010AccountTranscription) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ApiV2010AccountTranscription) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountTranscription) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountTranscription) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountTranscription) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountTranscription) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountTranscription) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountTranscription) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


