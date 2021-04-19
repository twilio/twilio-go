# ServerlessV1ServiceBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the Build resource | [optional] 
**AssetVersions** | Pointer to **[]map[string]interface{}** | The list of Asset Version resource SIDs that are included in the Build | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Build resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Build resource was last updated | [optional] 
**Dependencies** | Pointer to **[]map[string]interface{}** | A list of objects that describe the Dependencies included in the Build | [optional] 
**FunctionVersions** | Pointer to **[]map[string]interface{}** | The list of Function Version resource SIDs that are included in the Build | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Runtime** | Pointer to **NullableString** | The Runtime version that will be used to run the Build. | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the Build resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the Build resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Build | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Build resource | [optional] 

## Methods

### NewServerlessV1ServiceBuild

`func NewServerlessV1ServiceBuild() *ServerlessV1ServiceBuild`

NewServerlessV1ServiceBuild instantiates a new ServerlessV1ServiceBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessV1ServiceBuildWithDefaults

`func NewServerlessV1ServiceBuildWithDefaults() *ServerlessV1ServiceBuild`

NewServerlessV1ServiceBuildWithDefaults instantiates a new ServerlessV1ServiceBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ServerlessV1ServiceBuild) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ServerlessV1ServiceBuild) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ServerlessV1ServiceBuild) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ServerlessV1ServiceBuild) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ServerlessV1ServiceBuild) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ServerlessV1ServiceBuild) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssetVersions

`func (o *ServerlessV1ServiceBuild) GetAssetVersions() []map[string]interface{}`

GetAssetVersions returns the AssetVersions field if non-nil, zero value otherwise.

### GetAssetVersionsOk

`func (o *ServerlessV1ServiceBuild) GetAssetVersionsOk() (*[]map[string]interface{}, bool)`

GetAssetVersionsOk returns a tuple with the AssetVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersions

`func (o *ServerlessV1ServiceBuild) SetAssetVersions(v []map[string]interface{})`

SetAssetVersions sets AssetVersions field to given value.

### HasAssetVersions

`func (o *ServerlessV1ServiceBuild) HasAssetVersions() bool`

HasAssetVersions returns a boolean if a field has been set.

### SetAssetVersionsNil

`func (o *ServerlessV1ServiceBuild) SetAssetVersionsNil(b bool)`

 SetAssetVersionsNil sets the value for AssetVersions to be an explicit nil

### UnsetAssetVersions
`func (o *ServerlessV1ServiceBuild) UnsetAssetVersions()`

UnsetAssetVersions ensures that no value is present for AssetVersions, not even an explicit nil
### GetDateCreated

`func (o *ServerlessV1ServiceBuild) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServerlessV1ServiceBuild) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServerlessV1ServiceBuild) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServerlessV1ServiceBuild) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ServerlessV1ServiceBuild) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ServerlessV1ServiceBuild) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ServerlessV1ServiceBuild) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ServerlessV1ServiceBuild) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ServerlessV1ServiceBuild) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ServerlessV1ServiceBuild) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ServerlessV1ServiceBuild) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ServerlessV1ServiceBuild) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDependencies

`func (o *ServerlessV1ServiceBuild) GetDependencies() []map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ServerlessV1ServiceBuild) GetDependenciesOk() (*[]map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ServerlessV1ServiceBuild) SetDependencies(v []map[string]interface{})`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *ServerlessV1ServiceBuild) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### SetDependenciesNil

`func (o *ServerlessV1ServiceBuild) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *ServerlessV1ServiceBuild) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetFunctionVersions

`func (o *ServerlessV1ServiceBuild) GetFunctionVersions() []map[string]interface{}`

GetFunctionVersions returns the FunctionVersions field if non-nil, zero value otherwise.

### GetFunctionVersionsOk

`func (o *ServerlessV1ServiceBuild) GetFunctionVersionsOk() (*[]map[string]interface{}, bool)`

GetFunctionVersionsOk returns a tuple with the FunctionVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionVersions

`func (o *ServerlessV1ServiceBuild) SetFunctionVersions(v []map[string]interface{})`

SetFunctionVersions sets FunctionVersions field to given value.

### HasFunctionVersions

`func (o *ServerlessV1ServiceBuild) HasFunctionVersions() bool`

HasFunctionVersions returns a boolean if a field has been set.

### SetFunctionVersionsNil

`func (o *ServerlessV1ServiceBuild) SetFunctionVersionsNil(b bool)`

 SetFunctionVersionsNil sets the value for FunctionVersions to be an explicit nil

### UnsetFunctionVersions
`func (o *ServerlessV1ServiceBuild) UnsetFunctionVersions()`

UnsetFunctionVersions ensures that no value is present for FunctionVersions, not even an explicit nil
### GetLinks

`func (o *ServerlessV1ServiceBuild) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessV1ServiceBuild) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessV1ServiceBuild) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessV1ServiceBuild) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ServerlessV1ServiceBuild) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ServerlessV1ServiceBuild) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRuntime

`func (o *ServerlessV1ServiceBuild) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *ServerlessV1ServiceBuild) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *ServerlessV1ServiceBuild) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *ServerlessV1ServiceBuild) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *ServerlessV1ServiceBuild) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *ServerlessV1ServiceBuild) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil
### GetServiceSid

`func (o *ServerlessV1ServiceBuild) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ServerlessV1ServiceBuild) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ServerlessV1ServiceBuild) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ServerlessV1ServiceBuild) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ServerlessV1ServiceBuild) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ServerlessV1ServiceBuild) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ServerlessV1ServiceBuild) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ServerlessV1ServiceBuild) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ServerlessV1ServiceBuild) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ServerlessV1ServiceBuild) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ServerlessV1ServiceBuild) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ServerlessV1ServiceBuild) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ServerlessV1ServiceBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerlessV1ServiceBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerlessV1ServiceBuild) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerlessV1ServiceBuild) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServerlessV1ServiceBuild) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServerlessV1ServiceBuild) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *ServerlessV1ServiceBuild) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServerlessV1ServiceBuild) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServerlessV1ServiceBuild) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServerlessV1ServiceBuild) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ServerlessV1ServiceBuild) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ServerlessV1ServiceBuild) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


