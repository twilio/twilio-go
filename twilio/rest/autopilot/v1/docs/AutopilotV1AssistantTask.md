# AutopilotV1AssistantTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ActionsUrl** | Pointer to **NullableString** | The URL from which the Assistant can fetch actions | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**Links** | Pointer to **map[string]interface{}** | A list of the URLs of related resources | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Task resource | [optional] 

## Methods

### NewAutopilotV1AssistantTask

`func NewAutopilotV1AssistantTask() *AutopilotV1AssistantTask`

NewAutopilotV1AssistantTask instantiates a new AutopilotV1AssistantTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantTaskWithDefaults

`func NewAutopilotV1AssistantTaskWithDefaults() *AutopilotV1AssistantTask`

NewAutopilotV1AssistantTaskWithDefaults instantiates a new AutopilotV1AssistantTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantTask) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantTask) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantTask) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantTask) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantTask) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantTask) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetActionsUrl

`func (o *AutopilotV1AssistantTask) GetActionsUrl() string`

GetActionsUrl returns the ActionsUrl field if non-nil, zero value otherwise.

### GetActionsUrlOk

`func (o *AutopilotV1AssistantTask) GetActionsUrlOk() (*string, bool)`

GetActionsUrlOk returns a tuple with the ActionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsUrl

`func (o *AutopilotV1AssistantTask) SetActionsUrl(v string)`

SetActionsUrl sets ActionsUrl field to given value.

### HasActionsUrl

`func (o *AutopilotV1AssistantTask) HasActionsUrl() bool`

HasActionsUrl returns a boolean if a field has been set.

### SetActionsUrlNil

`func (o *AutopilotV1AssistantTask) SetActionsUrlNil(b bool)`

 SetActionsUrlNil sets the value for ActionsUrl to be an explicit nil

### UnsetActionsUrl
`func (o *AutopilotV1AssistantTask) UnsetActionsUrl()`

UnsetActionsUrl ensures that no value is present for ActionsUrl, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantTask) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantTask) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantTask) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantTask) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantTask) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantTask) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantTask) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantTask) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantTask) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantTask) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantTask) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantTask) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantTask) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantTask) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *AutopilotV1AssistantTask) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AutopilotV1AssistantTask) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AutopilotV1AssistantTask) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AutopilotV1AssistantTask) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AutopilotV1AssistantTask) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AutopilotV1AssistantTask) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *AutopilotV1AssistantTask) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AutopilotV1AssistantTask) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AutopilotV1AssistantTask) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AutopilotV1AssistantTask) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *AutopilotV1AssistantTask) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *AutopilotV1AssistantTask) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantTask) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantTask) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantTask) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantTask) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantTask) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantTask) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *AutopilotV1AssistantTask) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AutopilotV1AssistantTask) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AutopilotV1AssistantTask) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AutopilotV1AssistantTask) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *AutopilotV1AssistantTask) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *AutopilotV1AssistantTask) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantTask) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantTask) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantTask) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantTask) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantTask) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantTask) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


