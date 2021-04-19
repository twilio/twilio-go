# ServerlessV1ServiceEnvironmentLog

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Log resource | [optional] 
**BuildSid** | Pointer to **NullableString** | The SID of the build that corresponds to the log | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Log resource was created | [optional] 
**DeploymentSid** | Pointer to **NullableString** | The SID of the deployment that corresponds to the log | [optional] 
**EnvironmentSid** | Pointer to **NullableString** | The SID of the environment in which the log occurred | [optional] 
**FunctionSid** | Pointer to **NullableString** | The SID of the function whose invocation produced the log | [optional] 
**Level** | Pointer to **NullableString** | The log level | [optional] 
**Message** | Pointer to **NullableString** | The log message | [optional] 
**RequestSid** | Pointer to **NullableString** | The SID of the request associated with the log | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Log resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Log resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Log resource | [optional] 

## Methods

### NewServerlessV1ServiceEnvironmentLog

`func NewServerlessV1ServiceEnvironmentLog() *ServerlessV1ServiceEnvironmentLog`

NewServerlessV1ServiceEnvironmentLog instantiates a new ServerlessV1ServiceEnvironmentLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceEnvironmentLogWithDefaults

`func NewServerlessV1ServiceEnvironmentLogWithDefaults() *ServerlessV1ServiceEnvironmentLog`

NewServerlessV1ServiceEnvironmentLogWithDefaults instantiates a new ServerlessV1ServiceEnvironmentLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetBuildSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetBuildSid() string`

GetBuildSid returns the BuildSid field if non-nil, zero value otherwise.

### GetBuildSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetBuildSidOk() (*string, bool)`

GetBuildSidOk returns a tuple with the BuildSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetBuildSid(v string)`

SetBuildSid sets BuildSid field to given value.

### HasBuildSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasBuildSid() bool`

HasBuildSid returns a boolean if a field has been set.

### SetBuildSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetBuildSidNil(b bool)`

 SetBuildSidNil sets the value for BuildSid to be an explicit nil

### UnsetBuildSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetBuildSid()`

UnsetBuildSid ensures that no value is present for BuildSid, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceEnvironmentLog) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceEnvironmentLog) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceEnvironmentLog) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDeploymentSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetDeploymentSid() string`

GetDeploymentSid returns the DeploymentSid field if non-nil, zero value otherwise.

### GetDeploymentSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetDeploymentSidOk() (*string, bool)`

GetDeploymentSidOk returns a tuple with the DeploymentSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetDeploymentSid(v string)`

SetDeploymentSid sets DeploymentSid field to given value.

### HasDeploymentSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasDeploymentSid() bool`

HasDeploymentSid returns a boolean if a field has been set.

### SetDeploymentSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetDeploymentSidNil(b bool)`

 SetDeploymentSidNil sets the value for DeploymentSid to be an explicit nil

### UnsetDeploymentSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetDeploymentSid()`

UnsetDeploymentSid ensures that no value is present for DeploymentSid, not even an explicit nil
### GetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetEnvironmentSid() string`

GetEnvironmentSid returns the EnvironmentSid field if non-nil, zero value otherwise.

### GetEnvironmentSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetEnvironmentSidOk() (*string, bool)`

GetEnvironmentSidOk returns a tuple with the EnvironmentSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetEnvironmentSid(v string)`

SetEnvironmentSid sets EnvironmentSid field to given value.

### HasEnvironmentSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasEnvironmentSid() bool`

HasEnvironmentSid returns a boolean if a field has been set.

### SetEnvironmentSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetEnvironmentSidNil(b bool)`

 SetEnvironmentSidNil sets the value for EnvironmentSid to be an explicit nil

### UnsetEnvironmentSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetEnvironmentSid()`

UnsetEnvironmentSid ensures that no value is present for EnvironmentSid, not even an explicit nil
### GetFunctionSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetFunctionSid() string`

GetFunctionSid returns the FunctionSid field if non-nil, zero value otherwise.

### GetFunctionSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetFunctionSidOk() (*string, bool)`

GetFunctionSidOk returns a tuple with the FunctionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetFunctionSid(v string)`

SetFunctionSid sets FunctionSid field to given value.

### HasFunctionSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasFunctionSid() bool`

HasFunctionSid returns a boolean if a field has been set.

### SetFunctionSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetFunctionSidNil(b bool)`

 SetFunctionSidNil sets the value for FunctionSid to be an explicit nil

### UnsetFunctionSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetFunctionSid()`

UnsetFunctionSid ensures that no value is present for FunctionSid, not even an explicit nil
### GetLevel

`func (o *ServerlessV1ServiceEnvironmentLog) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ServerlessV1ServiceEnvironmentLog) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ServerlessV1ServiceEnvironmentLog) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetMessage

`func (o *ServerlessV1ServiceEnvironmentLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerlessV1ServiceEnvironmentLog) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServerlessV1ServiceEnvironmentLog) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRequestSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetRequestSid() string`

GetRequestSid returns the RequestSid field if non-nil, zero value otherwise.

### GetRequestSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetRequestSidOk() (*string, bool)`

GetRequestSidOk returns a tuple with the RequestSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetRequestSid(v string)`

SetRequestSid sets RequestSid field to given value.

### HasRequestSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasRequestSid() bool`

HasRequestSid returns a boolean if a field has been set.

### SetRequestSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetRequestSidNil(b bool)`

 SetRequestSidNil sets the value for RequestSid to be an explicit nil

### UnsetRequestSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetRequestSid()`

UnsetRequestSid ensures that no value is present for RequestSid, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceEnvironmentLog) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceEnvironmentLog) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceEnvironmentLog) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceEnvironmentLog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceEnvironmentLog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceEnvironmentLog) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceEnvironmentLog) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceEnvironmentLog) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceEnvironmentLog) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


