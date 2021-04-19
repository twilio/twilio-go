# ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this conversation. | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the resource is associated with. | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | The configuration of this webhook. | [optional] 
**ConversationSid** | Pointer to **NullableString** | The unique ID of the Conversation for this webhook. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Target** | Pointer to **NullableString** | The target of this webhook. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this webhook. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhook

`func NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhook() *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook`

NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhook instantiates a new ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhookWithDefaults

`func NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhookWithDefaults() *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook`

NewConversationsV1ServiceServiceConversationServiceConversationScopedWebhookWithDefaults instantiates a new ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetConfiguration

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetTarget

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConversationServiceConversationScopedWebhook) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


