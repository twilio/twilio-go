# TrunkingV1TrunkCredentialList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TrunkSid** | Pointer to **NullableString** | The SID of the Trunk the credential list in associated with | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewTrunkingV1TrunkCredentialList

`func NewTrunkingV1TrunkCredentialList() *TrunkingV1TrunkCredentialList`

NewTrunkingV1TrunkCredentialList instantiates a new TrunkingV1TrunkCredentialList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkingV1TrunkCredentialListWithDefaults

`func NewTrunkingV1TrunkCredentialListWithDefaults() *TrunkingV1TrunkCredentialList`

NewTrunkingV1TrunkCredentialListWithDefaults instantiates a new TrunkingV1TrunkCredentialList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TrunkingV1TrunkCredentialList) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TrunkingV1TrunkCredentialList) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TrunkingV1TrunkCredentialList) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TrunkingV1TrunkCredentialList) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TrunkingV1TrunkCredentialList) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TrunkingV1TrunkCredentialList) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TrunkingV1TrunkCredentialList) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TrunkingV1TrunkCredentialList) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TrunkingV1TrunkCredentialList) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TrunkingV1TrunkCredentialList) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TrunkingV1TrunkCredentialList) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TrunkingV1TrunkCredentialList) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TrunkingV1TrunkCredentialList) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TrunkingV1TrunkCredentialList) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TrunkingV1TrunkCredentialList) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TrunkingV1TrunkCredentialList) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TrunkingV1TrunkCredentialList) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TrunkingV1TrunkCredentialList) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *TrunkingV1TrunkCredentialList) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TrunkingV1TrunkCredentialList) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TrunkingV1TrunkCredentialList) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TrunkingV1TrunkCredentialList) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TrunkingV1TrunkCredentialList) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TrunkingV1TrunkCredentialList) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetSid

`func (o *TrunkingV1TrunkCredentialList) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TrunkingV1TrunkCredentialList) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TrunkingV1TrunkCredentialList) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TrunkingV1TrunkCredentialList) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TrunkingV1TrunkCredentialList) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TrunkingV1TrunkCredentialList) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTrunkSid

`func (o *TrunkingV1TrunkCredentialList) GetTrunkSid() string`

GetTrunkSid returns the TrunkSid field if non-nil, zero value otherwise.

### GetTrunkSidOk

`func (o *TrunkingV1TrunkCredentialList) GetTrunkSidOk() (*string, bool)`

GetTrunkSidOk returns a tuple with the TrunkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkSid

`func (o *TrunkingV1TrunkCredentialList) SetTrunkSid(v string)`

SetTrunkSid sets TrunkSid field to given value.

### HasTrunkSid

`func (o *TrunkingV1TrunkCredentialList) HasTrunkSid() bool`

HasTrunkSid returns a boolean if a field has been set.

### SetTrunkSidNil

`func (o *TrunkingV1TrunkCredentialList) SetTrunkSidNil(b bool)`

 SetTrunkSidNil sets the value for TrunkSid to be an explicit nil

### UnsetTrunkSid
`func (o *TrunkingV1TrunkCredentialList) UnsetTrunkSid()`

UnsetTrunkSid ensures that no value is present for TrunkSid, not even an explicit nil
### GetUrl

`func (o *TrunkingV1TrunkCredentialList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TrunkingV1TrunkCredentialList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TrunkingV1TrunkCredentialList) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TrunkingV1TrunkCredentialList) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TrunkingV1TrunkCredentialList) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TrunkingV1TrunkCredentialList) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


