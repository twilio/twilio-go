# AutopilotV1AssistantTaskTaskStatistics

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the Task associated with the resource | [optional] 
**FieldsCount** | Pointer to **NullableInt32** | The total number of Fields associated with the Task | [optional] 
**SamplesCount** | Pointer to **NullableInt32** | The total number of Samples associated with the Task | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the Task for which the statistics were collected | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the TaskStatistics resource | [optional] 

## Methods

### NewAutopilotV1AssistantTaskTaskStatistics

`func NewAutopilotV1AssistantTaskTaskStatistics() *AutopilotV1AssistantTaskTaskStatistics`

NewAutopilotV1AssistantTaskTaskStatistics instantiates a new AutopilotV1AssistantTaskTaskStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantTaskTaskStatisticsWithDefaults

`func NewAutopilotV1AssistantTaskTaskStatisticsWithDefaults() *AutopilotV1AssistantTaskTaskStatistics`

NewAutopilotV1AssistantTaskTaskStatisticsWithDefaults instantiates a new AutopilotV1AssistantTaskTaskStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetFieldsCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetFieldsCount() int32`

GetFieldsCount returns the FieldsCount field if non-nil, zero value otherwise.

### GetFieldsCountOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetFieldsCountOk() (*int32, bool)`

GetFieldsCountOk returns a tuple with the FieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetFieldsCount(v int32)`

SetFieldsCount sets FieldsCount field to given value.

### HasFieldsCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasFieldsCount() bool`

HasFieldsCount returns a boolean if a field has been set.

### SetFieldsCountNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetFieldsCountNil(b bool)`

 SetFieldsCountNil sets the value for FieldsCount to be an explicit nil

### UnsetFieldsCount
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetFieldsCount()`

UnsetFieldsCount ensures that no value is present for FieldsCount, not even an explicit nil
### GetSamplesCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetSamplesCount() int32`

GetSamplesCount returns the SamplesCount field if non-nil, zero value otherwise.

### GetSamplesCountOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetSamplesCountOk() (*int32, bool)`

GetSamplesCountOk returns a tuple with the SamplesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetSamplesCount(v int32)`

SetSamplesCount sets SamplesCount field to given value.

### HasSamplesCount

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasSamplesCount() bool`

HasSamplesCount returns a boolean if a field has been set.

### SetSamplesCountNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetSamplesCountNil(b bool)`

 SetSamplesCountNil sets the value for SamplesCount to be an explicit nil

### UnsetSamplesCount
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetSamplesCount()`

UnsetSamplesCount ensures that no value is present for SamplesCount, not even an explicit nil
### GetTaskSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantTaskTaskStatistics) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantTaskTaskStatistics) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantTaskTaskStatistics) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantTaskTaskStatistics) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


