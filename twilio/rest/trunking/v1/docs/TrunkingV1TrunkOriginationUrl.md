# TrunkingV1TrunkOriginationUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Enabled** | Pointer to **NullableBool** | Whether the URL is enabled | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Priority** | Pointer to **NullableInt32** | The relative importance of the URI | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SipUrl** | Pointer to **NullableString** | The SIP address you want Twilio to route your Origination calls to | [optional] 
**TrunkSid** | Pointer to **NullableString** | The SID of the Trunk that owns the Origination URL | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**Weight** | Pointer to **NullableInt32** | The value that determines the relative load the URI should receive compared to others with the same priority | [optional] 

## Methods

### NewTrunkingV1TrunkOriginationUrl

`func NewTrunkingV1TrunkOriginationUrl() *TrunkingV1TrunkOriginationUrl`

NewTrunkingV1TrunkOriginationUrl instantiates a new TrunkingV1TrunkOriginationUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkingV1TrunkOriginationUrlWithDefaults

`func NewTrunkingV1TrunkOriginationUrlWithDefaults() *TrunkingV1TrunkOriginationUrl`

NewTrunkingV1TrunkOriginationUrlWithDefaults instantiates a new TrunkingV1TrunkOriginationUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrunkingV1TrunkOriginationUrl) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrunkingV1TrunkOriginationUrl) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrunkingV1TrunkOriginationUrl) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrunkingV1TrunkOriginationUrl) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrunkingV1TrunkOriginationUrl) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrunkingV1TrunkOriginationUrl) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TrunkingV1TrunkOriginationUrl) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrunkingV1TrunkOriginationUrl) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrunkingV1TrunkOriginationUrl) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrunkingV1TrunkOriginationUrl) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrunkingV1TrunkOriginationUrl) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrunkingV1TrunkOriginationUrl) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrunkingV1TrunkOriginationUrl) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrunkingV1TrunkOriginationUrl) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrunkingV1TrunkOriginationUrl) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrunkingV1TrunkOriginationUrl) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrunkingV1TrunkOriginationUrl) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrunkingV1TrunkOriginationUrl) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnabled

`func (o *TrunkingV1TrunkOriginationUrl) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrunkingV1TrunkOriginationUrl) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrunkingV1TrunkOriginationUrl) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TrunkingV1TrunkOriginationUrl) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *TrunkingV1TrunkOriginationUrl) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *TrunkingV1TrunkOriginationUrl) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFriendlyName

`func (o *TrunkingV1TrunkOriginationUrl) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrunkingV1TrunkOriginationUrl) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrunkingV1TrunkOriginationUrl) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrunkingV1TrunkOriginationUrl) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrunkingV1TrunkOriginationUrl) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrunkingV1TrunkOriginationUrl) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetPriority

`func (o *TrunkingV1TrunkOriginationUrl) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TrunkingV1TrunkOriginationUrl) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TrunkingV1TrunkOriginationUrl) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TrunkingV1TrunkOriginationUrl) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TrunkingV1TrunkOriginationUrl) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TrunkingV1TrunkOriginationUrl) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSid

`func (o *TrunkingV1TrunkOriginationUrl) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrunkingV1TrunkOriginationUrl) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrunkingV1TrunkOriginationUrl) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrunkingV1TrunkOriginationUrl) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrunkingV1TrunkOriginationUrl) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrunkingV1TrunkOriginationUrl) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSipUrl

`func (o *TrunkingV1TrunkOriginationUrl) GetSipUrl() string`

GetSipUrl returns the SipUrl field if non-nil, zero value otherwise.

### GetSipUrlOk

`func (o *TrunkingV1TrunkOriginationUrl) GetSipUrlOk() (*string, bool)`

GetSipUrlOk returns a tuple with the SipUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUrl

`func (o *TrunkingV1TrunkOriginationUrl) SetSipUrl(v string)`

SetSipUrl sets SipUrl field to given value.

### HasSipUrl

`func (o *TrunkingV1TrunkOriginationUrl) HasSipUrl() bool`

HasSipUrl returns a boolean if a field has been set.

### SetSipUrlNil

`func (o *TrunkingV1TrunkOriginationUrl) SetSipUrlNil(b bool)`

 SetSipUrlNil sets the value for SipUrl to be an explicit nil

### UnsetSipUrl
`func (o *TrunkingV1TrunkOriginationUrl) UnsetSipUrl()`

UnsetSipUrl ensures that no value is present for SipUrl, not even an explicit nil
### GetTrunkSid

`func (o *TrunkingV1TrunkOriginationUrl) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *TrunkingV1TrunkOriginationUrl) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *TrunkingV1TrunkOriginationUrl) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *TrunkingV1TrunkOriginationUrl) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *TrunkingV1TrunkOriginationUrl) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *TrunkingV1TrunkOriginationUrl) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUrl

`func (o *TrunkingV1TrunkOriginationUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrunkingV1TrunkOriginationUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrunkingV1TrunkOriginationUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrunkingV1TrunkOriginationUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrunkingV1TrunkOriginationUrl) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrunkingV1TrunkOriginationUrl) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWeight

`func (o *TrunkingV1TrunkOriginationUrl) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *TrunkingV1TrunkOriginationUrl) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *TrunkingV1TrunkOriginationUrl) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *TrunkingV1TrunkOriginationUrl) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *TrunkingV1TrunkOriginationUrl) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *TrunkingV1TrunkOriginationUrl) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


