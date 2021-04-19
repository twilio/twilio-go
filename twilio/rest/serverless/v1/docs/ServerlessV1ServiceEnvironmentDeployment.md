# ServerlessV1ServiceEnvironmentDeployment

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Deployment resource | [optional] 
**BuildSid** | Pointer to **NullableString** | The SID of the Build for the deployment | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Deployment resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Deployment resource was last updated | [optional] 
**EnvironmentSid** | Pointer to **NullableString** | The SID of the Environment for the Deployment | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Deployment resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Deployment resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Deployment resource | [optional] 

## Methods

### NewServerlessV1ServiceEnvironmentDeployment

`func NewServerlessV1ServiceEnvironmentDeployment() *ServerlessV1ServiceEnvironmentDeployment`

NewServerlessV1ServiceEnvironmentDeployment instantiates a new ServerlessV1ServiceEnvironmentDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceEnvironmentDeploymentWithDefaults

`func NewServerlessV1ServiceEnvironmentDeploymentWithDefaults() *ServerlessV1ServiceEnvironmentDeployment`

NewServerlessV1ServiceEnvironmentDeploymentWithDefaults instantiates a new ServerlessV1ServiceEnvironmentDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBuildSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetBuildSid() string`

GetBuildSid returns the BuildSid field if non-nil, zero value otherwise.

### GetBuildSidOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetBuildSidOk() (*string, bool)`

GetBuildSidOk returns a tuple with the BuildSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetBuildSid(v string)`

SetBuildSid sets BuildSid field to given value.

### HasBuildSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasBuildSid() bool`

HasBuildSid returns a boolean if a field has been set.

### SetBuildSidNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetBuildSidNil(b bool)`

 SetBuildSidNil sets the value for BuildSid to be an explicit nil

### UnsetBuildSid
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetBuildSid()`

UnsetBuildSid ensures that no value is present for BuildSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetEnvironmentSid() string`

GetEnvironmentSid returns the EnvironmentSid field if non-nil, zero value otherwise.

### GetEnvironmentSidOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetEnvironmentSidOk() (*string, bool)`

GetEnvironmentSidOk returns a tuple with the EnvironmentSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetEnvironmentSid(v string)`

SetEnvironmentSid sets EnvironmentSid field to given value.

### HasEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasEnvironmentSid() bool`

HasEnvironmentSid returns a boolean if a field has been set.

### SetEnvironmentSidNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetEnvironmentSidNil(b bool)`

 SetEnvironmentSidNil sets the value for EnvironmentSid to be an explicit nil

### UnsetEnvironmentSid
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetEnvironmentSid()`

UnsetEnvironmentSid ensures that no value is present for EnvironmentSid, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceEnvironmentDeployment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceEnvironmentDeployment) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceEnvironmentDeployment) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceEnvironmentDeployment) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


