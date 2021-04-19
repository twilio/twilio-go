# EventsV1Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Schema Identifier. | [optional] 
**LastCreated** | Pointer to **NullableTime** | The date that the last schema version was created. | [optional] 
**LastVersion** | Pointer to **NullableInt32** | Last schema version. | [optional] 
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewEventsV1Schema

`func NewEventsV1Schema() *EventsV1Schema`

NewEventsV1Schema instantiates a new EventsV1Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsV1SchemaWithDefaults

`func NewEventsV1SchemaWithDefaults() *EventsV1Schema`

NewEventsV1SchemaWithDefaults instantiates a new EventsV1Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventsV1Schema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsV1Schema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsV1Schema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventsV1Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EventsV1Schema) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EventsV1Schema) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastCreated

`func (o *EventsV1Schema) GetLastCreated() time.Time`

GetLastCreated returns the LastCreated field if non-nil, zero value otherwise.

### GetLastCreatedOk

`func (o *EventsV1Schema) GetLastCreatedOk() (*time.Time, bool)`

GetLastCreatedOk returns a tuple with the LastCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCreated

`func (o *EventsV1Schema) SetLastCreated(v time.Time)`

SetLastCreated sets LastCreated field to given value.

### HasLastCreated

`func (o *EventsV1Schema) HasLastCreated() bool`

HasLastCreated returns a boolean if a field has been set.

### SetLastCreatedNil

`func (o *EventsV1Schema) SetLastCreatedNil(b bool)`

 SetLastCreatedNil sets the value for LastCreated to be an explicit nil

### UnsetLastCreated
`func (o *EventsV1Schema) UnsetLastCreated()`

UnsetLastCreated ensures that no value is present for LastCreated, not even an explicit nil
### GetLastVersion

`func (o *EventsV1Schema) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *EventsV1Schema) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *EventsV1Schema) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *EventsV1Schema) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### SetLastVersionNil

`func (o *EventsV1Schema) SetLastVersionNil(b bool)`

 SetLastVersionNil sets the value for LastVersion to be an explicit nil

### UnsetLastVersion
`func (o *EventsV1Schema) UnsetLastVersion()`

UnsetLastVersion ensures that no value is present for LastVersion, not even an explicit nil
### GetLinks

`func (o *EventsV1Schema) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsV1Schema) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsV1Schema) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventsV1Schema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EventsV1Schema) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EventsV1Schema) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetUrl

`func (o *EventsV1Schema) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsV1Schema) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsV1Schema) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsV1Schema) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EventsV1Schema) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EventsV1Schema) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


