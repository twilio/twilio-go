# TaskrouterV1WorkspaceWorkerWorkerReservation

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
**Url** | Pointer to **NullableString** | The absolute URL of the WorkerReservation resource | [optional] 
**WorkerName** | Pointer to **NullableString** | The friendly_name of the Worker that is reserved | [optional] 
**WorkerSid** | Pointer to **NullableString** | The SID of the reserved Worker resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that this worker is contained within. | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkerWorkerReservation

`func NewTaskrouterV1WorkspaceWorkerWorkerReservation() *TaskrouterV1WorkspaceWorkerWorkerReservation`

NewTaskrouterV1WorkspaceWorkerWorkerReservation instantiates a new TaskrouterV1WorkspaceWorkerWorkerReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWorkerReservationWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWorkerReservationWithDefaults() *TaskrouterV1WorkspaceWorkerWorkerReservation`

NewTaskrouterV1WorkspaceWorkerWorkerReservationWithDefaults instantiates a new TaskrouterV1WorkspaceWorkerWorkerReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetReservationStatus

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetReservationStatus() string`

GetReservationStatus returns the ReservationStatus field if non-nil, zero value otherwise.

### GetReservationStatusOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetReservationStatusOk() (*string, bool)`

GetReservationStatusOk returns a tuple with the ReservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationStatus

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetReservationStatus(v string)`

SetReservationStatus sets ReservationStatus field to given value.

### HasReservationStatus

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasReservationStatus() bool`

HasReservationStatus returns a boolean if a field has been set.

### SetReservationStatusNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetReservationStatusNil(b bool)`

 SetReservationStatusNil sets the value for ReservationStatus to be an explicit nil

### UnsetReservationStatus
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetReservationStatus()`

UnsetReservationStatus ensures that no value is present for ReservationStatus, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkerName

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkerName() string`

GetWorkerName returns the WorkerName field if non-nil, zero value otherwise.

### GetWorkerNameOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkerNameOk() (*string, bool)`

GetWorkerNameOk returns a tuple with the WorkerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerName

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkerName(v string)`

SetWorkerName sets WorkerName field to given value.

### HasWorkerName

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasWorkerName() bool`

HasWorkerName returns a boolean if a field has been set.

### SetWorkerNameNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkerNameNil(b bool)`

 SetWorkerNameNil sets the value for WorkerName to be an explicit nil

### UnsetWorkerName
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetWorkerName()`

UnsetWorkerName ensures that no value is present for WorkerName, not even an explicit nil
### GetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkerSid() string`

GetWorkerSid returns the WorkerSid field if non-nil, zero value otherwise.

### GetWorkerSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkerSidOk() (*string, bool)`

GetWorkerSidOk returns a tuple with the WorkerSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkerSid(v string)`

SetWorkerSid sets WorkerSid field to given value.

### HasWorkerSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasWorkerSid() bool`

HasWorkerSid returns a boolean if a field has been set.

### SetWorkerSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkerSidNil(b bool)`

 SetWorkerSidNil sets the value for WorkerSid to be an explicit nil

### UnsetWorkerSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetWorkerSid()`

UnsetWorkerSid ensures that no value is present for WorkerSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkerWorkerReservation) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


