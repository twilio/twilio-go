# TaskrouterV1WorkspaceWorkerWorkerStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Cumulative** | Pointer to **map[string]interface{}** | An object that contains the cumulative statistics for the Worker | [optional] 
**Realtime** | Pointer to **map[string]interface{}** | An object that contains the real-time statistics for the Worker | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Worker statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Worker | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkerStatistics

`func NewTaskrouterV1WorkspaceWorkerWorkerStatistics() *TaskrouterV1WorkspaceWorkerWorkerStatistics`

NewTaskrouterV1WorkspaceWorkerWorkerStatistics instantiates a new TaskrouterV1WorkspaceWorkerWorkerStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkerStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkerStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkerWorkerStatistics`

NewTaskrouterV1WorkspaceWorkerWorkerStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkerStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetCumulative() map[string]interface{}`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetCumulativeOk() (*map[string]interface{}, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetCumulative(v map[string]interface{})`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetRealtime

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetRealtime() map[string]interface{}`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetRealtimeOk() (*map[string]interface{}, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetRealtime(v map[string]interface{})`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### SetRealtimeNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetRealtimeNil(b bool)`

 SetRealtimeNil sets the value for Realtime to be an explicit nil

### UnsetRealtime
`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) UnsetRealtime()`

UnsetRealtime ensures that no value is present for Realtime, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


