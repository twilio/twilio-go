# StudioV2FlowExecutionExecutionStep

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Context** | Pointer to **map[string]interface{}** | The current state of the flow | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**ExecutionSid** | Pointer to **NullableString** | The SID of the Execution | [optional] 
**FlowSid** | Pointer to **NullableString** | The SID of the Flow | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources | [optional] 
**Name** | Pointer to **NullableString** | The event that caused the Flow to transition to the Step | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TransitionedFrom** | Pointer to **NullableString** | The Widget that preceded the Widget for the Step | [optional] 
**TransitionedTo** | Pointer to **NullableString** | The Widget that will follow the Widget for the Step | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewStudioV2FlowExecutionExecutionStep

`func NewStudioV2FlowExecutionExecutionStep() *StudioV2FlowExecutionExecutionStep`

NewStudioV2FlowExecutionExecutionStep instantiates a new StudioV2FlowExecutionExecutionStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV2FlowExecutionExecutionStepWithDefaults

`func NewStudioV2FlowExecutionExecutionStepWithDefaults() *StudioV2FlowExecutionExecutionStep`

NewStudioV2FlowExecutionExecutionStepWithDefaults instantiates a new StudioV2FlowExecutionExecutionStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV2FlowExecutionExecutionStep) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV2FlowExecutionExecutionStep) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV2FlowExecutionExecutionStep) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV2FlowExecutionExecutionStep) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV2FlowExecutionExecutionStep) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV2FlowExecutionExecutionStep) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetContext

`func (o *StudioV2FlowExecutionExecutionStep) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StudioV2FlowExecutionExecutionStep) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StudioV2FlowExecutionExecutionStep) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *StudioV2FlowExecutionExecutionStep) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *StudioV2FlowExecutionExecutionStep) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *StudioV2FlowExecutionExecutionStep) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetDateCreated

`func (o *StudioV2FlowExecutionExecutionStep) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV2FlowExecutionExecutionStep) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV2FlowExecutionExecutionStep) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV2FlowExecutionExecutionStep) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV2FlowExecutionExecutionStep) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV2FlowExecutionExecutionStep) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV2FlowExecutionExecutionStep) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV2FlowExecutionExecutionStep) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV2FlowExecutionExecutionStep) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV2FlowExecutionExecutionStep) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV2FlowExecutionExecutionStep) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV2FlowExecutionExecutionStep) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetExecutionSid

`func (o *StudioV2FlowExecutionExecutionStep) GetExecutionSid() string`

GetExecutionSid returns the ExecutionSid field if non-nil, zero value otherwise.

### GetExecutionSidOk

`func (o *StudioV2FlowExecutionExecutionStep) GetExecutionSidOk() (*string, bool)`

GetExecutionSidOk returns a tuple with the ExecutionSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionSid

`func (o *StudioV2FlowExecutionExecutionStep) SetExecutionSid(v string)`

SetExecutionSid sets ExecutionSid field to given value.

### HasExecutionSid

`func (o *StudioV2FlowExecutionExecutionStep) HasExecutionSid() bool`

HasExecutionSid returns a boolean if a field has been set.

### SetExecutionSidNil

`func (o *StudioV2FlowExecutionExecutionStep) SetExecutionSidNil(b bool)`

 SetExecutionSidNil sets the value for ExecutionSid to be an explicit nil

### UnsetExecutionSid
`func (o *StudioV2FlowExecutionExecutionStep) UnsetExecutionSid()`

UnsetExecutionSid ensures that no value is present for ExecutionSid, not even an explicit nil
### GetFlowSid

`func (o *StudioV2FlowExecutionExecutionStep) GetFlowSid() string`

GetFlowSid returns the FlowSid field if non-nil, zero value otherwise.

### GetFlowSidOk

`func (o *StudioV2FlowExecutionExecutionStep) GetFlowSidOk() (*string, bool)`

GetFlowSidOk returns a tuple with the FlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSid

`func (o *StudioV2FlowExecutionExecutionStep) SetFlowSid(v string)`

SetFlowSid sets FlowSid field to given value.

### HasFlowSid

`func (o *StudioV2FlowExecutionExecutionStep) HasFlowSid() bool`

HasFlowSid returns a boolean if a field has been set.

### SetFlowSidNil

`func (o *StudioV2FlowExecutionExecutionStep) SetFlowSidNil(b bool)`

 SetFlowSidNil sets the value for FlowSid to be an explicit nil

### UnsetFlowSid
`func (o *StudioV2FlowExecutionExecutionStep) UnsetFlowSid()`

UnsetFlowSid ensures that no value is present for FlowSid, not even an explicit nil
### GetLinks

`func (o *StudioV2FlowExecutionExecutionStep) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StudioV2FlowExecutionExecutionStep) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StudioV2FlowExecutionExecutionStep) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StudioV2FlowExecutionExecutionStep) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StudioV2FlowExecutionExecutionStep) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *StudioV2FlowExecutionExecutionStep) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetName

`func (o *StudioV2FlowExecutionExecutionStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StudioV2FlowExecutionExecutionStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StudioV2FlowExecutionExecutionStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StudioV2FlowExecutionExecutionStep) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StudioV2FlowExecutionExecutionStep) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StudioV2FlowExecutionExecutionStep) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSid

`func (o *StudioV2FlowExecutionExecutionStep) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV2FlowExecutionExecutionStep) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV2FlowExecutionExecutionStep) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV2FlowExecutionExecutionStep) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV2FlowExecutionExecutionStep) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV2FlowExecutionExecutionStep) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTransitionedFrom

`func (o *StudioV2FlowExecutionExecutionStep) GetTransitionedFrom() string`

GetTransitionedFrom returns the TransitionedFrom field if non-nil, zero value otherwise.

### GetTransitionedFromOk

`func (o *StudioV2FlowExecutionExecutionStep) GetTransitionedFromOk() (*string, bool)`

GetTransitionedFromOk returns a tuple with the TransitionedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionedFrom

`func (o *StudioV2FlowExecutionExecutionStep) SetTransitionedFrom(v string)`

SetTransitionedFrom sets TransitionedFrom field to given value.

### HasTransitionedFrom

`func (o *StudioV2FlowExecutionExecutionStep) HasTransitionedFrom() bool`

HasTransitionedFrom returns a boolean if a field has been set.

### SetTransitionedFromNil

`func (o *StudioV2FlowExecutionExecutionStep) SetTransitionedFromNil(b bool)`

 SetTransitionedFromNil sets the value for TransitionedFrom to be an explicit nil

### UnsetTransitionedFrom
`func (o *StudioV2FlowExecutionExecutionStep) UnsetTransitionedFrom()`

UnsetTransitionedFrom ensures that no value is present for TransitionedFrom, not even an explicit nil
### GetTransitionedTo

`func (o *StudioV2FlowExecutionExecutionStep) GetTransitionedTo() string`

GetTransitionedTo returns the TransitionedTo field if non-nil, zero value otherwise.

### GetTransitionedToOk

`func (o *StudioV2FlowExecutionExecutionStep) GetTransitionedToOk() (*string, bool)`

GetTransitionedToOk returns a tuple with the TransitionedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionedTo

`func (o *StudioV2FlowExecutionExecutionStep) SetTransitionedTo(v string)`

SetTransitionedTo sets TransitionedTo field to given value.

### HasTransitionedTo

`func (o *StudioV2FlowExecutionExecutionStep) HasTransitionedTo() bool`

HasTransitionedTo returns a boolean if a field has been set.

### SetTransitionedToNil

`func (o *StudioV2FlowExecutionExecutionStep) SetTransitionedToNil(b bool)`

 SetTransitionedToNil sets the value for TransitionedTo to be an explicit nil

### UnsetTransitionedTo
`func (o *StudioV2FlowExecutionExecutionStep) UnsetTransitionedTo()`

UnsetTransitionedTo ensures that no value is present for TransitionedTo, not even an explicit nil
### GetUrl

`func (o *StudioV2FlowExecutionExecutionStep) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV2FlowExecutionExecutionStep) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV2FlowExecutionExecutionStep) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV2FlowExecutionExecutionStep) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV2FlowExecutionExecutionStep) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV2FlowExecutionExecutionStep) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


