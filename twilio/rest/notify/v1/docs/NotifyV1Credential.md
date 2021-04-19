# NotifyV1Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Sandbox** | Pointer to **NullableString** | [APN only] Whether to send the credential to sandbox APNs | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Type** | Pointer to **NullableString** | The Credential type, one of &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60; | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Credential resource | [optional] 

## Methods

### NewNotifyV1Credential

`func NewNotifyV1Credential() *NotifyV1Credential`

NewNotifyV1Credential instantiates a new NotifyV1Credential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyV1CredentialWithDefaults

`func NewNotifyV1CredentialWithDefaults() *NotifyV1Credential`

NewNotifyV1CredentialWithDefaults instantiates a new NotifyV1Credential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *NotifyV1Credential) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *NotifyV1Credential) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *NotifyV1Credential) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *NotifyV1Credential) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *NotifyV1Credential) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *NotifyV1Credential) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *NotifyV1Credential) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NotifyV1Credential) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NotifyV1Credential) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NotifyV1Credential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *NotifyV1Credential) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *NotifyV1Credential) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *NotifyV1Credential) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *NotifyV1Credential) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *NotifyV1Credential) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *NotifyV1Credential) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *NotifyV1Credential) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *NotifyV1Credential) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *NotifyV1Credential) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *NotifyV1Credential) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *NotifyV1Credential) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *NotifyV1Credential) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *NotifyV1Credential) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *NotifyV1Credential) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetSandbox

`func (o *NotifyV1Credential) GetSandbox() string`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NotifyV1Credential) GetSandboxOk() (*string, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NotifyV1Credential) SetSandbox(v string)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *NotifyV1Credential) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandboxNil

`func (o *NotifyV1Credential) SetSandboxNil(b bool)`

 SetSandboxNil sets the value for Sandbox to be an explicit nil

### UnsetSandbox
`func (o *NotifyV1Credential) UnsetSandbox()`

UnsetSandbox ensures that no value is present for Sandbox, not even an explicit nil
### GetSid

`func (o *NotifyV1Credential) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotifyV1Credential) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotifyV1Credential) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *NotifyV1Credential) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *NotifyV1Credential) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *NotifyV1Credential) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *NotifyV1Credential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotifyV1Credential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotifyV1Credential) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotifyV1Credential) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NotifyV1Credential) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotifyV1Credential) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *NotifyV1Credential) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotifyV1Credential) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotifyV1Credential) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotifyV1Credential) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *NotifyV1Credential) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *NotifyV1Credential) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


