# TaskrouterV1WorkspaceTaskTaskReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**ReservationStatus** | Pointer to **NullableString** | The current status of the reservation | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the reserved Task resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskReservation reservation | [optional] 
**WorkerName** | Pointer to **NullableString** | The friendly_name of the Worker that is reserved | [optional] 
**WorkerSid** | Pointer to **NullableString** | The SID of the reserved Worker resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that this task is contained within. | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTaskTaskReservation

`func NewTaskrouterV1WorkspaceTaskTaskReservation() *TaskrouterV1WorkspaceTaskTaskReservation`

NewTaskrouterV1WorkspaceTaskTaskReservation instantiates a new TaskrouterV1WorkspaceTaskTaskReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskTaskReservationWithDefaults

`func NewTaskrouterV1WorkspaceTaskTaskReservationWithDefaults() *TaskrouterV1WorkspaceTaskTaskReservation`

NewTaskrouterV1WorkspaceTaskTaskReservationWithDefaults instantiates a new TaskrouterV1WorkspaceTaskTaskReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetReservationStatus

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetReservationStatus() string`

GetReservationStatus returns the ReservationStatus field if non-nil, zero value otherwise.

### GetReservationStatusOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetReservationStatusOk() (*string, bool)`

GetReservationStatusOk returns a tuple with the ReservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationStatus

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetReservationStatus(v string)`

SetReservationStatus sets ReservationStatus field to given value.

### HasReservationStatus

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasReservationStatus() bool`

HasReservationStatus returns a boolean if a field has been set.

### SetReservationStatusNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetReservationStatusNil(b bool)`

 SetReservationStatusNil sets the value for ReservationStatus to be an explicit nil

### UnsetReservationStatus
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetReservationStatus()`

UnsetReservationStatus ensures that no value is present for ReservationStatus, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkerName

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.

### HasWorkerName

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasWorkerName() bool`

HasWorkerName returns a boolean if a field has been set.

### SetWorkerNameNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkerNameNil(b bool)`

 SetWorkerNameNil sets the value for WorkerName to be an explicit nil

### UnsetWorkerName
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetWorkerName()`

UnsetWorkerName ensures that no value is present for WorkerName, not even an explicit nil
### GetWorkerSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkerSid() string`

GetWorkerSid returns the WorkerSid field if non-nil, zero value otherwise.

### GetWorkerSidOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkerSidOk() (*string, bool)`

GetWorkerSidOk returns a tuple with the WorkerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkerSid(v string)`

SetWorkerSid sets WorkerSid field to given value.

### HasWorkerSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasWorkerSid() bool`

HasWorkerSid returns a boolean if a field has been set.

### SetWorkerSidNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkerSidNil(b bool)`

 SetWorkerSidNil sets the value for WorkerSid to be an explicit nil

### UnsetWorkerSid
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetWorkerSid()`

UnsetWorkerSid ensures that no value is present for WorkerSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTaskTaskReservation) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTaskTaskReservation) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


