# ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFreeWithDefaults

`func NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFreeWithDefaults() *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree`

NewApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFreeWithDefaults instantiates a new ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAddressRequirements() string`

GetAddressRequirements returns the AddressRequirements field if non-nil, zero value otherwise.

### GetAddressRequirementsOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAddressRequirementsOk() (*string, bool)`

GetAddressRequirementsOk returns a tuple with the AddressRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAddressRequirements(v string)`

SetAddressRequirements sets AddressRequirements field to given value.

### HasAddressRequirements

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasAddressRequirements() bool`

HasAddressRequirements returns a boolean if a field has been set.

### SetAddressRequirementsNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAddressRequirementsNil(b bool)`

 SetAddressRequirementsNil sets the value for AddressRequirements to be an explicit nil

### UnsetAddressRequirements
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetAddressRequirements()`

UnsetAddressRequirements ensures that no value is present for AddressRequirements, not even an explicit nil
### GetAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAddressSid() string`

GetAddressSid returns the AddressSid field if non-nil, zero value otherwise.

### GetAddressSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetAddressSidOk() (*string, bool)`

GetAddressSidOk returns a tuple with the AddressSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAddressSid(v string)`

SetAddressSid sets AddressSid field to given value.

### HasAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasAddressSid() bool`

HasAddressSid returns a boolean if a field has been set.

### SetAddressSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetAddressSidNil(b bool)`

 SetAddressSidNil sets the value for AddressSid to be an explicit nil

### UnsetAddressSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetAddressSid()`

UnsetAddressSid ensures that no value is present for AddressSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetBeta() bool`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetBetaOk() (*bool, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetBeta(v bool)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### SetBetaNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetBetaNil(b bool)`

 SetBetaNil sets the value for Beta to be an explicit nil

### UnsetBeta
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetBeta()`

UnsetBeta ensures that no value is present for Beta, not even an explicit nil
### GetBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetBundleSid() string`

GetBundleSid returns the BundleSid field if non-nil, zero value otherwise.

### GetBundleSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetBundleSidOk() (*string, bool)`

GetBundleSidOk returns a tuple with the BundleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetBundleSid(v string)`

SetBundleSid sets BundleSid field to given value.

### HasBundleSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasBundleSid() bool`

HasBundleSid returns a boolean if a field has been set.

### SetBundleSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetBundleSidNil(b bool)`

 SetBundleSidNil sets the value for BundleSid to be an explicit nil

### UnsetBundleSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetBundleSid()`

UnsetBundleSid ensures that no value is present for BundleSid, not even an explicit nil
### GetCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetCapabilities() ApiV2010AccountIncomingPhoneNumberCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetCapabilitiesOk() (*ApiV2010AccountIncomingPhoneNumberCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetCapabilities(v ApiV2010AccountIncomingPhoneNumberCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetEmergencyAddressSid() string`

GetEmergencyAddressSid returns the EmergencyAddressSid field if non-nil, zero value otherwise.

### GetEmergencyAddressSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetEmergencyAddressSidOk() (*string, bool)`

GetEmergencyAddressSidOk returns a tuple with the EmergencyAddressSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetEmergencyAddressSid(v string)`

SetEmergencyAddressSid sets EmergencyAddressSid field to given value.

### HasEmergencyAddressSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasEmergencyAddressSid() bool`

HasEmergencyAddressSid returns a boolean if a field has been set.

### SetEmergencyAddressSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetEmergencyAddressSidNil(b bool)`

 SetEmergencyAddressSidNil sets the value for EmergencyAddressSid to be an explicit nil

### UnsetEmergencyAddressSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetEmergencyAddressSid()`

UnsetEmergencyAddressSid ensures that no value is present for EmergencyAddressSid, not even an explicit nil
### GetEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetEmergencyStatus() string`

GetEmergencyStatus returns the EmergencyStatus field if non-nil, zero value otherwise.

### GetEmergencyStatusOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetEmergencyStatusOk() (*string, bool)`

GetEmergencyStatusOk returns a tuple with the EmergencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetEmergencyStatus(v string)`

SetEmergencyStatus sets EmergencyStatus field to given value.

### HasEmergencyStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasEmergencyStatus() bool`

HasEmergencyStatus returns a boolean if a field has been set.

### SetEmergencyStatusNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetEmergencyStatusNil(b bool)`

 SetEmergencyStatusNil sets the value for EmergencyStatus to be an explicit nil

### UnsetEmergencyStatus
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetEmergencyStatus()`

UnsetEmergencyStatus ensures that no value is present for EmergencyStatus, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetIdentitySid() string`

GetIdentitySid returns the IdentitySid field if non-nil, zero value otherwise.

### GetIdentitySidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetIdentitySidOk() (*string, bool)`

GetIdentitySidOk returns a tuple with the IdentitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetIdentitySid(v string)`

SetIdentitySid sets IdentitySid field to given value.

### HasIdentitySid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasIdentitySid() bool`

HasIdentitySid returns a boolean if a field has been set.

### SetIdentitySidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetIdentitySidNil(b bool)`

 SetIdentitySidNil sets the value for IdentitySid to be an explicit nil

### UnsetIdentitySid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetIdentitySid()`

UnsetIdentitySid ensures that no value is present for IdentitySid, not even an explicit nil
### GetOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsApplicationSid() string`

GetSmsApplicationSid returns the SmsApplicationSid field if non-nil, zero value otherwise.

### GetSmsApplicationSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsApplicationSidOk() (*string, bool)`

GetSmsApplicationSidOk returns a tuple with the SmsApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsApplicationSid(v string)`

SetSmsApplicationSid sets SmsApplicationSid field to given value.

### HasSmsApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSmsApplicationSid() bool`

HasSmsApplicationSid returns a boolean if a field has been set.

### SetSmsApplicationSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsApplicationSidNil(b bool)`

 SetSmsApplicationSidNil sets the value for SmsApplicationSid to be an explicit nil

### UnsetSmsApplicationSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSmsApplicationSid()`

UnsetSmsApplicationSid ensures that no value is present for SmsApplicationSid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceApplicationSid() string`

GetVoiceApplicationSid returns the VoiceApplicationSid field if non-nil, zero value otherwise.

### GetVoiceApplicationSidOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceApplicationSidOk() (*string, bool)`

GetVoiceApplicationSidOk returns a tuple with the VoiceApplicationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceApplicationSid(v string)`

SetVoiceApplicationSid sets VoiceApplicationSid field to given value.

### HasVoiceApplicationSid

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceApplicationSid() bool`

HasVoiceApplicationSid returns a boolean if a field has been set.

### SetVoiceApplicationSidNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceApplicationSidNil(b bool)`

 SetVoiceApplicationSidNil sets the value for VoiceApplicationSid to be an explicit nil

### UnsetVoiceApplicationSid
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceApplicationSid()`

UnsetVoiceApplicationSid ensures that no value is present for VoiceApplicationSid, not even an explicit nil
### GetVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceCallerIdLookup() bool`

GetVoiceCallerIdLookup returns the VoiceCallerIdLookup field if non-nil, zero value otherwise.

### GetVoiceCallerIdLookupOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceCallerIdLookupOk() (*bool, bool)`

GetVoiceCallerIdLookupOk returns a tuple with the VoiceCallerIdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceCallerIdLookup(v bool)`

SetVoiceCallerIdLookup sets VoiceCallerIdLookup field to given value.

### HasVoiceCallerIdLookup

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceCallerIdLookup() bool`

HasVoiceCallerIdLookup returns a boolean if a field has been set.

### SetVoiceCallerIdLookupNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceCallerIdLookupNil(b bool)`

 SetVoiceCallerIdLookupNil sets the value for VoiceCallerIdLookup to be an explicit nil

### UnsetVoiceCallerIdLookup
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceCallerIdLookup()`

UnsetVoiceCallerIdLookup ensures that no value is present for VoiceCallerIdLookup, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceReceiveMode() string`

GetVoiceReceiveMode returns the VoiceReceiveMode field if non-nil, zero value otherwise.

### GetVoiceReceiveModeOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceReceiveModeOk() (*string, bool)`

GetVoiceReceiveModeOk returns a tuple with the VoiceReceiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceReceiveMode(v string)`

SetVoiceReceiveMode sets VoiceReceiveMode field to given value.

### HasVoiceReceiveMode

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceReceiveMode() bool`

HasVoiceReceiveMode returns a boolean if a field has been set.

### SetVoiceReceiveModeNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceReceiveModeNil(b bool)`

 SetVoiceReceiveModeNil sets the value for VoiceReceiveMode to be an explicit nil

### UnsetVoiceReceiveMode
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceReceiveMode()`

UnsetVoiceReceiveMode ensures that no value is present for VoiceReceiveMode, not even an explicit nil
### GetVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberTollFree) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


