# TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Cumulative** | Pointer to **map[string]interface{}** | An object that contains the cumulative statistics for the TaskQueues | [optional] 
**Realtime** | Pointer to **map[string]interface{}** | An object that contains the real-time statistics for the TaskQueues | [optional] 
**TaskQueueSid** | Pointer to **NullableString** | The SID of the TaskQueue from which these statistics were calculated | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the TaskQueues | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics() *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatisticsWithDefaults() *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics`

NewTaskrouterV1WorkspaceTaskQueueTaskQueuesStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCumulative

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetCumulative() map[string]interface{}`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetCumulativeOk() (*map[string]interface{}, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetCumulative(v map[string]interface{})`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetRealtime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetRealtime() map[string]interface{}`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetRealtimeOk() (*map[string]interface{}, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetRealtime(v map[string]interface{})`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### SetRealtimeNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetRealtimeNil(b bool)`

 SetRealtimeNil sets the value for Realtime to be an explicit nil

### UnsetRealtime
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) UnsetRealtime()`

UnsetRealtime ensures that no value is present for Realtime, not even an explicit nil
### GetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetTaskQueueSid() string`

GetTaskQueueSid returns the TaskQueueSid field if non-nil, zero value otherwise.

### GetTaskQueueSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetTaskQueueSidOk() (*string, bool)`

GetTaskQueueSidOk returns a tuple with the TaskQueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetTaskQueueSid(v string)`

SetTaskQueueSid sets TaskQueueSid field to given value.

### HasTaskQueueSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) HasTaskQueueSid() bool`

HasTaskQueueSid returns a boolean if a field has been set.

### SetTaskQueueSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetTaskQueueSidNil(b bool)`

 SetTaskQueueSidNil sets the value for TaskQueueSid to be an explicit nil

### UnsetTaskQueueSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) UnsetTaskQueueSid()`

UnsetTaskQueueSid ensures that no value is present for TaskQueueSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTaskQueueTaskQueuesStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


