# AutopilotV1AssistantWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**AssistantSid** | Pointer to **NullableString** | The SID of the Assistant that is the parent of the resource | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Events** | Pointer to **NullableString** | The list of space-separated events that this Webhook is subscribed to. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Webhook resource | [optional] 
**WebhookMethod** | Pointer to **NullableString** | The method used when calling the webhook&#39;s URL. | [optional] 
**WebhookUrl** | Pointer to **NullableString** | The URL associated with this Webhook. | [optional] 

## Methods

### NewAutopilotV1AssistantWebhook

`func NewAutopilotV1AssistantWebhook() *AutopilotV1AssistantWebhook`

NewAutopilotV1AssistantWebhook instantiates a new AutopilotV1AssistantWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutopilotV1AssistantWebhookWithDefaults

`func NewAutopilotV1AssistantWebhookWithDefaults() *AutopilotV1AssistantWebhook`

NewAutopilotV1AssistantWebhookWithDefaults instantiates a new AutopilotV1AssistantWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *AutopilotV1AssistantWebhook) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *AutopilotV1AssistantWebhook) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *AutopilotV1AssistantWebhook) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *AutopilotV1AssistantWebhook) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *AutopilotV1AssistantWebhook) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *AutopilotV1AssistantWebhook) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAssistantSid

`func (o *AutopilotV1AssistantWebhook) GetAssistantSid() string`

GetAssistantSid returns the AssistantSid field if non-nil, zero value otherwise.

### GetAssistantSidOk

`func (o *AutopilotV1AssistantWebhook) GetAssistantSidOk() (*string, bool)`

GetAssistantSidOk returns a tuple with the AssistantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantSid

`func (o *AutopilotV1AssistantWebhook) SetAssistantSid(v string)`

SetAssistantSid sets AssistantSid field to given value.

### HasAssistantSid

`func (o *AutopilotV1AssistantWebhook) HasAssistantSid() bool`

HasAssistantSid returns a boolean if a field has been set.

### SetAssistantSidNil

`func (o *AutopilotV1AssistantWebhook) SetAssistantSidNil(b bool)`

 SetAssistantSidNil sets the value for AssistantSid to be an explicit nil

### UnsetAssistantSid
`func (o *AutopilotV1AssistantWebhook) UnsetAssistantSid()`

UnsetAssistantSid ensures that no value is present for AssistantSid, not even an explicit nil
### GetDateCreated

`func (o *AutopilotV1AssistantWebhook) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AutopilotV1AssistantWebhook) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AutopilotV1AssistantWebhook) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AutopilotV1AssistantWebhook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AutopilotV1AssistantWebhook) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AutopilotV1AssistantWebhook) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *AutopilotV1AssistantWebhook) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AutopilotV1AssistantWebhook) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AutopilotV1AssistantWebhook) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AutopilotV1AssistantWebhook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *AutopilotV1AssistantWebhook) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *AutopilotV1AssistantWebhook) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEvents

`func (o *AutopilotV1AssistantWebhook) GetEvents() string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AutopilotV1AssistantWebhook) GetEventsOk() (*string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AutopilotV1AssistantWebhook) SetEvents(v string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AutopilotV1AssistantWebhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *AutopilotV1AssistantWebhook) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *AutopilotV1AssistantWebhook) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetSid

`func (o *AutopilotV1AssistantWebhook) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AutopilotV1AssistantWebhook) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AutopilotV1AssistantWebhook) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AutopilotV1AssistantWebhook) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *AutopilotV1AssistantWebhook) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *AutopilotV1AssistantWebhook) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *AutopilotV1AssistantWebhook) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *AutopilotV1AssistantWebhook) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *AutopilotV1AssistantWebhook) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *AutopilotV1AssistantWebhook) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *AutopilotV1AssistantWebhook) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *AutopilotV1AssistantWebhook) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *AutopilotV1AssistantWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutopilotV1AssistantWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutopilotV1AssistantWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutopilotV1AssistantWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AutopilotV1AssistantWebhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AutopilotV1AssistantWebhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookMethod

`func (o *AutopilotV1AssistantWebhook) GetWebhookMethod() string`

GetWebhookMethod returns the WebhookMethod field if non-nil, zero value otherwise.

### GetWebhookMethodOk

`func (o *AutopilotV1AssistantWebhook) GetWebhookMethodOk() (*string, bool)`

GetWebhookMethodOk returns a tuple with the WebhookMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookMethod

`func (o *AutopilotV1AssistantWebhook) SetWebhookMethod(v string)`

SetWebhookMethod sets WebhookMethod field to given value.

### HasWebhookMethod

`func (o *AutopilotV1AssistantWebhook) HasWebhookMethod() bool`

HasWebhookMethod returns a boolean if a field has been set.

### SetWebhookMethodNil

`func (o *AutopilotV1AssistantWebhook) SetWebhookMethodNil(b bool)`

 SetWebhookMethodNil sets the value for WebhookMethod to be an explicit nil

### UnsetWebhookMethod
`func (o *AutopilotV1AssistantWebhook) UnsetWebhookMethod()`

UnsetWebhookMethod ensures that no value is present for WebhookMethod, not even an explicit nil
### GetWebhookUrl

`func (o *AutopilotV1AssistantWebhook) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *AutopilotV1AssistantWebhook) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *AutopilotV1AssistantWebhook) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *AutopilotV1AssistantWebhook) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *AutopilotV1AssistantWebhook) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *AutopilotV1AssistantWebhook) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


