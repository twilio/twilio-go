# EventsV1Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | Account SID. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date this Subscription was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date this Subscription was updated | [optional] 
**Description** | Pointer to **NullableString** | Subscription description | [optional] 
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. | [optional] 
**Sid** | Pointer to **NullableString** | A string that uniquely identifies this Subscription. | [optional] 
**SinkSid** | Pointer to **NullableString** | Sink SID. | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewEventsV1Subscription

`func NewEventsV1Subscription() *EventsV1Subscription`

NewEventsV1Subscription instantiates a new EventsV1Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsV1SubscriptionWithDefaults

`func NewEventsV1SubscriptionWithDefaults() *EventsV1Subscription`

NewEventsV1SubscriptionWithDefaults instantiates a new EventsV1Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *EventsV1Subscription) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *EventsV1Subscription) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *EventsV1Subscription) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *EventsV1Subscription) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *EventsV1Subscription) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *EventsV1Subscription) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *EventsV1Subscription) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *EventsV1Subscription) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *EventsV1Subscription) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *EventsV1Subscription) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *EventsV1Subscription) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *EventsV1Subscription) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *EventsV1Subscription) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *EventsV1Subscription) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *EventsV1Subscription) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *EventsV1Subscription) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *EventsV1Subscription) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *EventsV1Subscription) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDescription

`func (o *EventsV1Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventsV1Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventsV1Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventsV1Subscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventsV1Subscription) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventsV1Subscription) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLinks

`func (o *EventsV1Subscription) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventsV1Subscription) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventsV1Subscription) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventsV1Subscription) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *EventsV1Subscription) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *EventsV1Subscription) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *EventsV1Subscription) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *EventsV1Subscription) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *EventsV1Subscription) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *EventsV1Subscription) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *EventsV1Subscription) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *EventsV1Subscription) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSinkSid

`func (o *EventsV1Subscription) GetSinkSid() string`

GetSinkSid returns the SinkSid field if non-nil, zero value otherwise.

### GetSinkSidOk

`func (o *EventsV1Subscription) GetSinkSidOk() (*string, bool)`

GetSinkSidOk returns a tuple with the SinkSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkSid

`func (o *EventsV1Subscription) SetSinkSid(v string)`

SetSinkSid sets SinkSid field to given value.

### HasSinkSid

`func (o *EventsV1Subscription) HasSinkSid() bool`

HasSinkSid returns a boolean if a field has been set.

### SetSinkSidNil

`func (o *EventsV1Subscription) SetSinkSidNil(b bool)`

 SetSinkSidNil sets the value for SinkSid to be an explicit nil

### UnsetSinkSid
`func (o *EventsV1Subscription) UnsetSinkSid()`

UnsetSinkSid ensures that no value is present for SinkSid, not even an explicit nil
### GetUrl

`func (o *EventsV1Subscription) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EventsV1Subscription) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EventsV1Subscription) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EventsV1Subscription) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EventsV1Subscription) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EventsV1Subscription) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


