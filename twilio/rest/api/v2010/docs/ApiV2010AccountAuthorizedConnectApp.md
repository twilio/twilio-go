# ApiV2010AccountAuthorizedConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ConnectAppCompanyName** | Pointer to **NullableString** | The company name set for the Connect App | [optional] 
**ConnectAppDescription** | Pointer to **NullableString** | A detailed description of the app | [optional] 
**ConnectAppFriendlyName** | Pointer to **NullableString** | The name of the Connect App | [optional] 
**ConnectAppHomepageUrl** | Pointer to **NullableString** | The public URL for the Connect App | [optional] 
**ConnectAppSid** | Pointer to **NullableString** | The SID that we assigned to the Connect App | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**Permissions** | Pointer to **[]string** | Permissions authorized to the app | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountAuthorizedConnectApp

`func NewApiV2010AccountAuthorizedConnectApp() *ApiV2010AccountAuthorizedConnectApp`

NewApiV2010AccountAuthorizedConnectApp instantiates a new ApiV2010AccountAuthorizedConnectApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountAuthorizedConnectAppWithDefaults

`func NewApiV2010AccountAuthorizedConnectAppWithDefaults() *ApiV2010AccountAuthorizedConnectApp`

NewApiV2010AccountAuthorizedConnectAppWithDefaults instantiates a new ApiV2010AccountAuthorizedConnectApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountAuthorizedConnectApp) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountAuthorizedConnectApp) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountAuthorizedConnectApp) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetConnectAppCompanyName

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppCompanyName() string`

GetConnectAppCompanyName returns the ConnectAppCompanyName field if non-nil, zero value otherwise.

### GetConnectAppCompanyNameOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppCompanyNameOk() (*string, bool)`

GetConnectAppCompanyNameOk returns a tuple with the ConnectAppCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAppCompanyName

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppCompanyName(v string)`

SetConnectAppCompanyName sets ConnectAppCompanyName field to given value.

### HasConnectAppCompanyName

`func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppCompanyName() bool`

HasConnectAppCompanyName returns a boolean if a field has been set.

### SetConnectAppCompanyNameNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppCompanyNameNil(b bool)`

 SetConnectAppCompanyNameNil sets the value for ConnectAppCompanyName to be an explicit nil

### UnsetConnectAppCompanyName
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppCompanyName()`

UnsetConnectAppCompanyName ensures that no value is present for ConnectAppCompanyName, not even an explicit nil
### GetConnectAppDescription

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppDescription() string`

GetConnectAppDescription returns the ConnectAppDescription field if non-nil, zero value otherwise.

### GetConnectAppDescriptionOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppDescriptionOk() (*string, bool)`

GetConnectAppDescriptionOk returns a tuple with the ConnectAppDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAppDescription

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppDescription(v string)`

SetConnectAppDescription sets ConnectAppDescription field to given value.

### HasConnectAppDescription

`func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppDescription() bool`

HasConnectAppDescription returns a boolean if a field has been set.

### SetConnectAppDescriptionNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppDescriptionNil(b bool)`

 SetConnectAppDescriptionNil sets the value for ConnectAppDescription to be an explicit nil

### UnsetConnectAppDescription
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppDescription()`

UnsetConnectAppDescription ensures that no value is present for ConnectAppDescription, not even an explicit nil
### GetConnectAppFriendlyName

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppFriendlyName() string`

GetConnectAppFriendlyName returns the ConnectAppFriendlyName field if non-nil, zero value otherwise.

### GetConnectAppFriendlyNameOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppFriendlyNameOk() (*string, bool)`

GetConnectAppFriendlyNameOk returns a tuple with the ConnectAppFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAppFriendlyName

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppFriendlyName(v string)`

SetConnectAppFriendlyName sets ConnectAppFriendlyName field to given value.

### HasConnectAppFriendlyName

`func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppFriendlyName() bool`

HasConnectAppFriendlyName returns a boolean if a field has been set.

### SetConnectAppFriendlyNameNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppFriendlyNameNil(b bool)`

 SetConnectAppFriendlyNameNil sets the value for ConnectAppFriendlyName to be an explicit nil

### UnsetConnectAppFriendlyName
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppFriendlyName()`

UnsetConnectAppFriendlyName ensures that no value is present for ConnectAppFriendlyName, not even an explicit nil
### GetConnectAppHomepageUrl

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppHomepageUrl() string`

GetConnectAppHomepageUrl returns the ConnectAppHomepageUrl field if non-nil, zero value otherwise.

### GetConnectAppHomepageUrlOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppHomepageUrlOk() (*string, bool)`

GetConnectAppHomepageUrlOk returns a tuple with the ConnectAppHomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAppHomepageUrl

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppHomepageUrl(v string)`

SetConnectAppHomepageUrl sets ConnectAppHomepageUrl field to given value.

### HasConnectAppHomepageUrl

`func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppHomepageUrl() bool`

HasConnectAppHomepageUrl returns a boolean if a field has been set.

### SetConnectAppHomepageUrlNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppHomepageUrlNil(b bool)`

 SetConnectAppHomepageUrlNil sets the value for ConnectAppHomepageUrl to be an explicit nil

### UnsetConnectAppHomepageUrl
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppHomepageUrl()`

UnsetConnectAppHomepageUrl ensures that no value is present for ConnectAppHomepageUrl, not even an explicit nil
### GetConnectAppSid

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppSid() string`

GetConnectAppSid returns the ConnectAppSid field if non-nil, zero value otherwise.

### GetConnectAppSidOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetConnectAppSidOk() (*string, bool)`

GetConnectAppSidOk returns a tuple with the ConnectAppSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectAppSid

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppSid(v string)`

SetConnectAppSid sets ConnectAppSid field to given value.

### HasConnectAppSid

`func (o *ApiV2010AccountAuthorizedConnectApp) HasConnectAppSid() bool`

HasConnectAppSid returns a boolean if a field has been set.

### SetConnectAppSidNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetConnectAppSidNil(b bool)`

 SetConnectAppSidNil sets the value for ConnectAppSid to be an explicit nil

### UnsetConnectAppSid
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetConnectAppSid()`

UnsetConnectAppSid ensures that no value is present for ConnectAppSid, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountAuthorizedConnectApp) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountAuthorizedConnectApp) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountAuthorizedConnectApp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountAuthorizedConnectApp) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountAuthorizedConnectApp) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountAuthorizedConnectApp) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetPermissions

`func (o *ApiV2010AccountAuthorizedConnectApp) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiV2010AccountAuthorizedConnectApp) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiV2010AccountAuthorizedConnectApp) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountAuthorizedConnectApp) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountAuthorizedConnectApp) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountAuthorizedConnectApp) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountAuthorizedConnectApp) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountAuthorizedConnectApp) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountAuthorizedConnectApp) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


