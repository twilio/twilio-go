# WirelessV1RatePlan

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DataEnabled** | Pointer to **NullableBool** | Whether SIMs can use GPRS/3G/4G/LTE data connectivity | [optional] 
**DataLimit** | Pointer to **NullableInt32** | The total data usage in Megabytes that the Network allows during one month on the home network | [optional] 
**DataMetering** | Pointer to **NullableString** | The model used to meter data usage | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date when the resource was created, given as GMT in ISO 8601 format | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date when the resource was last updated, given as GMT in ISO 8601 format | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**InternationalRoaming** | Pointer to **[]string** | The services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States | [optional] 
**InternationalRoamingDataLimit** | Pointer to **NullableInt32** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States | [optional] 
**MessagingEnabled** | Pointer to **NullableBool** | Whether SIMs can make, send, and receive SMS using Commands | [optional] 
**NationalRoamingDataLimit** | Pointer to **NullableInt32** | The total data usage in Megabytes that the Network allows during one month on non-home networks in the United States | [optional] 
**NationalRoamingEnabled** | Pointer to **NullableBool** | Whether SIMs can roam on networks other than the home network in the United States | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VoiceEnabled** | Pointer to **NullableBool** | Whether SIMs can make and receive voice calls | [optional] 

## Methods

### NewWirelessV1RatePlan

`func NewWirelessV1RatePlan() *WirelessV1RatePlan`

NewWirelessV1RatePlan instantiates a new WirelessV1RatePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1RatePlanWithDefaults

`func NewWirelessV1RatePlanWithDefaults() *WirelessV1RatePlan`

NewWirelessV1RatePlanWithDefaults instantiates a new WirelessV1RatePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1RatePlan) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1RatePlan) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1RatePlan) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1RatePlan) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1RatePlan) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1RatePlan) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDataEnabled

`func (o *WirelessV1RatePlan) GetDataEnabled() bool`

GetDataEnabled returns the DataEnabled field if non-nil, zero value otherwise.

### GetDataEnabledOk

`func (o *WirelessV1RatePlan) GetDataEnabledOk() (*bool, bool)`

GetDataEnabledOk returns a tuple with the DataEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataEnabled

`func (o *WirelessV1RatePlan) SetDataEnabled(v bool)`

SetDataEnabled sets DataEnabled field to given value.

### HasDataEnabled

`func (o *WirelessV1RatePlan) HasDataEnabled() bool`

HasDataEnabled returns a boolean if a field has been set.

### SetDataEnabledNil

`func (o *WirelessV1RatePlan) SetDataEnabledNil(b bool)`

 SetDataEnabledNil sets the value for DataEnabled to be an explicit nil

### UnsetDataEnabled
`func (o *WirelessV1RatePlan) UnsetDataEnabled()`

UnsetDataEnabled ensures that no value is present for DataEnabled, not even an explicit nil
### GetDataLimit

`func (o *WirelessV1RatePlan) GetDataLimit() int32`

GetDataLimit returns the DataLimit field if non-nil, zero value otherwise.

### GetDataLimitOk

`func (o *WirelessV1RatePlan) GetDataLimitOk() (*int32, bool)`

GetDataLimitOk returns a tuple with the DataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLimit

`func (o *WirelessV1RatePlan) SetDataLimit(v int32)`

SetDataLimit sets DataLimit field to given value.

### HasDataLimit

`func (o *WirelessV1RatePlan) HasDataLimit() bool`

HasDataLimit returns a boolean if a field has been set.

### SetDataLimitNil

`func (o *WirelessV1RatePlan) SetDataLimitNil(b bool)`

 SetDataLimitNil sets the value for DataLimit to be an explicit nil

### UnsetDataLimit
`func (o *WirelessV1RatePlan) UnsetDataLimit()`

UnsetDataLimit ensures that no value is present for DataLimit, not even an explicit nil
### GetDataMetering

`func (o *WirelessV1RatePlan) GetDataMetering() string`

GetDataMetering returns the DataMetering field if non-nil, zero value otherwise.

### GetDataMeteringOk

`func (o *WirelessV1RatePlan) GetDataMeteringOk() (*string, bool)`

GetDataMeteringOk returns a tuple with the DataMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMetering

`func (o *WirelessV1RatePlan) SetDataMetering(v string)`

SetDataMetering sets DataMetering field to given value.

### HasDataMetering

`func (o *WirelessV1RatePlan) HasDataMetering() bool`

HasDataMetering returns a boolean if a field has been set.

### SetDataMeteringNil

`func (o *WirelessV1RatePlan) SetDataMeteringNil(b bool)`

 SetDataMeteringNil sets the value for DataMetering to be an explicit nil

### UnsetDataMetering
`func (o *WirelessV1RatePlan) UnsetDataMetering()`

UnsetDataMetering ensures that no value is present for DataMetering, not even an explicit nil
### GetDateCreated

`func (o *WirelessV1RatePlan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WirelessV1RatePlan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WirelessV1RatePlan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WirelessV1RatePlan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *WirelessV1RatePlan) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *WirelessV1RatePlan) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *WirelessV1RatePlan) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *WirelessV1RatePlan) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *WirelessV1RatePlan) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *WirelessV1RatePlan) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *WirelessV1RatePlan) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *WirelessV1RatePlan) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *WirelessV1RatePlan) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *WirelessV1RatePlan) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *WirelessV1RatePlan) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *WirelessV1RatePlan) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *WirelessV1RatePlan) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *WirelessV1RatePlan) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetInternationalRoaming

`func (o *WirelessV1RatePlan) GetInternationalRoaming() []string`

GetInternationalRoaming returns the InternationalRoaming field if non-nil, zero value otherwise.

### GetInternationalRoamingOk

`func (o *WirelessV1RatePlan) GetInternationalRoamingOk() (*[]string, bool)`

GetInternationalRoamingOk returns a tuple with the InternationalRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalRoaming

`func (o *WirelessV1RatePlan) SetInternationalRoaming(v []string)`

SetInternationalRoaming sets InternationalRoaming field to given value.

### HasInternationalRoaming

`func (o *WirelessV1RatePlan) HasInternationalRoaming() bool`

HasInternationalRoaming returns a boolean if a field has been set.

### SetInternationalRoamingNil

`func (o *WirelessV1RatePlan) SetInternationalRoamingNil(b bool)`

 SetInternationalRoamingNil sets the value for InternationalRoaming to be an explicit nil

### UnsetInternationalRoaming
`func (o *WirelessV1RatePlan) UnsetInternationalRoaming()`

UnsetInternationalRoaming ensures that no value is present for InternationalRoaming, not even an explicit nil
### GetInternationalRoamingDataLimit

`func (o *WirelessV1RatePlan) GetInternationalRoamingDataLimit() int32`

GetInternationalRoamingDataLimit returns the InternationalRoamingDataLimit field if non-nil, zero value otherwise.

### GetInternationalRoamingDataLimitOk

`func (o *WirelessV1RatePlan) GetInternationalRoamingDataLimitOk() (*int32, bool)`

GetInternationalRoamingDataLimitOk returns a tuple with the InternationalRoamingDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalRoamingDataLimit

`func (o *WirelessV1RatePlan) SetInternationalRoamingDataLimit(v int32)`

SetInternationalRoamingDataLimit sets InternationalRoamingDataLimit field to given value.

### HasInternationalRoamingDataLimit

`func (o *WirelessV1RatePlan) HasInternationalRoamingDataLimit() bool`

HasInternationalRoamingDataLimit returns a boolean if a field has been set.

### SetInternationalRoamingDataLimitNil

`func (o *WirelessV1RatePlan) SetInternationalRoamingDataLimitNil(b bool)`

 SetInternationalRoamingDataLimitNil sets the value for InternationalRoamingDataLimit to be an explicit nil

### UnsetInternationalRoamingDataLimit
`func (o *WirelessV1RatePlan) UnsetInternationalRoamingDataLimit()`

UnsetInternationalRoamingDataLimit ensures that no value is present for InternationalRoamingDataLimit, not even an explicit nil
### GetMessagingEnabled

`func (o *WirelessV1RatePlan) GetMessagingEnabled() bool`

GetMessagingEnabled returns the MessagingEnabled field if non-nil, zero value otherwise.

### GetMessagingEnabledOk

`func (o *WirelessV1RatePlan) GetMessagingEnabledOk() (*bool, bool)`

GetMessagingEnabledOk returns a tuple with the MessagingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingEnabled

`func (o *WirelessV1RatePlan) SetMessagingEnabled(v bool)`

SetMessagingEnabled sets MessagingEnabled field to given value.

### HasMessagingEnabled

`func (o *WirelessV1RatePlan) HasMessagingEnabled() bool`

HasMessagingEnabled returns a boolean if a field has been set.

### SetMessagingEnabledNil

`func (o *WirelessV1RatePlan) SetMessagingEnabledNil(b bool)`

 SetMessagingEnabledNil sets the value for MessagingEnabled to be an explicit nil

### UnsetMessagingEnabled
`func (o *WirelessV1RatePlan) UnsetMessagingEnabled()`

UnsetMessagingEnabled ensures that no value is present for MessagingEnabled, not even an explicit nil
### GetNationalRoamingDataLimit

`func (o *WirelessV1RatePlan) GetNationalRoamingDataLimit() int32`

GetNationalRoamingDataLimit returns the NationalRoamingDataLimit field if non-nil, zero value otherwise.

### GetNationalRoamingDataLimitOk

`func (o *WirelessV1RatePlan) GetNationalRoamingDataLimitOk() (*int32, bool)`

GetNationalRoamingDataLimitOk returns a tuple with the NationalRoamingDataLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRoamingDataLimit

`func (o *WirelessV1RatePlan) SetNationalRoamingDataLimit(v int32)`

SetNationalRoamingDataLimit sets NationalRoamingDataLimit field to given value.

### HasNationalRoamingDataLimit

`func (o *WirelessV1RatePlan) HasNationalRoamingDataLimit() bool`

HasNationalRoamingDataLimit returns a boolean if a field has been set.

### SetNationalRoamingDataLimitNil

`func (o *WirelessV1RatePlan) SetNationalRoamingDataLimitNil(b bool)`

 SetNationalRoamingDataLimitNil sets the value for NationalRoamingDataLimit to be an explicit nil

### UnsetNationalRoamingDataLimit
`func (o *WirelessV1RatePlan) UnsetNationalRoamingDataLimit()`

UnsetNationalRoamingDataLimit ensures that no value is present for NationalRoamingDataLimit, not even an explicit nil
### GetNationalRoamingEnabled

`func (o *WirelessV1RatePlan) GetNationalRoamingEnabled() bool`

GetNationalRoamingEnabled returns the NationalRoamingEnabled field if non-nil, zero value otherwise.

### GetNationalRoamingEnabledOk

`func (o *WirelessV1RatePlan) GetNationalRoamingEnabledOk() (*bool, bool)`

GetNationalRoamingEnabledOk returns a tuple with the NationalRoamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalRoamingEnabled

`func (o *WirelessV1RatePlan) SetNationalRoamingEnabled(v bool)`

SetNationalRoamingEnabled sets NationalRoamingEnabled field to given value.

### HasNationalRoamingEnabled

`func (o *WirelessV1RatePlan) HasNationalRoamingEnabled() bool`

HasNationalRoamingEnabled returns a boolean if a field has been set.

### SetNationalRoamingEnabledNil

`func (o *WirelessV1RatePlan) SetNationalRoamingEnabledNil(b bool)`

 SetNationalRoamingEnabledNil sets the value for NationalRoamingEnabled to be an explicit nil

### UnsetNationalRoamingEnabled
`func (o *WirelessV1RatePlan) UnsetNationalRoamingEnabled()`

UnsetNationalRoamingEnabled ensures that no value is present for NationalRoamingEnabled, not even an explicit nil
### GetSid

`func (o *WirelessV1RatePlan) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *WirelessV1RatePlan) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *WirelessV1RatePlan) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *WirelessV1RatePlan) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *WirelessV1RatePlan) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *WirelessV1RatePlan) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *WirelessV1RatePlan) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WirelessV1RatePlan) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WirelessV1RatePlan) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WirelessV1RatePlan) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WirelessV1RatePlan) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WirelessV1RatePlan) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *WirelessV1RatePlan) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WirelessV1RatePlan) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WirelessV1RatePlan) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WirelessV1RatePlan) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WirelessV1RatePlan) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WirelessV1RatePlan) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVoiceEnabled

`func (o *WirelessV1RatePlan) GetVoiceEnabled() bool`

GetVoiceEnabled returns the VoiceEnabled field if non-nil, zero value otherwise.

### GetVoiceEnabledOk

`func (o *WirelessV1RatePlan) GetVoiceEnabledOk() (*bool, bool)`

GetVoiceEnabledOk returns a tuple with the VoiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceEnabled

`func (o *WirelessV1RatePlan) SetVoiceEnabled(v bool)`

SetVoiceEnabled sets VoiceEnabled field to given value.

### HasVoiceEnabled

`func (o *WirelessV1RatePlan) HasVoiceEnabled() bool`

HasVoiceEnabled returns a boolean if a field has been set.

### SetVoiceEnabledNil

`func (o *WirelessV1RatePlan) SetVoiceEnabledNil(b bool)`

 SetVoiceEnabledNil sets the value for VoiceEnabled to be an explicit nil

### UnsetVoiceEnabled
`func (o *WirelessV1RatePlan) UnsetVoiceEnabled()`

UnsetVoiceEnabled ensures that no value is present for VoiceEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


