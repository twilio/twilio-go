# AutopilotV1AssistantFieldTypeFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the FieldType associated with the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**FieldTypeSid** | Pointer to **NullableString** | The SID of the Field Type associated with the Field Value | [optional] 
**Language** | Pointer to **NullableString** | The ISO language-country tag that identifies the language of the value | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**SynonymOf** | Pointer to **NullableString** | The word for which the field value is a synonym of | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the FieldValue resource | [optional] 
**Value** | Pointer to **NullableString** | The Field Value data | [optional] 

## Methods

### NewAutopilotV1AssistantFieldTypeFieldValue

`func NewAutopilotV1AssistantFieldTypeFieldValue() *AutopilotV1AssistantFieldTypeFieldValue`

NewAutopilotV1AssistantFieldTypeFieldValue instantiates a new AutopilotV1AssistantFieldTypeFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantFieldTypeFieldValueWithDefaults

`func NewAutopilotV1AssistantFieldTypeFieldValueWithDefaults() *AutopilotV1AssistantFieldTypeFieldValue`

NewAutopilotV1AssistantFieldTypeFieldValueWithDefaults instantiates a new AutopilotV1AssistantFieldTypeFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFieldTypeSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetFieldTypeSid() string`

GetFieldTypeSid returns the FieldTypeSid field if non-nil, zero value otherwise.

### GetFieldTypeSidOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetFieldTypeSidOk() (*string, bool)`

GetFieldTypeSidOk returns a tuple with the FieldTypeSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypeSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetFieldTypeSid(v string)`

SetFieldTypeSid sets FieldTypeSid field to given value.

### HasFieldTypeSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasFieldTypeSid() bool`

HasFieldTypeSid returns a boolean if a field has been set.

### SetFieldTypeSidNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetFieldTypeSidNil(b bool)`

 SetFieldTypeSidNil sets the value for FieldTypeSid to be an explicit nil

### UnsetFieldTypeSid
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetFieldTypeSid()`

UnsetFieldTypeSid ensures that no value is present for FieldTypeSid, not even an explicit nil
### GetLanguage

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetSynonymOf

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetSynonymOf() string`

GetSynonymOf returns the SynonymOf field if non-nil, zero value otherwise.

### GetSynonymOfOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetSynonymOfOk() (*string, bool)`

GetSynonymOfOk returns a tuple with the SynonymOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonymOf

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetSynonymOf(v string)`

SetSynonymOf sets SynonymOf field to given value.

### HasSynonymOf

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasSynonymOf() bool`

HasSynonymOf returns a boolean if a field has been set.

### SetSynonymOfNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetSynonymOfNil(b bool)`

 SetSynonymOfNil sets the value for SynonymOf to be an explicit nil

### UnsetSynonymOf
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetSynonymOf()`

UnsetSynonymOf ensures that no value is present for SynonymOf, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetValue

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AutopilotV1AssistantFieldTypeFieldValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AutopilotV1AssistantFieldTypeFieldValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *AutopilotV1AssistantFieldTypeFieldValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AutopilotV1AssistantFieldTypeFieldValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


