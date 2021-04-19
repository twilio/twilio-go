# VerifyV2Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CodeLength** | Pointer to **NullableInt32** | The length of the verification code | [optional] 
**CustomCodeEnabled** | Pointer to **NullableBool** | Whether to allow sending verifications with a custom code. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DoNotShareWarningEnabled** | Pointer to **NullableBool** | Whether to add a security warning at the end of an SMS. | [optional] 
**DtmfInputRequired** | Pointer to **NullableBool** | Whether to ask the user to press a number before delivering the verify code in a phone call | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the verification service | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**LookupEnabled** | Pointer to **NullableBool** | Whether to perform a lookup with each verification | [optional] 
**Psd2Enabled** | Pointer to **NullableBool** | Whether to pass PSD2 transaction parameters when starting a verification | [optional] 
**Push** | Pointer to **map[string]interface{}** | The service level configuration of factor push type. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SkipSmsToLandlines** | Pointer to **NullableBool** | Whether to skip sending SMS verifications to landlines | [optional] 
**TtsName** | Pointer to **NullableString** | The name of an alternative text-to-speech service to use in phone calls | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVerifyV2Service

`func NewVerifyV2Service() *VerifyV2Service`

NewVerifyV2Service instantiates a new VerifyV2Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyV2ServiceWithDefaults

`func NewVerifyV2ServiceWithDefaults() *VerifyV2Service`

NewVerifyV2ServiceWithDefaults instantiates a new VerifyV2Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VerifyV2Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VerifyV2Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VerifyV2Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VerifyV2Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VerifyV2Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VerifyV2Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCodeLength

`func (o *VerifyV2Service) GetCodeLength() int32`

GetCodeLength returns the CodeLength field if non-nil, zero value otherwise.

### GetCodeLengthOk

`func (o *VerifyV2Service) GetCodeLengthOk() (*int32, bool)`

GetCodeLengthOk returns a tuple with the CodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLength

`func (o *VerifyV2Service) SetCodeLength(v int32)`

SetCodeLength sets CodeLength field to given value.

### HasCodeLength

`func (o *VerifyV2Service) HasCodeLength() bool`

HasCodeLength returns a boolean if a field has been set.

### SetCodeLengthNil

`func (o *VerifyV2Service) SetCodeLengthNil(b bool)`

 SetCodeLengthNil sets the value for CodeLength to be an explicit nil

### UnsetCodeLength
`func (o *VerifyV2Service) UnsetCodeLength()`

UnsetCodeLength ensures that no value is present for CodeLength, not even an explicit nil
### GetCustomCodeEnabled

`func (o *VerifyV2Service) GetCustomCodeEnabled() bool`

GetCustomCodeEnabled returns the CustomCodeEnabled field if non-nil, zero value otherwise.

### GetCustomCodeEnabledOk

`func (o *VerifyV2Service) GetCustomCodeEnabledOk() (*bool, bool)`

GetCustomCodeEnabledOk returns a tuple with the CustomCodeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCodeEnabled

`func (o *VerifyV2Service) SetCustomCodeEnabled(v bool)`

SetCustomCodeEnabled sets CustomCodeEnabled field to given value.

### HasCustomCodeEnabled

`func (o *VerifyV2Service) HasCustomCodeEnabled() bool`

HasCustomCodeEnabled returns a boolean if a field has been set.

### SetCustomCodeEnabledNil

`func (o *VerifyV2Service) SetCustomCodeEnabledNil(b bool)`

 SetCustomCodeEnabledNil sets the value for CustomCodeEnabled to be an explicit nil

### UnsetCustomCodeEnabled
`func (o *VerifyV2Service) UnsetCustomCodeEnabled()`

UnsetCustomCodeEnabled ensures that no value is present for CustomCodeEnabled, not even an explicit nil
### GetDateCreated

`func (o *VerifyV2Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VerifyV2Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VerifyV2Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VerifyV2Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VerifyV2Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VerifyV2Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VerifyV2Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VerifyV2Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VerifyV2Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VerifyV2Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VerifyV2Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VerifyV2Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDoNotShareWarningEnabled

`func (o *VerifyV2Service) GetDoNotShareWarningEnabled() bool`

GetDoNotShareWarningEnabled returns the DoNotShareWarningEnabled field if non-nil, zero value otherwise.

### GetDoNotShareWarningEnabledOk

`func (o *VerifyV2Service) GetDoNotShareWarningEnabledOk() (*bool, bool)`

GetDoNotShareWarningEnabledOk returns a tuple with the DoNotShareWarningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotShareWarningEnabled

`func (o *VerifyV2Service) SetDoNotShareWarningEnabled(v bool)`

SetDoNotShareWarningEnabled sets DoNotShareWarningEnabled field to given value.

### HasDoNotShareWarningEnabled

`func (o *VerifyV2Service) HasDoNotShareWarningEnabled() bool`

HasDoNotShareWarningEnabled returns a boolean if a field has been set.

### SetDoNotShareWarningEnabledNil

`func (o *VerifyV2Service) SetDoNotShareWarningEnabledNil(b bool)`

 SetDoNotShareWarningEnabledNil sets the value for DoNotShareWarningEnabled to be an explicit nil

### UnsetDoNotShareWarningEnabled
`func (o *VerifyV2Service) UnsetDoNotShareWarningEnabled()`

UnsetDoNotShareWarningEnabled ensures that no value is present for DoNotShareWarningEnabled, not even an explicit nil
### GetDtmfInputRequired

`func (o *VerifyV2Service) GetDtmfInputRequired() bool`

GetDtmfInputRequired returns the DtmfInputRequired field if non-nil, zero value otherwise.

### GetDtmfInputRequiredOk

`func (o *VerifyV2Service) GetDtmfInputRequiredOk() (*bool, bool)`

GetDtmfInputRequiredOk returns a tuple with the DtmfInputRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfInputRequired

`func (o *VerifyV2Service) SetDtmfInputRequired(v bool)`

SetDtmfInputRequired sets DtmfInputRequired field to given value.

### HasDtmfInputRequired

`func (o *VerifyV2Service) HasDtmfInputRequired() bool`

HasDtmfInputRequired returns a boolean if a field has been set.

### SetDtmfInputRequiredNil

`func (o *VerifyV2Service) SetDtmfInputRequiredNil(b bool)`

 SetDtmfInputRequiredNil sets the value for DtmfInputRequired to be an explicit nil

### UnsetDtmfInputRequired
`func (o *VerifyV2Service) UnsetDtmfInputRequired()`

UnsetDtmfInputRequired ensures that no value is present for DtmfInputRequired, not even an explicit nil
### GetFriendlyName

`func (o *VerifyV2Service) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VerifyV2Service) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VerifyV2Service) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VerifyV2Service) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VerifyV2Service) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VerifyV2Service) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *VerifyV2Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VerifyV2Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VerifyV2Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VerifyV2Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *VerifyV2Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *VerifyV2Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetLookupEnabled

`func (o *VerifyV2Service) GetLookupEnabled() bool`

GetLookupEnabled returns the LookupEnabled field if non-nil, zero value otherwise.

### GetLookupEnabledOk

`func (o *VerifyV2Service) GetLookupEnabledOk() (*bool, bool)`

GetLookupEnabledOk returns a tuple with the LookupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupEnabled

`func (o *VerifyV2Service) SetLookupEnabled(v bool)`

SetLookupEnabled sets LookupEnabled field to given value.

### HasLookupEnabled

`func (o *VerifyV2Service) HasLookupEnabled() bool`

HasLookupEnabled returns a boolean if a field has been set.

### SetLookupEnabledNil

`func (o *VerifyV2Service) SetLookupEnabledNil(b bool)`

 SetLookupEnabledNil sets the value for LookupEnabled to be an explicit nil

### UnsetLookupEnabled
`func (o *VerifyV2Service) UnsetLookupEnabled()`

UnsetLookupEnabled ensures that no value is present for LookupEnabled, not even an explicit nil
### GetPsd2Enabled

`func (o *VerifyV2Service) GetPsd2Enabled() bool`

GetPsd2Enabled returns the Psd2Enabled field if non-nil, zero value otherwise.

### GetPsd2EnabledOk

`func (o *VerifyV2Service) GetPsd2EnabledOk() (*bool, bool)`

GetPsd2EnabledOk returns a tuple with the Psd2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsd2Enabled

`func (o *VerifyV2Service) SetPsd2Enabled(v bool)`

SetPsd2Enabled sets Psd2Enabled field to given value.

### HasPsd2Enabled

`func (o *VerifyV2Service) HasPsd2Enabled() bool`

HasPsd2Enabled returns a boolean if a field has been set.

### SetPsd2EnabledNil

`func (o *VerifyV2Service) SetPsd2EnabledNil(b bool)`

 SetPsd2EnabledNil sets the value for Psd2Enabled to be an explicit nil

### UnsetPsd2Enabled
`func (o *VerifyV2Service) UnsetPsd2Enabled()`

UnsetPsd2Enabled ensures that no value is present for Psd2Enabled, not even an explicit nil
### GetPush

`func (o *VerifyV2Service) GetPush() map[string]interface{}`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *VerifyV2Service) GetPushOk() (*map[string]interface{}, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *VerifyV2Service) SetPush(v map[string]interface{})`

SetPush sets Push field to given value.

### HasPush

`func (o *VerifyV2Service) HasPush() bool`

HasPush returns a boolean if a field has been set.

### SetPushNil

`func (o *VerifyV2Service) SetPushNil(b bool)`

 SetPushNil sets the value for Push to be an explicit nil

### UnsetPush
`func (o *VerifyV2Service) UnsetPush()`

UnsetPush ensures that no value is present for Push, not even an explicit nil
### GetSid

`func (o *VerifyV2Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VerifyV2Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VerifyV2Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VerifyV2Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VerifyV2Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VerifyV2Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSkipSmsToLandlines

`func (o *VerifyV2Service) GetSkipSmsToLandlines() bool`

GetSkipSmsToLandlines returns the SkipSmsToLandlines field if non-nil, zero value otherwise.

### GetSkipSmsToLandlinesOk

`func (o *VerifyV2Service) GetSkipSmsToLandlinesOk() (*bool, bool)`

GetSkipSmsToLandlinesOk returns a tuple with the SkipSmsToLandlines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSmsToLandlines

`func (o *VerifyV2Service) SetSkipSmsToLandlines(v bool)`

SetSkipSmsToLandlines sets SkipSmsToLandlines field to given value.

### HasSkipSmsToLandlines

`func (o *VerifyV2Service) HasSkipSmsToLandlines() bool`

HasSkipSmsToLandlines returns a boolean if a field has been set.

### SetSkipSmsToLandlinesNil

`func (o *VerifyV2Service) SetSkipSmsToLandlinesNil(b bool)`

 SetSkipSmsToLandlinesNil sets the value for SkipSmsToLandlines to be an explicit nil

### UnsetSkipSmsToLandlines
`func (o *VerifyV2Service) UnsetSkipSmsToLandlines()`

UnsetSkipSmsToLandlines ensures that no value is present for SkipSmsToLandlines, not even an explicit nil
### GetTtsName

`func (o *VerifyV2Service) GetTtsName() string`

GetTtsName returns the TtsName field if non-nil, zero value otherwise.

### GetTtsNameOk

`func (o *VerifyV2Service) GetTtsNameOk() (*string, bool)`

GetTtsNameOk returns a tuple with the TtsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtsName

`func (o *VerifyV2Service) SetTtsName(v string)`

SetTtsName sets TtsName field to given value.

### HasTtsName

`func (o *VerifyV2Service) HasTtsName() bool`

HasTtsName returns a boolean if a field has been set.

### SetTtsNameNil

`func (o *VerifyV2Service) SetTtsNameNil(b bool)`

 SetTtsNameNil sets the value for TtsName to be an explicit nil

### UnsetTtsName
`func (o *VerifyV2Service) UnsetTtsName()`

UnsetTtsName ensures that no value is present for TtsName, not even an explicit nil
### GetUrl

`func (o *VerifyV2Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VerifyV2Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VerifyV2Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VerifyV2Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VerifyV2Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VerifyV2Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


