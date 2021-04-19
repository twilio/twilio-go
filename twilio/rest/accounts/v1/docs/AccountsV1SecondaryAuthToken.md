# AccountsV1SecondaryAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that the secondary Auth Token was created for | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in UTC when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in UTC when the resource was last updated | [optional] 
**SecondaryAuthToken** | Pointer to **NullableString** | The generated secondary Auth Token | [optional] 
**Url** | Pointer to **NullableString** | The URI for this resource, relative to &#x60;https://accounts.twilio.com&#x60; | [optional] 

## Methods

### NewAccountsV1SecondaryAuthToken

`func NewAccountsV1SecondaryAuthToken() *AccountsV1SecondaryAuthToken`

NewAccountsV1SecondaryAuthToken instantiates a new AccountsV1SecondaryAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsV1SecondaryAuthTokenWithDefaults

`func NewAccountsV1SecondaryAuthTokenWithDefaults() *AccountsV1SecondaryAuthToken`

NewAccountsV1SecondaryAuthTokenWithDefaults instantiates a new AccountsV1SecondaryAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AccountsV1SecondaryAuthToken) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AccountsV1SecondaryAuthToken) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AccountsV1SecondaryAuthToken) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AccountsV1SecondaryAuthToken) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AccountsV1SecondaryAuthToken) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AccountsV1SecondaryAuthToken) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *AccountsV1SecondaryAuthToken) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AccountsV1SecondaryAuthToken) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AccountsV1SecondaryAuthToken) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AccountsV1SecondaryAuthToken) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AccountsV1SecondaryAuthToken) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AccountsV1SecondaryAuthToken) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AccountsV1SecondaryAuthToken) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AccountsV1SecondaryAuthToken) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AccountsV1SecondaryAuthToken) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AccountsV1SecondaryAuthToken) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AccountsV1SecondaryAuthToken) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AccountsV1SecondaryAuthToken) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetSecondaryAuthToken

`func (o *AccountsV1SecondaryAuthToken) GetSecondaryAuthToken() string`

GetSecondaryAuthToken returns the SecondaryAuthToken field if non-nil, zero value otherwise.

### GetSecondaryAuthTokenOk

`func (o *AccountsV1SecondaryAuthToken) GetSecondaryAuthTokenOk() (*string, bool)`

GetSecondaryAuthTokenOk returns a tuple with the SecondaryAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAuthToken

`func (o *AccountsV1SecondaryAuthToken) SetSecondaryAuthToken(v string)`

SetSecondaryAuthToken sets SecondaryAuthToken field to given value.

### HasSecondaryAuthToken

`func (o *AccountsV1SecondaryAuthToken) HasSecondaryAuthToken() bool`

HasSecondaryAuthToken returns a boolean if a field has been set.

### SetSecondaryAuthTokenNil

`func (o *AccountsV1SecondaryAuthToken) SetSecondaryAuthTokenNil(b bool)`

 SetSecondaryAuthTokenNil sets the value for SecondaryAuthToken to be an explicit nil

### UnsetSecondaryAuthToken
`func (o *AccountsV1SecondaryAuthToken) UnsetSecondaryAuthToken()`

UnsetSecondaryAuthToken ensures that no value is present for SecondaryAuthToken, not even an explicit nil
### GetUrl

`func (o *AccountsV1SecondaryAuthToken) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountsV1SecondaryAuthToken) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountsV1SecondaryAuthToken) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccountsV1SecondaryAuthToken) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AccountsV1SecondaryAuthToken) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AccountsV1SecondaryAuthToken) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


