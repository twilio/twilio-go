# VoiceV1ByocTrunk

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CnamLookupEnabled** | Pointer to **NullableBool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk | [optional] 
**ConnectionPolicySid** | Pointer to **NullableString** | Origination Connection Policy (to your Carrier) | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**FromDomainSid** | Pointer to **NullableString** | The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call status_callback_url | [optional] 
**StatusCallbackUrl** | Pointer to **NullableString** | The URL that we call with status updates | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method used with voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs while executing TwiML | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method to use with voice_url | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when receiving a call | [optional] 

## Methods

### NewVoiceV1ByocTrunk

`func NewVoiceV1ByocTrunk() *VoiceV1ByocTrunk`

NewVoiceV1ByocTrunk instantiates a new VoiceV1ByocTrunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceV1ByocTrunkWithDefaults

`func NewVoiceV1ByocTrunkWithDefaults() *VoiceV1ByocTrunk`

NewVoiceV1ByocTrunkWithDefaults instantiates a new VoiceV1ByocTrunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VoiceV1ByocTrunk) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VoiceV1ByocTrunk) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VoiceV1ByocTrunk) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VoiceV1ByocTrunk) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VoiceV1ByocTrunk) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VoiceV1ByocTrunk) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCnamLookupEnabled

`func (o *VoiceV1ByocTrunk) GetCnamLookupEnabled() bool`

GetCnamLookupEnabled returns the CnamLookupEnabled field if non-nil, zero value otherwise.

### GetCnamLookupEnabledOk

`func (o *VoiceV1ByocTrunk) GetCnamLookupEnabledOk() (*bool, bool)`

GetCnamLookupEnabledOk returns a tuple with the CnamLookupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamLookupEnabled

`func (o *VoiceV1ByocTrunk) SetCnamLookupEnabled(v bool)`

SetCnamLookupEnabled sets CnamLookupEnabled field to given value.

### HasCnamLookupEnabled

`func (o *VoiceV1ByocTrunk) HasCnamLookupEnabled() bool`

HasCnamLookupEnabled returns a boolean if a field has been set.

### SetCnamLookupEnabledNil

`func (o *VoiceV1ByocTrunk) SetCnamLookupEnabledNil(b bool)`

 SetCnamLookupEnabledNil sets the value for CnamLookupEnabled to be an explicit nil

### UnsetCnamLookupEnabled
`func (o *VoiceV1ByocTrunk) UnsetCnamLookupEnabled()`

UnsetCnamLookupEnabled ensures that no value is present for CnamLookupEnabled, not even an explicit nil
### GetConnectionPolicySid

`func (o *VoiceV1ByocTrunk) GetConnectionPolicySid() string`

GetConnectionPolicySid returns the ConnectionPolicySid field if non-nil, zero value otherwise.

### GetConnectionPolicySidOk

`func (o *VoiceV1ByocTrunk) GetConnectionPolicySidOk() (*string, bool)`

GetConnectionPolicySidOk returns a tuple with the ConnectionPolicySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPolicySid

`func (o *VoiceV1ByocTrunk) SetConnectionPolicySid(v string)`

SetConnectionPolicySid sets ConnectionPolicySid field to given value.

### HasConnectionPolicySid

`func (o *VoiceV1ByocTrunk) HasConnectionPolicySid() bool`

HasConnectionPolicySid returns a boolean if a field has been set.

### SetConnectionPolicySidNil

`func (o *VoiceV1ByocTrunk) SetConnectionPolicySidNil(b bool)`

 SetConnectionPolicySidNil sets the value for ConnectionPolicySid to be an explicit nil

### UnsetConnectionPolicySid
`func (o *VoiceV1ByocTrunk) UnsetConnectionPolicySid()`

UnsetConnectionPolicySid ensures that no value is present for ConnectionPolicySid, not even an explicit nil
### GetDateCreated

`func (o *VoiceV1ByocTrunk) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VoiceV1ByocTrunk) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VoiceV1ByocTrunk) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VoiceV1ByocTrunk) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VoiceV1ByocTrunk) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VoiceV1ByocTrunk) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VoiceV1ByocTrunk) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VoiceV1ByocTrunk) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VoiceV1ByocTrunk) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VoiceV1ByocTrunk) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VoiceV1ByocTrunk) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VoiceV1ByocTrunk) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *VoiceV1ByocTrunk) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VoiceV1ByocTrunk) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VoiceV1ByocTrunk) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VoiceV1ByocTrunk) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VoiceV1ByocTrunk) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VoiceV1ByocTrunk) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetFromDomainSid

`func (o *VoiceV1ByocTrunk) GetFromDomainSid() string`

GetFromDomainSid returns the FromDomainSid field if non-nil, zero value otherwise.

### GetFromDomainSidOk

`func (o *VoiceV1ByocTrunk) GetFromDomainSidOk() (*string, bool)`

GetFromDomainSidOk returns a tuple with the FromDomainSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDomainSid

`func (o *VoiceV1ByocTrunk) SetFromDomainSid(v string)`

SetFromDomainSid sets FromDomainSid field to given value.

### HasFromDomainSid

`func (o *VoiceV1ByocTrunk) HasFromDomainSid() bool`

HasFromDomainSid returns a boolean if a field has been set.

### SetFromDomainSidNil

`func (o *VoiceV1ByocTrunk) SetFromDomainSidNil(b bool)`

 SetFromDomainSidNil sets the value for FromDomainSid to be an explicit nil

### UnsetFromDomainSid
`func (o *VoiceV1ByocTrunk) UnsetFromDomainSid()`

UnsetFromDomainSid ensures that no value is present for FromDomainSid, not even an explicit nil
### GetSid

`func (o *VoiceV1ByocTrunk) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VoiceV1ByocTrunk) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VoiceV1ByocTrunk) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VoiceV1ByocTrunk) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VoiceV1ByocTrunk) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VoiceV1ByocTrunk) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatusCallbackMethod

`func (o *VoiceV1ByocTrunk) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *VoiceV1ByocTrunk) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *VoiceV1ByocTrunk) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *VoiceV1ByocTrunk) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *VoiceV1ByocTrunk) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *VoiceV1ByocTrunk) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetStatusCallbackUrl

`func (o *VoiceV1ByocTrunk) GetStatusCallbackUrl() string`

GetStatusCallbackUrl returns the StatusCallbackUrl field if non-nil, zero value otherwise.

### GetStatusCallbackUrlOk

`func (o *VoiceV1ByocTrunk) GetStatusCallbackUrlOk() (*string, bool)`

GetStatusCallbackUrlOk returns a tuple with the StatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackUrl

`func (o *VoiceV1ByocTrunk) SetStatusCallbackUrl(v string)`

SetStatusCallbackUrl sets StatusCallbackUrl field to given value.

### HasStatusCallbackUrl

`func (o *VoiceV1ByocTrunk) HasStatusCallbackUrl() bool`

HasStatusCallbackUrl returns a boolean if a field has been set.

### SetStatusCallbackUrlNil

`func (o *VoiceV1ByocTrunk) SetStatusCallbackUrlNil(b bool)`

 SetStatusCallbackUrlNil sets the value for StatusCallbackUrl to be an explicit nil

### UnsetStatusCallbackUrl
`func (o *VoiceV1ByocTrunk) UnsetStatusCallbackUrl()`

UnsetStatusCallbackUrl ensures that no value is present for StatusCallbackUrl, not even an explicit nil
### GetUrl

`func (o *VoiceV1ByocTrunk) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VoiceV1ByocTrunk) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VoiceV1ByocTrunk) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VoiceV1ByocTrunk) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VoiceV1ByocTrunk) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VoiceV1ByocTrunk) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *VoiceV1ByocTrunk) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *VoiceV1ByocTrunk) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *VoiceV1ByocTrunk) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *VoiceV1ByocTrunk) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *VoiceV1ByocTrunk) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *VoiceV1ByocTrunk) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *VoiceV1ByocTrunk) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *VoiceV1ByocTrunk) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *VoiceV1ByocTrunk) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *VoiceV1ByocTrunk) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *VoiceV1ByocTrunk) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *VoiceV1ByocTrunk) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *VoiceV1ByocTrunk) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *VoiceV1ByocTrunk) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *VoiceV1ByocTrunk) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *VoiceV1ByocTrunk) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *VoiceV1ByocTrunk) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *VoiceV1ByocTrunk) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceUrl

`func (o *VoiceV1ByocTrunk) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *VoiceV1ByocTrunk) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *VoiceV1ByocTrunk) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *VoiceV1ByocTrunk) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *VoiceV1ByocTrunk) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *VoiceV1ByocTrunk) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


