# ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AddressRequirements** | Pointer to **NullableString** | Whether the phone number requires an Address registered with Twilio. | [optional] 
**AddressSid** | Pointer to **NullableString** | The SID of the Address resource associated with the phone number | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to start a new TwiML session | [optional] 
**Beta** | Pointer to **NullableBool** | Whether the phone number is new to the Twilio platform | [optional] 
**BundleSid** | Pointer to **NullableString** | The SID of the Bundle resource associated with number | [optional] 
**Capabilities** | Pointer to [**NullableApiV2010AccountIncomingPhoneNumberCapabilities**](api_v2010_account_incoming_phone_number_capabilities.md) |  | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**EmergencyAddressSid** | Pointer to **NullableString** | The emergency address configuration to use for emergency calling | [optional] 
**EmergencyStatus** | Pointer to **NullableString** | Whether the phone number is enabled for emergency calling | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**IdentitySid** | Pointer to **NullableString** | The SID of the Identity resource associated with number | [optional] 
**Origin** | Pointer to **NullableString** | The phone number&#39;s origin. Can be twilio or hosted. | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number in E.164 format | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SmsApplicationSid** | Pointer to **NullableString** | The SID of the application that handles SMS messages sent to the phone number | [optional] 
**SmsFallbackMethod** | Pointer to **NullableString** | The HTTP method used with sms_fallback_url | [optional] 
**SmsFallbackUrl** | Pointer to **NullableString** | The URL that we call when an error occurs while retrieving or executing the TwiML | [optional] 
**SmsMethod** | Pointer to **NullableString** | The HTTP method to use with sms_url | [optional] 
**SmsUrl** | Pointer to **NullableString** | The URL we call when the phone number receives an incoming SMS message | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call status_callback | [optional] 
**TrunkSid** | Pointer to **NullableString** | The SID of the Trunk that handles calls to the phone number | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**VoiceApplicationSid** | Pointer to **NullableString** | The SID of the application that handles calls to the phone number | [optional] 
**VoiceCallerIdLookup** | Pointer to **NullableBool** | Whether to lookup the caller&#39;s name | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method used with voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when an error occurs in TwiML | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method used with the voice_url | [optional] 
**VoiceReceiveMode** | Pointer to **NullableString** |  | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when the phone number receives a call | [optional] 

## Methods

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobileWithDefaults

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobileWithDefaults() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobileWithDefaults instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAddressRequirements() string`

GetAddressRequirements returns the AddressRequirements field if non-nil, zero value otherwise.

### GetAddressRequirementsOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAddressRequirementsOk() (*string, bool)`

GetAddressRequirementsOk returns a tuple with the AddressRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAddressRequirements(v string)`

SetAddressRequirements sets AddressRequirements field to given value.

### HasAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasAddressRequirements() bool`

HasAddressRequirements returns a boolean if a field has been set.

### SetAddressRequirementsNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAddressRequirementsNil(b bool)`

 SetAddressRequirementsNil sets the value for AddressRequirements to be an explicit nil

### UnsetAddressRequirements
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetAddressRequirements()`

UnsetAddressRequirements ensures that no value is present for AddressRequirements, not even an explicit nil
### GetAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAddressSid() string`

GetAddressSid returns the AddressSid field if non-nil, zero value otherwise.

### GetAddressSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetAddressSidOk() (*string, bool)`

GetAddressSidOk returns a tuple with the AddressSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAddressSid(v string)`

SetAddressSid sets AddressSid field to given value.

### HasAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasAddressSid() bool`

HasAddressSid returns a boolean if a field has been set.

### SetAddressSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetAddressSidNil(b bool)`

 SetAddressSidNil sets the value for AddressSid to be an explicit nil

### UnsetAddressSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetAddressSid()`

UnsetAddressSid ensures that no value is present for AddressSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### SetBetaNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetBetaNil(b bool)`

 SetBetaNil sets the value for Beta to be an explicit nil

### UnsetBeta
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetBeta()`

UnsetBeta ensures that no value is present for Beta, not even an explicit nil
### GetBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetBundleSid() string`

GetBundleSid returns the BundleSid field if non-nil, zero value otherwise.

### GetBundleSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetBundleSidOk() (*string, bool)`

GetBundleSidOk returns a tuple with the BundleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetBundleSid(v string)`

SetBundleSid sets BundleSid field to given value.

### HasBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasBundleSid() bool`

HasBundleSid returns a boolean if a field has been set.

### SetBundleSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetBundleSidNil(b bool)`

 SetBundleSidNil sets the value for BundleSid to be an explicit nil

### UnsetBundleSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetBundleSid()`

UnsetBundleSid ensures that no value is present for BundleSid, not even an explicit nil
### GetCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetCapabilities() ApiV2010AccountIncomingPhoneNumberCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetCapabilitiesOk() (*ApiV2010AccountIncomingPhoneNumberCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetCapabilities(v ApiV2010AccountIncomingPhoneNumberCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetEmergencyAddressSid() string`

GetEmergencyAddressSid returns the EmergencyAddressSid field if non-nil, zero value otherwise.

### GetEmergencyAddressSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetEmergencyAddressSidOk() (*string, bool)`

GetEmergencyAddressSidOk returns a tuple with the EmergencyAddressSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetEmergencyAddressSid(v string)`

SetEmergencyAddressSid sets EmergencyAddressSid field to given value.

### HasEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasEmergencyAddressSid() bool`

HasEmergencyAddressSid returns a boolean if a field has been set.

### SetEmergencyAddressSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetEmergencyAddressSidNil(b bool)`

 SetEmergencyAddressSidNil sets the value for EmergencyAddressSid to be an explicit nil

### UnsetEmergencyAddressSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetEmergencyAddressSid()`

UnsetEmergencyAddressSid ensures that no value is present for EmergencyAddressSid, not even an explicit nil
### GetEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetEmergencyStatus() string`

GetEmergencyStatus returns the EmergencyStatus field if non-nil, zero value otherwise.

### GetEmergencyStatusOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetEmergencyStatusOk() (*string, bool)`

GetEmergencyStatusOk returns a tuple with the EmergencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetEmergencyStatus(v string)`

SetEmergencyStatus sets EmergencyStatus field to given value.

### HasEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasEmergencyStatus() bool`

HasEmergencyStatus returns a boolean if a field has been set.

### SetEmergencyStatusNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetEmergencyStatusNil(b bool)`

 SetEmergencyStatusNil sets the value for EmergencyStatus to be an explicit nil

### UnsetEmergencyStatus
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetEmergencyStatus()`

UnsetEmergencyStatus ensures that no value is present for EmergencyStatus, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetIdentitySid() string`

GetIdentitySid returns the IdentitySid field if non-nil, zero value otherwise.

### GetIdentitySidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetIdentitySidOk() (*string, bool)`

GetIdentitySidOk returns a tuple with the IdentitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetIdentitySid(v string)`

SetIdentitySid sets IdentitySid field to given value.

### HasIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasIdentitySid() bool`

HasIdentitySid returns a boolean if a field has been set.

### SetIdentitySidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetIdentitySidNil(b bool)`

 SetIdentitySidNil sets the value for IdentitySid to be an explicit nil

### UnsetIdentitySid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetIdentitySid()`

UnsetIdentitySid ensures that no value is present for IdentitySid, not even an explicit nil
### GetOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsApplicationSid() string`

GetSmsApplicationSid returns the SmsApplicationSid field if non-nil, zero value otherwise.

### GetSmsApplicationSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsApplicationSidOk() (*string, bool)`

GetSmsApplicationSidOk returns a tuple with the SmsApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsApplicationSid(v string)`

SetSmsApplicationSid sets SmsApplicationSid field to given value.

### HasSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSmsApplicationSid() bool`

HasSmsApplicationSid returns a boolean if a field has been set.

### SetSmsApplicationSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsApplicationSidNil(b bool)`

 SetSmsApplicationSidNil sets the value for SmsApplicationSid to be an explicit nil

### UnsetSmsApplicationSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSmsApplicationSid()`

UnsetSmsApplicationSid ensures that no value is present for SmsApplicationSid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceApplicationSid() string`

GetVoiceApplicationSid returns the VoiceApplicationSid field if non-nil, zero value otherwise.

### GetVoiceApplicationSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceApplicationSidOk() (*string, bool)`

GetVoiceApplicationSidOk returns a tuple with the VoiceApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceApplicationSid(v string)`

SetVoiceApplicationSid sets VoiceApplicationSid field to given value.

### HasVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceApplicationSid() bool`

HasVoiceApplicationSid returns a boolean if a field has been set.

### SetVoiceApplicationSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceApplicationSidNil(b bool)`

 SetVoiceApplicationSidNil sets the value for VoiceApplicationSid to be an explicit nil

### UnsetVoiceApplicationSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceApplicationSid()`

UnsetVoiceApplicationSid ensures that no value is present for VoiceApplicationSid, not even an explicit nil
### GetVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceCallerIdLookup() bool`

GetVoiceCallerIdLookup returns the VoiceCallerIdLookup field if non-nil, zero value otherwise.

### GetVoiceCallerIdLookupOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceCallerIdLookupOk() (*bool, bool)`

GetVoiceCallerIdLookupOk returns a tuple with the VoiceCallerIdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceCallerIdLookup(v bool)`

SetVoiceCallerIdLookup sets VoiceCallerIdLookup field to given value.

### HasVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceCallerIdLookup() bool`

HasVoiceCallerIdLookup returns a boolean if a field has been set.

### SetVoiceCallerIdLookupNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceCallerIdLookupNil(b bool)`

 SetVoiceCallerIdLookupNil sets the value for VoiceCallerIdLookup to be an explicit nil

### UnsetVoiceCallerIdLookup
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceCallerIdLookup()`

UnsetVoiceCallerIdLookup ensures that no value is present for VoiceCallerIdLookup, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceReceiveMode() string`

GetVoiceReceiveMode returns the VoiceReceiveMode field if non-nil, zero value otherwise.

### GetVoiceReceiveModeOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceReceiveModeOk() (*string, bool)`

GetVoiceReceiveModeOk returns a tuple with the VoiceReceiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceReceiveMode(v string)`

SetVoiceReceiveMode sets VoiceReceiveMode field to given value.

### HasVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceReceiveMode() bool`

HasVoiceReceiveMode returns a boolean if a field has been set.

### SetVoiceReceiveModeNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceReceiveModeNil(b bool)`

 SetVoiceReceiveModeNil sets the value for VoiceReceiveMode to be an explicit nil

### UnsetVoiceReceiveMode
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceReceiveMode()`

UnsetVoiceReceiveMode ensures that no value is present for VoiceReceiveMode, not even an explicit nil
### GetVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberMobile) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


