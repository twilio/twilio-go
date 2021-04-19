# ApiV2010AccountToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**IceServers** | Pointer to [**[]ApiV2010AccountTokenIceServers**](ApiV2010AccountTokenIceServers.md) | An array representing the ephemeral credentials | [optional] 
**Password** | Pointer to **NullableString** | The temporary password used for authenticating | [optional] 
**Ttl** | Pointer to **NullableString** | The duration in seconds the credentials are valid | [optional] 
**Username** | Pointer to **NullableString** | The temporary username that uniquely identifies a Token | [optional] 

## Methods

### NewApiV2010AccountToken

`func NewApiV2010AccountToken() *ApiV2010AccountToken`

NewApiV2010AccountToken instantiates a new ApiV2010AccountToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountTokenWithDefaults

`func NewApiV2010AccountTokenWithDefaults() *ApiV2010AccountToken`

NewApiV2010AccountTokenWithDefaults instantiates a new ApiV2010AccountToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountToken) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountToken) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountToken) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountToken) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountToken) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountToken) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountToken) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountToken) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountToken) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountToken) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountToken) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountToken) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountToken) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountToken) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountToken) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountToken) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountToken) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountToken) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIceServers

`func (o *ApiV2010AccountToken) GetIceServers() []ApiV2010AccountTokenIceServers`

GetIceServers returns the IceServers field if non-nil, zero value otherwise.

### GetIceServersOk

`func (o *ApiV2010AccountToken) GetIceServersOk() (*[]ApiV2010AccountTokenIceServers, bool)`

GetIceServersOk returns a tuple with the IceServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIceServers

`func (o *ApiV2010AccountToken) SetIceServers(v []ApiV2010AccountTokenIceServers)`

SetIceServers sets IceServers field to given value.

### HasIceServers

`func (o *ApiV2010AccountToken) HasIceServers() bool`

HasIceServers returns a boolean if a field has been set.

### SetIceServersNil

`func (o *ApiV2010AccountToken) SetIceServersNil(b bool)`

 SetIceServersNil sets the value for IceServers to be an explicit nil

### UnsetIceServers
`func (o *ApiV2010AccountToken) UnsetIceServers()`

UnsetIceServers ensures that no value is present for IceServers, not even an explicit nil
### GetPassword

`func (o *ApiV2010AccountToken) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiV2010AccountToken) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiV2010AccountToken) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiV2010AccountToken) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ApiV2010AccountToken) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ApiV2010AccountToken) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetTtl

`func (o *ApiV2010AccountToken) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ApiV2010AccountToken) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ApiV2010AccountToken) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ApiV2010AccountToken) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *ApiV2010AccountToken) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *ApiV2010AccountToken) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetUsername

`func (o *ApiV2010AccountToken) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiV2010AccountToken) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiV2010AccountToken) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiV2010AccountToken) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ApiV2010AccountToken) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ApiV2010AccountToken) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


