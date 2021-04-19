# AccountsV1AuthTokenPromotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that the secondary Auth Token was created for | [optional] 
**AuthToken** | Pointer to **NullableString** | The promoted Auth Token | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in UTC when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 formatted date and time in UTC when the resource was last updated | [optional] 
**Url** | Pointer to **NullableString** | The URI for this resource, relative to &#x60;https://accounts.twilio.com&#x60; | [optional] 

## Methods

### NewAccountsV1AuthTokenPromotion

`func NewAccountsV1AuthTokenPromotion() *AccountsV1AuthTokenPromotion`

NewAccountsV1AuthTokenPromotion instantiates a new AccountsV1AuthTokenPromotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsV1AuthTokenPromotionWithDefaults

`func NewAccountsV1AuthTokenPromotionWithDefaults() *AccountsV1AuthTokenPromotion`

NewAccountsV1AuthTokenPromotionWithDefaults instantiates a new AccountsV1AuthTokenPromotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AccountsV1AuthTokenPromotion) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AccountsV1AuthTokenPromotion) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AccountsV1AuthTokenPromotion) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AccountsV1AuthTokenPromotion) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AccountsV1AuthTokenPromotion) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AccountsV1AuthTokenPromotion) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAuthToken

`func (o *AccountsV1AuthTokenPromotion) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *AccountsV1AuthTokenPromotion) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *AccountsV1AuthTokenPromotion) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *AccountsV1AuthTokenPromotion) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *AccountsV1AuthTokenPromotion) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *AccountsV1AuthTokenPromotion) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetDateCreated

`func (o *AccountsV1AuthTokenPromotion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AccountsV1AuthTokenPromotion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AccountsV1AuthTokenPromotion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AccountsV1AuthTokenPromotion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AccountsV1AuthTokenPromotion) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AccountsV1AuthTokenPromotion) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AccountsV1AuthTokenPromotion) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AccountsV1AuthTokenPromotion) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AccountsV1AuthTokenPromotion) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AccountsV1AuthTokenPromotion) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AccountsV1AuthTokenPromotion) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AccountsV1AuthTokenPromotion) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetUrl

`func (o *AccountsV1AuthTokenPromotion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountsV1AuthTokenPromotion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountsV1AuthTokenPromotion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccountsV1AuthTokenPromotion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AccountsV1AuthTokenPromotion) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AccountsV1AuthTokenPromotion) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


