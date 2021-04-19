# ServerlessV1ServiceAssetAssetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Asset Version resource | [optional] 
**AssetSid** | Pointer to **NullableString** | The SID of the Asset resource that is the parent of the Asset Version | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Asset Version resource was created | [optional] 
**Path** | Pointer to **NullableString** | The URL-friendly string by which the Asset Version can be referenced | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Asset Version resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Asset Version resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Asset Version resource | [optional] 
**Visibility** | Pointer to **NullableString** | The access control that determines how the Asset Version can be accessed | [optional] 

## Methods

### NewServerlessV1ServiceAssetAssetVersion

`func NewServerlessV1ServiceAssetAssetVersion() *ServerlessV1ServiceAssetAssetVersion`

NewServerlessV1ServiceAssetAssetVersion instantiates a new ServerlessV1ServiceAssetAssetVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceAssetAssetVersionWithDefaults

`func NewServerlessV1ServiceAssetAssetVersionWithDefaults() *ServerlessV1ServiceAssetAssetVersion`

NewServerlessV1ServiceAssetAssetVersionWithDefaults instantiates a new ServerlessV1ServiceAssetAssetVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceAssetAssetVersion) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceAssetAssetVersion) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceAssetAssetVersion) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssetSid

`func (o *ServerlessV1ServiceAssetAssetVersion) GetAssetSid() string`

GetAssetSid returns the AssetSid field if non-nil, zero value otherwise.

### GetAssetSidOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetAssetSidOk() (*string, bool)`

GetAssetSidOk returns a tuple with the AssetSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSid

`func (o *ServerlessV1ServiceAssetAssetVersion) SetAssetSid(v string)`

SetAssetSid sets AssetSid field to given value.

### HasAssetSid

`func (o *ServerlessV1ServiceAssetAssetVersion) HasAssetSid() bool`

HasAssetSid returns a boolean if a field has been set.

### SetAssetSidNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetAssetSidNil(b bool)`

 SetAssetSidNil sets the value for AssetSid to be an explicit nil

### UnsetAssetSid
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetAssetSid()`

UnsetAssetSid ensures that no value is present for AssetSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceAssetAssetVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceAssetAssetVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceAssetAssetVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetPath

`func (o *ServerlessV1ServiceAssetAssetVersion) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ServerlessV1ServiceAssetAssetVersion) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ServerlessV1ServiceAssetAssetVersion) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceAssetAssetVersion) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceAssetAssetVersion) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceAssetAssetVersion) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceAssetAssetVersion) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceAssetAssetVersion) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceAssetAssetVersion) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceAssetAssetVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceAssetAssetVersion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceAssetAssetVersion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVisibility

`func (o *ServerlessV1ServiceAssetAssetVersion) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServerlessV1ServiceAssetAssetVersion) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServerlessV1ServiceAssetAssetVersion) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServerlessV1ServiceAssetAssetVersion) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *ServerlessV1ServiceAssetAssetVersion) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *ServerlessV1ServiceAssetAssetVersion) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


