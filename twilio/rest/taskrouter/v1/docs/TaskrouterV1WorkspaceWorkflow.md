# TaskrouterV1WorkspaceWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssignmentCallbackUrl** | Pointer to **NullableString** | The URL that we call when a task managed by the Workflow is assigned to a Worker | [optional] 
**Configuration** | Pointer to **NullableString** | A JSON string that contains the Workflow&#39;s configuration | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DocumentContentType** | Pointer to **NullableString** | The MIME type of the document | [optional] 
**FallbackAssignmentCallbackUrl** | Pointer to **NullableString** | The URL that we call when a call to the &#x60;assignment_callback_url&#x60; fails | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the Workflow resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskReservationTimeout** | Pointer to **NullableInt32** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workflow resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Workflow | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorkflow

`func NewTaskrouterV1WorkspaceWorkflow() *TaskrouterV1WorkspaceWorkflow`

NewTaskrouterV1WorkspaceWorkflow instantiates a new TaskrouterV1WorkspaceWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkflowWithDefaults

`func NewTaskrouterV1WorkspaceWorkflowWithDefaults() *TaskrouterV1WorkspaceWorkflow`

NewTaskrouterV1WorkspaceWorkflowWithDefaults instantiates a new TaskrouterV1WorkspaceWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflow) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorkflow) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorkflow) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) GetAssignmentCallbackUrl() string`

GetAssignmentCallbackUrl returns the AssignmentCallbackUrl field if non-nil, zero value otherwise.

### GetAssignmentCallbackUrlOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetAssignmentCallbackUrlOk() (*string, bool)`

GetAssignmentCallbackUrlOk returns a tuple with the AssignmentCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) SetAssignmentCallbackUrl(v string)`

SetAssignmentCallbackUrl sets AssignmentCallbackUrl field to given value.

### HasAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) HasAssignmentCallbackUrl() bool`

HasAssignmentCallbackUrl returns a boolean if a field has been set.

### SetAssignmentCallbackUrlNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetAssignmentCallbackUrlNil(b bool)`

 SetAssignmentCallbackUrlNil sets the value for AssignmentCallbackUrl to be an explicit nil

### UnsetAssignmentCallbackUrl
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetAssignmentCallbackUrl()`

UnsetAssignmentCallbackUrl ensures that no value is present for AssignmentCallbackUrl, not even an explicit nil
### GetConfiguration

`func (o *TaskrouterV1WorkspaceWorkflow) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *TaskrouterV1WorkspaceWorkflow) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *TaskrouterV1WorkspaceWorkflow) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceWorkflow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceWorkflow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceWorkflow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkflow) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceWorkflow) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceWorkflow) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDocumentContentType

`func (o *TaskrouterV1WorkspaceWorkflow) GetDocumentContentType() string`

GetDocumentContentType returns the DocumentContentType field if non-nil, zero value otherwise.

### GetDocumentContentTypeOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetDocumentContentTypeOk() (*string, bool)`

GetDocumentContentTypeOk returns a tuple with the DocumentContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentContentType

`func (o *TaskrouterV1WorkspaceWorkflow) SetDocumentContentType(v string)`

SetDocumentContentType sets DocumentContentType field to given value.

### HasDocumentContentType

`func (o *TaskrouterV1WorkspaceWorkflow) HasDocumentContentType() bool`

HasDocumentContentType returns a boolean if a field has been set.

### SetDocumentContentTypeNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetDocumentContentTypeNil(b bool)`

 SetDocumentContentTypeNil sets the value for DocumentContentType to be an explicit nil

### UnsetDocumentContentType
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetDocumentContentType()`

UnsetDocumentContentType ensures that no value is present for DocumentContentType, not even an explicit nil
### GetFallbackAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) GetFallbackAssignmentCallbackUrl() string`

GetFallbackAssignmentCallbackUrl returns the FallbackAssignmentCallbackUrl field if non-nil, zero value otherwise.

### GetFallbackAssignmentCallbackUrlOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetFallbackAssignmentCallbackUrlOk() (*string, bool)`

GetFallbackAssignmentCallbackUrlOk returns a tuple with the FallbackAssignmentCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) SetFallbackAssignmentCallbackUrl(v string)`

SetFallbackAssignmentCallbackUrl sets FallbackAssignmentCallbackUrl field to given value.

### HasFallbackAssignmentCallbackUrl

`func (o *TaskrouterV1WorkspaceWorkflow) HasFallbackAssignmentCallbackUrl() bool`

HasFallbackAssignmentCallbackUrl returns a boolean if a field has been set.

### SetFallbackAssignmentCallbackUrlNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetFallbackAssignmentCallbackUrlNil(b bool)`

 SetFallbackAssignmentCallbackUrlNil sets the value for FallbackAssignmentCallbackUrl to be an explicit nil

### UnsetFallbackAssignmentCallbackUrl
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetFallbackAssignmentCallbackUrl()`

UnsetFallbackAssignmentCallbackUrl ensures that no value is present for FallbackAssignmentCallbackUrl, not even an explicit nil
### GetFriendlyName

`func (o *TaskrouterV1WorkspaceWorkflow) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TaskrouterV1WorkspaceWorkflow) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TaskrouterV1WorkspaceWorkflow) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceWorkflow) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceWorkflow) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceWorkflow) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceWorkflow) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceWorkflow) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceWorkflow) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskReservationTimeout

`func (o *TaskrouterV1WorkspaceWorkflow) GetTaskReservationTimeout() int32`

GetTaskReservationTimeout returns the TaskReservationTimeout field if non-nil, zero value otherwise.

### GetTaskReservationTimeoutOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetTaskReservationTimeoutOk() (*int32, bool)`

GetTaskReservationTimeoutOk returns a tuple with the TaskReservationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReservationTimeout

`func (o *TaskrouterV1WorkspaceWorkflow) SetTaskReservationTimeout(v int32)`

SetTaskReservationTimeout sets TaskReservationTimeout field to given value.

### HasTaskReservationTimeout

`func (o *TaskrouterV1WorkspaceWorkflow) HasTaskReservationTimeout() bool`

HasTaskReservationTimeout returns a boolean if a field has been set.

### SetTaskReservationTimeoutNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetTaskReservationTimeoutNil(b bool)`

 SetTaskReservationTimeoutNil sets the value for TaskReservationTimeout to be an explicit nil

### UnsetTaskReservationTimeout
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetTaskReservationTimeout()`

UnsetTaskReservationTimeout ensures that no value is present for TaskReservationTimeout, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorkflow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorkflow) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorkflow) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflow) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorkflow) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflow) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorkflow) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorkflow) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorkflow) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


