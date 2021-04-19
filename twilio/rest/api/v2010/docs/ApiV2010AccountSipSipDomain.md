# ApiV2010AccountSipSipDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to process the call | [optional] 
**AuthType** | Pointer to **NullableString** | The types of authentication mapped to the domain | [optional] 
**ByocTrunkSid** | Pointer to **NullableString** | The SID of the BYOC Trunk resource. | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**DomainName** | Pointer to **NullableString** | The unique address on Twilio to route SIP traffic | [optional] 
**EmergencyCallerSid** | Pointer to **NullableString** | Whether an emergency caller sid is configured for the domain. | [optional] 
**EmergencyCallingEnabled** | Pointer to **NullableBool** | Whether emergency calling is enabled for the domain. | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Secure** | Pointer to **NullableBool** | Whether secure SIP is enabled for the domain | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SipRegistration** | Pointer to **NullableBool** | Whether SIP registration is allowed | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list mapping resources associated with the SIP Domain resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method used with voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs while executing TwiML | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method to use with voice_url | [optional] 
**VoiceStatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call voice_status_callback_url | [optional] 
**VoiceStatusCallbackUrl** | Pointer to **NullableString** | The URL that we call with status updates | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when receiving a call | [optional] 

## Methods

### NewApiV2010AccountSipSipDomain

`func NewApiV2010AccountSipSipDomain() *ApiV2010AccountSipSipDomain`

NewApiV2010AccountSipSipDomain instantiates a new ApiV2010AccountSipSipDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountSipSipDomainWithDefaults

`func NewApiV2010AccountSipSipDomainWithDefaults() *ApiV2010AccountSipSipDomain`

NewApiV2010AccountSipSipDomainWithDefaults instantiates a new ApiV2010AccountSipSipDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountSipSipDomain) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountSipSipDomain) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountSipSipDomain) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountSipSipDomain) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountSipSipDomain) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountSipSipDomain) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountSipSipDomain) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountSipSipDomain) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountSipSipDomain) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountSipSipDomain) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountSipSipDomain) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountSipSipDomain) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetAuthType

`func (o *ApiV2010AccountSipSipDomain) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ApiV2010AccountSipSipDomain) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ApiV2010AccountSipSipDomain) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ApiV2010AccountSipSipDomain) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *ApiV2010AccountSipSipDomain) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *ApiV2010AccountSipSipDomain) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetByocTrunkSid

`func (o *ApiV2010AccountSipSipDomain) GetByocTrunkSid() string`

GetByocTrunkSid returns the ByocTrunkSid field if non-nil, zero value otherwise.

### GetByocTrunkSidOk

`func (o *ApiV2010AccountSipSipDomain) GetByocTrunkSidOk() (*string, bool)`

GetByocTrunkSidOk returns a tuple with the ByocTrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByocTrunkSid

`func (o *ApiV2010AccountSipSipDomain) SetByocTrunkSid(v string)`

SetByocTrunkSid sets ByocTrunkSid field to given value.

### HasByocTrunkSid

`func (o *ApiV2010AccountSipSipDomain) HasByocTrunkSid() bool`

HasByocTrunkSid returns a boolean if a field has been set.

### SetByocTrunkSidNil

`func (o *ApiV2010AccountSipSipDomain) SetByocTrunkSidNil(b bool)`

 SetByocTrunkSidNil sets the value for ByocTrunkSid to be an explicit nil

### UnsetByocTrunkSid
`func (o *ApiV2010AccountSipSipDomain) UnsetByocTrunkSid()`

UnsetByocTrunkSid ensures that no value is present for ByocTrunkSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountSipSipDomain) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountSipSipDomain) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountSipSipDomain) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountSipSipDomain) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountSipSipDomain) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountSipSipDomain) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountSipSipDomain) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountSipSipDomain) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountSipSipDomain) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountSipSipDomain) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountSipSipDomain) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountSipSipDomain) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDomainName

`func (o *ApiV2010AccountSipSipDomain) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ApiV2010AccountSipSipDomain) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ApiV2010AccountSipSipDomain) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ApiV2010AccountSipSipDomain) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *ApiV2010AccountSipSipDomain) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *ApiV2010AccountSipSipDomain) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetEmergencyCallerSid

`func (o *ApiV2010AccountSipSipDomain) GetEmergencyCallerSid() string`

GetEmergencyCallerSid returns the EmergencyCallerSid field if non-nil, zero value otherwise.

### GetEmergencyCallerSidOk

`func (o *ApiV2010AccountSipSipDomain) GetEmergencyCallerSidOk() (*string, bool)`

GetEmergencyCallerSidOk returns a tuple with the EmergencyCallerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyCallerSid

`func (o *ApiV2010AccountSipSipDomain) SetEmergencyCallerSid(v string)`

SetEmergencyCallerSid sets EmergencyCallerSid field to given value.

### HasEmergencyCallerSid

`func (o *ApiV2010AccountSipSipDomain) HasEmergencyCallerSid() bool`

HasEmergencyCallerSid returns a boolean if a field has been set.

### SetEmergencyCallerSidNil

`func (o *ApiV2010AccountSipSipDomain) SetEmergencyCallerSidNil(b bool)`

 SetEmergencyCallerSidNil sets the value for EmergencyCallerSid to be an explicit nil

### UnsetEmergencyCallerSid
`func (o *ApiV2010AccountSipSipDomain) UnsetEmergencyCallerSid()`

UnsetEmergencyCallerSid ensures that no value is present for EmergencyCallerSid, not even an explicit nil
### GetEmergencyCallingEnabled

`func (o *ApiV2010AccountSipSipDomain) GetEmergencyCallingEnabled() bool`

GetEmergencyCallingEnabled returns the EmergencyCallingEnabled field if non-nil, zero value otherwise.

### GetEmergencyCallingEnabledOk

`func (o *ApiV2010AccountSipSipDomain) GetEmergencyCallingEnabledOk() (*bool, bool)`

GetEmergencyCallingEnabledOk returns a tuple with the EmergencyCallingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyCallingEnabled

`func (o *ApiV2010AccountSipSipDomain) SetEmergencyCallingEnabled(v bool)`

SetEmergencyCallingEnabled sets EmergencyCallingEnabled field to given value.

### HasEmergencyCallingEnabled

`func (o *ApiV2010AccountSipSipDomain) HasEmergencyCallingEnabled() bool`

HasEmergencyCallingEnabled returns a boolean if a field has been set.

### SetEmergencyCallingEnabledNil

`func (o *ApiV2010AccountSipSipDomain) SetEmergencyCallingEnabledNil(b bool)`

 SetEmergencyCallingEnabledNil sets the value for EmergencyCallingEnabled to be an explicit nil

### UnsetEmergencyCallingEnabled
`func (o *ApiV2010AccountSipSipDomain) UnsetEmergencyCallingEnabled()`

UnsetEmergencyCallingEnabled ensures that no value is present for EmergencyCallingEnabled, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountSipSipDomain) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountSipSipDomain) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountSipSipDomain) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountSipSipDomain) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountSipSipDomain) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountSipSipDomain) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetSecure

`func (o *ApiV2010AccountSipSipDomain) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *ApiV2010AccountSipSipDomain) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *ApiV2010AccountSipSipDomain) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *ApiV2010AccountSipSipDomain) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### SetSecureNil

`func (o *ApiV2010AccountSipSipDomain) SetSecureNil(b bool)`

 SetSecureNil sets the value for Secure to be an explicit nil

### UnsetSecure
`func (o *ApiV2010AccountSipSipDomain) UnsetSecure()`

UnsetSecure ensures that no value is present for Secure, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountSipSipDomain) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountSipSipDomain) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountSipSipDomain) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountSipSipDomain) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountSipSipDomain) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountSipSipDomain) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSipRegistration

`func (o *ApiV2010AccountSipSipDomain) GetSipRegistration() bool`

GetSipRegistration returns the SipRegistration field if non-nil, zero value otherwise.

### GetSipRegistrationOk

`func (o *ApiV2010AccountSipSipDomain) GetSipRegistrationOk() (*bool, bool)`

GetSipRegistrationOk returns a tuple with the SipRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipRegistration

`func (o *ApiV2010AccountSipSipDomain) SetSipRegistration(v bool)`

SetSipRegistration sets SipRegistration field to given value.

### HasSipRegistration

`func (o *ApiV2010AccountSipSipDomain) HasSipRegistration() bool`

HasSipRegistration returns a boolean if a field has been set.

### SetSipRegistrationNil

`func (o *ApiV2010AccountSipSipDomain) SetSipRegistrationNil(b bool)`

 SetSipRegistrationNil sets the value for SipRegistration to be an explicit nil

### UnsetSipRegistration
`func (o *ApiV2010AccountSipSipDomain) UnsetSipRegistration()`

UnsetSipRegistration ensures that no value is present for SipRegistration, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountSipSipDomain) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountSipSipDomain) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountSipSipDomain) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountSipSipDomain) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountSipSipDomain) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountSipSipDomain) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountSipSipDomain) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountSipSipDomain) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountSipSipDomain) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountSipSipDomain) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountSipSipDomain) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountSipSipDomain) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *ApiV2010AccountSipSipDomain) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *ApiV2010AccountSipSipDomain) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *ApiV2010AccountSipSipDomain) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *ApiV2010AccountSipSipDomain) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *ApiV2010AccountSipSipDomain) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *ApiV2010AccountSipSipDomain) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *ApiV2010AccountSipSipDomain) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *ApiV2010AccountSipSipDomain) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *ApiV2010AccountSipSipDomain) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceStatusCallbackMethod

`func (o *ApiV2010AccountSipSipDomain) GetVoiceStatusCallbackMethod() string`

GetVoiceStatusCallbackMethod returns the VoiceStatusCallbackMethod field if non-nil, zero value otherwise.

### GetVoiceStatusCallbackMethodOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceStatusCallbackMethodOk() (*string, bool)`

GetVoiceStatusCallbackMethodOk returns a tuple with the VoiceStatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceStatusCallbackMethod

`func (o *ApiV2010AccountSipSipDomain) SetVoiceStatusCallbackMethod(v string)`

SetVoiceStatusCallbackMethod sets VoiceStatusCallbackMethod field to given value.

### HasVoiceStatusCallbackMethod

`func (o *ApiV2010AccountSipSipDomain) HasVoiceStatusCallbackMethod() bool`

HasVoiceStatusCallbackMethod returns a boolean if a field has been set.

### SetVoiceStatusCallbackMethodNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceStatusCallbackMethodNil(b bool)`

 SetVoiceStatusCallbackMethodNil sets the value for VoiceStatusCallbackMethod to be an explicit nil

### UnsetVoiceStatusCallbackMethod
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceStatusCallbackMethod()`

UnsetVoiceStatusCallbackMethod ensures that no value is present for VoiceStatusCallbackMethod, not even an explicit nil
### GetVoiceStatusCallbackUrl

`func (o *ApiV2010AccountSipSipDomain) GetVoiceStatusCallbackUrl() string`

GetVoiceStatusCallbackUrl returns the VoiceStatusCallbackUrl field if non-nil, zero value otherwise.

### GetVoiceStatusCallbackUrlOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceStatusCallbackUrlOk() (*string, bool)`

GetVoiceStatusCallbackUrlOk returns a tuple with the VoiceStatusCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceStatusCallbackUrl

`func (o *ApiV2010AccountSipSipDomain) SetVoiceStatusCallbackUrl(v string)`

SetVoiceStatusCallbackUrl sets VoiceStatusCallbackUrl field to given value.

### HasVoiceStatusCallbackUrl

`func (o *ApiV2010AccountSipSipDomain) HasVoiceStatusCallbackUrl() bool`

HasVoiceStatusCallbackUrl returns a boolean if a field has been set.

### SetVoiceStatusCallbackUrlNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceStatusCallbackUrlNil(b bool)`

 SetVoiceStatusCallbackUrlNil sets the value for VoiceStatusCallbackUrl to be an explicit nil

### UnsetVoiceStatusCallbackUrl
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceStatusCallbackUrl()`

UnsetVoiceStatusCallbackUrl ensures that no value is present for VoiceStatusCallbackUrl, not even an explicit nil
### GetVoiceUrl

`func (o *ApiV2010AccountSipSipDomain) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *ApiV2010AccountSipSipDomain) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *ApiV2010AccountSipSipDomain) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *ApiV2010AccountSipSipDomain) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *ApiV2010AccountSipSipDomain) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *ApiV2010AccountSipSipDomain) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


