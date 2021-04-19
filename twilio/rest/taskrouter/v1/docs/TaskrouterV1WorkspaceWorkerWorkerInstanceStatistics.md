# TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Cumulative** | Pointer to **map[string]interface{}** | An object that contains the cumulative statistics for the Worker | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the WorkerChannel statistics resource | [optional] 
**WorkerSid** | Pointer to **NullableString** | The SID of the Worker that contains the WorkerChannel | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the WorkerChannel | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatistics

`func NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatistics() *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics`

NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatistics instantiates a new TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics`

NewTaskrouterV1WorkspaceWorkerWorkerInstanceStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetCumulative() map[string]interface{}`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetCumulativeOk() (*map[string]interface{}, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetCumulative(v map[string]interface{})`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetWorkerSid() string`

GetWorkerSid returns the WorkerSid field if non-nil, zero value otherwise.

### GetWorkerSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetWorkerSidOk() (*string, bool)`

GetWorkerSidOk returns a tuple with the WorkerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetWorkerSid(v string)`

SetWorkerSid sets WorkerSid field to given value.

### HasWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) HasWorkerSid() bool`

HasWorkerSid returns a boolean if a field has been set.

### SetWorkerSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetWorkerSidNil(b bool)`

 SetWorkerSidNil sets the value for WorkerSid to be an explicit nil

### UnsetWorkerSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) UnsetWorkerSid()`

UnsetWorkerSid ensures that no value is present for WorkerSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


