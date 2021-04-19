# VerifyV2VerificationAttempt

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account Sid | [optional] 
**Channel** | Pointer to **NullableString** | Channel used for the attempt | [optional] 
**ChannelData** | Pointer to **map[string]interface{}** | Object with the channel information for an attempt | [optional] 
**ConversionStatus** | Pointer to **NullableString** | Status of a conversion | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date this Attempt was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Attempt was updated | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Verification Attempt | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVerifyV2VerificationAttempt

`func NewVerifyV2VerificationAttempt() *VerifyV2VerificationAttempt`

NewVerifyV2VerificationAttempt instantiates a new VerifyV2VerificationAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2VerificationAttemptWithDefaults

`func NewVerifyV2VerificationAttemptWithDefaults() *VerifyV2VerificationAttempt`

NewVerifyV2VerificationAttemptWithDefaults instantiates a new VerifyV2VerificationAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2VerificationAttempt) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2VerificationAttempt) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2VerificationAttempt) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2VerificationAttempt) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2VerificationAttempt) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2VerificationAttempt) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannel

`func (o *VerifyV2VerificationAttempt) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VerifyV2VerificationAttempt) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VerifyV2VerificationAttempt) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VerifyV2VerificationAttempt) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### SetChannelNil

`func (o *VerifyV2VerificationAttempt) SetChannelNil(b bool)`

 SetChannelNil sets the value for Channel to be an explicit nil

### UnsetChannel
`func (o *VerifyV2VerificationAttempt) UnsetChannel()`

UnsetChannel ensures that no value is present for Channel, not even an explicit nil
### GetChannelData

`func (o *VerifyV2VerificationAttempt) GetChannelData() map[string]interface{}`

GetChannelData returns the ChannelData field if non-nil, zero value otherwise.

### GetChannelDataOk

`func (o *VerifyV2VerificationAttempt) GetChannelDataOk() (*map[string]interface{}, bool)`

GetChannelDataOk returns a tuple with the ChannelData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelData

`func (o *VerifyV2VerificationAttempt) SetChannelData(v map[string]interface{})`

SetChannelData sets ChannelData field to given value.

### HasChannelData

`func (o *VerifyV2VerificationAttempt) HasChannelData() bool`

HasChannelData returns a boolean if a field has been set.

### SetChannelDataNil

`func (o *VerifyV2VerificationAttempt) SetChannelDataNil(b bool)`

 SetChannelDataNil sets the value for ChannelData to be an explicit nil

### UnsetChannelData
`func (o *VerifyV2VerificationAttempt) UnsetChannelData()`

UnsetChannelData ensures that no value is present for ChannelData, not even an explicit nil
### GetConversionStatus

`func (o *VerifyV2VerificationAttempt) GetConversionStatus() string`

GetConversionStatus returns the ConversionStatus field if non-nil, zero value otherwise.

### GetConversionStatusOk

`func (o *VerifyV2VerificationAttempt) GetConversionStatusOk() (*string, bool)`

GetConversionStatusOk returns a tuple with the ConversionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionStatus

`func (o *VerifyV2VerificationAttempt) SetConversionStatus(v string)`

SetConversionStatus sets ConversionStatus field to given value.

### HasConversionStatus

`func (o *VerifyV2VerificationAttempt) HasConversionStatus() bool`

HasConversionStatus returns a boolean if a field has been set.

### SetConversionStatusNil

`func (o *VerifyV2VerificationAttempt) SetConversionStatusNil(b bool)`

 SetConversionStatusNil sets the value for ConversionStatus to be an explicit nil

### UnsetConversionStatus
`func (o *VerifyV2VerificationAttempt) UnsetConversionStatus()`

UnsetConversionStatus ensures that no value is present for ConversionStatus, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2VerificationAttempt) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2VerificationAttempt) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2VerificationAttempt) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2VerificationAttempt) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2VerificationAttempt) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2VerificationAttempt) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2VerificationAttempt) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2VerificationAttempt) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2VerificationAttempt) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2VerificationAttempt) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2VerificationAttempt) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2VerificationAttempt) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetServiceSid

`func (o *VerifyV2VerificationAttempt) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *VerifyV2VerificationAttempt) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *VerifyV2VerificationAttempt) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *VerifyV2VerificationAttempt) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *VerifyV2VerificationAttempt) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *VerifyV2VerificationAttempt) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *VerifyV2VerificationAttempt) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2VerificationAttempt) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2VerificationAttempt) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2VerificationAttempt) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2VerificationAttempt) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2VerificationAttempt) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *VerifyV2VerificationAttempt) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2VerificationAttempt) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2VerificationAttempt) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2VerificationAttempt) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2VerificationAttempt) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2VerificationAttempt) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


