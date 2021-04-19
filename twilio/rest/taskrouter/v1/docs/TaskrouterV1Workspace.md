# TaskrouterV1Workspace

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**DefaultActivityName** | Pointer to **NullableString** | The name of the default activity | [optional] 
**DefaultActivitySid** | Pointer to **NullableString** | The SID of the Activity that will be used when new Workers are created in the Workspace | [optional] 
**EventCallbackUrl** | Pointer to **NullableString** | The URL we call when an event occurs | [optional] 
**EventsFilter** | Pointer to **NullableString** | The list of Workspace events for which to call event_callback_url | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the Workspace resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**MultiTaskEnabled** | Pointer to **NullableBool** | Whether multi-tasking is enabled | [optional] 
**PrioritizeQueueOrder** | Pointer to **NullableString** | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TimeoutActivityName** | Pointer to **NullableString** | The name of the timeout activity | [optional] 
**TimeoutActivitySid** | Pointer to **NullableString** | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Workspace resource | [optional] 

## Methods

### NewTaskrouterV1Workspace

`func NewTaskrouterV1Workspace() *TaskrouterV1Workspace`

NewTaskrouterV1Workspace instantiates a new TaskrouterV1Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceWithDefaults

`func NewTaskrouterV1WorkspaceWithDefaults() *TaskrouterV1Workspace`

NewTaskrouterV1WorkspaceWithDefaults instantiates a new TaskrouterV1Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1Workspace) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1Workspace) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1Workspace) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1Workspace) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1Workspace) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1Workspace) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *TaskrouterV1Workspace) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskrouterV1Workspace) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskrouterV1Workspace) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskrouterV1Workspace) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TaskrouterV1Workspace) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TaskrouterV1Workspace) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *TaskrouterV1Workspace) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TaskrouterV1Workspace) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TaskrouterV1Workspace) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TaskrouterV1Workspace) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *TaskrouterV1Workspace) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *TaskrouterV1Workspace) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefaultActivityName

`func (o *TaskrouterV1Workspace) GetDefaultActivityName() string`

GetDefaultActivityName returns the DefaultActivityName field if non-nil, zero value otherwise.

### GetDefaultActivityNameOk

`func (o *TaskrouterV1Workspace) GetDefaultActivityNameOk() (*string, bool)`

GetDefaultActivityNameOk returns a tuple with the DefaultActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActivityName

`func (o *TaskrouterV1Workspace) SetDefaultActivityName(v string)`

SetDefaultActivityName sets DefaultActivityName field to given value.

### HasDefaultActivityName

`func (o *TaskrouterV1Workspace) HasDefaultActivityName() bool`

HasDefaultActivityName returns a boolean if a field has been set.

### SetDefaultActivityNameNil

`func (o *TaskrouterV1Workspace) SetDefaultActivityNameNil(b bool)`

 SetDefaultActivityNameNil sets the value for DefaultActivityName to be an explicit nil

### UnsetDefaultActivityName
`func (o *TaskrouterV1Workspace) UnsetDefaultActivityName()`

UnsetDefaultActivityName ensures that no value is present for DefaultActivityName, not even an explicit nil
### GetDefaultActivitySid

`func (o *TaskrouterV1Workspace) GetDefaultActivitySid() string`

GetDefaultActivitySid returns the DefaultActivitySid field if non-nil, zero value otherwise.

### GetDefaultActivitySidOk

`func (o *TaskrouterV1Workspace) GetDefaultActivitySidOk() (*string, bool)`

GetDefaultActivitySidOk returns a tuple with the DefaultActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultActivitySid

`func (o *TaskrouterV1Workspace) SetDefaultActivitySid(v string)`

SetDefaultActivitySid sets DefaultActivitySid field to given value.

### HasDefaultActivitySid

`func (o *TaskrouterV1Workspace) HasDefaultActivitySid() bool`

HasDefaultActivitySid returns a boolean if a field has been set.

### SetDefaultActivitySidNil

`func (o *TaskrouterV1Workspace) SetDefaultActivitySidNil(b bool)`

 SetDefaultActivitySidNil sets the value for DefaultActivitySid to be an explicit nil

### UnsetDefaultActivitySid
`func (o *TaskrouterV1Workspace) UnsetDefaultActivitySid()`

UnsetDefaultActivitySid ensures that no value is present for DefaultActivitySid, not even an explicit nil
### GetEventCallbackUrl

`func (o *TaskrouterV1Workspace) GetEventCallbackUrl() string`

GetEventCallbackUrl returns the EventCallbackUrl field if non-nil, zero value otherwise.

### GetEventCallbackUrlOk

`func (o *TaskrouterV1Workspace) GetEventCallbackUrlOk() (*string, bool)`

GetEventCallbackUrlOk returns a tuple with the EventCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCallbackUrl

`func (o *TaskrouterV1Workspace) SetEventCallbackUrl(v string)`

SetEventCallbackUrl sets EventCallbackUrl field to given value.

### HasEventCallbackUrl

`func (o *TaskrouterV1Workspace) HasEventCallbackUrl() bool`

HasEventCallbackUrl returns a boolean if a field has been set.

### SetEventCallbackUrlNil

`func (o *TaskrouterV1Workspace) SetEventCallbackUrlNil(b bool)`

 SetEventCallbackUrlNil sets the value for EventCallbackUrl to be an explicit nil

### UnsetEventCallbackUrl
`func (o *TaskrouterV1Workspace) UnsetEventCallbackUrl()`

UnsetEventCallbackUrl ensures that no value is present for EventCallbackUrl, not even an explicit nil
### GetEventsFilter

`func (o *TaskrouterV1Workspace) GetEventsFilter() string`

GetEventsFilter returns the EventsFilter field if non-nil, zero value otherwise.

### GetEventsFilterOk

`func (o *TaskrouterV1Workspace) GetEventsFilterOk() (*string, bool)`

GetEventsFilterOk returns a tuple with the EventsFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsFilter

`func (o *TaskrouterV1Workspace) SetEventsFilter(v string)`

SetEventsFilter sets EventsFilter field to given value.

### HasEventsFilter

`func (o *TaskrouterV1Workspace) HasEventsFilter() bool`

HasEventsFilter returns a boolean if a field has been set.

### SetEventsFilterNil

`func (o *TaskrouterV1Workspace) SetEventsFilterNil(b bool)`

 SetEventsFilterNil sets the value for EventsFilter to be an explicit nil

### UnsetEventsFilter
`func (o *TaskrouterV1Workspace) UnsetEventsFilter()`

UnsetEventsFilter ensures that no value is present for EventsFilter, not even an explicit nil
### GetFriendlyName

`func (o *TaskrouterV1Workspace) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TaskrouterV1Workspace) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TaskrouterV1Workspace) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TaskrouterV1Workspace) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TaskrouterV1Workspace) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TaskrouterV1Workspace) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *TaskrouterV1Workspace) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TaskrouterV1Workspace) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TaskrouterV1Workspace) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TaskrouterV1Workspace) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *TaskrouterV1Workspace) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *TaskrouterV1Workspace) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMultiTaskEnabled

`func (o *TaskrouterV1Workspace) GetMultiTaskEnabled() bool`

GetMultiTaskEnabled returns the MultiTaskEnabled field if non-nil, zero value otherwise.

### GetMultiTaskEnabledOk

`func (o *TaskrouterV1Workspace) GetMultiTaskEnabledOk() (*bool, bool)`

GetMultiTaskEnabledOk returns a tuple with the MultiTaskEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTaskEnabled

`func (o *TaskrouterV1Workspace) SetMultiTaskEnabled(v bool)`

SetMultiTaskEnabled sets MultiTaskEnabled field to given value.

### HasMultiTaskEnabled

`func (o *TaskrouterV1Workspace) HasMultiTaskEnabled() bool`

HasMultiTaskEnabled returns a boolean if a field has been set.

### SetMultiTaskEnabledNil

`func (o *TaskrouterV1Workspace) SetMultiTaskEnabledNil(b bool)`

 SetMultiTaskEnabledNil sets the value for MultiTaskEnabled to be an explicit nil

### UnsetMultiTaskEnabled
`func (o *TaskrouterV1Workspace) UnsetMultiTaskEnabled()`

UnsetMultiTaskEnabled ensures that no value is present for MultiTaskEnabled, not even an explicit nil
### GetPrioritizeQueueOrder

`func (o *TaskrouterV1Workspace) GetPrioritizeQueueOrder() string`

GetPrioritizeQueueOrder returns the PrioritizeQueueOrder field if non-nil, zero value otherwise.

### GetPrioritizeQueueOrderOk

`func (o *TaskrouterV1Workspace) GetPrioritizeQueueOrderOk() (*string, bool)`

GetPrioritizeQueueOrderOk returns a tuple with the PrioritizeQueueOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizeQueueOrder

`func (o *TaskrouterV1Workspace) SetPrioritizeQueueOrder(v string)`

SetPrioritizeQueueOrder sets PrioritizeQueueOrder field to given value.

### HasPrioritizeQueueOrder

`func (o *TaskrouterV1Workspace) HasPrioritizeQueueOrder() bool`

HasPrioritizeQueueOrder returns a boolean if a field has been set.

### SetPrioritizeQueueOrderNil

`func (o *TaskrouterV1Workspace) SetPrioritizeQueueOrderNil(b bool)`

 SetPrioritizeQueueOrderNil sets the value for PrioritizeQueueOrder to be an explicit nil

### UnsetPrioritizeQueueOrder
`func (o *TaskrouterV1Workspace) UnsetPrioritizeQueueOrder()`

UnsetPrioritizeQueueOrder ensures that no value is present for PrioritizeQueueOrder, not even an explicit nil
### GetSid

`func (o *TaskrouterV1Workspace) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1Workspace) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1Workspace) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1Workspace) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1Workspace) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1Workspace) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTimeoutActivityName

`func (o *TaskrouterV1Workspace) GetTimeoutActivityName() string`

GetTimeoutActivityName returns the TimeoutActivityName field if non-nil, zero value otherwise.

### GetTimeoutActivityNameOk

`func (o *TaskrouterV1Workspace) GetTimeoutActivityNameOk() (*string, bool)`

GetTimeoutActivityNameOk returns a tuple with the TimeoutActivityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutActivityName

`func (o *TaskrouterV1Workspace) SetTimeoutActivityName(v string)`

SetTimeoutActivityName sets TimeoutActivityName field to given value.

### HasTimeoutActivityName

`func (o *TaskrouterV1Workspace) HasTimeoutActivityName() bool`

HasTimeoutActivityName returns a boolean if a field has been set.

### SetTimeoutActivityNameNil

`func (o *TaskrouterV1Workspace) SetTimeoutActivityNameNil(b bool)`

 SetTimeoutActivityNameNil sets the value for TimeoutActivityName to be an explicit nil

### UnsetTimeoutActivityName
`func (o *TaskrouterV1Workspace) UnsetTimeoutActivityName()`

UnsetTimeoutActivityName ensures that no value is present for TimeoutActivityName, not even an explicit nil
### GetTimeoutActivitySid

`func (o *TaskrouterV1Workspace) GetTimeoutActivitySid() string`

GetTimeoutActivitySid returns the TimeoutActivitySid field if non-nil, zero value otherwise.

### GetTimeoutActivitySidOk

`func (o *TaskrouterV1Workspace) GetTimeoutActivitySidOk() (*string, bool)`

GetTimeoutActivitySidOk returns a tuple with the TimeoutActivitySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutActivitySid

`func (o *TaskrouterV1Workspace) SetTimeoutActivitySid(v string)`

SetTimeoutActivitySid sets TimeoutActivitySid field to given value.

### HasTimeoutActivitySid

`func (o *TaskrouterV1Workspace) HasTimeoutActivitySid() bool`

HasTimeoutActivitySid returns a boolean if a field has been set.

### SetTimeoutActivitySidNil

`func (o *TaskrouterV1Workspace) SetTimeoutActivitySidNil(b bool)`

 SetTimeoutActivitySidNil sets the value for TimeoutActivitySid to be an explicit nil

### UnsetTimeoutActivitySid
`func (o *TaskrouterV1Workspace) UnsetTimeoutActivitySid()`

UnsetTimeoutActivitySid ensures that no value is present for TimeoutActivitySid, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1Workspace) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1Workspace) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1Workspace) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1Workspace) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1Workspace) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1Workspace) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


