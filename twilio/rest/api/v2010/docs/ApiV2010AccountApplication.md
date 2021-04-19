# ApiV2010AccountApplication

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to start a new TwiML session | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**MessageStatusCallback** | Pointer to **NullableString** | The URL to send message status information to your application | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SmsFallbackMethod** | Pointer to **NullableString** | The HTTP method used with sms_fallback_url | [optional] 
**SmsFallbackUrl** | Pointer to **NullableString** | The URL that we call when an error occurs while retrieving or executing the TwiML | [optional] 
**SmsMethod** | Pointer to **NullableString** | The HTTP method to use with sms_url | [optional] 
**SmsStatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**SmsUrl** | Pointer to **NullableString** | The URL we call when the phone number receives an incoming SMS message | [optional] 
**StatusCallback** | Pointer to **NullableString** | The URL to send status information to your application | [optional] 
**StatusCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call status_callback | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**VoiceCallerIdLookup** | Pointer to **NullableBool** | Whether to lookup the caller&#39;s name | [optional] 
**VoiceFallbackMethod** | Pointer to **NullableString** | The HTTP method used with voice_fallback_url | [optional] 
**VoiceFallbackUrl** | Pointer to **NullableString** | The URL we call when a TwiML error occurs | [optional] 
**VoiceMethod** | Pointer to **NullableString** | The HTTP method used with the voice_url | [optional] 
**VoiceUrl** | Pointer to **NullableString** | The URL we call when the phone number receives a call | [optional] 

## Methods

### NewApiV2010AccountApplication

`func NewApiV2010AccountApplication() *ApiV2010AccountApplication`

NewApiV2010AccountApplication instantiates a new ApiV2010AccountApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountApplicationWithDefaults

`func NewApiV2010AccountApplicationWithDefaults() *ApiV2010AccountApplication`

NewApiV2010AccountApplicationWithDefaults instantiates a new ApiV2010AccountApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountApplication) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountApplication) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountApplication) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountApplication) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountApplication) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountApplication) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountApplication) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountApplication) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountApplication) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountApplication) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountApplication) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountApplication) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountApplication) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountApplication) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountApplication) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountApplication) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountApplication) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountApplication) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountApplication) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountApplication) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountApplication) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountApplication) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountApplication) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountApplication) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountApplication) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountApplication) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountApplication) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountApplication) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountApplication) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountApplication) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetMessageStatusCallback

`func (o *ApiV2010AccountApplication) GetMessageStatusCallback() string`

GetMessageStatusCallback returns the MessageStatusCallback field if non-nil, zero value otherwise.

### GetMessageStatusCallbackOk

`func (o *ApiV2010AccountApplication) GetMessageStatusCallbackOk() (*string, bool)`

GetMessageStatusCallbackOk returns a tuple with the MessageStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageStatusCallback

`func (o *ApiV2010AccountApplication) SetMessageStatusCallback(v string)`

SetMessageStatusCallback sets MessageStatusCallback field to given value.

### HasMessageStatusCallback

`func (o *ApiV2010AccountApplication) HasMessageStatusCallback() bool`

HasMessageStatusCallback returns a boolean if a field has been set.

### SetMessageStatusCallbackNil

`func (o *ApiV2010AccountApplication) SetMessageStatusCallbackNil(b bool)`

 SetMessageStatusCallbackNil sets the value for MessageStatusCallback to be an explicit nil

### UnsetMessageStatusCallback
`func (o *ApiV2010AccountApplication) UnsetMessageStatusCallback()`

UnsetMessageStatusCallback ensures that no value is present for MessageStatusCallback, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountApplication) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountApplication) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountApplication) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountApplication) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountApplication) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountApplication) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *ApiV2010AccountApplication) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *ApiV2010AccountApplication) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *ApiV2010AccountApplication) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *ApiV2010AccountApplication) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *ApiV2010AccountApplication) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *ApiV2010AccountApplication) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *ApiV2010AccountApplication) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *ApiV2010AccountApplication) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *ApiV2010AccountApplication) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *ApiV2010AccountApplication) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *ApiV2010AccountApplication) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *ApiV2010AccountApplication) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *ApiV2010AccountApplication) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *ApiV2010AccountApplication) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *ApiV2010AccountApplication) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *ApiV2010AccountApplication) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *ApiV2010AccountApplication) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *ApiV2010AccountApplication) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsStatusCallback

`func (o *ApiV2010AccountApplication) GetSmsStatusCallback() string`

GetSmsStatusCallback returns the SmsStatusCallback field if non-nil, zero value otherwise.

### GetSmsStatusCallbackOk

`func (o *ApiV2010AccountApplication) GetSmsStatusCallbackOk() (*string, bool)`

GetSmsStatusCallbackOk returns a tuple with the SmsStatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsStatusCallback

`func (o *ApiV2010AccountApplication) SetSmsStatusCallback(v string)`

SetSmsStatusCallback sets SmsStatusCallback field to given value.

### HasSmsStatusCallback

`func (o *ApiV2010AccountApplication) HasSmsStatusCallback() bool`

HasSmsStatusCallback returns a boolean if a field has been set.

### SetSmsStatusCallbackNil

`func (o *ApiV2010AccountApplication) SetSmsStatusCallbackNil(b bool)`

 SetSmsStatusCallbackNil sets the value for SmsStatusCallback to be an explicit nil

### UnsetSmsStatusCallback
`func (o *ApiV2010AccountApplication) UnsetSmsStatusCallback()`

UnsetSmsStatusCallback ensures that no value is present for SmsStatusCallback, not even an explicit nil
### GetSmsUrl

`func (o *ApiV2010AccountApplication) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *ApiV2010AccountApplication) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *ApiV2010AccountApplication) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *ApiV2010AccountApplication) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *ApiV2010AccountApplication) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *ApiV2010AccountApplication) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetStatusCallback

`func (o *ApiV2010AccountApplication) GetStatusCallback() string`

GetStatusCallback returns the StatusCallback field if non-nil, zero value otherwise.

### GetStatusCallbackOk

`func (o *ApiV2010AccountApplication) GetStatusCallbackOk() (*string, bool)`

GetStatusCallbackOk returns a tuple with the StatusCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallback

`func (o *ApiV2010AccountApplication) SetStatusCallback(v string)`

SetStatusCallback sets StatusCallback field to given value.

### HasStatusCallback

`func (o *ApiV2010AccountApplication) HasStatusCallback() bool`

HasStatusCallback returns a boolean if a field has been set.

### SetStatusCallbackNil

`func (o *ApiV2010AccountApplication) SetStatusCallbackNil(b bool)`

 SetStatusCallbackNil sets the value for StatusCallback to be an explicit nil

### UnsetStatusCallback
`func (o *ApiV2010AccountApplication) UnsetStatusCallback()`

UnsetStatusCallback ensures that no value is present for StatusCallback, not even an explicit nil
### GetStatusCallbackMethod

`func (o *ApiV2010AccountApplication) GetStatusCallbackMethod() string`

GetStatusCallbackMethod returns the StatusCallbackMethod field if non-nil, zero value otherwise.

### GetStatusCallbackMethodOk

`func (o *ApiV2010AccountApplication) GetStatusCallbackMethodOk() (*string, bool)`

GetStatusCallbackMethodOk returns a tuple with the StatusCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCallbackMethod

`func (o *ApiV2010AccountApplication) SetStatusCallbackMethod(v string)`

SetStatusCallbackMethod sets StatusCallbackMethod field to given value.

### HasStatusCallbackMethod

`func (o *ApiV2010AccountApplication) HasStatusCallbackMethod() bool`

HasStatusCallbackMethod returns a boolean if a field has been set.

### SetStatusCallbackMethodNil

`func (o *ApiV2010AccountApplication) SetStatusCallbackMethodNil(b bool)`

 SetStatusCallbackMethodNil sets the value for StatusCallbackMethod to be an explicit nil

### UnsetStatusCallbackMethod
`func (o *ApiV2010AccountApplication) UnsetStatusCallbackMethod()`

UnsetStatusCallbackMethod ensures that no value is present for StatusCallbackMethod, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountApplication) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountApplication) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountApplication) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountApplication) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountApplication) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountApplication) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetVoiceCallerIdLookup

`func (o *ApiV2010AccountApplication) GetVoiceCallerIdLookup() bool`

GetVoiceCallerIdLookup returns the VoiceCallerIdLookup field if non-nil, zero value otherwise.

### GetVoiceCallerIdLookupOk

`func (o *ApiV2010AccountApplication) GetVoiceCallerIdLookupOk() (*bool, bool)`

GetVoiceCallerIdLookupOk returns a tuple with the VoiceCallerIdLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceCallerIdLookup

`func (o *ApiV2010AccountApplication) SetVoiceCallerIdLookup(v bool)`

SetVoiceCallerIdLookup sets VoiceCallerIdLookup field to given value.

### HasVoiceCallerIdLookup

`func (o *ApiV2010AccountApplication) HasVoiceCallerIdLookup() bool`

HasVoiceCallerIdLookup returns a boolean if a field has been set.

### SetVoiceCallerIdLookupNil

`func (o *ApiV2010AccountApplication) SetVoiceCallerIdLookupNil(b bool)`

 SetVoiceCallerIdLookupNil sets the value for VoiceCallerIdLookup to be an explicit nil

### UnsetVoiceCallerIdLookup
`func (o *ApiV2010AccountApplication) UnsetVoiceCallerIdLookup()`

UnsetVoiceCallerIdLookup ensures that no value is present for VoiceCallerIdLookup, not even an explicit nil
### GetVoiceFallbackMethod

`func (o *ApiV2010AccountApplication) GetVoiceFallbackMethod() string`

GetVoiceFallbackMethod returns the VoiceFallbackMethod field if non-nil, zero value otherwise.

### GetVoiceFallbackMethodOk

`func (o *ApiV2010AccountApplication) GetVoiceFallbackMethodOk() (*string, bool)`

GetVoiceFallbackMethodOk returns a tuple with the VoiceFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackMethod

`func (o *ApiV2010AccountApplication) SetVoiceFallbackMethod(v string)`

SetVoiceFallbackMethod sets VoiceFallbackMethod field to given value.

### HasVoiceFallbackMethod

`func (o *ApiV2010AccountApplication) HasVoiceFallbackMethod() bool`

HasVoiceFallbackMethod returns a boolean if a field has been set.

### SetVoiceFallbackMethodNil

`func (o *ApiV2010AccountApplication) SetVoiceFallbackMethodNil(b bool)`

 SetVoiceFallbackMethodNil sets the value for VoiceFallbackMethod to be an explicit nil

### UnsetVoiceFallbackMethod
`func (o *ApiV2010AccountApplication) UnsetVoiceFallbackMethod()`

UnsetVoiceFallbackMethod ensures that no value is present for VoiceFallbackMethod, not even an explicit nil
### GetVoiceFallbackUrl

`func (o *ApiV2010AccountApplication) GetVoiceFallbackUrl() string`

GetVoiceFallbackUrl returns the VoiceFallbackUrl field if non-nil, zero value otherwise.

### GetVoiceFallbackUrlOk

`func (o *ApiV2010AccountApplication) GetVoiceFallbackUrlOk() (*string, bool)`

GetVoiceFallbackUrlOk returns a tuple with the VoiceFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceFallbackUrl

`func (o *ApiV2010AccountApplication) SetVoiceFallbackUrl(v string)`

SetVoiceFallbackUrl sets VoiceFallbackUrl field to given value.

### HasVoiceFallbackUrl

`func (o *ApiV2010AccountApplication) HasVoiceFallbackUrl() bool`

HasVoiceFallbackUrl returns a boolean if a field has been set.

### SetVoiceFallbackUrlNil

`func (o *ApiV2010AccountApplication) SetVoiceFallbackUrlNil(b bool)`

 SetVoiceFallbackUrlNil sets the value for VoiceFallbackUrl to be an explicit nil

### UnsetVoiceFallbackUrl
`func (o *ApiV2010AccountApplication) UnsetVoiceFallbackUrl()`

UnsetVoiceFallbackUrl ensures that no value is present for VoiceFallbackUrl, not even an explicit nil
### GetVoiceMethod

`func (o *ApiV2010AccountApplication) GetVoiceMethod() string`

GetVoiceMethod returns the VoiceMethod field if non-nil, zero value otherwise.

### GetVoiceMethodOk

`func (o *ApiV2010AccountApplication) GetVoiceMethodOk() (*string, bool)`

GetVoiceMethodOk returns a tuple with the VoiceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceMethod

`func (o *ApiV2010AccountApplication) SetVoiceMethod(v string)`

SetVoiceMethod sets VoiceMethod field to given value.

### HasVoiceMethod

`func (o *ApiV2010AccountApplication) HasVoiceMethod() bool`

HasVoiceMethod returns a boolean if a field has been set.

### SetVoiceMethodNil

`func (o *ApiV2010AccountApplication) SetVoiceMethodNil(b bool)`

 SetVoiceMethodNil sets the value for VoiceMethod to be an explicit nil

### UnsetVoiceMethod
`func (o *ApiV2010AccountApplication) UnsetVoiceMethod()`

UnsetVoiceMethod ensures that no value is present for VoiceMethod, not even an explicit nil
### GetVoiceUrl

`func (o *ApiV2010AccountApplication) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *ApiV2010AccountApplication) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *ApiV2010AccountApplication) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *ApiV2010AccountApplication) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *ApiV2010AccountApplication) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *ApiV2010AccountApplication) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


