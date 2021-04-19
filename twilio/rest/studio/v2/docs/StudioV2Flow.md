# StudioV2Flow

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
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs | [optional] 
**Revision** | Pointer to **NullableInt32** | The latest revision number of the Flow&#39;s definition | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Status** | Pointer to **NullableString** | The status of the Flow | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**Valid** | Pointer to **NullableBool** | Boolean if the flow definition is valid | [optional] 
**Warnings** | Pointer to **[]map[string]interface{}** | List of warnings in the flow definition | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStudioV2Flow

`func NewStudioV2Flow() *StudioV2Flow`

NewStudioV2Flow instantiates a new StudioV2Flow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioV2FlowWithDefaults

`func NewStudioV2FlowWithDefaults() *StudioV2Flow`

NewStudioV2FlowWithDefaults instantiates a new StudioV2Flow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *StudioV2Flow) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *StudioV2Flow) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *StudioV2Flow) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *StudioV2Flow) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *StudioV2Flow) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *StudioV2Flow) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCommitMessage

`func (o *StudioV2Flow) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *StudioV2Flow) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *StudioV2Flow) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *StudioV2Flow) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.

### SetCommitMessageNil

`func (o *StudioV2Flow) SetCommitMessageNil(b bool)`

 SetCommitMessageNil sets the value for CommitMessage to be an explicit nil

### UnsetCommitMessage
`func (o *StudioV2Flow) UnsetCommitMessage()`

UnsetCommitMessage ensures that no value is present for CommitMessage, not even an explicit nil
### GetDateCreated

`func (o *StudioV2Flow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StudioV2Flow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StudioV2Flow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StudioV2Flow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *StudioV2Flow) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *StudioV2Flow) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *StudioV2Flow) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *StudioV2Flow) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *StudioV2Flow) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *StudioV2Flow) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *StudioV2Flow) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *StudioV2Flow) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefinition

`func (o *StudioV2Flow) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *StudioV2Flow) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *StudioV2Flow) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *StudioV2Flow) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *StudioV2Flow) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *StudioV2Flow) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetErrors

`func (o *StudioV2Flow) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StudioV2Flow) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StudioV2Flow) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *StudioV2Flow) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *StudioV2Flow) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *StudioV2Flow) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetFriendlyName

`func (o *StudioV2Flow) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *StudioV2Flow) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *StudioV2Flow) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *StudioV2Flow) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *StudioV2Flow) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *StudioV2Flow) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *StudioV2Flow) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StudioV2Flow) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StudioV2Flow) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StudioV2Flow) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StudioV2Flow) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *StudioV2Flow) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRevision

`func (o *StudioV2Flow) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *StudioV2Flow) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *StudioV2Flow) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *StudioV2Flow) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *StudioV2Flow) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *StudioV2Flow) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetSid

`func (o *StudioV2Flow) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *StudioV2Flow) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *StudioV2Flow) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *StudioV2Flow) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *StudioV2Flow) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *StudioV2Flow) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *StudioV2Flow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioV2Flow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioV2Flow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioV2Flow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *StudioV2Flow) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *StudioV2Flow) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *StudioV2Flow) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StudioV2Flow) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StudioV2Flow) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StudioV2Flow) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StudioV2Flow) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StudioV2Flow) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValid

`func (o *StudioV2Flow) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *StudioV2Flow) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *StudioV2Flow) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *StudioV2Flow) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *StudioV2Flow) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *StudioV2Flow) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetWarnings

`func (o *StudioV2Flow) GetWarnings() []map[string]interface{}`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *StudioV2Flow) GetWarningsOk() (*[]map[string]interface{}, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *StudioV2Flow) SetWarnings(v []map[string]interface{})`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *StudioV2Flow) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *StudioV2Flow) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *StudioV2Flow) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil
### GetWebhookUrl

`func (o *StudioV2Flow) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *StudioV2Flow) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *StudioV2Flow) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *StudioV2Flow) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *StudioV2Flow) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *StudioV2Flow) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


