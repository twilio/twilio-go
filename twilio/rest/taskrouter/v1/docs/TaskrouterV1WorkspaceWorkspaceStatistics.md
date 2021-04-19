# TaskrouterV1WorkspaceWorkspaceStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Cumulative** | Pointer to **map[string]interface{}** | An object that contains the cumulative statistics for the Workspace | [optional] 
**Realtime** | Pointer to **map[string]interface{}** | n object that contains the real-time statistics for the Workspace | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workspace statistics resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkspaceStatistics

`func NewTaskrouterV1WorkspaceWorkspaceStatistics() *TaskrouterV1WorkspaceWorkspaceStatistics`

NewTaskrouterV1WorkspaceWorkspaceStatistics instantiates a new TaskrouterV1WorkspaceWorkspaceStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkspaceStatisticsWithDefaults

`func NewTaskrouterV1WorkspaceWorkspaceStatisticsWithDefaults() *TaskrouterV1WorkspaceWorkspaceStatistics`

NewTaskrouterV1WorkspaceWorkspaceStatisticsWithDefaults instantiates a new TaskrouterV1WorkspaceWorkspaceStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCumulative

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetCumulative() map[string]interface{}`

GetCumulative returns the Cumulative field if non-nil, zero value otherwise.

### GetCumulativeOk

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetCumulativeOk() (*map[string]interface{}, bool)`

GetCumulativeOk returns a tuple with the Cumulative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulative

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetCumulative(v map[string]interface{})`

SetCumulative sets Cumulative field to given value.

### HasCumulative

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) HasCumulative() bool`

HasCumulative returns a boolean if a field has been set.

### SetCumulativeNil

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetCumulativeNil(b bool)`

 SetCumulativeNil sets the value for Cumulative to be an explicit nil

### UnsetCumulative
`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) UnsetCumulative()`

UnsetCumulative ensures that no value is present for Cumulative, not even an explicit nil
### GetRealtime

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetRealtime() map[string]interface{}`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetRealtimeOk() (*map[string]interface{}, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetRealtime(v map[string]interface{})`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### SetRealtimeNil

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetRealtimeNil(b bool)`

 SetRealtimeNil sets the value for Realtime to be an explicit nil

### UnsetRealtime
`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) UnsetRealtime()`

UnsetRealtime ensures that no value is present for Realtime, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkspaceStatistics) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


