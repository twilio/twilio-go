# WirelessV1Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account to which the Sim resource belongs | [optional] 
**CommandsCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call commands_callback_url | [optional] 
**CommandsCallbackUrl** | Pointer to **NullableString** | The URL we call when the SIM originates a machine-to-machine Command | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Sim resource was last updated | [optional] 
**EId** | Pointer to **NullableString** | Deprecated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the Sim resource | [optional] 
**Iccid** | Pointer to **NullableString** | The ICCID associated with the SIM | [optional] 
**IpAddress** | Pointer to **NullableString** | Deprecated | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related subresources | [optional] 
**RatePlanSid** | Pointer to **NullableString** | The SID of the RatePlan resource to which the Sim resource is assigned. | [optional] 
**ResetStatus** | Pointer to **NullableString** | The connectivity reset status of the SIM | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Sim resource | [optional] 
**SmsFallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call sms_fallback_url | [optional] 
**SmsFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs while retrieving or executing the TwiML requested from the sms_url | [optional] 
**SmsMethod** | Pointer to **NullableString** | The HTTP method we use to call sms_url | [optional] 
**SmsUrl** | Pointer to **NullableString** | The URL we call when the SIM-connected device sends an SMS message that is not a Command | [optional] 
**Status** | Pointer to **NullableString** | The status of the Sim resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs while retrieving or executing the TwiML requested from voice_url | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method we use to call voice_url | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when the SIM-connected device makes a voice call | [optional] 

## Methods

### NewWirelessV1Sim

`func NewWirelessV1Sim() *WirelessV1Sim`

NewWirelessV1Sim instantiates a new WirelessV1Sim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWirelessV1SimWithDefaults

`func NewWirelessV1SimWithDefaults() *WirelessV1Sim`

NewWirelessV1SimWithDefaults instantiates a new WirelessV1Sim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *WirelessV1Sim) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *WirelessV1Sim) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *WirelessV1Sim) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *WirelessV1Sim) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *WirelessV1Sim) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *WirelessV1Sim) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommandsCallbackMethod

`func (o *WirelessV1Sim) GetCommandsCallbackMethod() string`

GetCommandsCallbackMethod returns the CommandsCallbackMethod field if non-nil, zero value otherwise.

### GetCommandsCallbackMethodOk

`func (o *WirelessV1Sim) GetCommandsCallbackMethodOk() (*string, bool)`

GetCommandsCallbackMethodOk returns a tuple with the CommandsCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandsCallbackMethod

`func (o *WirelessV1Sim) SetCommandsCallbackMethod(v string)`

SetCommandsCallbackMethod sets CommandsCallbackMethod field to given value.

### HasCommandsCallbackMethod

`func (o *WirelessV1Sim) HasCommandsCallbackMethod() bool`

HasCommandsCallbackMethod returns a boolean if a field has been set.

### SetCommandsCallbackMethodNil

`func (o *WirelessV1Sim) SetCommandsCallbackMethodNil(b bool)`

 SetCommandsCallbackMethodNil sets the value for CommandsCallbackMethod to be an explicit nil

### UnsetCommandsCallbackMethod
`func (o *WirelessV1Sim) UnsetCommandsCallbackMethod()`

UnsetCommandsCallbackMethod ensures that no value is present for CommandsCallbackMethod, not even an explicit nil
### GetCommandsCallbackUrl

`func (o *WirelessV1Sim) GetCommandsCallbackUrl() string`

GetCommandsCallbackUrl returns the CommandsCallbackUrl field if non-nil, zero value otherwise.

### GetCommandsCallbackUrlOk

`func (o *WirelessV1Sim) GetCommandsCallbackUrlOk() (*string, bool)`

GetCommandsCallbackUrlOk returns a tuple with the CommandsCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandsCallbackUrl

`func (o *WirelessV1Sim) SetCommandsCallbackUrl(v string)`

SetCommandsCallbackUrl sets CommandsCallbackUrl field to given value.

### HasCommandsCallbackUrl

`func (o *WirelessV1Sim) HasCommandsCallbackUrl() bool`

HasCommandsCallbackUrl returns a boolean if a field has been set.

### SetCommandsCallbackUrlNil

`func (o *WirelessV1Sim) SetCommandsCallbackUrlNil(b bool)`

 SetCommandsCallbackUrlNil sets the value for CommandsCallbackUrl to be an explicit nil

### UnsetCommandsCallbackUrl
`func (o *WirelessV1Sim) UnsetCommandsCallbackUrl()`

UnsetCommandsCallbackUrl ensures that no value is present for CommandsCallbackUrl, not even an explicit nil
### GetDateCreated

`func (o *WirelessV1Sim) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *WirelessV1Sim) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *WirelessV1Sim) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *WirelessV1Sim) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *WirelessV1Sim) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *WirelessV1Sim) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *WirelessV1Sim) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *WirelessV1Sim) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *WirelessV1Sim) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *WirelessV1Sim) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *WirelessV1Sim) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *WirelessV1Sim) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEId

`func (o *WirelessV1Sim) GetEId() string`

GetEId returns the EId field if non-nil, zero value otherwise.

### GetEIdOk

`func (o *WirelessV1Sim) GetEIdOk() (*string, bool)`

GetEIdOk returns a tuple with the EId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEId

`func (o *WirelessV1Sim) SetEId(v string)`

SetEId sets EId field to given value.

### HasEId

`func (o *WirelessV1Sim) HasEId() bool`

HasEId returns a boolean if a field has been set.

### SetEIdNil

`func (o *WirelessV1Sim) SetEIdNil(b bool)`

 SetEIdNil sets the value for EId to be an explicit nil

### UnsetEId
`func (o *WirelessV1Sim) UnsetEId()`

UnsetEId ensures that no value is present for EId, not even an explicit nil
### GetFriendlyName

`func (o *WirelessV1Sim) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *WirelessV1Sim) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *WirelessV1Sim) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *WirelessV1Sim) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *WirelessV1Sim) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *WirelessV1Sim) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIccid

`func (o *WirelessV1Sim) GetIccid() string`

GetIccid returns the Iccid field if non-nil, zero value otherwise.

### GetIccidOk

`func (o *WirelessV1Sim) GetIccidOk() (*string, bool)`

GetIccidOk returns a tuple with the Iccid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIccid

`func (o *WirelessV1Sim) SetIccid(v string)`

SetIccid sets Iccid field to given value.

### HasIccid

`func (o *WirelessV1Sim) HasIccid() bool`

HasIccid returns a boolean if a field has been set.

### SetIccidNil

`func (o *WirelessV1Sim) SetIccidNil(b bool)`

 SetIccidNil sets the value for Iccid to be an explicit nil

### UnsetIccid
`func (o *WirelessV1Sim) UnsetIccid()`

UnsetIccid ensures that no value is present for Iccid, not even an explicit nil
### GetIpAddress

`func (o *WirelessV1Sim) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *WirelessV1Sim) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *WirelessV1Sim) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *WirelessV1Sim) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *WirelessV1Sim) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *WirelessV1Sim) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetLinks

`func (o *WirelessV1Sim) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WirelessV1Sim) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WirelessV1Sim) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WirelessV1Sim) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *WirelessV1Sim) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *WirelessV1Sim) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRatePlanSid

`func (o *WirelessV1Sim) GetRatePlanSid() string`

GetRatePlanSid returns the RatePlanSid field if non-nil, zero value otherwise.

### GetRatePlanSidOk

`func (o *WirelessV1Sim) GetRatePlanSidOk() (*string, bool)`

GetRatePlanSidOk returns a tuple with the RatePlanSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePlanSid

`func (o *WirelessV1Sim) SetRatePlanSid(v string)`

SetRatePlanSid sets RatePlanSid field to given value.

### HasRatePlanSid

`func (o *WirelessV1Sim) HasRatePlanSid() bool`

HasRatePlanSid returns a boolean if a field has been set.

### SetRatePlanSidNil

`func (o *WirelessV1Sim) SetRatePlanSidNil(b bool)`

 SetRatePlanSidNil sets the value for RatePlanSid to be an explicit nil

### UnsetRatePlanSid
`func (o *WirelessV1Sim) UnsetRatePlanSid()`

UnsetRatePlanSid ensures that no value is present for RatePlanSid, not even an explicit nil
### GetResetStatus

`func (o *WirelessV1Sim) GetResetStatus() string`

GetResetStatus returns the ResetStatus field if non-nil, zero value otherwise.

### GetResetStatusOk

`func (o *WirelessV1Sim) GetResetStatusOk() (*string, bool)`

GetResetStatusOk returns a tuple with the ResetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetStatus

`func (o *WirelessV1Sim) SetResetStatus(v string)`

SetResetStatus sets ResetStatus field to given value.

### HasResetStatus

`func (o *WirelessV1Sim) HasResetStatus() bool`

HasResetStatus returns a boolean if a field has been set.

### SetResetStatusNil

`func (o *WirelessV1Sim) SetResetStatusNil(b bool)`

 SetResetStatusNil sets the value for ResetStatus to be an explicit nil

### UnsetResetStatus
`func (o *WirelessV1Sim) UnsetResetStatus()`

UnsetResetStatus ensures that no value is present for ResetStatus, not even an explicit nil
### GetSid

`func (o *WirelessV1Sim) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *WirelessV1Sim) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *WirelessV1Sim) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *WirelessV1Sim) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *WirelessV1Sim) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *WirelessV1Sim) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *WirelessV1Sim) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *WirelessV1Sim) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *WirelessV1Sim) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *WirelessV1Sim) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *WirelessV1Sim) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *WirelessV1Sim) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *WirelessV1Sim) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *WirelessV1Sim) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *WirelessV1Sim) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *WirelessV1Sim) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *WirelessV1Sim) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *WirelessV1Sim) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *WirelessV1Sim) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *WirelessV1Sim) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *WirelessV1Sim) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *WirelessV1Sim) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *WirelessV1Sim) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *WirelessV1Sim) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsUrl

`func (o *WirelessV1Sim) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *WirelessV1Sim) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *WirelessV1Sim) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *WirelessV1Sim) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *WirelessV1Sim) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *WirelessV1Sim) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetStatus

`func (o *WirelessV1Sim) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WirelessV1Sim) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WirelessV1Sim) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WirelessV1Sim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *WirelessV1Sim) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WirelessV1Sim) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUniqueName

`func (o *WirelessV1Sim) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *WirelessV1Sim) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *WirelessV1Sim) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *WirelessV1Sim) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *WirelessV1Sim) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *WirelessV1Sim) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *WirelessV1Sim) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WirelessV1Sim) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WirelessV1Sim) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WirelessV1Sim) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WirelessV1Sim) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WirelessV1Sim) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *WirelessV1Sim) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *WirelessV1Sim) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *WirelessV1Sim) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *WirelessV1Sim) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *WirelessV1Sim) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *WirelessV1Sim) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *WirelessV1Sim) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *WirelessV1Sim) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *WirelessV1Sim) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *WirelessV1Sim) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *WirelessV1Sim) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *WirelessV1Sim) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *WirelessV1Sim) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *WirelessV1Sim) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *WirelessV1Sim) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *WirelessV1Sim) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *WirelessV1Sim) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *WirelessV1Sim) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceUrl

`func (o *WirelessV1Sim) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *WirelessV1Sim) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *WirelessV1Sim) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *WirelessV1Sim) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *WirelessV1Sim) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *WirelessV1Sim) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


