# EventsV1EventType

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The date this Event Type was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Event Type was updated. | [optional] 
**Description** | Pointer to **NullableString** | Event Type description. | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**SchemaId** | Pointer to **NullableString** | The Schema identifier for this Event Type. | [optional] 
**Type** | Pointer to **NullableString** | The Event Type identifier. | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewEventsV1EventType

`func NewEventsV1EventType() *EventsV1EventType`

NewEventsV1EventType instantiates a new EventsV1EventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsV1EventTypeWithDefaults

`func NewEventsV1EventTypeWithDefaults() *EventsV1EventType`

NewEventsV1EventTypeWithDefaults instantiates a new EventsV1EventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *EventsV1EventType) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EventsV1EventType) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EventsV1EventType) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EventsV1EventType) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *EventsV1EventType) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *EventsV1EventType) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *EventsV1EventType) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *EventsV1EventType) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *EventsV1EventType) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *EventsV1EventType) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *EventsV1EventType) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *EventsV1EventType) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDescription

`func (o *EventsV1EventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventsV1EventType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventsV1EventType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventsV1EventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventsV1EventType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventsV1EventType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinks

`func (o *EventsV1EventType) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsV1EventType) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsV1EventType) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventsV1EventType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EventsV1EventType) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EventsV1EventType) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSchemaId

`func (o *EventsV1EventType) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *EventsV1EventType) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *EventsV1EventType) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *EventsV1EventType) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### SetSchemaIdNil

`func (o *EventsV1EventType) SetSchemaIdNil(b bool)`

 SetSchemaIdNil sets the value for SchemaId to be an explicit nil

### UnsetSchemaId
`func (o *EventsV1EventType) UnsetSchemaId()`

UnsetSchemaId ensures that no value is present for SchemaId, not even an explicit nil
### GetType

`func (o *EventsV1EventType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsV1EventType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsV1EventType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventsV1EventType) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *EventsV1EventType) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *EventsV1EventType) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *EventsV1EventType) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsV1EventType) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsV1EventType) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsV1EventType) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EventsV1EventType) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EventsV1EventType) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


