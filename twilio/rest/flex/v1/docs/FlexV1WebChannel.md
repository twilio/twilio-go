# FlexV1WebChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource and owns this Workflow | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FlexFlowSid** | Pointer to **NullableString** | The SID of the Flex Flow | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the WebChannel resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the WebChannel resource | [optional] 

## Methods

### NewFlexV1WebChannel

`func NewFlexV1WebChannel() *FlexV1WebChannel`

NewFlexV1WebChannel instantiates a new FlexV1WebChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexV1WebChannelWithDefaults

`func NewFlexV1WebChannelWithDefaults() *FlexV1WebChannel`

NewFlexV1WebChannelWithDefaults instantiates a new FlexV1WebChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FlexV1WebChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FlexV1WebChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FlexV1WebChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FlexV1WebChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FlexV1WebChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FlexV1WebChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *FlexV1WebChannel) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FlexV1WebChannel) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FlexV1WebChannel) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FlexV1WebChannel) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FlexV1WebChannel) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FlexV1WebChannel) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FlexV1WebChannel) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FlexV1WebChannel) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FlexV1WebChannel) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FlexV1WebChannel) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FlexV1WebChannel) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FlexV1WebChannel) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFlexFlowSid

`func (o *FlexV1WebChannel) GetFlexFlowSid() string`

GetFlexFlowSid returns the FlexFlowSid field if non-nil, zero value otherwise.

### GetFlexFlowSidOk

`func (o *FlexV1WebChannel) GetFlexFlowSidOk() (*string, bool)`

GetFlexFlowSidOk returns a tuple with the FlexFlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlowSid

`func (o *FlexV1WebChannel) SetFlexFlowSid(v string)`

SetFlexFlowSid sets FlexFlowSid field to given value.

### HasFlexFlowSid

`func (o *FlexV1WebChannel) HasFlexFlowSid() bool`

HasFlexFlowSid returns a boolean if a field has been set.

### SetFlexFlowSidNil

`func (o *FlexV1WebChannel) SetFlexFlowSidNil(b bool)`

 SetFlexFlowSidNil sets the value for FlexFlowSid to be an explicit nil

### UnsetFlexFlowSid
`func (o *FlexV1WebChannel) UnsetFlexFlowSid()`

UnsetFlexFlowSid ensures that no value is present for FlexFlowSid, not even an explicit nil
### GetSid

`func (o *FlexV1WebChannel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *FlexV1WebChannel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *FlexV1WebChannel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *FlexV1WebChannel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *FlexV1WebChannel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *FlexV1WebChannel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *FlexV1WebChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FlexV1WebChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FlexV1WebChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FlexV1WebChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FlexV1WebChannel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FlexV1WebChannel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


