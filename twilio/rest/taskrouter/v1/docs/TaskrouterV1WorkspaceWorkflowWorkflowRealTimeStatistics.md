# TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**LongestTaskWaitingAge** | Pointer to **NullableInt32** | The age of the longest waiting Task | [optional] 
**LongestTaskWaitingSid** | Pointer to **NullableString** | The SID of the longest waiting Task | [optional] 
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority | [optional] 
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status | [optional] 
**TotalTasks** | Pointer to **NullableInt32** | The total number of Tasks | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workflow statistics resource | [optional] 
**WorkflowSid** | Pointer to **NullableString** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Workflow. | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics

`func NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics() *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics instantiates a new TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics`

NewTaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetLongestTaskWaitingAge() int32`

GetLongestTaskWaitingAge returns the LongestTaskWaitingAge field if non-nil, zero value otherwise.

### GetLongestTaskWaitingAgeOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetLongestTaskWaitingAgeOk() (*int32, bool)`

GetLongestTaskWaitingAgeOk returns a tuple with the LongestTaskWaitingAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetLongestTaskWaitingAge(v int32)`

SetLongestTaskWaitingAge sets LongestTaskWaitingAge field to given value.

### HasLongestTaskWaitingAge

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasLongestTaskWaitingAge() bool`

HasLongestTaskWaitingAge returns a boolean if a field has been set.

### SetLongestTaskWaitingAgeNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetLongestTaskWaitingAgeNil(b bool)`

 SetLongestTaskWaitingAgeNil sets the value for LongestTaskWaitingAge to be an explicit nil

### UnsetLongestTaskWaitingAge
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetLongestTaskWaitingAge()`

UnsetLongestTaskWaitingAge ensures that no value is present for LongestTaskWaitingAge, not even an explicit nil
### GetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetLongestTaskWaitingSid() string`

GetLongestTaskWaitingSid returns the LongestTaskWaitingSid field if non-nil, zero value otherwise.

### GetLongestTaskWaitingSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetLongestTaskWaitingSidOk() (*string, bool)`

GetLongestTaskWaitingSidOk returns a tuple with the LongestTaskWaitingSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetLongestTaskWaitingSid(v string)`

SetLongestTaskWaitingSid sets LongestTaskWaitingSid field to given value.

### HasLongestTaskWaitingSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasLongestTaskWaitingSid() bool`

HasLongestTaskWaitingSid returns a boolean if a field has been set.

### SetLongestTaskWaitingSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetLongestTaskWaitingSidNil(b bool)`

 SetLongestTaskWaitingSidNil sets the value for LongestTaskWaitingSid to be an explicit nil

### UnsetLongestTaskWaitingSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetLongestTaskWaitingSid()`

UnsetLongestTaskWaitingSid ensures that no value is present for LongestTaskWaitingSid, not even an explicit nil
### GetTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTasksByPriority() map[string]interface{}`

GetTasksByPriority returns the TasksByPriority field if non-nil, zero value otherwise.

### GetTasksByPriorityOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTasksByPriorityOk() (*map[string]interface{}, bool)`

GetTasksByPriorityOk returns a tuple with the TasksByPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTasksByPriority(v map[string]interface{})`

SetTasksByPriority sets TasksByPriority field to given value.

### HasTasksByPriority

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasTasksByPriority() bool`

HasTasksByPriority returns a boolean if a field has been set.

### SetTasksByPriorityNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTasksByPriorityNil(b bool)`

 SetTasksByPriorityNil sets the value for TasksByPriority to be an explicit nil

### UnsetTasksByPriority
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetTasksByPriority()`

UnsetTasksByPriority ensures that no value is present for TasksByPriority, not even an explicit nil
### GetTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTasksByStatus() map[string]interface{}`

GetTasksByStatus returns the TasksByStatus field if non-nil, zero value otherwise.

### GetTasksByStatusOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTasksByStatusOk() (*map[string]interface{}, bool)`

GetTasksByStatusOk returns a tuple with the TasksByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTasksByStatus(v map[string]interface{})`

SetTasksByStatus sets TasksByStatus field to given value.

### HasTasksByStatus

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasTasksByStatus() bool`

HasTasksByStatus returns a boolean if a field has been set.

### SetTasksByStatusNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTasksByStatusNil(b bool)`

 SetTasksByStatusNil sets the value for TasksByStatus to be an explicit nil

### UnsetTasksByStatus
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetTasksByStatus()`

UnsetTasksByStatus ensures that no value is present for TasksByStatus, not even an explicit nil
### GetTotalTasks

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTotalTasks() int32`

GetTotalTasks returns the TotalTasks field if non-nil, zero value otherwise.

### GetTotalTasksOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetTotalTasksOk() (*int32, bool)`

GetTotalTasksOk returns a tuple with the TotalTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasks

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTotalTasks(v int32)`

SetTotalTasks sets TotalTasks field to given value.

### HasTotalTasks

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasTotalTasks() bool`

HasTotalTasks returns a boolean if a field has been set.

### SetTotalTasksNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetTotalTasksNil(b bool)`

 SetTotalTasksNil sets the value for TotalTasks to be an explicit nil

### UnsetTotalTasks
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetTotalTasks()`

UnsetTotalTasks ensures that no value is present for TotalTasks, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetWorkflowSid() string`

GetWorkflowSid returns the WorkflowSid field if non-nil, zero value otherwise.

### GetWorkflowSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetWorkflowSidOk() (*string, bool)`

GetWorkflowSidOk returns a tuple with the WorkflowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetWorkflowSid(v string)`

SetWorkflowSid sets WorkflowSid field to given value.

### HasWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasWorkflowSid() bool`

HasWorkflowSid returns a boolean if a field has been set.

### SetWorkflowSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetWorkflowSidNil(b bool)`

 SetWorkflowSidNil sets the value for WorkflowSid to be an explicit nil

### UnsetWorkflowSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetWorkflowSid()`

UnsetWorkflowSid ensures that no value is present for WorkflowSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


