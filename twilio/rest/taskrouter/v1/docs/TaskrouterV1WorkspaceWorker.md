# TaskrouterV1WorkspaceWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActivityName** | Pointer to **NullableString** | The friendly_name of the Worker&#39;s current Activity | [optional] 
**ActivitySid** | Pointer to **NullableString** | The SID of the Worker&#39;s current Activity | [optional] 
**Attributes** | Pointer to **NullableString** | The JSON string that describes the Worker | [optional] 
**Available** | Pointer to **NullableBool** | Whether the Worker is available to perform tasks | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateStatusChanged** | Pointer to **NullableTime** | The date and time in GMT of the last change to the Worker&#39;s activity | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Worker resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Worker | [optional] 

## Methods

### NewTaskrouterV1WorkspaceWorker

`func NewTaskrouterV1WorkspaceWorker() *TaskrouterV1WorkspaceWorker`

NewTaskrouterV1WorkspaceWorker instantiates a new TaskrouterV1WorkspaceWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWorkerWithDefaults

`func NewTaskrouterV1WorkspaceWorkerWithDefaults() *TaskrouterV1WorkspaceWorker`

NewTaskrouterV1WorkspaceWorkerWithDefaults instantiates a new TaskrouterV1WorkspaceWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceWorker) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceWorker) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceWorker) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceWorker) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceWorker) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceWorker) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActivityName

`func (o *TaskrouterV1WorkspaceWorker) GetActivityName() string`

GetActivityName returns the ActivityName field if non-nil, zero value otherwise.

### GetActivityNameOk

`func (o *TaskrouterV1WorkspaceWorker) GetActivityNameOk() (*string, bool)`

GetActivityNameOk returns a tuple with the ActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityName

`func (o *TaskrouterV1WorkspaceWorker) SetActivityName(v string)`

SetActivityName sets ActivityName field to given value.

### HasActivityName

`func (o *TaskrouterV1WorkspaceWorker) HasActivityName() bool`

HasActivityName returns a boolean if a field has been set.

### SetActivityNameNil

`func (o *TaskrouterV1WorkspaceWorker) SetActivityNameNil(b bool)`

 SetActivityNameNil sets the value for ActivityName to be an explicit nil

### UnsetActivityName
`func (o *TaskrouterV1WorkspaceWorker) UnsetActivityName()`

UnsetActivityName ensures that no value is present for ActivityName, not even an explicit nil
### GetActivitySid

`func (o *TaskrouterV1WorkspaceWorker) GetActivitySid() string`

GetActivitySid returns the ActivitySid field if non-nil, zero value otherwise.

### GetActivitySidOk

`func (o *TaskrouterV1WorkspaceWorker) GetActivitySidOk() (*string, bool)`

GetActivitySidOk returns a tuple with the ActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySid

`func (o *TaskrouterV1WorkspaceWorker) SetActivitySid(v string)`

SetActivitySid sets ActivitySid field to given value.

### HasActivitySid

`func (o *TaskrouterV1WorkspaceWorker) HasActivitySid() bool`

HasActivitySid returns a boolean if a field has been set.

### SetActivitySidNil

`func (o *TaskrouterV1WorkspaceWorker) SetActivitySidNil(b bool)`

 SetActivitySidNil sets the value for ActivitySid to be an explicit nil

### UnsetActivitySid
`func (o *TaskrouterV1WorkspaceWorker) UnsetActivitySid()`

UnsetActivitySid ensures that no value is present for ActivitySid, not even an explicit nil
### GetAttributes

`func (o *TaskrouterV1WorkspaceWorker) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaskrouterV1WorkspaceWorker) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaskrouterV1WorkspaceWorker) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TaskrouterV1WorkspaceWorker) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *TaskrouterV1WorkspaceWorker) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *TaskrouterV1WorkspaceWorker) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetAvailable

`func (o *TaskrouterV1WorkspaceWorker) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TaskrouterV1WorkspaceWorker) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TaskrouterV1WorkspaceWorker) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *TaskrouterV1WorkspaceWorker) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *TaskrouterV1WorkspaceWorker) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *TaskrouterV1WorkspaceWorker) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1WorkspaceWorker) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1WorkspaceWorker) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1WorkspaceWorker) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1WorkspaceWorker) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1WorkspaceWorker) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1WorkspaceWorker) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateStatusChanged

`func (o *TaskrouterV1WorkspaceWorker) GetDateStatusChanged() time.Time`

GetDateStatusChanged returns the DateStatusChanged field if non-nil, zero value otherwise.

### GetDateStatusChangedOk

`func (o *TaskrouterV1WorkspaceWorker) GetDateStatusChangedOk() (*time.Time, bool)`

GetDateStatusChangedOk returns a tuple with the DateStatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStatusChanged

`func (o *TaskrouterV1WorkspaceWorker) SetDateStatusChanged(v time.Time)`

SetDateStatusChanged sets DateStatusChanged field to given value.

### HasDateStatusChanged

`func (o *TaskrouterV1WorkspaceWorker) HasDateStatusChanged() bool`

HasDateStatusChanged returns a boolean if a field has been set.

### SetDateStatusChangedNil

`func (o *TaskrouterV1WorkspaceWorker) SetDateStatusChangedNil(b bool)`

 SetDateStatusChangedNil sets the value for DateStatusChanged to be an explicit nil

### UnsetDateStatusChanged
`func (o *TaskrouterV1WorkspaceWorker) UnsetDateStatusChanged()`

UnsetDateStatusChanged ensures that no value is present for DateStatusChanged, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1WorkspaceWorker) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1WorkspaceWorker) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1WorkspaceWorker) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1WorkspaceWorker) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1WorkspaceWorker) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1WorkspaceWorker) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *TaskrouterV1WorkspaceWorker) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TaskrouterV1WorkspaceWorker) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TaskrouterV1WorkspaceWorker) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TaskrouterV1WorkspaceWorker) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TaskrouterV1WorkspaceWorker) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TaskrouterV1WorkspaceWorker) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1WorkspaceWorker) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1WorkspaceWorker) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1WorkspaceWorker) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1WorkspaceWorker) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1WorkspaceWorker) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1WorkspaceWorker) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceWorker) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceWorker) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceWorker) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceWorker) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceWorker) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceWorker) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceWorker) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceWorker) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceWorker) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceWorker) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceWorker) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceWorker) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorker) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceWorker) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorker) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceWorker) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceWorker) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceWorker) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


