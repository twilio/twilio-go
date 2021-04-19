# TaskrouterV1WorkspaceWorkspaceRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActivityStatistics** | Pointer to **[]map[string]interface{}** | The number of current Workers by Activity | [optional] 
**LongestTaskWaitingAge** | Pointer to **NullableInt32** | The age of the longest waiting Task | [optional] 
**LongestTaskWaitingSid** | Pointer to **NullableString** | The SID of the longest waiting Task | [optional] 
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority | [optional] 
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status | [optional] 
**TotalTasks** | Pointer to **NullableInt32** | The total number of Tasks | [optional] 
**TotalWorkers** | Pointer to **NullableInt32** | The total number of Workers in the Workspace | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workspace statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkspaceRealTimeStatistics

`func NewTaskrouterV1WorkspaceWorkspaceRealTimeStatistics() *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkspaceRealTimeStatistics instantiates a new TaskrouterV1WorkspaceWorkspaceRealTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkspaceRealTimeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkspaceRealTimeStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkspaceRealTimeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkspaceRealTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetActivityStatistics() []map[string]interface{}`

GetActivityStatistics returns the ActivityStatistics field if non-nil, zero value otherwise.

### GetActivityStatisticsOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetActivityStatisticsOk() (*[]map[string]interface{}, bool)`

GetActivityStatisticsOk returns a tuple with the ActivityStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetActivityStatistics(v []map[string]interface{})`

SetActivityStatistics sets ActivityStatistics field to given value.

### HasActivityStatistics

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasActivityStatistics() bool`

HasActivityStatistics returns a boolean if a field has been set.

### SetActivityStatisticsNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetActivityStatisticsNil(b bool)`

 SetActivityStatisticsNil sets the value for ActivityStatistics to be an explicit nil

### UnsetActivityStatistics
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetActivityStatistics()`

UnsetActivityStatistics ensures that no value is present for ActivityStatistics, not even an explicit nil
### GetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetLongestTaskWaitingAge() int32`

GetLongestTaskWaitingAge returns the LongestTaskWaitingAge field if non-nil, zero value otherwise.

### GetLongestTaskWaitingAgeOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetLongestTaskWaitingAgeOk() (*int32, bool)`

GetLongestTaskWaitingAgeOk returns a tuple with the LongestTaskWaitingAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetLongestTaskWaitingAge(v int32)`

SetLongestTaskWaitingAge sets LongestTaskWaitingAge field to given value.

### HasLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasLongestTaskWaitingAge() bool`

HasLongestTaskWaitingAge returns a boolean if a field has been set.

### SetLongestTaskWaitingAgeNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetLongestTaskWaitingAgeNil(b bool)`

 SetLongestTaskWaitingAgeNil sets the value for LongestTaskWaitingAge to be an explicit nil

### UnsetLongestTaskWaitingAge
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetLongestTaskWaitingAge()`

UnsetLongestTaskWaitingAge ensures that no value is present for LongestTaskWaitingAge, not even an explicit nil
### GetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetLongestTaskWaitingSid() string`

GetLongestTaskWaitingSid returns the LongestTaskWaitingSid field if non-nil, zero value otherwise.

### GetLongestTaskWaitingSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetLongestTaskWaitingSidOk() (*string, bool)`

GetLongestTaskWaitingSidOk returns a tuple with the LongestTaskWaitingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetLongestTaskWaitingSid(v string)`

SetLongestTaskWaitingSid sets LongestTaskWaitingSid field to given value.

### HasLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasLongestTaskWaitingSid() bool`

HasLongestTaskWaitingSid returns a boolean if a field has been set.

### SetLongestTaskWaitingSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetLongestTaskWaitingSidNil(b bool)`

 SetLongestTaskWaitingSidNil sets the value for LongestTaskWaitingSid to be an explicit nil

### UnsetLongestTaskWaitingSid
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetLongestTaskWaitingSid()`

UnsetLongestTaskWaitingSid ensures that no value is present for LongestTaskWaitingSid, not even an explicit nil
### GetTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTasksByPriority() map[string]interface{}`

GetTasksByPriority returns the TasksByPriority field if non-nil, zero value otherwise.

### GetTasksByPriorityOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTasksByPriorityOk() (*map[string]interface{}, bool)`

GetTasksByPriorityOk returns a tuple with the TasksByPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTasksByPriority(v map[string]interface{})`

SetTasksByPriority sets TasksByPriority field to given value.

### HasTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasTasksByPriority() bool`

HasTasksByPriority returns a boolean if a field has been set.

### SetTasksByPriorityNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTasksByPriorityNil(b bool)`

 SetTasksByPriorityNil sets the value for TasksByPriority to be an explicit nil

### UnsetTasksByPriority
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetTasksByPriority()`

UnsetTasksByPriority ensures that no value is present for TasksByPriority, not even an explicit nil
### GetTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTasksByStatus() map[string]interface{}`

GetTasksByStatus returns the TasksByStatus field if non-nil, zero value otherwise.

### GetTasksByStatusOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTasksByStatusOk() (*map[string]interface{}, bool)`

GetTasksByStatusOk returns a tuple with the TasksByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTasksByStatus(v map[string]interface{})`

SetTasksByStatus sets TasksByStatus field to given value.

### HasTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasTasksByStatus() bool`

HasTasksByStatus returns a boolean if a field has been set.

### SetTasksByStatusNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTasksByStatusNil(b bool)`

 SetTasksByStatusNil sets the value for TasksByStatus to be an explicit nil

### UnsetTasksByStatus
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetTasksByStatus()`

UnsetTasksByStatus ensures that no value is present for TasksByStatus, not even an explicit nil
### GetTotalTasks

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTotalTasks() int32`

GetTotalTasks returns the TotalTasks field if non-nil, zero value otherwise.

### GetTotalTasksOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTotalTasksOk() (*int32, bool)`

GetTotalTasksOk returns a tuple with the TotalTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasks

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTotalTasks(v int32)`

SetTotalTasks sets TotalTasks field to given value.

### HasTotalTasks

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasTotalTasks() bool`

HasTotalTasks returns a boolean if a field has been set.

### SetTotalTasksNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTotalTasksNil(b bool)`

 SetTotalTasksNil sets the value for TotalTasks to be an explicit nil

### UnsetTotalTasks
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetTotalTasks()`

UnsetTotalTasks ensures that no value is present for TotalTasks, not even an explicit nil
### GetTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTotalWorkers() int32`

GetTotalWorkers returns the TotalWorkers field if non-nil, zero value otherwise.

### GetTotalWorkersOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetTotalWorkersOk() (*int32, bool)`

GetTotalWorkersOk returns a tuple with the TotalWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTotalWorkers(v int32)`

SetTotalWorkers sets TotalWorkers field to given value.

### HasTotalWorkers

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasTotalWorkers() bool`

HasTotalWorkers returns a boolean if a field has been set.

### SetTotalWorkersNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetTotalWorkersNil(b bool)`

 SetTotalWorkersNil sets the value for TotalWorkers to be an explicit nil

### UnsetTotalWorkers
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetTotalWorkers()`

UnsetTotalWorkers ensures that no value is present for TotalWorkers, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkspaceRealTimeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


