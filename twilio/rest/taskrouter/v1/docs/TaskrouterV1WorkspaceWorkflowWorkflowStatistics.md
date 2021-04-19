# TaskrouterV1WorkspaceWorkflowWorkflowStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Cumulative** | Pointer to **map[string]interface{}** | An object that contains the cumulative statistics for the Workflow | [optional] 
**Realtime** | Pointer to **map[string]interface{}** | An object that contains the real-time statistics for the Workflow | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workflow statistics resource | [optional] 
**WorkflowSid** | Pointer to **NullableString** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Workflow | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkflowWorkflowStatistics

`func NewTaskrouterV1WorkspaceWorkflowWorkflowStatistics() *TaskrouterV1WorkspaceWorkflowWorkflowStatistics`

NewTaskrouterV1WorkspaceWorkflowWorkflowStatistics instantiates a new TaskrouterV1WorkspaceWorkflowWorkflowStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkflowWorkflowStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkflowWorkflowStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkflowWorkflowStatistics`

NewTaskrouterV1WorkspaceWorkflowWorkflowStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkflowWorkflowStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCumulative

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetCumulative() map[string]interface{}`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetCumulativeOk() (*map[string]interface{}, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetCumulative(v map[string]interface{})`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetRealtime

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetRealtime() map[string]interface{}`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetRealtimeOk() (*map[string]interface{}, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetRealtime(v map[string]interface{})`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### SetRealtimeNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetRealtimeNil(b bool)`

 SetRealtimeNil sets the value for Realtime to be an explicit nil

### UnsetRealtime
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetRealtime()`

UnsetRealtime ensures that no value is present for Realtime, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetWorkflowSid() string`

GetWorkflowSid returns the WorkflowSid field if non-nil, zero value otherwise.

### GetWorkflowSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetWorkflowSidOk() (*string, bool)`

GetWorkflowSidOk returns a tuple with the WorkflowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetWorkflowSid(v string)`

SetWorkflowSid sets WorkflowSid field to given value.

### HasWorkflowSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasWorkflowSid() bool`

HasWorkflowSid returns a boolean if a field has been set.

### SetWorkflowSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetWorkflowSidNil(b bool)`

 SetWorkflowSidNil sets the value for WorkflowSid to be an explicit nil

### UnsetWorkflowSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetWorkflowSid()`

UnsetWorkflowSid ensures that no value is present for WorkflowSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkflowWorkflowStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


