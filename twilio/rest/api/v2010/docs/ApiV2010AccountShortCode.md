# ApiV2010AccountShortCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created this resource | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to start a new TwiML session | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that this resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | A string that you assigned to describe this resource | [optional] 
**ShortCode** | Pointer to **NullableString** | The short code. e.g., 894546. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies this resource | [optional] 
**SmsFallbackMethod** | Pointer to **NullableString** | HTTP method we use to call the sms_fallback_url | [optional] 
**SmsFallbackUrl** | Pointer to **NullableString** | URL Twilio will request if an error occurs in executing TwiML | [optional] 
**SmsMethod** | Pointer to **NullableString** | HTTP method to use when requesting the sms url | [optional] 
**SmsUrl** | Pointer to **NullableString** | URL we call when receiving an incoming SMS message to this short code | [optional] 
**Uri** | Pointer to **NullableString** | The URI of this resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountShortCode

`func NewApiV2010AccountShortCode() *ApiV2010AccountShortCode`

NewApiV2010AccountShortCode instantiates a new ApiV2010AccountShortCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountShortCodeWithDefaults

`func NewApiV2010AccountShortCodeWithDefaults() *ApiV2010AccountShortCode`

NewApiV2010AccountShortCodeWithDefaults instantiates a new ApiV2010AccountShortCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountShortCode) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountShortCode) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountShortCode) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountShortCode) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountShortCode) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountShortCode) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountShortCode) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountShortCode) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountShortCode) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountShortCode) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountShortCode) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountShortCode) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountShortCode) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountShortCode) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountShortCode) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountShortCode) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountShortCode) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountShortCode) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountShortCode) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountShortCode) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountShortCode) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountShortCode) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountShortCode) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountShortCode) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountShortCode) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountShortCode) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountShortCode) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountShortCode) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountShortCode) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountShortCode) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetShortCode

`func (o *ApiV2010AccountShortCode) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *ApiV2010AccountShortCode) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *ApiV2010AccountShortCode) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *ApiV2010AccountShortCode) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### SetShortCodeNil

`func (o *ApiV2010AccountShortCode) SetShortCodeNil(b bool)`

 SetShortCodeNil sets the value for ShortCode to be an explicit nil

### UnsetShortCode
`func (o *ApiV2010AccountShortCode) UnsetShortCode()`

UnsetShortCode ensures that no value is present for ShortCode, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountShortCode) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountShortCode) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountShortCode) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountShortCode) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountShortCode) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountShortCode) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSmsFallbackMethod

`func (o *ApiV2010AccountShortCode) GetSmsFallbackMethod() string`

GetSmsFallbackMethod returns the SmsFallbackMethod field if non-nil, zero value otherwise.

### GetSmsFallbackMethodOk

`func (o *ApiV2010AccountShortCode) GetSmsFallbackMethodOk() (*string, bool)`

GetSmsFallbackMethodOk returns a tuple with the SmsFallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackMethod

`func (o *ApiV2010AccountShortCode) SetSmsFallbackMethod(v string)`

SetSmsFallbackMethod sets SmsFallbackMethod field to given value.

### HasSmsFallbackMethod

`func (o *ApiV2010AccountShortCode) HasSmsFallbackMethod() bool`

HasSmsFallbackMethod returns a boolean if a field has been set.

### SetSmsFallbackMethodNil

`func (o *ApiV2010AccountShortCode) SetSmsFallbackMethodNil(b bool)`

 SetSmsFallbackMethodNil sets the value for SmsFallbackMethod to be an explicit nil

### UnsetSmsFallbackMethod
`func (o *ApiV2010AccountShortCode) UnsetSmsFallbackMethod()`

UnsetSmsFallbackMethod ensures that no value is present for SmsFallbackMethod, not even an explicit nil
### GetSmsFallbackUrl

`func (o *ApiV2010AccountShortCode) GetSmsFallbackUrl() string`

GetSmsFallbackUrl returns the SmsFallbackUrl field if non-nil, zero value otherwise.

### GetSmsFallbackUrlOk

`func (o *ApiV2010AccountShortCode) GetSmsFallbackUrlOk() (*string, bool)`

GetSmsFallbackUrlOk returns a tuple with the SmsFallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsFallbackUrl

`func (o *ApiV2010AccountShortCode) SetSmsFallbackUrl(v string)`

SetSmsFallbackUrl sets SmsFallbackUrl field to given value.

### HasSmsFallbackUrl

`func (o *ApiV2010AccountShortCode) HasSmsFallbackUrl() bool`

HasSmsFallbackUrl returns a boolean if a field has been set.

### SetSmsFallbackUrlNil

`func (o *ApiV2010AccountShortCode) SetSmsFallbackUrlNil(b bool)`

 SetSmsFallbackUrlNil sets the value for SmsFallbackUrl to be an explicit nil

### UnsetSmsFallbackUrl
`func (o *ApiV2010AccountShortCode) UnsetSmsFallbackUrl()`

UnsetSmsFallbackUrl ensures that no value is present for SmsFallbackUrl, not even an explicit nil
### GetSmsMethod

`func (o *ApiV2010AccountShortCode) GetSmsMethod() string`

GetSmsMethod returns the SmsMethod field if non-nil, zero value otherwise.

### GetSmsMethodOk

`func (o *ApiV2010AccountShortCode) GetSmsMethodOk() (*string, bool)`

GetSmsMethodOk returns a tuple with the SmsMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMethod

`func (o *ApiV2010AccountShortCode) SetSmsMethod(v string)`

SetSmsMethod sets SmsMethod field to given value.

### HasSmsMethod

`func (o *ApiV2010AccountShortCode) HasSmsMethod() bool`

HasSmsMethod returns a boolean if a field has been set.

### SetSmsMethodNil

`func (o *ApiV2010AccountShortCode) SetSmsMethodNil(b bool)`

 SetSmsMethodNil sets the value for SmsMethod to be an explicit nil

### UnsetSmsMethod
`func (o *ApiV2010AccountShortCode) UnsetSmsMethod()`

UnsetSmsMethod ensures that no value is present for SmsMethod, not even an explicit nil
### GetSmsUrl

`func (o *ApiV2010AccountShortCode) GetSmsUrl() string`

GetSmsUrl returns the SmsUrl field if non-nil, zero value otherwise.

### GetSmsUrlOk

`func (o *ApiV2010AccountShortCode) GetSmsUrlOk() (*string, bool)`

GetSmsUrlOk returns a tuple with the SmsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsUrl

`func (o *ApiV2010AccountShortCode) SetSmsUrl(v string)`

SetSmsUrl sets SmsUrl field to given value.

### HasSmsUrl

`func (o *ApiV2010AccountShortCode) HasSmsUrl() bool`

HasSmsUrl returns a boolean if a field has been set.

### SetSmsUrlNil

`func (o *ApiV2010AccountShortCode) SetSmsUrlNil(b bool)`

 SetSmsUrlNil sets the value for SmsUrl to be an explicit nil

### UnsetSmsUrl
`func (o *ApiV2010AccountShortCode) UnsetSmsUrl()`

UnsetSmsUrl ensures that no value is present for SmsUrl, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountShortCode) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountShortCode) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountShortCode) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountShortCode) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountShortCode) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountShortCode) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


