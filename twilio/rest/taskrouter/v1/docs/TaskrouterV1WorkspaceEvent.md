# TaskrouterV1WorkspaceEvent

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActorSid** | Pointer to **NullableString** | The SID of the resource that triggered the event | [optional] 
**ActorType** | Pointer to **NullableString** | The type of resource that triggered the event | [optional] 
**ActorUrl** | Pointer to **NullableString** | The absolute URL of the resource that triggered the event | [optional] 
**Description** | Pointer to **NullableString** | A description of the event | [optional] 
**EventData** | Pointer to **map[string]interface{}** | Data about the event | [optional] 
**EventDate** | Pointer to **NullableTime** | The time the event was sent | [optional] 
**EventDateMs** | Pointer to **NullableInt32** | The time the event was sent in milliseconds | [optional] 
**EventType** | Pointer to **NullableString** | The identifier for the event | [optional] 
**ResourceSid** | Pointer to **NullableString** | The SID of the object the event is most relevant to | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of object the event is most relevant to | [optional] 
**ResourceUrl** | Pointer to **NullableString** | The URL of the resource the event is most relevant to | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Source** | Pointer to **NullableString** | Where the Event originated | [optional] 
**SourceIpAddress** | Pointer to **NullableString** | The IP from which the Event originated | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Event resource | [optional] 
**WorkspaceSid** | Pointer to **NullableString** | The SID of the Workspace that contains the Event | [optional] 

## Methods

### NewTaskrouterV1WorkspaceEvent

`func NewTaskrouterV1WorkspaceEvent() *TaskrouterV1WorkspaceEvent`

NewTaskrouterV1WorkspaceEvent instantiates a new TaskrouterV1WorkspaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskrouterV1WorkspaceEventWithDefaults

`func NewTaskrouterV1WorkspaceEventWithDefaults() *TaskrouterV1WorkspaceEvent`

NewTaskrouterV1WorkspaceEventWithDefaults instantiates a new TaskrouterV1WorkspaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TaskrouterV1WorkspaceEvent) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TaskrouterV1WorkspaceEvent) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TaskrouterV1WorkspaceEvent) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TaskrouterV1WorkspaceEvent) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *TaskrouterV1WorkspaceEvent) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *TaskrouterV1WorkspaceEvent) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActorSid

`func (o *TaskrouterV1WorkspaceEvent) GetActorSid() string`

GetActorSid returns the ActorSid field if non-nil, zero value otherwise.

### GetActorSidOk

`func (o *TaskrouterV1WorkspaceEvent) GetActorSidOk() (*string, bool)`

GetActorSidOk returns a tuple with the ActorSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSid

`func (o *TaskrouterV1WorkspaceEvent) SetActorSid(v string)`

SetActorSid sets ActorSid field to given value.

### HasActorSid

`func (o *TaskrouterV1WorkspaceEvent) HasActorSid() bool`

HasActorSid returns a boolean if a field has been set.

### SetActorSidNil

`func (o *TaskrouterV1WorkspaceEvent) SetActorSidNil(b bool)`

 SetActorSidNil sets the value for ActorSid to be an explicit nil

### UnsetActorSid
`func (o *TaskrouterV1WorkspaceEvent) UnsetActorSid()`

UnsetActorSid ensures that no value is present for ActorSid, not even an explicit nil
### GetActorType

`func (o *TaskrouterV1WorkspaceEvent) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *TaskrouterV1WorkspaceEvent) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *TaskrouterV1WorkspaceEvent) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *TaskrouterV1WorkspaceEvent) HasActorType() bool`

HasActorType returns a boolean if a field has been set.

### SetActorTypeNil

`func (o *TaskrouterV1WorkspaceEvent) SetActorTypeNil(b bool)`

 SetActorTypeNil sets the value for ActorType to be an explicit nil

### UnsetActorType
`func (o *TaskrouterV1WorkspaceEvent) UnsetActorType()`

UnsetActorType ensures that no value is present for ActorType, not even an explicit nil
### GetActorUrl

`func (o *TaskrouterV1WorkspaceEvent) GetActorUrl() string`

GetActorUrl returns the ActorUrl field if non-nil, zero value otherwise.

### GetActorUrlOk

`func (o *TaskrouterV1WorkspaceEvent) GetActorUrlOk() (*string, bool)`

GetActorUrlOk returns a tuple with the ActorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUrl

`func (o *TaskrouterV1WorkspaceEvent) SetActorUrl(v string)`

SetActorUrl sets ActorUrl field to given value.

### HasActorUrl

`func (o *TaskrouterV1WorkspaceEvent) HasActorUrl() bool`

HasActorUrl returns a boolean if a field has been set.

### SetActorUrlNil

`func (o *TaskrouterV1WorkspaceEvent) SetActorUrlNil(b bool)`

 SetActorUrlNil sets the value for ActorUrl to be an explicit nil

### UnsetActorUrl
`func (o *TaskrouterV1WorkspaceEvent) UnsetActorUrl()`

UnsetActorUrl ensures that no value is present for ActorUrl, not even an explicit nil
### GetDescription

`func (o *TaskrouterV1WorkspaceEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskrouterV1WorkspaceEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskrouterV1WorkspaceEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskrouterV1WorkspaceEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskrouterV1WorkspaceEvent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskrouterV1WorkspaceEvent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventData

`func (o *TaskrouterV1WorkspaceEvent) GetEventData() map[string]interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *TaskrouterV1WorkspaceEvent) GetEventDataOk() (*map[string]interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *TaskrouterV1WorkspaceEvent) SetEventData(v map[string]interface{})`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *TaskrouterV1WorkspaceEvent) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### SetEventDataNil

`func (o *TaskrouterV1WorkspaceEvent) SetEventDataNil(b bool)`

 SetEventDataNil sets the value for EventData to be an explicit nil

### UnsetEventData
`func (o *TaskrouterV1WorkspaceEvent) UnsetEventData()`

UnsetEventData ensures that no value is present for EventData, not even an explicit nil
### GetEventDate

`func (o *TaskrouterV1WorkspaceEvent) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *TaskrouterV1WorkspaceEvent) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *TaskrouterV1WorkspaceEvent) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *TaskrouterV1WorkspaceEvent) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *TaskrouterV1WorkspaceEvent) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *TaskrouterV1WorkspaceEvent) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetEventDateMs

`func (o *TaskrouterV1WorkspaceEvent) GetEventDateMs() int32`

GetEventDateMs returns the EventDateMs field if non-nil, zero value otherwise.

### GetEventDateMsOk

`func (o *TaskrouterV1WorkspaceEvent) GetEventDateMsOk() (*int32, bool)`

GetEventDateMsOk returns a tuple with the EventDateMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateMs

`func (o *TaskrouterV1WorkspaceEvent) SetEventDateMs(v int32)`

SetEventDateMs sets EventDateMs field to given value.

### HasEventDateMs

`func (o *TaskrouterV1WorkspaceEvent) HasEventDateMs() bool`

HasEventDateMs returns a boolean if a field has been set.

### SetEventDateMsNil

`func (o *TaskrouterV1WorkspaceEvent) SetEventDateMsNil(b bool)`

 SetEventDateMsNil sets the value for EventDateMs to be an explicit nil

### UnsetEventDateMs
`func (o *TaskrouterV1WorkspaceEvent) UnsetEventDateMs()`

UnsetEventDateMs ensures that no value is present for EventDateMs, not even an explicit nil
### GetEventType

`func (o *TaskrouterV1WorkspaceEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TaskrouterV1WorkspaceEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TaskrouterV1WorkspaceEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TaskrouterV1WorkspaceEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *TaskrouterV1WorkspaceEvent) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *TaskrouterV1WorkspaceEvent) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetResourceSid

`func (o *TaskrouterV1WorkspaceEvent) GetResourceSid() string`

GetResourceSid returns the ResourceSid field if non-nil, zero value otherwise.

### GetResourceSidOk

`func (o *TaskrouterV1WorkspaceEvent) GetResourceSidOk() (*string, bool)`

GetResourceSidOk returns a tuple with the ResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSid

`func (o *TaskrouterV1WorkspaceEvent) SetResourceSid(v string)`

SetResourceSid sets ResourceSid field to given value.

### HasResourceSid

`func (o *TaskrouterV1WorkspaceEvent) HasResourceSid() bool`

HasResourceSid returns a boolean if a field has been set.

### SetResourceSidNil

`func (o *TaskrouterV1WorkspaceEvent) SetResourceSidNil(b bool)`

 SetResourceSidNil sets the value for ResourceSid to be an explicit nil

### UnsetResourceSid
`func (o *TaskrouterV1WorkspaceEvent) UnsetResourceSid()`

UnsetResourceSid ensures that no value is present for ResourceSid, not even an explicit nil
### GetResourceType

`func (o *TaskrouterV1WorkspaceEvent) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TaskrouterV1WorkspaceEvent) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TaskrouterV1WorkspaceEvent) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *TaskrouterV1WorkspaceEvent) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *TaskrouterV1WorkspaceEvent) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *TaskrouterV1WorkspaceEvent) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceUrl

`func (o *TaskrouterV1WorkspaceEvent) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *TaskrouterV1WorkspaceEvent) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *TaskrouterV1WorkspaceEvent) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *TaskrouterV1WorkspaceEvent) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### SetResourceUrlNil

`func (o *TaskrouterV1WorkspaceEvent) SetResourceUrlNil(b bool)`

 SetResourceUrlNil sets the value for ResourceUrl to be an explicit nil

### UnsetResourceUrl
`func (o *TaskrouterV1WorkspaceEvent) UnsetResourceUrl()`

UnsetResourceUrl ensures that no value is present for ResourceUrl, not even an explicit nil
### GetSid

`func (o *TaskrouterV1WorkspaceEvent) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TaskrouterV1WorkspaceEvent) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TaskrouterV1WorkspaceEvent) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TaskrouterV1WorkspaceEvent) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *TaskrouterV1WorkspaceEvent) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *TaskrouterV1WorkspaceEvent) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSource

`func (o *TaskrouterV1WorkspaceEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaskrouterV1WorkspaceEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaskrouterV1WorkspaceEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TaskrouterV1WorkspaceEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *TaskrouterV1WorkspaceEvent) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *TaskrouterV1WorkspaceEvent) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceIpAddress

`func (o *TaskrouterV1WorkspaceEvent) GetSourceIpAddress() string`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *TaskrouterV1WorkspaceEvent) GetSourceIpAddressOk() (*string, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *TaskrouterV1WorkspaceEvent) SetSourceIpAddress(v string)`

SetSourceIpAddress sets SourceIpAddress field to given value.

### HasSourceIpAddress

`func (o *TaskrouterV1WorkspaceEvent) HasSourceIpAddress() bool`

HasSourceIpAddress returns a boolean if a field has been set.

### SetSourceIpAddressNil

`func (o *TaskrouterV1WorkspaceEvent) SetSourceIpAddressNil(b bool)`

 SetSourceIpAddressNil sets the value for SourceIpAddress to be an explicit nil

### UnsetSourceIpAddress
`func (o *TaskrouterV1WorkspaceEvent) UnsetSourceIpAddress()`

UnsetSourceIpAddress ensures that no value is present for SourceIpAddress, not even an explicit nil
### GetUrl

`func (o *TaskrouterV1WorkspaceEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TaskrouterV1WorkspaceEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TaskrouterV1WorkspaceEvent) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TaskrouterV1WorkspaceEvent) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TaskrouterV1WorkspaceEvent) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TaskrouterV1WorkspaceEvent) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWorkspaceSid

`func (o *TaskrouterV1WorkspaceEvent) GetWorkspaceSid() string`

GetWorkspaceSid returns the WorkspaceSid field if non-nil, zero value otherwise.

### GetWorkspaceSidOk

`func (o *TaskrouterV1WorkspaceEvent) GetWorkspaceSidOk() (*string, bool)`

GetWorkspaceSidOk returns a tuple with the WorkspaceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSid

`func (o *TaskrouterV1WorkspaceEvent) SetWorkspaceSid(v string)`

SetWorkspaceSid sets WorkspaceSid field to given value.

### HasWorkspaceSid

`func (o *TaskrouterV1WorkspaceEvent) HasWorkspaceSid() bool`

HasWorkspaceSid returns a boolean if a field has been set.

### SetWorkspaceSidNil

`func (o *TaskrouterV1WorkspaceEvent) SetWorkspaceSidNil(b bool)`

 SetWorkspaceSidNil sets the value for WorkspaceSid to be an explicit nil

### UnsetWorkspaceSid
`func (o *TaskrouterV1WorkspaceEvent) UnsetWorkspaceSid()`

UnsetWorkspaceSid ensures that no value is present for WorkspaceSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


