# AutopilotV1AssistantQuery

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**DialogueSid** | Pointer to **NullableString** | The SID of the [Dialogue](https://www.twilio.com/docs/autopilot/api/dialogue). | [optional] 
**Language** | Pointer to **NullableString** | The ISO language-country string that specifies the language used by the Query | [optional] 
**ModelBuildSid** | Pointer to **NullableString** | The SID of the [Model Build](https://www.twilio.com/docs/autopilot/api/model-build) queried | [optional] 
**Query** | Pointer to **NullableString** | The end-user&#39;s natural language input | [optional] 
**Results** | Pointer to **map[string]interface{}** | The natural language analysis results that include the Task recognized and a list of identified Fields | [optional] 
**SampleSid** | Pointer to **NullableString** | The SID of an optional reference to the Sample created from the query | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SourceChannel** | Pointer to **NullableString** | The communication channel from where the end-user input came | [optional] 
**Status** | Pointer to **NullableString** | The status of the Query | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Query resource | [optional] 

## Methods

### NewAutopilotV1AssistantQuery

`func NewAutopilotV1AssistantQuery() *AutopilotV1AssistantQuery`

NewAutopilotV1AssistantQuery instantiates a new AutopilotV1AssistantQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantQueryWithDefaults

`func NewAutopilotV1AssistantQueryWithDefaults() *AutopilotV1AssistantQuery`

NewAutopilotV1AssistantQueryWithDefaults instantiates a new AutopilotV1AssistantQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantQuery) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantQuery) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantQuery) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantQuery) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantQuery) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantQuery) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantQuery) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantQuery) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantQuery) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantQuery) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantQuery) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantQuery) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantQuery) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantQuery) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantQuery) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantQuery) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantQuery) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantQuery) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantQuery) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantQuery) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantQuery) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantQuery) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantQuery) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantQuery) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDialogueSid

`func (o *AutopilotV1AssistantQuery) GetDialogueSid() string`

GetDialogueSid returns the DialogueSid field if non-nil, zero value otherwise.

### GetDialogueSidOk

`func (o *AutopilotV1AssistantQuery) GetDialogueSidOk() (*string, bool)`

GetDialogueSidOk returns a tuple with the DialogueSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialogueSid

`func (o *AutopilotV1AssistantQuery) SetDialogueSid(v string)`

SetDialogueSid sets DialogueSid field to given value.

### HasDialogueSid

`func (o *AutopilotV1AssistantQuery) HasDialogueSid() bool`

HasDialogueSid returns a boolean if a field has been set.

### SetDialogueSidNil

`func (o *AutopilotV1AssistantQuery) SetDialogueSidNil(b bool)`

 SetDialogueSidNil sets the value for DialogueSid to be an explicit nil

### UnsetDialogueSid
`func (o *AutopilotV1AssistantQuery) UnsetDialogueSid()`

UnsetDialogueSid ensures that no value is present for DialogueSid, not even an explicit nil
### GetLanguage

`func (o *AutopilotV1AssistantQuery) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AutopilotV1AssistantQuery) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AutopilotV1AssistantQuery) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AutopilotV1AssistantQuery) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AutopilotV1AssistantQuery) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AutopilotV1AssistantQuery) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetModelBuildSid

`func (o *AutopilotV1AssistantQuery) GetModelBuildSid() string`

GetModelBuildSid returns the ModelBuildSid field if non-nil, zero value otherwise.

### GetModelBuildSidOk

`func (o *AutopilotV1AssistantQuery) GetModelBuildSidOk() (*string, bool)`

GetModelBuildSidOk returns a tuple with the ModelBuildSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelBuildSid

`func (o *AutopilotV1AssistantQuery) SetModelBuildSid(v string)`

SetModelBuildSid sets ModelBuildSid field to given value.

### HasModelBuildSid

`func (o *AutopilotV1AssistantQuery) HasModelBuildSid() bool`

HasModelBuildSid returns a boolean if a field has been set.

### SetModelBuildSidNil

`func (o *AutopilotV1AssistantQuery) SetModelBuildSidNil(b bool)`

 SetModelBuildSidNil sets the value for ModelBuildSid to be an explicit nil

### UnsetModelBuildSid
`func (o *AutopilotV1AssistantQuery) UnsetModelBuildSid()`

UnsetModelBuildSid ensures that no value is present for ModelBuildSid, not even an explicit nil
### GetQuery

`func (o *AutopilotV1AssistantQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AutopilotV1AssistantQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AutopilotV1AssistantQuery) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AutopilotV1AssistantQuery) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *AutopilotV1AssistantQuery) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *AutopilotV1AssistantQuery) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetResults

`func (o *AutopilotV1AssistantQuery) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AutopilotV1AssistantQuery) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AutopilotV1AssistantQuery) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *AutopilotV1AssistantQuery) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *AutopilotV1AssistantQuery) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *AutopilotV1AssistantQuery) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSampleSid

`func (o *AutopilotV1AssistantQuery) GetSampleSid() string`

GetSampleSid returns the SampleSid field if non-nil, zero value otherwise.

### GetSampleSidOk

`func (o *AutopilotV1AssistantQuery) GetSampleSidOk() (*string, bool)`

GetSampleSidOk returns a tuple with the SampleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSid

`func (o *AutopilotV1AssistantQuery) SetSampleSid(v string)`

SetSampleSid sets SampleSid field to given value.

### HasSampleSid

`func (o *AutopilotV1AssistantQuery) HasSampleSid() bool`

HasSampleSid returns a boolean if a field has been set.

### SetSampleSidNil

`func (o *AutopilotV1AssistantQuery) SetSampleSidNil(b bool)`

 SetSampleSidNil sets the value for SampleSid to be an explicit nil

### UnsetSampleSid
`func (o *AutopilotV1AssistantQuery) UnsetSampleSid()`

UnsetSampleSid ensures that no value is present for SampleSid, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantQuery) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantQuery) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantQuery) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantQuery) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantQuery) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantQuery) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSourceChannel

`func (o *AutopilotV1AssistantQuery) GetSourceChannel() string`

GetSourceChannel returns the SourceChannel field if non-nil, zero value otherwise.

### GetSourceChannelOk

`func (o *AutopilotV1AssistantQuery) GetSourceChannelOk() (*string, bool)`

GetSourceChannelOk returns a tuple with the SourceChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChannel

`func (o *AutopilotV1AssistantQuery) SetSourceChannel(v string)`

SetSourceChannel sets SourceChannel field to given value.

### HasSourceChannel

`func (o *AutopilotV1AssistantQuery) HasSourceChannel() bool`

HasSourceChannel returns a boolean if a field has been set.

### SetSourceChannelNil

`func (o *AutopilotV1AssistantQuery) SetSourceChannelNil(b bool)`

 SetSourceChannelNil sets the value for SourceChannel to be an explicit nil

### UnsetSourceChannel
`func (o *AutopilotV1AssistantQuery) UnsetSourceChannel()`

UnsetSourceChannel ensures that no value is present for SourceChannel, not even an explicit nil
### GetStatus

`func (o *AutopilotV1AssistantQuery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutopilotV1AssistantQuery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutopilotV1AssistantQuery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutopilotV1AssistantQuery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AutopilotV1AssistantQuery) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AutopilotV1AssistantQuery) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantQuery) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantQuery) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantQuery) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantQuery) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantQuery) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantQuery) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


