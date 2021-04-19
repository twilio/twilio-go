# TrunkingV1TrunkPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AddressRequirements** | Pointer to **NullableString** | Whether the phone number requires an Address registered with Twilio | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to start a new TwiML session | [optional] 
**Beta** | Pointer to **NullableBool** | Whether the phone number is new to the Twilio platform | [optional] 
**Capabilities** | Pointer to **map[string]interface{}** | Indicate if a phone can receive calls or messages | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number in E.164 format | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SmsApplicationSid** | Pointer to **NullableString** | The SID of the application that handles SMS messages sent to the phone number | [optional] 
**SmsFallbackMethod** | Pointer to **NullableString** | The HTTP method used with sms_fallback_url | [optional] 
**SmsFallbackUrl** | Pointer to **NullableString** | The URL that we call when an error occurs while retrieving or executing the TwiML | [optional] 
**SmsMethod** | Pointer to **NullableString** | The HTTP method to use with sms_url | [optional] 
**SmsUrl** | Pointer to **NullableString** | The URL we call when the phone number receives an incoming SMS message | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call status_callback | [optional] 
**TrunkSid** | Pointer to **NullableString** | The SID of the Trunk that handles calls to the phone number | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**VoiceApplicationSid** | Pointer to **NullableString** | The SID of the application that handles calls to the phone number | [optional] 
**VoiceCallerIdLookup** | Pointer to **NullableBool** | Whether to lookup the caller&#39;s name | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method that we use to call voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs in TwiML | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method used with the voice_url | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when the phone number receives a call | [optional] 

## Methods

### NewTrunkingV1TrunkPhoneNumber

`func NewTrunkingV1TrunkPhoneNumber() *TrunkingV1TrunkPhoneNumber`

NewTrunkingV1TrunkPhoneNumber instantiates a new TrunkingV1TrunkPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkingV1TrunkPhoneNumberWithDefaults

`func NewTrunkingV1TrunkPhoneNumberWithDefaults() *TrunkingV1TrunkPhoneNumber`

NewTrunkingV1TrunkPhoneNumberWithDefaults instantiates a new TrunkingV1TrunkPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrunkingV1TrunkPhoneNumber) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrunkingV1TrunkPhoneNumber) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrunkingV1TrunkPhoneNumber) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrunkingV1TrunkPhoneNumber) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrunkingV1TrunkPhoneNumber) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrunkingV1TrunkPhoneNumber) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddressRequirements

`func (o *TrunkingV1TrunkPhoneNumber) GetAddressRequirements() string`

GetAddressRequirements returns the AddressRequirements field if non-nil, zero value otherwise.

### GetAddressRequirementsOk

`func (o *TrunkingV1TrunkPhoneNumber) GetAddressRequirementsOk() (*string, bool)`

GetAddressRequirementsOk returns a tuple with the AddressRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRequirements

`func (o *TrunkingV1TrunkPhoneNumber) SetAddressRequirements(v string)`

SetAddressRequirements sets AddressRequirements field to given value.

### HasAddressRequirements

`func (o *TrunkingV1TrunkPhoneNumber) HasAddressRequirements() bool`

HasAddressRequirements returns a boolean if a field has been set.

### SetAddressRequirementsNil

`func (o *TrunkingV1TrunkPhoneNumber) SetAddressRequirementsNil(b bool)`

 SetAddressRequirementsNil sets the value for AddressRequirements to be an explicit nil

### UnsetAddressRequirements
`func (o *TrunkingV1TrunkPhoneNumber) UnsetAddressRequirements()`

UnsetAddressRequirements ensures that no value is present for AddressRequirements, not even an explicit nil
### GetApiVersion

`func (o *TrunkingV1TrunkPhoneNumber) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *TrunkingV1TrunkPhoneNumber) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *TrunkingV1TrunkPhoneNumber) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *TrunkingV1TrunkPhoneNumber) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *TrunkingV1TrunkPhoneNumber) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *TrunkingV1TrunkPhoneNumber) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetBeta

`func (o *TrunkingV1TrunkPhoneNumber) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *TrunkingV1TrunkPhoneNumber) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *TrunkingV1TrunkPhoneNumber) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *TrunkingV1TrunkPhoneNumber) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### SetBetaNil

`func (o *TrunkingV1TrunkPhoneNumber) SetBetaNil(b bool)`

 SetBetaNil sets the value for Beta to be an explicit nil

### UnsetBeta
`func (o *TrunkingV1TrunkPhoneNumber) UnsetBeta()`

UnsetBeta ensures that no value is present for Beta, not even an explicit nil
### GetCapabilities

`func (o *TrunkingV1TrunkPhoneNumber) GetCapabilities() map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *TrunkingV1TrunkPhoneNumber) GetCapabilitiesOk() (*map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *TrunkingV1TrunkPhoneNumber) SetCapabilities(v map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *TrunkingV1TrunkPhoneNumber) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *TrunkingV1TrunkPhoneNumber) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *TrunkingV1TrunkPhoneNumber) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDateCreated

`func (o *TrunkingV1TrunkPhoneNumber) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrunkingV1TrunkPhoneNumber) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrunkingV1TrunkPhoneNumber) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrunkingV1TrunkPhoneNumber) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrunkingV1TrunkPhoneNumber) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrunkingV1TrunkPhoneNumber) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrunkingV1TrunkPhoneNumber) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrunkingV1TrunkPhoneNumber) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrunkingV1TrunkPhoneNumber) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrunkingV1TrunkPhoneNumber) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrunkingV1TrunkPhoneNumber) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrunkingV1TrunkPhoneNumber) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *TrunkingV1TrunkPhoneNumber) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrunkingV1TrunkPhoneNumber) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrunkingV1TrunkPhoneNumber) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrunkingV1TrunkPhoneNumber) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrunkingV1TrunkPhoneNumber) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrunkingV1TrunkPhoneNumber) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TrunkingV1TrunkPhoneNumber) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TrunkingV1TrunkPhoneNumber) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TrunkingV1TrunkPhoneNumber) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TrunkingV1TrunkPhoneNumber) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TrunkingV1TrunkPhoneNumber) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TrunkingV1TrunkPhoneNumber) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPhoneNumber

`func (o *TrunkingV1TrunkPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TrunkingV1TrunkPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TrunkingV1TrunkPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TrunkingV1TrunkPhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *TrunkingV1TrunkPhoneNumber) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *TrunkingV1TrunkPhoneNumber) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSid

`func (o *TrunkingV1TrunkPhoneNumber) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrunkingV1TrunkPhoneNumber) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrunkingV1TrunkPhoneNumber) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsApplicationSid() string`

GetSmsApplicationSid returns the SmsApplicationSid field if non-nil, zero value otherwise.

### GetSmsApplicationSidOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsApplicationSidOk() (*string, bool)`

GetSmsApplicationSidOk returns a tuple with the SmsApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsApplicationSid(v string)`

SetSmsApplicationSid sets SmsApplicationSid field to given value.

### HasSmsApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) HasSmsApplicationSid() bool`

HasSmsApplicationSid returns a boolean if a field has been set.

### SetSmsApplicationSidNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsApplicationSidNil(b bool)`

 SetSmsApplicationSidNil sets the value for SmsApplicationSid to be an explicit nil

### UnsetSmsApplicationSid
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSmsApplicationSid()`

UnsetSmsApplicationSid ensures that no value is present for SmsApplicationSid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *TrunkingV1TrunkPhoneNumber) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsUrl

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *TrunkingV1TrunkPhoneNumber) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *TrunkingV1TrunkPhoneNumber) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *TrunkingV1TrunkPhoneNumber) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *TrunkingV1TrunkPhoneNumber) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetStatusCallback

`func (o *TrunkingV1TrunkPhoneNumber) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *TrunkingV1TrunkPhoneNumber) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *TrunkingV1TrunkPhoneNumber) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *TrunkingV1TrunkPhoneNumber) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *TrunkingV1TrunkPhoneNumber) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *TrunkingV1TrunkPhoneNumber) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *TrunkingV1TrunkPhoneNumber) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *TrunkingV1TrunkPhoneNumber) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *TrunkingV1TrunkPhoneNumber) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetTrunkSid

`func (o *TrunkingV1TrunkPhoneNumber) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *TrunkingV1TrunkPhoneNumber) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *TrunkingV1TrunkPhoneNumber) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *TrunkingV1TrunkPhoneNumber) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *TrunkingV1TrunkPhoneNumber) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *TrunkingV1TrunkPhoneNumber) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUrl

`func (o *TrunkingV1TrunkPhoneNumber) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrunkingV1TrunkPhoneNumber) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrunkingV1TrunkPhoneNumber) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrunkingV1TrunkPhoneNumber) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrunkingV1TrunkPhoneNumber) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrunkingV1TrunkPhoneNumber) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVoiceApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceApplicationSid() string`

GetVoiceApplicationSid returns the VoiceApplicationSid field if non-nil, zero value otherwise.

### GetVoiceApplicationSidOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceApplicationSidOk() (*string, bool)`

GetVoiceApplicationSidOk returns a tuple with the VoiceApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceApplicationSid(v string)`

SetVoiceApplicationSid sets VoiceApplicationSid field to given value.

### HasVoiceApplicationSid

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceApplicationSid() bool`

HasVoiceApplicationSid returns a boolean if a field has been set.

### SetVoiceApplicationSidNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceApplicationSidNil(b bool)`

 SetVoiceApplicationSidNil sets the value for VoiceApplicationSid to be an explicit nil

### UnsetVoiceApplicationSid
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceApplicationSid()`

UnsetVoiceApplicationSid ensures that no value is present for VoiceApplicationSid, not even an explicit nil
### GetVoiceCallerIdLookup

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceCallerIdLookup() bool`

GetVoiceCallerIdLookup returns the VoiceCallerIdLookup field if non-nil, zero value otherwise.

### GetVoiceCallerIdLookupOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceCallerIdLookupOk() (*bool, bool)`

GetVoiceCallerIdLookupOk returns a tuple with the VoiceCallerIdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceCallerIdLookup

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceCallerIdLookup(v bool)`

SetVoiceCallerIdLookup sets VoiceCallerIdLookup field to given value.

### HasVoiceCallerIdLookup

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceCallerIdLookup() bool`

HasVoiceCallerIdLookup returns a boolean if a field has been set.

### SetVoiceCallerIdLookupNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceCallerIdLookupNil(b bool)`

 SetVoiceCallerIdLookupNil sets the value for VoiceCallerIdLookup to be an explicit nil

### UnsetVoiceCallerIdLookup
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceCallerIdLookup()`

UnsetVoiceCallerIdLookup ensures that no value is present for VoiceCallerIdLookup, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceUrl

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *TrunkingV1TrunkPhoneNumber) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *TrunkingV1TrunkPhoneNumber) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *TrunkingV1TrunkPhoneNumber) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *TrunkingV1TrunkPhoneNumber) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


