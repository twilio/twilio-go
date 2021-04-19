# AutopilotV1AssistantTaskSample

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the Task associated with the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Language** | Pointer to **NullableString** | An ISO language-country string that specifies the language used for the sample | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SourceChannel** | Pointer to **NullableString** | The communication channel from which the sample was captured | [optional] 
**TaggedText** | Pointer to **NullableString** | The text example of how end users might express the task | [optional] 
**TaskSid** | Pointer to **NullableString** | The SID of the Task associated with the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Sample resource | [optional] 

## Methods

### NewAutopilotV1AssistantTaskSample

`func NewAutopilotV1AssistantTaskSample() *AutopilotV1AssistantTaskSample`

NewAutopilotV1AssistantTaskSample instantiates a new AutopilotV1AssistantTaskSample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantTaskSampleWithDefaults

`func NewAutopilotV1AssistantTaskSampleWithDefaults() *AutopilotV1AssistantTaskSample`

NewAutopilotV1AssistantTaskSampleWithDefaults instantiates a new AutopilotV1AssistantTaskSample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantTaskSample) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantTaskSample) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantTaskSample) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantTaskSample) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantTaskSample) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantTaskSample) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantTaskSample) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantTaskSample) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantTaskSample) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantTaskSample) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantTaskSample) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantTaskSample) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantTaskSample) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantTaskSample) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantTaskSample) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantTaskSample) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantTaskSample) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantTaskSample) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantTaskSample) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantTaskSample) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantTaskSample) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantTaskSample) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantTaskSample) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantTaskSample) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetLanguage

`func (o *AutopilotV1AssistantTaskSample) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AutopilotV1AssistantTaskSample) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AutopilotV1AssistantTaskSample) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AutopilotV1AssistantTaskSample) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AutopilotV1AssistantTaskSample) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AutopilotV1AssistantTaskSample) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantTaskSample) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantTaskSample) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantTaskSample) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantTaskSample) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantTaskSample) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantTaskSample) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSourceChannel

`func (o *AutopilotV1AssistantTaskSample) GetSourceChannel() string`

GetSourceChannel returns the SourceChannel field if non-nil, zero value otherwise.

### GetSourceChannelOk

`func (o *AutopilotV1AssistantTaskSample) GetSourceChannelOk() (*string, bool)`

GetSourceChannelOk returns a tuple with the SourceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChannel

`func (o *AutopilotV1AssistantTaskSample) SetSourceChannel(v string)`

SetSourceChannel sets SourceChannel field to given value.

### HasSourceChannel

`func (o *AutopilotV1AssistantTaskSample) HasSourceChannel() bool`

HasSourceChannel returns a boolean if a field has been set.

### SetSourceChannelNil

`func (o *AutopilotV1AssistantTaskSample) SetSourceChannelNil(b bool)`

 SetSourceChannelNil sets the value for SourceChannel to be an explicit nil

### UnsetSourceChannel
`func (o *AutopilotV1AssistantTaskSample) UnsetSourceChannel()`

UnsetSourceChannel ensures that no value is present for SourceChannel, not even an explicit nil
### GetTaggedText

`func (o *AutopilotV1AssistantTaskSample) GetTaggedText() string`

GetTaggedText returns the TaggedText field if non-nil, zero value otherwise.

### GetTaggedTextOk

`func (o *AutopilotV1AssistantTaskSample) GetTaggedTextOk() (*string, bool)`

GetTaggedTextOk returns a tuple with the TaggedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedText

`func (o *AutopilotV1AssistantTaskSample) SetTaggedText(v string)`

SetTaggedText sets TaggedText field to given value.

### HasTaggedText

`func (o *AutopilotV1AssistantTaskSample) HasTaggedText() bool`

HasTaggedText returns a boolean if a field has been set.

### SetTaggedTextNil

`func (o *AutopilotV1AssistantTaskSample) SetTaggedTextNil(b bool)`

 SetTaggedTextNil sets the value for TaggedText to be an explicit nil

### UnsetTaggedText
`func (o *AutopilotV1AssistantTaskSample) UnsetTaggedText()`

UnsetTaggedText ensures that no value is present for TaggedText, not even an explicit nil
### GetTaskSid

`func (o *AutopilotV1AssistantTaskSample) GetTaskSid() string`

GetTaskSid returns the TaskSid field if non-nil, zero value otherwise.

### GetTaskSidOk

`func (o *AutopilotV1AssistantTaskSample) GetTaskSidOk() (*string, bool)`

GetTaskSidOk returns a tuple with the TaskSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSid

`func (o *AutopilotV1AssistantTaskSample) SetTaskSid(v string)`

SetTaskSid sets TaskSid field to given value.

### HasTaskSid

`func (o *AutopilotV1AssistantTaskSample) HasTaskSid() bool`

HasTaskSid returns a boolean if a field has been set.

### SetTaskSidNil

`func (o *AutopilotV1AssistantTaskSample) SetTaskSidNil(b bool)`

 SetTaskSidNil sets the value for TaskSid to be an explicit nil

### UnsetTaskSid
`func (o *AutopilotV1AssistantTaskSample) UnsetTaskSid()`

UnsetTaskSid ensures that no value is present for TaskSid, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantTaskSample) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantTaskSample) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantTaskSample) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantTaskSample) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantTaskSample) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantTaskSample) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


