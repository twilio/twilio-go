# ConversationsV1Configuration

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account responsible for this configuration. | [optional] 
**DefaultChatServiceSid** | Pointer to **NullableString** | The SID of the default Conversation Service that every new conversation is associated with. | [optional] 
**DefaultClosedTimer** | Pointer to **NullableString** | Default ISO8601 duration when conversation will be switched to &#x60;closed&#x60; state. | [optional] 
**DefaultInactiveTimer** | Pointer to **NullableString** | Default ISO8601 duration when conversation will be switched to &#x60;inactive&#x60; state. | [optional] 
**DefaultMessagingServiceSid** | Pointer to **NullableString** | The SID of the default Messaging Service that every new conversation is associated with. | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the webhook and default service configurations. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this global configuration. | [optional] 

## Methods

### NewConversationsV1Configuration

`func NewConversationsV1Configuration() *ConversationsV1Configuration`

NewConversationsV1Configuration instantiates a new ConversationsV1Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ConfigurationWithDefaults

`func NewConversationsV1ConfigurationWithDefaults() *ConversationsV1Configuration`

NewConversationsV1ConfigurationWithDefaults instantiates a new ConversationsV1Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1Configuration) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1Configuration) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1Configuration) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1Configuration) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1Configuration) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1Configuration) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetDefaultChatServiceSid

`func (o *ConversationsV1Configuration) GetDefaultChatServiceSid() string`

GetDefaultChatServiceSid returns the DefaultChatServiceSid field if non-nil, zero value otherwise.

### GetDefaultChatServiceSidOk

`func (o *ConversationsV1Configuration) GetDefaultChatServiceSidOk() (*string, bool)`

GetDefaultChatServiceSidOk returns a tuple with the DefaultChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultChatServiceSid

`func (o *ConversationsV1Configuration) SetDefaultChatServiceSid(v string)`

SetDefaultChatServiceSid sets DefaultChatServiceSid field to given value.

### HasDefaultChatServiceSid

`func (o *ConversationsV1Configuration) HasDefaultChatServiceSid() bool`

HasDefaultChatServiceSid returns a boolean if a field has been set.

### SetDefaultChatServiceSidNil

`func (o *ConversationsV1Configuration) SetDefaultChatServiceSidNil(b bool)`

 SetDefaultChatServiceSidNil sets the value for DefaultChatServiceSid to be an explicit nil

### UnsetDefaultChatServiceSid
`func (o *ConversationsV1Configuration) UnsetDefaultChatServiceSid()`

UnsetDefaultChatServiceSid ensures that no value is present for DefaultChatServiceSid, not even an explicit nil
### GetDefaultClosedTimer

`func (o *ConversationsV1Configuration) GetDefaultClosedTimer() string`

GetDefaultClosedTimer returns the DefaultClosedTimer field if non-nil, zero value otherwise.

### GetDefaultClosedTimerOk

`func (o *ConversationsV1Configuration) GetDefaultClosedTimerOk() (*string, bool)`

GetDefaultClosedTimerOk returns a tuple with the DefaultClosedTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClosedTimer

`func (o *ConversationsV1Configuration) SetDefaultClosedTimer(v string)`

SetDefaultClosedTimer sets DefaultClosedTimer field to given value.

### HasDefaultClosedTimer

`func (o *ConversationsV1Configuration) HasDefaultClosedTimer() bool`

HasDefaultClosedTimer returns a boolean if a field has been set.

### SetDefaultClosedTimerNil

`func (o *ConversationsV1Configuration) SetDefaultClosedTimerNil(b bool)`

 SetDefaultClosedTimerNil sets the value for DefaultClosedTimer to be an explicit nil

### UnsetDefaultClosedTimer
`func (o *ConversationsV1Configuration) UnsetDefaultClosedTimer()`

UnsetDefaultClosedTimer ensures that no value is present for DefaultClosedTimer, not even an explicit nil
### GetDefaultInactiveTimer

`func (o *ConversationsV1Configuration) GetDefaultInactiveTimer() string`

GetDefaultInactiveTimer returns the DefaultInactiveTimer field if non-nil, zero value otherwise.

### GetDefaultInactiveTimerOk

`func (o *ConversationsV1Configuration) GetDefaultInactiveTimerOk() (*string, bool)`

GetDefaultInactiveTimerOk returns a tuple with the DefaultInactiveTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInactiveTimer

`func (o *ConversationsV1Configuration) SetDefaultInactiveTimer(v string)`

SetDefaultInactiveTimer sets DefaultInactiveTimer field to given value.

### HasDefaultInactiveTimer

`func (o *ConversationsV1Configuration) HasDefaultInactiveTimer() bool`

HasDefaultInactiveTimer returns a boolean if a field has been set.

### SetDefaultInactiveTimerNil

`func (o *ConversationsV1Configuration) SetDefaultInactiveTimerNil(b bool)`

 SetDefaultInactiveTimerNil sets the value for DefaultInactiveTimer to be an explicit nil

### UnsetDefaultInactiveTimer
`func (o *ConversationsV1Configuration) UnsetDefaultInactiveTimer()`

UnsetDefaultInactiveTimer ensures that no value is present for DefaultInactiveTimer, not even an explicit nil
### GetDefaultMessagingServiceSid

`func (o *ConversationsV1Configuration) GetDefaultMessagingServiceSid() string`

GetDefaultMessagingServiceSid returns the DefaultMessagingServiceSid field if non-nil, zero value otherwise.

### GetDefaultMessagingServiceSidOk

`func (o *ConversationsV1Configuration) GetDefaultMessagingServiceSidOk() (*string, bool)`

GetDefaultMessagingServiceSidOk returns a tuple with the DefaultMessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessagingServiceSid

`func (o *ConversationsV1Configuration) SetDefaultMessagingServiceSid(v string)`

SetDefaultMessagingServiceSid sets DefaultMessagingServiceSid field to given value.

### HasDefaultMessagingServiceSid

`func (o *ConversationsV1Configuration) HasDefaultMessagingServiceSid() bool`

HasDefaultMessagingServiceSid returns a boolean if a field has been set.

### SetDefaultMessagingServiceSidNil

`func (o *ConversationsV1Configuration) SetDefaultMessagingServiceSidNil(b bool)`

 SetDefaultMessagingServiceSidNil sets the value for DefaultMessagingServiceSid to be an explicit nil

### UnsetDefaultMessagingServiceSid
`func (o *ConversationsV1Configuration) UnsetDefaultMessagingServiceSid()`

UnsetDefaultMessagingServiceSid ensures that no value is present for DefaultMessagingServiceSid, not even an explicit nil
### GetLinks

`func (o *ConversationsV1Configuration) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConversationsV1Configuration) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConversationsV1Configuration) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConversationsV1Configuration) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ConversationsV1Configuration) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ConversationsV1Configuration) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetUrl

`func (o *ConversationsV1Configuration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1Configuration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1Configuration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1Configuration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1Configuration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1Configuration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


