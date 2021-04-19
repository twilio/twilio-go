# ConversationsV1ServiceServiceConfigurationServiceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this configuration. | [optional] 
**AddedToConversation** | Pointer to **map[string]interface{}** | The Push Notification configuration for being added to a Conversation. | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the Configuration applies to. | [optional] 
**LogEnabled** | Pointer to **NullableBool** | Weather the notification logging is enabled. | [optional] 
**NewMessage** | Pointer to **map[string]interface{}** | The Push Notification configuration for New Messages. | [optional] 
**RemovedFromConversation** | Pointer to **map[string]interface{}** | The Push Notification configuration for being removed from a Conversation. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this configuration. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConfigurationServiceNotification

`func NewConversationsV1ServiceServiceConfigurationServiceNotification() *ConversationsV1ServiceServiceConfigurationServiceNotification`

NewConversationsV1ServiceServiceConfigurationServiceNotification instantiates a new ConversationsV1ServiceServiceConfigurationServiceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConfigurationServiceNotificationWithDefaults

`func NewConversationsV1ServiceServiceConfigurationServiceNotificationWithDefaults() *ConversationsV1ServiceServiceConfigurationServiceNotification`

NewConversationsV1ServiceServiceConfigurationServiceNotificationWithDefaults instantiates a new ConversationsV1ServiceServiceConfigurationServiceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAddedToConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetAddedToConversation() map[string]interface{}`

GetAddedToConversation returns the AddedToConversation field if non-nil, zero value otherwise.

### GetAddedToConversationOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetAddedToConversationOk() (*map[string]interface{}, bool)`

GetAddedToConversationOk returns a tuple with the AddedToConversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedToConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetAddedToConversation(v map[string]interface{})`

SetAddedToConversation sets AddedToConversation field to given value.

### HasAddedToConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasAddedToConversation() bool`

HasAddedToConversation returns a boolean if a field has been set.

### SetAddedToConversationNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetAddedToConversationNil(b bool)`

 SetAddedToConversationNil sets the value for AddedToConversation to be an explicit nil

### UnsetAddedToConversation
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetAddedToConversation()`

UnsetAddedToConversation ensures that no value is present for AddedToConversation, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetLogEnabled

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetLogEnabled() bool`

GetLogEnabled returns the LogEnabled field if non-nil, zero value otherwise.

### GetLogEnabledOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetLogEnabledOk() (*bool, bool)`

GetLogEnabledOk returns a tuple with the LogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogEnabled

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetLogEnabled(v bool)`

SetLogEnabled sets LogEnabled field to given value.

### HasLogEnabled

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasLogEnabled() bool`

HasLogEnabled returns a boolean if a field has been set.

### SetLogEnabledNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetLogEnabledNil(b bool)`

 SetLogEnabledNil sets the value for LogEnabled to be an explicit nil

### UnsetLogEnabled
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetLogEnabled()`

UnsetLogEnabled ensures that no value is present for LogEnabled, not even an explicit nil
### GetNewMessage

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetNewMessage() map[string]interface{}`

GetNewMessage returns the NewMessage field if non-nil, zero value otherwise.

### GetNewMessageOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetNewMessageOk() (*map[string]interface{}, bool)`

GetNewMessageOk returns a tuple with the NewMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMessage

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetNewMessage(v map[string]interface{})`

SetNewMessage sets NewMessage field to given value.

### HasNewMessage

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasNewMessage() bool`

HasNewMessage returns a boolean if a field has been set.

### SetNewMessageNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetNewMessageNil(b bool)`

 SetNewMessageNil sets the value for NewMessage to be an explicit nil

### UnsetNewMessage
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetNewMessage()`

UnsetNewMessage ensures that no value is present for NewMessage, not even an explicit nil
### GetRemovedFromConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetRemovedFromConversation() map[string]interface{}`

GetRemovedFromConversation returns the RemovedFromConversation field if non-nil, zero value otherwise.

### GetRemovedFromConversationOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetRemovedFromConversationOk() (*map[string]interface{}, bool)`

GetRemovedFromConversationOk returns a tuple with the RemovedFromConversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedFromConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetRemovedFromConversation(v map[string]interface{})`

SetRemovedFromConversation sets RemovedFromConversation field to given value.

### HasRemovedFromConversation

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasRemovedFromConversation() bool`

HasRemovedFromConversation returns a boolean if a field has been set.

### SetRemovedFromConversationNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetRemovedFromConversationNil(b bool)`

 SetRemovedFromConversationNil sets the value for RemovedFromConversation to be an explicit nil

### UnsetRemovedFromConversation
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetRemovedFromConversation()`

UnsetRemovedFromConversation ensures that no value is present for RemovedFromConversation, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConfigurationServiceNotification) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


