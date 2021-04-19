# FlexV1Channel

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource and owns this Workflow | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Flex chat channel was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the Flex chat channel was last updated | [optional] 
**FlexFlowSid** | Pointer to **NullableString** | The SID of the Flex Flow | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the TaskRouter Task | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Flex chat channel resource | [optional] 
**UserSid** | Pointer to **NullableString** | The SID of the chat user | [optional] 

## Methods

### NewFlexV1Channel

`func NewFlexV1Channel() *FlexV1Channel`

NewFlexV1Channel instantiates a new FlexV1Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexV1ChannelWithDefaults

`func NewFlexV1ChannelWithDefaults() *FlexV1Channel`

NewFlexV1ChannelWithDefaults instantiates a new FlexV1Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *FlexV1Channel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *FlexV1Channel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *FlexV1Channel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *FlexV1Channel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *FlexV1Channel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *FlexV1Channel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDateCreated

`func (o *FlexV1Channel) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FlexV1Channel) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FlexV1Channel) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FlexV1Channel) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *FlexV1Channel) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *FlexV1Channel) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *FlexV1Channel) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *FlexV1Channel) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *FlexV1Channel) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *FlexV1Channel) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *FlexV1Channel) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *FlexV1Channel) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFlexFlowSid

`func (o *FlexV1Channel) GetFlexFlowSid() string`

GetFlexFlowSid returns the FlexFlowSid field if non-nil, zero value otherwise.

### GetFlexFlowSidOk

`func (o *FlexV1Channel) GetFlexFlowSidOk() (*string, bool)`

GetFlexFlowSidOk returns a tuple with the FlexFlowSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlexFlowSid

`func (o *FlexV1Channel) SetFlexFlowSid(v string)`

SetFlexFlowSid sets FlexFlowSid field to given value.

### HasFlexFlowSid

`func (o *FlexV1Channel) HasFlexFlowSid() bool`

HasFlexFlowSid returns a boolean if a field has been set.

### SetFlexFlowSidNil

`func (o *FlexV1Channel) SetFlexFlowSidNil(b bool)`

 SetFlexFlowSidNil sets the value for FlexFlowSid to be an explicit nil

### UnsetFlexFlowSid
`func (o *FlexV1Channel) UnsetFlexFlowSid()`

UnsetFlexFlowSid ensures that no value is present for FlexFlowSid, not even an explicit nil
### GetSid

`func (o *FlexV1Channel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *FlexV1Channel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *FlexV1Channel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *FlexV1Channel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *FlexV1Channel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *FlexV1Channel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskSid

`func (o *FlexV1Channel) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *FlexV1Channel) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *FlexV1Channel) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *FlexV1Channel) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *FlexV1Channel) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *FlexV1Channel) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *FlexV1Channel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FlexV1Channel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FlexV1Channel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FlexV1Channel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FlexV1Channel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FlexV1Channel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUserSid

`func (o *FlexV1Channel) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *FlexV1Channel) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *FlexV1Channel) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *FlexV1Channel) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *FlexV1Channel) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *FlexV1Channel) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


