# AutopilotV1AssistantTaskField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the Task associated with the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FieldType** | Pointer to **NullableString** | The Field Type of the field | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the [Task](https://www.twilio.com/docs/autopilot/api/task) resource associated with this Field | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Field resource | [optional] 

## Methods

### NewAutopilotV1AssistantTaskField

`func NewAutopilotV1AssistantTaskField() *AutopilotV1AssistantTaskField`

NewAutopilotV1AssistantTaskField instantiates a new AutopilotV1AssistantTaskField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantTaskFieldWithDefaults

`func NewAutopilotV1AssistantTaskFieldWithDefaults() *AutopilotV1AssistantTaskField`

NewAutopilotV1AssistantTaskFieldWithDefaults instantiates a new AutopilotV1AssistantTaskField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantTaskField) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantTaskField) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantTaskField) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantTaskField) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantTaskField) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantTaskField) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantTaskField) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantTaskField) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantTaskField) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantTaskField) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantTaskField) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantTaskField) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantTaskField) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantTaskField) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantTaskField) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantTaskField) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantTaskField) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantTaskField) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantTaskField) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantTaskField) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantTaskField) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantTaskField) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantTaskField) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantTaskField) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFieldType

`func (o *AutopilotV1AssistantTaskField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *AutopilotV1AssistantTaskField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *AutopilotV1AssistantTaskField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *AutopilotV1AssistantTaskField) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### SetFieldTypeNil

`func (o *AutopilotV1AssistantTaskField) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *AutopilotV1AssistantTaskField) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantTaskField) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantTaskField) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantTaskField) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantTaskField) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantTaskField) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantTaskField) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTaskSid

`func (o *AutopilotV1AssistantTaskField) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *AutopilotV1AssistantTaskField) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *AutopilotV1AssistantTaskField) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *AutopilotV1AssistantTaskField) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *AutopilotV1AssistantTaskField) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *AutopilotV1AssistantTaskField) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUniqueName

`func (o *AutopilotV1AssistantTaskField) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AutopilotV1AssistantTaskField) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AutopilotV1AssistantTaskField) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AutopilotV1AssistantTaskField) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *AutopilotV1AssistantTaskField) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *AutopilotV1AssistantTaskField) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantTaskField) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantTaskField) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantTaskField) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantTaskField) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantTaskField) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantTaskField) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


