# TaskrouterV1WorkspaceWorkerWorkerChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssignedTasks** | Pointer to **NullableInt32** | The total number of Tasks assigned to Worker for the TaskChannel type | [optional] 
**Available** | Pointer to **NullableBool** | Whether the Worker should receive Tasks of the TaskChannel type | [optional] 
**AvailableCapacityPercentage** | Pointer to **NullableInt32** | The current available capacity between 0 to 100 for the TaskChannel | [optional] 
**ConfiguredCapacity** | Pointer to **NullableInt32** | The current configured capacity for the WorkerChannel | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskChannelSid** | Pointer to **NullableString** | The SID of the TaskChannel | [optional] 
**TaskChannelUniqueName** | Pointer to **NullableString** | The unique name of the TaskChannel, such as &#39;voice&#39; or &#39;sms&#39; | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the WorkerChannel resource | [optional] 
**WorkerSid** | Pointer to **NullableString** | The SID of the Worker that contains the WorkerChannel | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the WorkerChannel | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkerChannel

`func NewTaskrouterV1WorkspaceWorkerWorkerChannel() *TaskrouterV1WorkspaceWorkerWorkerChannel`

NewTaskrouterV1WorkspaceWorkerWorkerChannel instantiates a new TaskrouterV1WorkspaceWorkerWorkerChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkerChannelWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkerChannelWithDefaults() *TaskrouterV1WorkspaceWorkerWorkerChannel`

NewTaskrouterV1WorkspaceWorkerWorkerChannelWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkerChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssignedTasks

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAssignedTasks() int32`

GetAssignedTasks returns the AssignedTasks field if non-nil, zero value otherwise.

### GetAssignedTasksOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAssignedTasksOk() (*int32, bool)`

GetAssignedTasksOk returns a tuple with the AssignedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTasks

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAssignedTasks(v int32)`

SetAssignedTasks sets AssignedTasks field to given value.

### HasAssignedTasks

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasAssignedTasks() bool`

HasAssignedTasks returns a boolean if a field has been set.

### SetAssignedTasksNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAssignedTasksNil(b bool)`

 SetAssignedTasksNil sets the value for AssignedTasks to be an explicit nil

### UnsetAssignedTasks
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetAssignedTasks()`

UnsetAssignedTasks ensures that no value is present for AssignedTasks, not even an explicit nil
### GetAvailable

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetAvailableCapacityPercentage

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAvailableCapacityPercentage() int32`

GetAvailableCapacityPercentage returns the AvailableCapacityPercentage field if non-nil, zero value otherwise.

### GetAvailableCapacityPercentageOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetAvailableCapacityPercentageOk() (*int32, bool)`

GetAvailableCapacityPercentageOk returns a tuple with the AvailableCapacityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCapacityPercentage

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAvailableCapacityPercentage(v int32)`

SetAvailableCapacityPercentage sets AvailableCapacityPercentage field to given value.

### HasAvailableCapacityPercentage

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasAvailableCapacityPercentage() bool`

HasAvailableCapacityPercentage returns a boolean if a field has been set.

### SetAvailableCapacityPercentageNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetAvailableCapacityPercentageNil(b bool)`

 SetAvailableCapacityPercentageNil sets the value for AvailableCapacityPercentage to be an explicit nil

### UnsetAvailableCapacityPercentage
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetAvailableCapacityPercentage()`

UnsetAvailableCapacityPercentage ensures that no value is present for AvailableCapacityPercentage, not even an explicit nil
### GetConfiguredCapacity

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetConfiguredCapacity() int32`

GetConfiguredCapacity returns the ConfiguredCapacity field if non-nil, zero value otherwise.

### GetConfiguredCapacityOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetConfiguredCapacityOk() (*int32, bool)`

GetConfiguredCapacityOk returns a tuple with the ConfiguredCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredCapacity

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetConfiguredCapacity(v int32)`

SetConfiguredCapacity sets ConfiguredCapacity field to given value.

### HasConfiguredCapacity

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasConfiguredCapacity() bool`

HasConfiguredCapacity returns a boolean if a field has been set.

### SetConfiguredCapacityNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetConfiguredCapacityNil(b bool)`

 SetConfiguredCapacityNil sets the value for ConfiguredCapacity to be an explicit nil

### UnsetConfiguredCapacity
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetConfiguredCapacity()`

UnsetConfiguredCapacity ensures that no value is present for ConfiguredCapacity, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskChannelSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetTaskChannelSid() string`

GetTaskChannelSid returns the TaskChannelSid field if non-nil, zero value otherwise.

### GetTaskChannelSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetTaskChannelSidOk() (*string, bool)`

GetTaskChannelSidOk returns a tuple with the TaskChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskChannelSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetTaskChannelSid(v string)`

SetTaskChannelSid sets TaskChannelSid field to given value.

### HasTaskChannelSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasTaskChannelSid() bool`

HasTaskChannelSid returns a boolean if a field has been set.

### SetTaskChannelSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetTaskChannelSidNil(b bool)`

 SetTaskChannelSidNil sets the value for TaskChannelSid to be an explicit nil

### UnsetTaskChannelSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetTaskChannelSid()`

UnsetTaskChannelSid ensures that no value is present for TaskChannelSid, not even an explicit nil
### GetTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetTaskChannelUniqueName() string`

GetTaskChannelUniqueName returns the TaskChannelUniqueName field if non-nil, zero value otherwise.

### GetTaskChannelUniqueNameOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetTaskChannelUniqueNameOk() (*string, bool)`

GetTaskChannelUniqueNameOk returns a tuple with the TaskChannelUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetTaskChannelUniqueName(v string)`

SetTaskChannelUniqueName sets TaskChannelUniqueName field to given value.

### HasTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasTaskChannelUniqueName() bool`

HasTaskChannelUniqueName returns a boolean if a field has been set.

### SetTaskChannelUniqueNameNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetTaskChannelUniqueNameNil(b bool)`

 SetTaskChannelUniqueNameNil sets the value for TaskChannelUniqueName to be an explicit nil

### UnsetTaskChannelUniqueName
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetTaskChannelUniqueName()`

UnsetTaskChannelUniqueName ensures that no value is present for TaskChannelUniqueName, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetWorkerSid() string`

GetWorkerSid returns the WorkerSid field if non-nil, zero value otherwise.

### GetWorkerSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetWorkerSidOk() (*string, bool)`

GetWorkerSidOk returns a tuple with the WorkerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetWorkerSid(v string)`

SetWorkerSid sets WorkerSid field to given value.

### HasWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasWorkerSid() bool`

HasWorkerSid returns a boolean if a field has been set.

### SetWorkerSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetWorkerSidNil(b bool)`

 SetWorkerSidNil sets the value for WorkerSid to be an explicit nil

### UnsetWorkerSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetWorkerSid()`

UnsetWorkerSid ensures that no value is present for WorkerSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerChannel) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


