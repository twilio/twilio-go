# StudioV1FlowEngagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ContactChannelAddress** | Pointer to **NullableString** | The phone number, SIP address or Client identifier that triggered this Engagement | [optional] 
**ContactSid** | Pointer to **NullableString** | The SID of the Contact | [optional] 
**Context** | Pointer to **map[string]interface{}** | The current state of the execution flow | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Engagement was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Engagement was last updated | [optional] 
**FlowSid** | Pointer to **NullableString** | The SID of the Flow | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of the Engagement&#39;s nested resources | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Engagement | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewStudioV1FlowEngagement

`func NewStudioV1FlowEngagement() *StudioV1FlowEngagement`

NewStudioV1FlowEngagement instantiates a new StudioV1FlowEngagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV1FlowEngagementWithDefaults

`func NewStudioV1FlowEngagementWithDefaults() *StudioV1FlowEngagement`

NewStudioV1FlowEngagementWithDefaults instantiates a new StudioV1FlowEngagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV1FlowEngagement) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV1FlowEngagement) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV1FlowEngagement) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV1FlowEngagement) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV1FlowEngagement) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV1FlowEngagement) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContactChannelAddress

`func (o *StudioV1FlowEngagement) GetContactChannelAddress() string`

GetContactChannelAddress returns the ContactChannelAddress field if non-nil, zero value otherwise.

### GetContactChannelAddressOk

`func (o *StudioV1FlowEngagement) GetContactChannelAddressOk() (*string, bool)`

GetContactChannelAddressOk returns a tuple with the ContactChannelAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactChannelAddress

`func (o *StudioV1FlowEngagement) SetContactChannelAddress(v string)`

SetContactChannelAddress sets ContactChannelAddress field to given value.

### HasContactChannelAddress

`func (o *StudioV1FlowEngagement) HasContactChannelAddress() bool`

HasContactChannelAddress returns a boolean if a field has been set.

### SetContactChannelAddressNil

`func (o *StudioV1FlowEngagement) SetContactChannelAddressNil(b bool)`

 SetContactChannelAddressNil sets the value for ContactChannelAddress to be an explicit nil

### UnsetContactChannelAddress
`func (o *StudioV1FlowEngagement) UnsetContactChannelAddress()`

UnsetContactChannelAddress ensures that no value is present for ContactChannelAddress, not even an explicit nil
### GetContactSid

`func (o *StudioV1FlowEngagement) GetContactSid() string`

GetContactSid returns the ContactSid field if non-nil, zero value otherwise.

### GetContactSidOk

`func (o *StudioV1FlowEngagement) GetContactSidOk() (*string, bool)`

GetContactSidOk returns a tuple with the ContactSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSid

`func (o *StudioV1FlowEngagement) SetContactSid(v string)`

SetContactSid sets ContactSid field to given value.

### HasContactSid

`func (o *StudioV1FlowEngagement) HasContactSid() bool`

HasContactSid returns a boolean if a field has been set.

### SetContactSidNil

`func (o *StudioV1FlowEngagement) SetContactSidNil(b bool)`

 SetContactSidNil sets the value for ContactSid to be an explicit nil

### UnsetContactSid
`func (o *StudioV1FlowEngagement) UnsetContactSid()`

UnsetContactSid ensures that no value is present for ContactSid, not even an explicit nil
### GetContext

`func (o *StudioV1FlowEngagement) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StudioV1FlowEngagement) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StudioV1FlowEngagement) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *StudioV1FlowEngagement) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *StudioV1FlowEngagement) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *StudioV1FlowEngagement) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDateCreated

`func (o *StudioV1FlowEngagement) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV1FlowEngagement) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV1FlowEngagement) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV1FlowEngagement) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV1FlowEngagement) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV1FlowEngagement) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV1FlowEngagement) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV1FlowEngagement) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV1FlowEngagement) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV1FlowEngagement) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV1FlowEngagement) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV1FlowEngagement) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFlowSid

`func (o *StudioV1FlowEngagement) GetFlowSid() string`

GetFlowSid returns the FlowSid field if non-nil, zero value otherwise.

### GetFlowSidOk

`func (o *StudioV1FlowEngagement) GetFlowSidOk() (*string, bool)`

GetFlowSidOk returns a tuple with the FlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSid

`func (o *StudioV1FlowEngagement) SetFlowSid(v string)`

SetFlowSid sets FlowSid field to given value.

### HasFlowSid

`func (o *StudioV1FlowEngagement) HasFlowSid() bool`

HasFlowSid returns a boolean if a field has been set.

### SetFlowSidNil

`func (o *StudioV1FlowEngagement) SetFlowSidNil(b bool)`

 SetFlowSidNil sets the value for FlowSid to be an explicit nil

### UnsetFlowSid
`func (o *StudioV1FlowEngagement) UnsetFlowSid()`

UnsetFlowSid ensures that no value is present for FlowSid, not even an explicit nil
### GetLinks

`func (o *StudioV1FlowEngagement) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StudioV1FlowEngagement) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StudioV1FlowEngagement) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StudioV1FlowEngagement) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StudioV1FlowEngagement) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *StudioV1FlowEngagement) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *StudioV1FlowEngagement) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV1FlowEngagement) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV1FlowEngagement) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV1FlowEngagement) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV1FlowEngagement) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV1FlowEngagement) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *StudioV1FlowEngagement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioV1FlowEngagement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioV1FlowEngagement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioV1FlowEngagement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StudioV1FlowEngagement) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StudioV1FlowEngagement) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *StudioV1FlowEngagement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV1FlowEngagement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV1FlowEngagement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV1FlowEngagement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV1FlowEngagement) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV1FlowEngagement) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


