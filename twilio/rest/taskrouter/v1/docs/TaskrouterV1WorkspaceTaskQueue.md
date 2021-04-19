# TaskrouterV1WorkspaceTaskQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssignmentActivityName** | Pointer to **NullableString** | The name of the Activity to assign Workers when a task is assigned for them | [optional] 
**AssignmentActivitySid** | Pointer to **NullableString** | The SID of the Activity to assign Workers when a task is assigned for them | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**MaxReservedWorkers** | Pointer to **NullableInt32** | The maximum number of Workers to reserve | [optional] 
**ReservationActivityName** | Pointer to **NullableString** | The name of the Activity to assign Workers once a task is reserved for them | [optional] 
**ReservationActivitySid** | Pointer to **NullableString** | The SID of the Activity to assign Workers once a task is reserved for them | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TargetWorkers** | Pointer to **NullableString** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue | [optional] 
**TaskOrder** | Pointer to **NullableString** | How Tasks will be assigned to Workers | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskQueue resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the TaskQueue | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTaskQueue

`func NewTaskrouterV1WorkspaceTaskQueue() *TaskrouterV1WorkspaceTaskQueue`

NewTaskrouterV1WorkspaceTaskQueue instantiates a new TaskrouterV1WorkspaceTaskQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskQueueWithDefaults

`func NewTaskrouterV1WorkspaceTaskQueueWithDefaults() *TaskrouterV1WorkspaceTaskQueue`

NewTaskrouterV1WorkspaceTaskQueueWithDefaults instantiates a new TaskrouterV1WorkspaceTaskQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTaskQueue) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssignmentActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAssignmentActivityName() string`

GetAssignmentActivityName returns the AssignmentActivityName field if non-nil, zero value otherwise.

### GetAssignmentActivityNameOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAssignmentActivityNameOk() (*string, bool)`

GetAssignmentActivityNameOk returns a tuple with the AssignmentActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAssignmentActivityName(v string)`

SetAssignmentActivityName sets AssignmentActivityName field to given value.

### HasAssignmentActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) HasAssignmentActivityName() bool`

HasAssignmentActivityName returns a boolean if a field has been set.

### SetAssignmentActivityNameNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAssignmentActivityNameNil(b bool)`

 SetAssignmentActivityNameNil sets the value for AssignmentActivityName to be an explicit nil

### UnsetAssignmentActivityName
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetAssignmentActivityName()`

UnsetAssignmentActivityName ensures that no value is present for AssignmentActivityName, not even an explicit nil
### GetAssignmentActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAssignmentActivitySid() string`

GetAssignmentActivitySid returns the AssignmentActivitySid field if non-nil, zero value otherwise.

### GetAssignmentActivitySidOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetAssignmentActivitySidOk() (*string, bool)`

GetAssignmentActivitySidOk returns a tuple with the AssignmentActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAssignmentActivitySid(v string)`

SetAssignmentActivitySid sets AssignmentActivitySid field to given value.

### HasAssignmentActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) HasAssignmentActivitySid() bool`

HasAssignmentActivitySid returns a boolean if a field has been set.

### SetAssignmentActivitySidNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetAssignmentActivitySidNil(b bool)`

 SetAssignmentActivitySidNil sets the value for AssignmentActivitySid to be an explicit nil

### UnsetAssignmentActivitySid
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetAssignmentActivitySid()`

UnsetAssignmentActivitySid ensures that no value is present for AssignmentActivitySid, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceTaskQueue) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceTaskQueue) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceTaskQueue) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceTaskQueue) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceTaskQueue) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceTaskQueue) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *TaskrouterV1WorkspaceTaskQueue) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TaskrouterV1WorkspaceTaskQueue) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TaskrouterV1WorkspaceTaskQueue) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceTaskQueue) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceTaskQueue) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceTaskQueue) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMaxReservedWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) GetMaxReservedWorkers() int32`

GetMaxReservedWorkers returns the MaxReservedWorkers field if non-nil, zero value otherwise.

### GetMaxReservedWorkersOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetMaxReservedWorkersOk() (*int32, bool)`

GetMaxReservedWorkersOk returns a tuple with the MaxReservedWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReservedWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) SetMaxReservedWorkers(v int32)`

SetMaxReservedWorkers sets MaxReservedWorkers field to given value.

### HasMaxReservedWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) HasMaxReservedWorkers() bool`

HasMaxReservedWorkers returns a boolean if a field has been set.

### SetMaxReservedWorkersNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetMaxReservedWorkersNil(b bool)`

 SetMaxReservedWorkersNil sets the value for MaxReservedWorkers to be an explicit nil

### UnsetMaxReservedWorkers
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetMaxReservedWorkers()`

UnsetMaxReservedWorkers ensures that no value is present for MaxReservedWorkers, not even an explicit nil
### GetReservationActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) GetReservationActivityName() string`

GetReservationActivityName returns the ReservationActivityName field if non-nil, zero value otherwise.

### GetReservationActivityNameOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetReservationActivityNameOk() (*string, bool)`

GetReservationActivityNameOk returns a tuple with the ReservationActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) SetReservationActivityName(v string)`

SetReservationActivityName sets ReservationActivityName field to given value.

### HasReservationActivityName

`func (o *TaskrouterV1WorkspaceTaskQueue) HasReservationActivityName() bool`

HasReservationActivityName returns a boolean if a field has been set.

### SetReservationActivityNameNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetReservationActivityNameNil(b bool)`

 SetReservationActivityNameNil sets the value for ReservationActivityName to be an explicit nil

### UnsetReservationActivityName
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetReservationActivityName()`

UnsetReservationActivityName ensures that no value is present for ReservationActivityName, not even an explicit nil
### GetReservationActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) GetReservationActivitySid() string`

GetReservationActivitySid returns the ReservationActivitySid field if non-nil, zero value otherwise.

### GetReservationActivitySidOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetReservationActivitySidOk() (*string, bool)`

GetReservationActivitySidOk returns a tuple with the ReservationActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) SetReservationActivitySid(v string)`

SetReservationActivitySid sets ReservationActivitySid field to given value.

### HasReservationActivitySid

`func (o *TaskrouterV1WorkspaceTaskQueue) HasReservationActivitySid() bool`

HasReservationActivitySid returns a boolean if a field has been set.

### SetReservationActivitySidNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetReservationActivitySidNil(b bool)`

 SetReservationActivitySidNil sets the value for ReservationActivitySid to be an explicit nil

### UnsetReservationActivitySid
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetReservationActivitySid()`

UnsetReservationActivitySid ensures that no value is present for ReservationActivitySid, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceTaskQueue) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceTaskQueue) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceTaskQueue) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTargetWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) GetTargetWorkers() string`

GetTargetWorkers returns the TargetWorkers field if non-nil, zero value otherwise.

### GetTargetWorkersOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetTargetWorkersOk() (*string, bool)`

GetTargetWorkersOk returns a tuple with the TargetWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) SetTargetWorkers(v string)`

SetTargetWorkers sets TargetWorkers field to given value.

### HasTargetWorkers

`func (o *TaskrouterV1WorkspaceTaskQueue) HasTargetWorkers() bool`

HasTargetWorkers returns a boolean if a field has been set.

### SetTargetWorkersNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetTargetWorkersNil(b bool)`

 SetTargetWorkersNil sets the value for TargetWorkers to be an explicit nil

### UnsetTargetWorkers
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetTargetWorkers()`

UnsetTargetWorkers ensures that no value is present for TargetWorkers, not even an explicit nil
### GetTaskOrder

`func (o *TaskrouterV1WorkspaceTaskQueue) GetTaskOrder() string`

GetTaskOrder returns the TaskOrder field if non-nil, zero value otherwise.

### GetTaskOrderOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetTaskOrderOk() (*string, bool)`

GetTaskOrderOk returns a tuple with the TaskOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOrder

`func (o *TaskrouterV1WorkspaceTaskQueue) SetTaskOrder(v string)`

SetTaskOrder sets TaskOrder field to given value.

### HasTaskOrder

`func (o *TaskrouterV1WorkspaceTaskQueue) HasTaskOrder() bool`

HasTaskOrder returns a boolean if a field has been set.

### SetTaskOrderNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetTaskOrderNil(b bool)`

 SetTaskOrderNil sets the value for TaskOrder to be an explicit nil

### UnsetTaskOrder
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetTaskOrder()`

UnsetTaskOrder ensures that no value is present for TaskOrder, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceTaskQueue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceTaskQueue) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceTaskQueue) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueue) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTaskQueue) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueue) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskQueue) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTaskQueue) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTaskQueue) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


