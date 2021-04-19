# MonitorV1Event

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActorSid** | Pointer to **NullableString** | The SID of the actor that caused the event, if available | [optional] 
**ActorType** | Pointer to **NullableString** | The type of actor that caused the event | [optional] 
**Description** | Pointer to **NullableString** | A description of the event | [optional] 
**EventData** | Pointer to **map[string]interface{}** | A JSON string that represents an object with additional data about the event | [optional] 
**EventDate** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the event was recorded | [optional] 
**EventType** | Pointer to **NullableString** | The event&#39;s type | [optional] 
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources | [optional] 
**ResourceSid** | Pointer to **NullableString** | The SID of the resource that was affected | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of resource that was affected | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Source** | Pointer to **NullableString** | The originating system or interface that caused the event | [optional] 
**SourceIpAddress** | Pointer to **NullableString** | The IP address of the source | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource that was affected | [optional] 

## Methods

### NewMonitorV1Event

`func NewMonitorV1Event() *MonitorV1Event`

NewMonitorV1Event instantiates a new MonitorV1Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorV1EventWithDefaults

`func NewMonitorV1EventWithDefaults() *MonitorV1Event`

NewMonitorV1EventWithDefaults instantiates a new MonitorV1Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *MonitorV1Event) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *MonitorV1Event) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *MonitorV1Event) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *MonitorV1Event) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *MonitorV1Event) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *MonitorV1Event) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActorSid

`func (o *MonitorV1Event) GetActorSid() string`

GetActorSid returns the ActorSid field if non-nil, zero value otherwise.

### GetActorSidOk

`func (o *MonitorV1Event) GetActorSidOk() (*string, bool)`

GetActorSidOk returns a tuple with the ActorSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorSid

`func (o *MonitorV1Event) SetActorSid(v string)`

SetActorSid sets ActorSid field to given value.

### HasActorSid

`func (o *MonitorV1Event) HasActorSid() bool`

HasActorSid returns a boolean if a field has been set.

### SetActorSidNil

`func (o *MonitorV1Event) SetActorSidNil(b bool)`

 SetActorSidNil sets the value for ActorSid to be an explicit nil

### UnsetActorSid
`func (o *MonitorV1Event) UnsetActorSid()`

UnsetActorSid ensures that no value is present for ActorSid, not even an explicit nil
### GetActorType

`func (o *MonitorV1Event) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *MonitorV1Event) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *MonitorV1Event) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *MonitorV1Event) HasActorType() bool`

HasActorType returns a boolean if a field has been set.

### SetActorTypeNil

`func (o *MonitorV1Event) SetActorTypeNil(b bool)`

 SetActorTypeNil sets the value for ActorType to be an explicit nil

### UnsetActorType
`func (o *MonitorV1Event) UnsetActorType()`

UnsetActorType ensures that no value is present for ActorType, not even an explicit nil
### GetDescription

`func (o *MonitorV1Event) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorV1Event) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorV1Event) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorV1Event) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MonitorV1Event) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MonitorV1Event) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEventData

`func (o *MonitorV1Event) GetEventData() map[string]interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *MonitorV1Event) GetEventDataOk() (*map[string]interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *MonitorV1Event) SetEventData(v map[string]interface{})`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *MonitorV1Event) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### SetEventDataNil

`func (o *MonitorV1Event) SetEventDataNil(b bool)`

 SetEventDataNil sets the value for EventData to be an explicit nil

### UnsetEventData
`func (o *MonitorV1Event) UnsetEventData()`

UnsetEventData ensures that no value is present for EventData, not even an explicit nil
### GetEventDate

`func (o *MonitorV1Event) GetEventDate() time.Time`

GetEventDate returns the EventDate field if non-nil, zero value otherwise.

### GetEventDateOk

`func (o *MonitorV1Event) GetEventDateOk() (*time.Time, bool)`

GetEventDateOk returns a tuple with the EventDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDate

`func (o *MonitorV1Event) SetEventDate(v time.Time)`

SetEventDate sets EventDate field to given value.

### HasEventDate

`func (o *MonitorV1Event) HasEventDate() bool`

HasEventDate returns a boolean if a field has been set.

### SetEventDateNil

`func (o *MonitorV1Event) SetEventDateNil(b bool)`

 SetEventDateNil sets the value for EventDate to be an explicit nil

### UnsetEventDate
`func (o *MonitorV1Event) UnsetEventDate()`

UnsetEventDate ensures that no value is present for EventDate, not even an explicit nil
### GetEventType

`func (o *MonitorV1Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MonitorV1Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MonitorV1Event) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MonitorV1Event) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *MonitorV1Event) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *MonitorV1Event) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetLinks

`func (o *MonitorV1Event) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MonitorV1Event) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MonitorV1Event) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MonitorV1Event) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MonitorV1Event) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MonitorV1Event) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetResourceSid

`func (o *MonitorV1Event) GetResourceSid() string`

GetResourceSid returns the ResourceSid field if non-nil, zero value otherwise.

### GetResourceSidOk

`func (o *MonitorV1Event) GetResourceSidOk() (*string, bool)`

GetResourceSidOk returns a tuple with the ResourceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSid

`func (o *MonitorV1Event) SetResourceSid(v string)`

SetResourceSid sets ResourceSid field to given value.

### HasResourceSid

`func (o *MonitorV1Event) HasResourceSid() bool`

HasResourceSid returns a boolean if a field has been set.

### SetResourceSidNil

`func (o *MonitorV1Event) SetResourceSidNil(b bool)`

 SetResourceSidNil sets the value for ResourceSid to be an explicit nil

### UnsetResourceSid
`func (o *MonitorV1Event) UnsetResourceSid()`

UnsetResourceSid ensures that no value is present for ResourceSid, not even an explicit nil
### GetResourceType

`func (o *MonitorV1Event) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MonitorV1Event) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MonitorV1Event) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *MonitorV1Event) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *MonitorV1Event) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *MonitorV1Event) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetSid

`func (o *MonitorV1Event) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MonitorV1Event) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MonitorV1Event) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MonitorV1Event) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *MonitorV1Event) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *MonitorV1Event) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSource

`func (o *MonitorV1Event) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MonitorV1Event) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MonitorV1Event) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MonitorV1Event) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MonitorV1Event) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MonitorV1Event) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceIpAddress

`func (o *MonitorV1Event) GetSourceIpAddress() string`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *MonitorV1Event) GetSourceIpAddressOk() (*string, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *MonitorV1Event) SetSourceIpAddress(v string)`

SetSourceIpAddress sets SourceIpAddress field to given value.

### HasSourceIpAddress

`func (o *MonitorV1Event) HasSourceIpAddress() bool`

HasSourceIpAddress returns a boolean if a field has been set.

### SetSourceIpAddressNil

`func (o *MonitorV1Event) SetSourceIpAddressNil(b bool)`

 SetSourceIpAddressNil sets the value for SourceIpAddress to be an explicit nil

### UnsetSourceIpAddress
`func (o *MonitorV1Event) UnsetSourceIpAddress()`

UnsetSourceIpAddress ensures that no value is present for SourceIpAddress, not even an explicit nil
### GetUrl

`func (o *MonitorV1Event) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonitorV1Event) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonitorV1Event) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MonitorV1Event) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MonitorV1Event) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MonitorV1Event) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


