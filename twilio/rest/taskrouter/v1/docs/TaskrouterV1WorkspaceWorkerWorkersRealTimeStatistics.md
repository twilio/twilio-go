# TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActivityStatistics** | Pointer to **[]map[string]interface{}** | The number of current Workers by Activity | [optional] 
**TotalWorkers** | Pointer to **NullableInt32** | The total number of Workers | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workers statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Workers | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics

`func NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics() *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics instantiates a new TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkerWorkersRealTimeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetActivityStatistics() []map[string]interface{}`

GetActivityStatistics returns the ActivityStatistics field if non-nil, zero value otherwise.

### GetActivityStatisticsOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetActivityStatisticsOk() (*[]map[string]interface{}, bool)`

GetActivityStatisticsOk returns a tuple with the ActivityStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetActivityStatistics(v []map[string]interface{})`

SetActivityStatistics sets ActivityStatistics field to given value.

### HasActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) HasActivityStatistics() bool`

HasActivityStatistics returns a boolean if a field has been set.

### SetActivityStatisticsNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetActivityStatisticsNil(b bool)`

 SetActivityStatisticsNil sets the value for ActivityStatistics to be an explicit nil

### UnsetActivityStatistics
`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) UnsetActivityStatistics()`

UnsetActivityStatistics ensures that no value is present for ActivityStatistics, not even an explicit nil
### GetTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetTotalWorkers() int32`

GetTotalWorkers returns the TotalWorkers field if non-nil, zero value otherwise.

### GetTotalWorkersOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetTotalWorkersOk() (*int32, bool)`

GetTotalWorkersOk returns a tuple with the TotalWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetTotalWorkers(v int32)`

SetTotalWorkers sets TotalWorkers field to given value.

### HasTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) HasTotalWorkers() bool`

HasTotalWorkers returns a boolean if a field has been set.

### SetTotalWorkersNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetTotalWorkersNil(b bool)`

 SetTotalWorkersNil sets the value for TotalWorkers to be an explicit nil

### UnsetTotalWorkers
`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) UnsetTotalWorkers()`

UnsetTotalWorkers ensures that no value is present for TotalWorkers, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


