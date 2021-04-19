# EventsV1Sink

## Properties

Name | Type | Description
------------ | ------------- | -------------
**DateCreated** | Pointer to **NullableTime** | The date this Sink was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Sink was updated | [optional] 
**Description** | Pointer to **NullableString** | Sink Description | [optional] 
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Sink. | [optional] 
**SinkConfiguration** | Pointer to **map[string]interface{}** | JSON Sink configuration. | [optional] 
**SinkType** | Pointer to **NullableString** | Sink type. | [optional] 
**Status** | Pointer to **NullableString** | The Status of this Sink | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewEventsV1Sink

`func NewEventsV1Sink() *EventsV1Sink`

NewEventsV1Sink instantiates a new EventsV1Sink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsV1SinkWithDefaults

`func NewEventsV1SinkWithDefaults() *EventsV1Sink`

NewEventsV1SinkWithDefaults instantiates a new EventsV1Sink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *EventsV1Sink) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EventsV1Sink) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EventsV1Sink) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EventsV1Sink) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *EventsV1Sink) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *EventsV1Sink) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *EventsV1Sink) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *EventsV1Sink) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *EventsV1Sink) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *EventsV1Sink) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *EventsV1Sink) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *EventsV1Sink) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDescription

`func (o *EventsV1Sink) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventsV1Sink) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventsV1Sink) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventsV1Sink) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventsV1Sink) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventsV1Sink) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinks

`func (o *EventsV1Sink) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsV1Sink) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsV1Sink) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventsV1Sink) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EventsV1Sink) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EventsV1Sink) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *EventsV1Sink) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *EventsV1Sink) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *EventsV1Sink) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *EventsV1Sink) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *EventsV1Sink) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *EventsV1Sink) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSinkConfiguration

`func (o *EventsV1Sink) GetSinkConfiguration() map[string]interface{}`

GetSinkConfiguration returns the SinkConfiguration field if non-nil, zero value otherwise.

### GetSinkConfigurationOk

`func (o *EventsV1Sink) GetSinkConfigurationOk() (*map[string]interface{}, bool)`

GetSinkConfigurationOk returns a tuple with the SinkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkConfiguration

`func (o *EventsV1Sink) SetSinkConfiguration(v map[string]interface{})`

SetSinkConfiguration sets SinkConfiguration field to given value.

### HasSinkConfiguration

`func (o *EventsV1Sink) HasSinkConfiguration() bool`

HasSinkConfiguration returns a boolean if a field has been set.

### SetSinkConfigurationNil

`func (o *EventsV1Sink) SetSinkConfigurationNil(b bool)`

 SetSinkConfigurationNil sets the value for SinkConfiguration to be an explicit nil

### UnsetSinkConfiguration
`func (o *EventsV1Sink) UnsetSinkConfiguration()`

UnsetSinkConfiguration ensures that no value is present for SinkConfiguration, not even an explicit nil
### GetSinkType

`func (o *EventsV1Sink) GetSinkType() string`

GetSinkType returns the SinkType field if non-nil, zero value otherwise.

### GetSinkTypeOk

`func (o *EventsV1Sink) GetSinkTypeOk() (*string, bool)`

GetSinkTypeOk returns a tuple with the SinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkType

`func (o *EventsV1Sink) SetSinkType(v string)`

SetSinkType sets SinkType field to given value.

### HasSinkType

`func (o *EventsV1Sink) HasSinkType() bool`

HasSinkType returns a boolean if a field has been set.

### SetSinkTypeNil

`func (o *EventsV1Sink) SetSinkTypeNil(b bool)`

 SetSinkTypeNil sets the value for SinkType to be an explicit nil

### UnsetSinkType
`func (o *EventsV1Sink) UnsetSinkType()`

UnsetSinkType ensures that no value is present for SinkType, not even an explicit nil
### GetStatus

`func (o *EventsV1Sink) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventsV1Sink) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventsV1Sink) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventsV1Sink) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EventsV1Sink) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EventsV1Sink) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *EventsV1Sink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsV1Sink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsV1Sink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsV1Sink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EventsV1Sink) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EventsV1Sink) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


