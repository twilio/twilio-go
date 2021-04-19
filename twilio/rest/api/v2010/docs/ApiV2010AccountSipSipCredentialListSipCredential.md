# ApiV2010AccountSipSipCredentialListSipCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique id of the Account that is responsible for this resource. | [optional] 
**CredentialListSid** | Pointer to **NullableString** | The unique id that identifies the credential list that includes this credential | [optional] 
**DateCreated** | Pointer to **NullableString** | The date that this resource was created, given as GMT in RFC 2822 format. | [optional] 
**DateUpdated** | Pointer to **NullableString** | The date that this resource was last updated, given as GMT in RFC 2822 format. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Uri** | Pointer to **NullableString** | The URI for this resource, relative to https://api.twilio.com | [optional] 
**Username** | Pointer to **NullableString** | The username for this credential. | [optional] 

## Methods

### NewApiV2010AccountSipSipCredentialListSipCredential

`func NewApiV2010AccountSipSipCredentialListSipCredential() *ApiV2010AccountSipSipCredentialListSipCredential`

NewApiV2010AccountSipSipCredentialListSipCredential instantiates a new ApiV2010AccountSipSipCredentialListSipCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountSipSipCredentialListSipCredentialWithDefaults

`func NewApiV2010AccountSipSipCredentialListSipCredentialWithDefaults() *ApiV2010AccountSipSipCredentialListSipCredential`

NewApiV2010AccountSipSipCredentialListSipCredentialWithDefaults instantiates a new ApiV2010AccountSipSipCredentialListSipCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCredentialListSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetCredentialListSid() string`

GetCredentialListSid returns the CredentialListSid field if non-nil, zero value otherwise.

### GetCredentialListSidOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetCredentialListSidOk() (*string, bool)`

GetCredentialListSidOk returns a tuple with the CredentialListSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialListSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetCredentialListSid(v string)`

SetCredentialListSid sets CredentialListSid field to given value.

### HasCredentialListSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasCredentialListSid() bool`

HasCredentialListSid returns a boolean if a field has been set.

### SetCredentialListSidNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetCredentialListSidNil(b bool)`

 SetCredentialListSidNil sets the value for CredentialListSid to be an explicit nil

### UnsetCredentialListSid
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetCredentialListSid()`

UnsetCredentialListSid ensures that no value is present for CredentialListSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetUsername

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ApiV2010AccountSipSipCredentialListSipCredential) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ApiV2010AccountSipSipCredentialListSipCredential) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


