# StudioV1FlowExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ContactChannelAddress** | Pointer to **NullableString** | The phone number, SIP address or Client identifier that triggered the Execution | [optional] 
**ContactSid** | Pointer to **NullableString** | The SID of the Contact | [optional] 
**Context** | Pointer to **map[string]interface{}** | The current state of the flow | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FlowSid** | Pointer to **NullableString** | The SID of the Flow | [optional] 
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Execution | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewStudioV1FlowExecution

`func NewStudioV1FlowExecution() *StudioV1FlowExecution`

NewStudioV1FlowExecution instantiates a new StudioV1FlowExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV1FlowExecutionWithDefaults

`func NewStudioV1FlowExecutionWithDefaults() *StudioV1FlowExecution`

NewStudioV1FlowExecutionWithDefaults instantiates a new StudioV1FlowExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV1FlowExecution) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV1FlowExecution) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV1FlowExecution) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV1FlowExecution) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV1FlowExecution) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV1FlowExecution) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContactChannelAddress

`func (o *StudioV1FlowExecution) GetContactChannelAddress() string`

GetContactChannelAddress returns the ContactChannelAddress field if non-nil, zero value otherwise.

### GetContactChannelAddressOk

`func (o *StudioV1FlowExecution) GetContactChannelAddressOk() (*string, bool)`

GetContactChannelAddressOk returns a tuple with the ContactChannelAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactChannelAddress

`func (o *StudioV1FlowExecution) SetContactChannelAddress(v string)`

SetContactChannelAddress sets ContactChannelAddress field to given value.

### HasContactChannelAddress

`func (o *StudioV1FlowExecution) HasContactChannelAddress() bool`

HasContactChannelAddress returns a boolean if a field has been set.

### SetContactChannelAddressNil

`func (o *StudioV1FlowExecution) SetContactChannelAddressNil(b bool)`

 SetContactChannelAddressNil sets the value for ContactChannelAddress to be an explicit nil

### UnsetContactChannelAddress
`func (o *StudioV1FlowExecution) UnsetContactChannelAddress()`

UnsetContactChannelAddress ensures that no value is present for ContactChannelAddress, not even an explicit nil
### GetContactSid

`func (o *StudioV1FlowExecution) GetContactSid() string`

GetContactSid returns the ContactSid field if non-nil, zero value otherwise.

### GetContactSidOk

`func (o *StudioV1FlowExecution) GetContactSidOk() (*string, bool)`

GetContactSidOk returns a tuple with the ContactSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSid

`func (o *StudioV1FlowExecution) SetContactSid(v string)`

SetContactSid sets ContactSid field to given value.

### HasContactSid

`func (o *StudioV1FlowExecution) HasContactSid() bool`

HasContactSid returns a boolean if a field has been set.

### SetContactSidNil

`func (o *StudioV1FlowExecution) SetContactSidNil(b bool)`

 SetContactSidNil sets the value for ContactSid to be an explicit nil

### UnsetContactSid
`func (o *StudioV1FlowExecution) UnsetContactSid()`

UnsetContactSid ensures that no value is present for ContactSid, not even an explicit nil
### GetContext

`func (o *StudioV1FlowExecution) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StudioV1FlowExecution) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StudioV1FlowExecution) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *StudioV1FlowExecution) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *StudioV1FlowExecution) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *StudioV1FlowExecution) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDateCreated

`func (o *StudioV1FlowExecution) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV1FlowExecution) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV1FlowExecution) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV1FlowExecution) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV1FlowExecution) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV1FlowExecution) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV1FlowExecution) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV1FlowExecution) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV1FlowExecution) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV1FlowExecution) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV1FlowExecution) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV1FlowExecution) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFlowSid

`func (o *StudioV1FlowExecution) GetFlowSid() string`

GetFlowSid returns the FlowSid field if non-nil, zero value otherwise.

### GetFlowSidOk

`func (o *StudioV1FlowExecution) GetFlowSidOk() (*string, bool)`

GetFlowSidOk returns a tuple with the FlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSid

`func (o *StudioV1FlowExecution) SetFlowSid(v string)`

SetFlowSid sets FlowSid field to given value.

### HasFlowSid

`func (o *StudioV1FlowExecution) HasFlowSid() bool`

HasFlowSid returns a boolean if a field has been set.

### SetFlowSidNil

`func (o *StudioV1FlowExecution) SetFlowSidNil(b bool)`

 SetFlowSidNil sets the value for FlowSid to be an explicit nil

### UnsetFlowSid
`func (o *StudioV1FlowExecution) UnsetFlowSid()`

UnsetFlowSid ensures that no value is present for FlowSid, not even an explicit nil
### GetLinks

`func (o *StudioV1FlowExecution) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StudioV1FlowExecution) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StudioV1FlowExecution) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StudioV1FlowExecution) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StudioV1FlowExecution) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *StudioV1FlowExecution) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *StudioV1FlowExecution) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV1FlowExecution) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV1FlowExecution) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV1FlowExecution) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV1FlowExecution) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV1FlowExecution) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *StudioV1FlowExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioV1FlowExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioV1FlowExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioV1FlowExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StudioV1FlowExecution) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StudioV1FlowExecution) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *StudioV1FlowExecution) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV1FlowExecution) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV1FlowExecution) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV1FlowExecution) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV1FlowExecution) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV1FlowExecution) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


