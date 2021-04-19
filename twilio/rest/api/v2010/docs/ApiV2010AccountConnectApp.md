# ApiV2010AccountConnectApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AuthorizeRedirectUrl** | Pointer to **NullableString** | The URL to redirect the user to after authorization | [optional] 
**CompanyName** | Pointer to **NullableString** | The company name set for the Connect App | [optional] 
**DeauthorizeCallbackMethod** | Pointer to **NullableString** | The HTTP method we use to call deauthorize_callback_url | [optional] 
**DeauthorizeCallbackUrl** | Pointer to **NullableString** | The URL we call to de-authorize the Connect App | [optional] 
**Description** | Pointer to **NullableString** | The description of the Connect App | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**HomepageUrl** | Pointer to **NullableString** | The URL users can obtain more information | [optional] 
**Permissions** | Pointer to **[]string** | The set of permissions that your ConnectApp requests | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 

## Methods

### NewApiV2010AccountConnectApp

`func NewApiV2010AccountConnectApp() *ApiV2010AccountConnectApp`

NewApiV2010AccountConnectApp instantiates a new ApiV2010AccountConnectApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountConnectAppWithDefaults

`func NewApiV2010AccountConnectAppWithDefaults() *ApiV2010AccountConnectApp`

NewApiV2010AccountConnectAppWithDefaults instantiates a new ApiV2010AccountConnectApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountConnectApp) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountConnectApp) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountConnectApp) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountConnectApp) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountConnectApp) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountConnectApp) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAuthorizeRedirectUrl

`func (o *ApiV2010AccountConnectApp) GetAuthorizeRedirectUrl() string`

GetAuthorizeRedirectUrl returns the AuthorizeRedirectUrl field if non-nil, zero value otherwise.

### GetAuthorizeRedirectUrlOk

`func (o *ApiV2010AccountConnectApp) GetAuthorizeRedirectUrlOk() (*string, bool)`

GetAuthorizeRedirectUrlOk returns a tuple with the AuthorizeRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeRedirectUrl

`func (o *ApiV2010AccountConnectApp) SetAuthorizeRedirectUrl(v string)`

SetAuthorizeRedirectUrl sets AuthorizeRedirectUrl field to given value.

### HasAuthorizeRedirectUrl

`func (o *ApiV2010AccountConnectApp) HasAuthorizeRedirectUrl() bool`

HasAuthorizeRedirectUrl returns a boolean if a field has been set.

### SetAuthorizeRedirectUrlNil

`func (o *ApiV2010AccountConnectApp) SetAuthorizeRedirectUrlNil(b bool)`

 SetAuthorizeRedirectUrlNil sets the value for AuthorizeRedirectUrl to be an explicit nil

### UnsetAuthorizeRedirectUrl
`func (o *ApiV2010AccountConnectApp) UnsetAuthorizeRedirectUrl()`

UnsetAuthorizeRedirectUrl ensures that no value is present for AuthorizeRedirectUrl, not even an explicit nil
### GetCompanyName

`func (o *ApiV2010AccountConnectApp) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ApiV2010AccountConnectApp) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ApiV2010AccountConnectApp) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ApiV2010AccountConnectApp) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *ApiV2010AccountConnectApp) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *ApiV2010AccountConnectApp) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDeauthorizeCallbackMethod

`func (o *ApiV2010AccountConnectApp) GetDeauthorizeCallbackMethod() string`

GetDeauthorizeCallbackMethod returns the DeauthorizeCallbackMethod field if non-nil, zero value otherwise.

### GetDeauthorizeCallbackMethodOk

`func (o *ApiV2010AccountConnectApp) GetDeauthorizeCallbackMethodOk() (*string, bool)`

GetDeauthorizeCallbackMethodOk returns a tuple with the DeauthorizeCallbackMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeauthorizeCallbackMethod

`func (o *ApiV2010AccountConnectApp) SetDeauthorizeCallbackMethod(v string)`

SetDeauthorizeCallbackMethod sets DeauthorizeCallbackMethod field to given value.

### HasDeauthorizeCallbackMethod

`func (o *ApiV2010AccountConnectApp) HasDeauthorizeCallbackMethod() bool`

HasDeauthorizeCallbackMethod returns a boolean if a field has been set.

### SetDeauthorizeCallbackMethodNil

`func (o *ApiV2010AccountConnectApp) SetDeauthorizeCallbackMethodNil(b bool)`

 SetDeauthorizeCallbackMethodNil sets the value for DeauthorizeCallbackMethod to be an explicit nil

### UnsetDeauthorizeCallbackMethod
`func (o *ApiV2010AccountConnectApp) UnsetDeauthorizeCallbackMethod()`

UnsetDeauthorizeCallbackMethod ensures that no value is present for DeauthorizeCallbackMethod, not even an explicit nil
### GetDeauthorizeCallbackUrl

`func (o *ApiV2010AccountConnectApp) GetDeauthorizeCallbackUrl() string`

GetDeauthorizeCallbackUrl returns the DeauthorizeCallbackUrl field if non-nil, zero value otherwise.

### GetDeauthorizeCallbackUrlOk

`func (o *ApiV2010AccountConnectApp) GetDeauthorizeCallbackUrlOk() (*string, bool)`

GetDeauthorizeCallbackUrlOk returns a tuple with the DeauthorizeCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeauthorizeCallbackUrl

`func (o *ApiV2010AccountConnectApp) SetDeauthorizeCallbackUrl(v string)`

SetDeauthorizeCallbackUrl sets DeauthorizeCallbackUrl field to given value.

### HasDeauthorizeCallbackUrl

`func (o *ApiV2010AccountConnectApp) HasDeauthorizeCallbackUrl() bool`

HasDeauthorizeCallbackUrl returns a boolean if a field has been set.

### SetDeauthorizeCallbackUrlNil

`func (o *ApiV2010AccountConnectApp) SetDeauthorizeCallbackUrlNil(b bool)`

 SetDeauthorizeCallbackUrlNil sets the value for DeauthorizeCallbackUrl to be an explicit nil

### UnsetDeauthorizeCallbackUrl
`func (o *ApiV2010AccountConnectApp) UnsetDeauthorizeCallbackUrl()`

UnsetDeauthorizeCallbackUrl ensures that no value is present for DeauthorizeCallbackUrl, not even an explicit nil
### GetDescription

`func (o *ApiV2010AccountConnectApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2010AccountConnectApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2010AccountConnectApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2010AccountConnectApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2010AccountConnectApp) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2010AccountConnectApp) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountConnectApp) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountConnectApp) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountConnectApp) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountConnectApp) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountConnectApp) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountConnectApp) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetHomepageUrl

`func (o *ApiV2010AccountConnectApp) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *ApiV2010AccountConnectApp) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *ApiV2010AccountConnectApp) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *ApiV2010AccountConnectApp) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### SetHomepageUrlNil

`func (o *ApiV2010AccountConnectApp) SetHomepageUrlNil(b bool)`

 SetHomepageUrlNil sets the value for HomepageUrl to be an explicit nil

### UnsetHomepageUrl
`func (o *ApiV2010AccountConnectApp) UnsetHomepageUrl()`

UnsetHomepageUrl ensures that no value is present for HomepageUrl, not even an explicit nil
### GetPermissions

`func (o *ApiV2010AccountConnectApp) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiV2010AccountConnectApp) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiV2010AccountConnectApp) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiV2010AccountConnectApp) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *ApiV2010AccountConnectApp) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *ApiV2010AccountConnectApp) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountConnectApp) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountConnectApp) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountConnectApp) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountConnectApp) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountConnectApp) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountConnectApp) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountConnectApp) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountConnectApp) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountConnectApp) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountConnectApp) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountConnectApp) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountConnectApp) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


