# AutopilotV1AssistantDialogue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the resource | [optional] 
**Data** | Pointer to **map[string]interface{}** | The JSON string that describes the dialogue session object | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Dialogue resource | [optional] 

## Methods

### NewAutopilotV1AssistantDialogue

`func NewAutopilotV1AssistantDialogue() *AutopilotV1AssistantDialogue`

NewAutopilotV1AssistantDialogue instantiates a new AutopilotV1AssistantDialogue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantDialogueWithDefaults

`func NewAutopilotV1AssistantDialogueWithDefaults() *AutopilotV1AssistantDialogue`

NewAutopilotV1AssistantDialogueWithDefaults instantiates a new AutopilotV1AssistantDialogue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantDialogue) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantDialogue) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantDialogue) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantDialogue) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantDialogue) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantDialogue) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantDialogue) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantDialogue) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantDialogue) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantDialogue) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantDialogue) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantDialogue) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetData

`func (o *AutopilotV1AssistantDialogue) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AutopilotV1AssistantDialogue) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AutopilotV1AssistantDialogue) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AutopilotV1AssistantDialogue) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *AutopilotV1AssistantDialogue) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *AutopilotV1AssistantDialogue) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantDialogue) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantDialogue) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantDialogue) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantDialogue) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantDialogue) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantDialogue) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantDialogue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantDialogue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantDialogue) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantDialogue) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantDialogue) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantDialogue) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


