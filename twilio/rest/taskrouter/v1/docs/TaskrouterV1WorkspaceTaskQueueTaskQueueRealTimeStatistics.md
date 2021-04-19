# TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActivityStatistics** | Pointer to **[]map[string]interface{}** | The number of current Workers by Activity | [optional] 
**LongestRelativeTaskAgeInQueue** | Pointer to **NullableInt32** | The relative age in the TaskQueue for the longest waiting Task. | [optional] 
**LongestRelativeTaskSidInQueue** | Pointer to **NullableString** | The SID of the Task waiting in the TaskQueue the longest. | [optional] 
**LongestTaskWaitingAge** | Pointer to **NullableInt32** | The age of the longest waiting Task | [optional] 
**LongestTaskWaitingSid** | Pointer to **NullableString** | The SID of the longest waiting Task | [optional] 
**TaskQueueSid** | Pointer to **NullableString** | The SID of the TaskQueue from which these statistics were calculated | [optional] 
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority | [optional] 
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status | [optional] 
**TotalAvailableWorkers** | Pointer to **NullableInt32** | The total number of Workers available for Tasks in the TaskQueue | [optional] 
**TotalEligibleWorkers** | Pointer to **NullableInt32** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state | [optional] 
**TotalTasks** | Pointer to **NullableInt32** | The total number of Tasks | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskQueue statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the TaskQueue | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics() *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatisticsWithDefaults() *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActivityStatistics

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetActivityStatistics() []map[string]interface{}`

GetActivityStatistics returns the ActivityStatistics field if non-nil, zero value otherwise.

### GetActivityStatisticsOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetActivityStatisticsOk() (*[]map[string]interface{}, bool)`

GetActivityStatisticsOk returns a tuple with the ActivityStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityStatistics

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetActivityStatistics(v []map[string]interface{})`

SetActivityStatistics sets ActivityStatistics field to given value.

### HasActivityStatistics

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasActivityStatistics() bool`

HasActivityStatistics returns a boolean if a field has been set.

### SetActivityStatisticsNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetActivityStatisticsNil(b bool)`

 SetActivityStatisticsNil sets the value for ActivityStatistics to be an explicit nil

### UnsetActivityStatistics
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetActivityStatistics()`

UnsetActivityStatistics ensures that no value is present for ActivityStatistics, not even an explicit nil
### GetLongestRelativeTaskAgeInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestRelativeTaskAgeInQueue() int32`

GetLongestRelativeTaskAgeInQueue returns the LongestRelativeTaskAgeInQueue field if non-nil, zero value otherwise.

### GetLongestRelativeTaskAgeInQueueOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestRelativeTaskAgeInQueueOk() (*int32, bool)`

GetLongestRelativeTaskAgeInQueueOk returns a tuple with the LongestRelativeTaskAgeInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestRelativeTaskAgeInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestRelativeTaskAgeInQueue(v int32)`

SetLongestRelativeTaskAgeInQueue sets LongestRelativeTaskAgeInQueue field to given value.

### HasLongestRelativeTaskAgeInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasLongestRelativeTaskAgeInQueue() bool`

HasLongestRelativeTaskAgeInQueue returns a boolean if a field has been set.

### SetLongestRelativeTaskAgeInQueueNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestRelativeTaskAgeInQueueNil(b bool)`

 SetLongestRelativeTaskAgeInQueueNil sets the value for LongestRelativeTaskAgeInQueue to be an explicit nil

### UnsetLongestRelativeTaskAgeInQueue
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetLongestRelativeTaskAgeInQueue()`

UnsetLongestRelativeTaskAgeInQueue ensures that no value is present for LongestRelativeTaskAgeInQueue, not even an explicit nil
### GetLongestRelativeTaskSidInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestRelativeTaskSidInQueue() string`

GetLongestRelativeTaskSidInQueue returns the LongestRelativeTaskSidInQueue field if non-nil, zero value otherwise.

### GetLongestRelativeTaskSidInQueueOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestRelativeTaskSidInQueueOk() (*string, bool)`

GetLongestRelativeTaskSidInQueueOk returns a tuple with the LongestRelativeTaskSidInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestRelativeTaskSidInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestRelativeTaskSidInQueue(v string)`

SetLongestRelativeTaskSidInQueue sets LongestRelativeTaskSidInQueue field to given value.

### HasLongestRelativeTaskSidInQueue

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasLongestRelativeTaskSidInQueue() bool`

HasLongestRelativeTaskSidInQueue returns a boolean if a field has been set.

### SetLongestRelativeTaskSidInQueueNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestRelativeTaskSidInQueueNil(b bool)`

 SetLongestRelativeTaskSidInQueueNil sets the value for LongestRelativeTaskSidInQueue to be an explicit nil

### UnsetLongestRelativeTaskSidInQueue
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetLongestRelativeTaskSidInQueue()`

UnsetLongestRelativeTaskSidInQueue ensures that no value is present for LongestRelativeTaskSidInQueue, not even an explicit nil
### GetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestTaskWaitingAge() int32`

GetLongestTaskWaitingAge returns the LongestTaskWaitingAge field if non-nil, zero value otherwise.

### GetLongestTaskWaitingAgeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestTaskWaitingAgeOk() (*int32, bool)`

GetLongestTaskWaitingAgeOk returns a tuple with the LongestTaskWaitingAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestTaskWaitingAge(v int32)`

SetLongestTaskWaitingAge sets LongestTaskWaitingAge field to given value.

### HasLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasLongestTaskWaitingAge() bool`

HasLongestTaskWaitingAge returns a boolean if a field has been set.

### SetLongestTaskWaitingAgeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestTaskWaitingAgeNil(b bool)`

 SetLongestTaskWaitingAgeNil sets the value for LongestTaskWaitingAge to be an explicit nil

### UnsetLongestTaskWaitingAge
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetLongestTaskWaitingAge()`

UnsetLongestTaskWaitingAge ensures that no value is present for LongestTaskWaitingAge, not even an explicit nil
### GetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestTaskWaitingSid() string`

GetLongestTaskWaitingSid returns the LongestTaskWaitingSid field if non-nil, zero value otherwise.

### GetLongestTaskWaitingSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetLongestTaskWaitingSidOk() (*string, bool)`

GetLongestTaskWaitingSidOk returns a tuple with the LongestTaskWaitingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestTaskWaitingSid(v string)`

SetLongestTaskWaitingSid sets LongestTaskWaitingSid field to given value.

### HasLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasLongestTaskWaitingSid() bool`

HasLongestTaskWaitingSid returns a boolean if a field has been set.

### SetLongestTaskWaitingSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetLongestTaskWaitingSidNil(b bool)`

 SetLongestTaskWaitingSidNil sets the value for LongestTaskWaitingSid to be an explicit nil

### UnsetLongestTaskWaitingSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetLongestTaskWaitingSid()`

UnsetLongestTaskWaitingSid ensures that no value is present for LongestTaskWaitingSid, not even an explicit nil
### GetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTaskQueueSid() string`

GetTaskQueueSid returns the TaskQueueSid field if non-nil, zero value otherwise.

### GetTaskQueueSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTaskQueueSidOk() (*string, bool)`

GetTaskQueueSidOk returns a tuple with the TaskQueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTaskQueueSid(v string)`

SetTaskQueueSid sets TaskQueueSid field to given value.

### HasTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTaskQueueSid() bool`

HasTaskQueueSid returns a boolean if a field has been set.

### SetTaskQueueSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTaskQueueSidNil(b bool)`

 SetTaskQueueSidNil sets the value for TaskQueueSid to be an explicit nil

### UnsetTaskQueueSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTaskQueueSid()`

UnsetTaskQueueSid ensures that no value is present for TaskQueueSid, not even an explicit nil
### GetTasksByPriority

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTasksByPriority() map[string]interface{}`

GetTasksByPriority returns the TasksByPriority field if non-nil, zero value otherwise.

### GetTasksByPriorityOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTasksByPriorityOk() (*map[string]interface{}, bool)`

GetTasksByPriorityOk returns a tuple with the TasksByPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByPriority

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTasksByPriority(v map[string]interface{})`

SetTasksByPriority sets TasksByPriority field to given value.

### HasTasksByPriority

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTasksByPriority() bool`

HasTasksByPriority returns a boolean if a field has been set.

### SetTasksByPriorityNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTasksByPriorityNil(b bool)`

 SetTasksByPriorityNil sets the value for TasksByPriority to be an explicit nil

### UnsetTasksByPriority
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTasksByPriority()`

UnsetTasksByPriority ensures that no value is present for TasksByPriority, not even an explicit nil
### GetTasksByStatus

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTasksByStatus() map[string]interface{}`

GetTasksByStatus returns the TasksByStatus field if non-nil, zero value otherwise.

### GetTasksByStatusOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTasksByStatusOk() (*map[string]interface{}, bool)`

GetTasksByStatusOk returns a tuple with the TasksByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByStatus

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTasksByStatus(v map[string]interface{})`

SetTasksByStatus sets TasksByStatus field to given value.

### HasTasksByStatus

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTasksByStatus() bool`

HasTasksByStatus returns a boolean if a field has been set.

### SetTasksByStatusNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTasksByStatusNil(b bool)`

 SetTasksByStatusNil sets the value for TasksByStatus to be an explicit nil

### UnsetTasksByStatus
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTasksByStatus()`

UnsetTasksByStatus ensures that no value is present for TasksByStatus, not even an explicit nil
### GetTotalAvailableWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalAvailableWorkers() int32`

GetTotalAvailableWorkers returns the TotalAvailableWorkers field if non-nil, zero value otherwise.

### GetTotalAvailableWorkersOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalAvailableWorkersOk() (*int32, bool)`

GetTotalAvailableWorkersOk returns a tuple with the TotalAvailableWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailableWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalAvailableWorkers(v int32)`

SetTotalAvailableWorkers sets TotalAvailableWorkers field to given value.

### HasTotalAvailableWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTotalAvailableWorkers() bool`

HasTotalAvailableWorkers returns a boolean if a field has been set.

### SetTotalAvailableWorkersNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalAvailableWorkersNil(b bool)`

 SetTotalAvailableWorkersNil sets the value for TotalAvailableWorkers to be an explicit nil

### UnsetTotalAvailableWorkers
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTotalAvailableWorkers()`

UnsetTotalAvailableWorkers ensures that no value is present for TotalAvailableWorkers, not even an explicit nil
### GetTotalEligibleWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalEligibleWorkers() int32`

GetTotalEligibleWorkers returns the TotalEligibleWorkers field if non-nil, zero value otherwise.

### GetTotalEligibleWorkersOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalEligibleWorkersOk() (*int32, bool)`

GetTotalEligibleWorkersOk returns a tuple with the TotalEligibleWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEligibleWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalEligibleWorkers(v int32)`

SetTotalEligibleWorkers sets TotalEligibleWorkers field to given value.

### HasTotalEligibleWorkers

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTotalEligibleWorkers() bool`

HasTotalEligibleWorkers returns a boolean if a field has been set.

### SetTotalEligibleWorkersNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalEligibleWorkersNil(b bool)`

 SetTotalEligibleWorkersNil sets the value for TotalEligibleWorkers to be an explicit nil

### UnsetTotalEligibleWorkers
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTotalEligibleWorkers()`

UnsetTotalEligibleWorkers ensures that no value is present for TotalEligibleWorkers, not even an explicit nil
### GetTotalTasks

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalTasks() int32`

GetTotalTasks returns the TotalTasks field if non-nil, zero value otherwise.

### GetTotalTasksOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetTotalTasksOk() (*int32, bool)`

GetTotalTasksOk returns a tuple with the TotalTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasks

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalTasks(v int32)`

SetTotalTasks sets TotalTasks field to given value.

### HasTotalTasks

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasTotalTasks() bool`

HasTotalTasks returns a boolean if a field has been set.

### SetTotalTasksNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetTotalTasksNil(b bool)`

 SetTotalTasksNil sets the value for TotalTasks to be an explicit nil

### UnsetTotalTasks
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetTotalTasks()`

UnsetTotalTasks ensures that no value is present for TotalTasks, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


