# AutopilotV1AssistantTaskTaskActions

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the Task associated with the resource | [optional] 
**Data** | Pointer to **map[string]interface{}** | The JSON string that specifies the actions that instruct the Assistant on how to perform the task | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the Task associated with the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskActions resource | [optional] 

## Methods

### NewAutopilotV1AssistantTaskTaskActions

`func NewAutopilotV1AssistantTaskTaskActions() *AutopilotV1AssistantTaskTaskActions`

NewAutopilotV1AssistantTaskTaskActions instantiates a new AutopilotV1AssistantTaskTaskActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantTaskTaskActionsWithDefaults

`func NewAutopilotV1AssistantTaskTaskActionsWithDefaults() *AutopilotV1AssistantTaskTaskActions`

NewAutopilotV1AssistantTaskTaskActionsWithDefaults instantiates a new AutopilotV1AssistantTaskTaskActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantTaskTaskActions) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantTaskTaskActions) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantTaskTaskActions) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantTaskTaskActions) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantTaskTaskActions) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantTaskTaskActions) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantTaskTaskActions) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantTaskTaskActions) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantTaskTaskActions) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantTaskTaskActions) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantTaskTaskActions) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantTaskTaskActions) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetData

`func (o *AutopilotV1AssistantTaskTaskActions) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutopilotV1AssistantTaskTaskActions) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutopilotV1AssistantTaskTaskActions) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AutopilotV1AssistantTaskTaskActions) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AutopilotV1AssistantTaskTaskActions) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AutopilotV1AssistantTaskTaskActions) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTaskSid

`func (o *AutopilotV1AssistantTaskTaskActions) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *AutopilotV1AssistantTaskTaskActions) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *AutopilotV1AssistantTaskTaskActions) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *AutopilotV1AssistantTaskTaskActions) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *AutopilotV1AssistantTaskTaskActions) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *AutopilotV1AssistantTaskTaskActions) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantTaskTaskActions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantTaskTaskActions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantTaskTaskActions) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantTaskTaskActions) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantTaskTaskActions) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantTaskTaskActions) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


