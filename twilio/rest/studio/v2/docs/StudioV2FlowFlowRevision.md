# StudioV2FlowFlowRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CommitMessage** | Pointer to **NullableString** | Description of change made in the revision | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Definition** | Pointer to **map[string]interface{}** | JSON representation of flow definition | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** | List of error in the flow definition | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the Flow | [optional] 
**Revision** | Pointer to **NullableInt32** | The latest revision number of the Flow&#39;s definition | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Flow | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**Valid** | Pointer to **NullableBool** | Boolean if the flow definition is valid | [optional] 

## Methods

### NewStudioV2FlowFlowRevision

`func NewStudioV2FlowFlowRevision() *StudioV2FlowFlowRevision`

NewStudioV2FlowFlowRevision instantiates a new StudioV2FlowFlowRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV2FlowFlowRevisionWithDefaults

`func NewStudioV2FlowFlowRevisionWithDefaults() *StudioV2FlowFlowRevision`

NewStudioV2FlowFlowRevisionWithDefaults instantiates a new StudioV2FlowFlowRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV2FlowFlowRevision) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV2FlowFlowRevision) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV2FlowFlowRevision) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV2FlowFlowRevision) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV2FlowFlowRevision) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV2FlowFlowRevision) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommitMessage

`func (o *StudioV2FlowFlowRevision) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *StudioV2FlowFlowRevision) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *StudioV2FlowFlowRevision) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *StudioV2FlowFlowRevision) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### SetCommitMessageNil

`func (o *StudioV2FlowFlowRevision) SetCommitMessageNil(b bool)`

 SetCommitMessageNil sets the value for CommitMessage to be an explicit nil

### UnsetCommitMessage
`func (o *StudioV2FlowFlowRevision) UnsetCommitMessage()`

UnsetCommitMessage ensures that no value is present for CommitMessage, not even an explicit nil
### GetDateCreated

`func (o *StudioV2FlowFlowRevision) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV2FlowFlowRevision) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV2FlowFlowRevision) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV2FlowFlowRevision) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV2FlowFlowRevision) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV2FlowFlowRevision) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV2FlowFlowRevision) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV2FlowFlowRevision) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV2FlowFlowRevision) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV2FlowFlowRevision) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV2FlowFlowRevision) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV2FlowFlowRevision) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefinition

`func (o *StudioV2FlowFlowRevision) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *StudioV2FlowFlowRevision) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *StudioV2FlowFlowRevision) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *StudioV2FlowFlowRevision) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *StudioV2FlowFlowRevision) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *StudioV2FlowFlowRevision) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetErrors

`func (o *StudioV2FlowFlowRevision) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StudioV2FlowFlowRevision) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StudioV2FlowFlowRevision) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *StudioV2FlowFlowRevision) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *StudioV2FlowFlowRevision) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *StudioV2FlowFlowRevision) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetFriendlyName

`func (o *StudioV2FlowFlowRevision) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *StudioV2FlowFlowRevision) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *StudioV2FlowFlowRevision) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *StudioV2FlowFlowRevision) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *StudioV2FlowFlowRevision) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *StudioV2FlowFlowRevision) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetRevision

`func (o *StudioV2FlowFlowRevision) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StudioV2FlowFlowRevision) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StudioV2FlowFlowRevision) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StudioV2FlowFlowRevision) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *StudioV2FlowFlowRevision) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *StudioV2FlowFlowRevision) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetSid

`func (o *StudioV2FlowFlowRevision) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV2FlowFlowRevision) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV2FlowFlowRevision) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV2FlowFlowRevision) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV2FlowFlowRevision) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV2FlowFlowRevision) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *StudioV2FlowFlowRevision) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioV2FlowFlowRevision) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioV2FlowFlowRevision) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioV2FlowFlowRevision) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StudioV2FlowFlowRevision) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StudioV2FlowFlowRevision) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *StudioV2FlowFlowRevision) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV2FlowFlowRevision) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV2FlowFlowRevision) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV2FlowFlowRevision) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV2FlowFlowRevision) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV2FlowFlowRevision) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValid

`func (o *StudioV2FlowFlowRevision) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *StudioV2FlowFlowRevision) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *StudioV2FlowFlowRevision) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *StudioV2FlowFlowRevision) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *StudioV2FlowFlowRevision) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *StudioV2FlowFlowRevision) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


