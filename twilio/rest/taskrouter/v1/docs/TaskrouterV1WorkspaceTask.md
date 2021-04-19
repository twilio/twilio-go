# TaskrouterV1WorkspaceTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Addons** | Pointer to **NullableString** | An object that contains the addon data for all installed addons | [optional] 
**Age** | Pointer to **NullableInt32** | The number of seconds since the Task was created | [optional] 
**AssignmentStatus** | Pointer to **NullableString** | The current status of the Task&#39;s assignment | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON string with custom attributes of the work | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Priority** | Pointer to **NullableInt32** | Retrieve the list of all Tasks in the Workspace with the specified priority | [optional] 
**Reason** | Pointer to **NullableString** | The reason the Task was canceled or completed | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskChannelSid** | Pointer to **NullableString** | The SID of the TaskChannel | [optional] 
**TaskChannelUniqueName** | Pointer to **NullableString** | The unique name of the TaskChannel | [optional] 
**TaskQueueEnteredDate** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Task entered the TaskQueue. | [optional] 
**TaskQueueFriendlyName** | Pointer to **NullableString** | The friendly name of the TaskQueue | [optional] 
**TaskQueueSid** | Pointer to **NullableString** | The SID of the TaskQueue | [optional] 
**Timeout** | Pointer to **NullableInt32** | The amount of time in seconds that the Task can live before being assigned | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Task resource | [optional] 
**WorkflowFriendlyName** | Pointer to **NullableString** | The friendly name of the Workflow that is controlling the Task | [optional] 
**WorkflowSid** | Pointer to **NullableString** | The SID of the Workflow that is controlling the Task | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Task | [optional] 

## Methods

### NewTaskrouterV1WorkspaceTask

`func NewTaskrouterV1WorkspaceTask() *TaskrouterV1WorkspaceTask`

NewTaskrouterV1WorkspaceTask instantiates a new TaskrouterV1WorkspaceTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceTaskWithDefaults

`func NewTaskrouterV1WorkspaceTaskWithDefaults() *TaskrouterV1WorkspaceTask`

NewTaskrouterV1WorkspaceTaskWithDefaults instantiates a new TaskrouterV1WorkspaceTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceTask) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceTask) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceTask) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceTask) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceTask) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceTask) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddons

`func (o *TaskrouterV1WorkspaceTask) GetAddons() string`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *TaskrouterV1WorkspaceTask) GetAddonsOk() (*string, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *TaskrouterV1WorkspaceTask) SetAddons(v string)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *TaskrouterV1WorkspaceTask) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### SetAddonsNil

`func (o *TaskrouterV1WorkspaceTask) SetAddonsNil(b bool)`

 SetAddonsNil sets the value for Addons to be an explicit nil

### UnsetAddons
`func (o *TaskrouterV1WorkspaceTask) UnsetAddons()`

UnsetAddons ensures that no value is present for Addons, not even an explicit nil
### GetAge

`func (o *TaskrouterV1WorkspaceTask) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *TaskrouterV1WorkspaceTask) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *TaskrouterV1WorkspaceTask) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *TaskrouterV1WorkspaceTask) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *TaskrouterV1WorkspaceTask) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *TaskrouterV1WorkspaceTask) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAssignmentStatus

`func (o *TaskrouterV1WorkspaceTask) GetAssignmentStatus() string`

GetAssignmentStatus returns the AssignmentStatus field if non-nil, zero value otherwise.

### GetAssignmentStatusOk

`func (o *TaskrouterV1WorkspaceTask) GetAssignmentStatusOk() (*string, bool)`

GetAssignmentStatusOk returns a tuple with the AssignmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentStatus

`func (o *TaskrouterV1WorkspaceTask) SetAssignmentStatus(v string)`

SetAssignmentStatus sets AssignmentStatus field to given value.

### HasAssignmentStatus

`func (o *TaskrouterV1WorkspaceTask) HasAssignmentStatus() bool`

HasAssignmentStatus returns a boolean if a field has been set.

### SetAssignmentStatusNil

`func (o *TaskrouterV1WorkspaceTask) SetAssignmentStatusNil(b bool)`

 SetAssignmentStatusNil sets the value for AssignmentStatus to be an explicit nil

### UnsetAssignmentStatus
`func (o *TaskrouterV1WorkspaceTask) UnsetAssignmentStatus()`

UnsetAssignmentStatus ensures that no value is present for AssignmentStatus, not even an explicit nil
### GetAttributes

`func (o *TaskrouterV1WorkspaceTask) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaskrouterV1WorkspaceTask) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaskrouterV1WorkspaceTask) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TaskrouterV1WorkspaceTask) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TaskrouterV1WorkspaceTask) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TaskrouterV1WorkspaceTask) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceTask) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceTask) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceTask) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceTask) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceTask) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceTask) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceTask) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceTask) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceTask) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceTask) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceTask) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceTask) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceTask) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceTask) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPriority

`func (o *TaskrouterV1WorkspaceTask) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TaskrouterV1WorkspaceTask) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TaskrouterV1WorkspaceTask) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TaskrouterV1WorkspaceTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TaskrouterV1WorkspaceTask) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TaskrouterV1WorkspaceTask) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetReason

`func (o *TaskrouterV1WorkspaceTask) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaskrouterV1WorkspaceTask) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaskrouterV1WorkspaceTask) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaskrouterV1WorkspaceTask) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *TaskrouterV1WorkspaceTask) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TaskrouterV1WorkspaceTask) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceTask) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceTask) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceTask) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceTask) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceTask) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceTask) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskChannelSid

`func (o *TaskrouterV1WorkspaceTask) GetTaskChannelSid() string`

GetTaskChannelSid returns the TaskChannelSid field if non-nil, zero value otherwise.

### GetTaskChannelSidOk

`func (o *TaskrouterV1WorkspaceTask) GetTaskChannelSidOk() (*string, bool)`

GetTaskChannelSidOk returns a tuple with the TaskChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskChannelSid

`func (o *TaskrouterV1WorkspaceTask) SetTaskChannelSid(v string)`

SetTaskChannelSid sets TaskChannelSid field to given value.

### HasTaskChannelSid

`func (o *TaskrouterV1WorkspaceTask) HasTaskChannelSid() bool`

HasTaskChannelSid returns a boolean if a field has been set.

### SetTaskChannelSidNil

`func (o *TaskrouterV1WorkspaceTask) SetTaskChannelSidNil(b bool)`

 SetTaskChannelSidNil sets the value for TaskChannelSid to be an explicit nil

### UnsetTaskChannelSid
`func (o *TaskrouterV1WorkspaceTask) UnsetTaskChannelSid()`

UnsetTaskChannelSid ensures that no value is present for TaskChannelSid, not even an explicit nil
### GetTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceTask) GetTaskChannelUniqueName() string`

GetTaskChannelUniqueName returns the TaskChannelUniqueName field if non-nil, zero value otherwise.

### GetTaskChannelUniqueNameOk

`func (o *TaskrouterV1WorkspaceTask) GetTaskChannelUniqueNameOk() (*string, bool)`

GetTaskChannelUniqueNameOk returns a tuple with the TaskChannelUniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceTask) SetTaskChannelUniqueName(v string)`

SetTaskChannelUniqueName sets TaskChannelUniqueName field to given value.

### HasTaskChannelUniqueName

`func (o *TaskrouterV1WorkspaceTask) HasTaskChannelUniqueName() bool`

HasTaskChannelUniqueName returns a boolean if a field has been set.

### SetTaskChannelUniqueNameNil

`func (o *TaskrouterV1WorkspaceTask) SetTaskChannelUniqueNameNil(b bool)`

 SetTaskChannelUniqueNameNil sets the value for TaskChannelUniqueName to be an explicit nil

### UnsetTaskChannelUniqueName
`func (o *TaskrouterV1WorkspaceTask) UnsetTaskChannelUniqueName()`

UnsetTaskChannelUniqueName ensures that no value is present for TaskChannelUniqueName, not even an explicit nil
### GetTaskQueueEnteredDate

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueEnteredDate() time.Time`

GetTaskQueueEnteredDate returns the TaskQueueEnteredDate field if non-nil, zero value otherwise.

### GetTaskQueueEnteredDateOk

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueEnteredDateOk() (*time.Time, bool)`

GetTaskQueueEnteredDateOk returns a tuple with the TaskQueueEnteredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueEnteredDate

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueEnteredDate(v time.Time)`

SetTaskQueueEnteredDate sets TaskQueueEnteredDate field to given value.

### HasTaskQueueEnteredDate

`func (o *TaskrouterV1WorkspaceTask) HasTaskQueueEnteredDate() bool`

HasTaskQueueEnteredDate returns a boolean if a field has been set.

### SetTaskQueueEnteredDateNil

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueEnteredDateNil(b bool)`

 SetTaskQueueEnteredDateNil sets the value for TaskQueueEnteredDate to be an explicit nil

### UnsetTaskQueueEnteredDate
`func (o *TaskrouterV1WorkspaceTask) UnsetTaskQueueEnteredDate()`

UnsetTaskQueueEnteredDate ensures that no value is present for TaskQueueEnteredDate, not even an explicit nil
### GetTaskQueueFriendlyName

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueFriendlyName() string`

GetTaskQueueFriendlyName returns the TaskQueueFriendlyName field if non-nil, zero value otherwise.

### GetTaskQueueFriendlyNameOk

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueFriendlyNameOk() (*string, bool)`

GetTaskQueueFriendlyNameOk returns a tuple with the TaskQueueFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueFriendlyName

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueFriendlyName(v string)`

SetTaskQueueFriendlyName sets TaskQueueFriendlyName field to given value.

### HasTaskQueueFriendlyName

`func (o *TaskrouterV1WorkspaceTask) HasTaskQueueFriendlyName() bool`

HasTaskQueueFriendlyName returns a boolean if a field has been set.

### SetTaskQueueFriendlyNameNil

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueFriendlyNameNil(b bool)`

 SetTaskQueueFriendlyNameNil sets the value for TaskQueueFriendlyName to be an explicit nil

### UnsetTaskQueueFriendlyName
`func (o *TaskrouterV1WorkspaceTask) UnsetTaskQueueFriendlyName()`

UnsetTaskQueueFriendlyName ensures that no value is present for TaskQueueFriendlyName, not even an explicit nil
### GetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueSid() string`

GetTaskQueueSid returns the TaskQueueSid field if non-nil, zero value otherwise.

### GetTaskQueueSidOk

`func (o *TaskrouterV1WorkspaceTask) GetTaskQueueSidOk() (*string, bool)`

GetTaskQueueSidOk returns a tuple with the TaskQueueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueSid

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueSid(v string)`

SetTaskQueueSid sets TaskQueueSid field to given value.

### HasTaskQueueSid

`func (o *TaskrouterV1WorkspaceTask) HasTaskQueueSid() bool`

HasTaskQueueSid returns a boolean if a field has been set.

### SetTaskQueueSidNil

`func (o *TaskrouterV1WorkspaceTask) SetTaskQueueSidNil(b bool)`

 SetTaskQueueSidNil sets the value for TaskQueueSid to be an explicit nil

### UnsetTaskQueueSid
`func (o *TaskrouterV1WorkspaceTask) UnsetTaskQueueSid()`

UnsetTaskQueueSid ensures that no value is present for TaskQueueSid, not even an explicit nil
### GetTimeout

`func (o *TaskrouterV1WorkspaceTask) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TaskrouterV1WorkspaceTask) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TaskrouterV1WorkspaceTask) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TaskrouterV1WorkspaceTask) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *TaskrouterV1WorkspaceTask) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *TaskrouterV1WorkspaceTask) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceTask) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceTask) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceTask) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceTask) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkflowFriendlyName

`func (o *TaskrouterV1WorkspaceTask) GetWorkflowFriendlyName() string`

GetWorkflowFriendlyName returns the WorkflowFriendlyName field if non-nil, zero value otherwise.

### GetWorkflowFriendlyNameOk

`func (o *TaskrouterV1WorkspaceTask) GetWorkflowFriendlyNameOk() (*string, bool)`

GetWorkflowFriendlyNameOk returns a tuple with the WorkflowFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFriendlyName

`func (o *TaskrouterV1WorkspaceTask) SetWorkflowFriendlyName(v string)`

SetWorkflowFriendlyName sets WorkflowFriendlyName field to given value.

### HasWorkflowFriendlyName

`func (o *TaskrouterV1WorkspaceTask) HasWorkflowFriendlyName() bool`

HasWorkflowFriendlyName returns a boolean if a field has been set.

### SetWorkflowFriendlyNameNil

`func (o *TaskrouterV1WorkspaceTask) SetWorkflowFriendlyNameNil(b bool)`

 SetWorkflowFriendlyNameNil sets the value for WorkflowFriendlyName to be an explicit nil

### UnsetWorkflowFriendlyName
`func (o *TaskrouterV1WorkspaceTask) UnsetWorkflowFriendlyName()`

UnsetWorkflowFriendlyName ensures that no value is present for WorkflowFriendlyName, not even an explicit nil
### GetWorkflowSid

`func (o *TaskrouterV1WorkspaceTask) GetWorkflowSid() string`

GetWorkflowSid returns the WorkflowSid field if non-nil, zero value otherwise.

### GetWorkflowSidOk

`func (o *TaskrouterV1WorkspaceTask) GetWorkflowSidOk() (*string, bool)`

GetWorkflowSidOk returns a tuple with the WorkflowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSid

`func (o *TaskrouterV1WorkspaceTask) SetWorkflowSid(v string)`

SetWorkflowSid sets WorkflowSid field to given value.

### HasWorkflowSid

`func (o *TaskrouterV1WorkspaceTask) HasWorkflowSid() bool`

HasWorkflowSid returns a boolean if a field has been set.

### SetWorkflowSidNil

`func (o *TaskrouterV1WorkspaceTask) SetWorkflowSidNil(b bool)`

 SetWorkflowSidNil sets the value for WorkflowSid to be an explicit nil

### UnsetWorkflowSid
`func (o *TaskrouterV1WorkspaceTask) UnsetWorkflowSid()`

UnsetWorkflowSid ensures that no value is present for WorkflowSid, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTask) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceTask) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceTask) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceTask) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceTask) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceTask) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


