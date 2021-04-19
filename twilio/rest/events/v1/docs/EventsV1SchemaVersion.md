# EventsV1SchemaVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The date the schema version was created. | [optional] 
**Id** | Pointer to **NullableString** | The unique identifier of the schema. | [optional] 
**Raw** | Pointer to **NullableString** |  | [optional] 
**SchemaVersion** | Pointer to **NullableInt32** | The version of this schema. | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewEventsV1SchemaVersion

`func NewEventsV1SchemaVersion() *EventsV1SchemaVersion`

NewEventsV1SchemaVersion instantiates a new EventsV1SchemaVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsV1SchemaVersionWithDefaults

`func NewEventsV1SchemaVersionWithDefaults() *EventsV1SchemaVersion`

NewEventsV1SchemaVersionWithDefaults instantiates a new EventsV1SchemaVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *EventsV1SchemaVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EventsV1SchemaVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EventsV1SchemaVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EventsV1SchemaVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *EventsV1SchemaVersion) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *EventsV1SchemaVersion) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetId

`func (o *EventsV1SchemaVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsV1SchemaVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsV1SchemaVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventsV1SchemaVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventsV1SchemaVersion) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventsV1SchemaVersion) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRaw

`func (o *EventsV1SchemaVersion) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *EventsV1SchemaVersion) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *EventsV1SchemaVersion) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *EventsV1SchemaVersion) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *EventsV1SchemaVersion) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *EventsV1SchemaVersion) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetSchemaVersion

`func (o *EventsV1SchemaVersion) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *EventsV1SchemaVersion) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *EventsV1SchemaVersion) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *EventsV1SchemaVersion) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### SetSchemaVersionNil

`func (o *EventsV1SchemaVersion) SetSchemaVersionNil(b bool)`

 SetSchemaVersionNil sets the value for SchemaVersion to be an explicit nil

### UnsetSchemaVersion
`func (o *EventsV1SchemaVersion) UnsetSchemaVersion()`

UnsetSchemaVersion ensures that no value is present for SchemaVersion, not even an explicit nil
### GetUrl

`func (o *EventsV1SchemaVersion) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsV1SchemaVersion) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsV1SchemaVersion) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsV1SchemaVersion) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EventsV1SchemaVersion) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EventsV1SchemaVersion) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


