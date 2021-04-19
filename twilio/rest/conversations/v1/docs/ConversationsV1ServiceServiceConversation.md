# ConversationsV1ServiceServiceConversation

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this conversation. | [optional] 
**Attributes** | Pointer to **NullableString** | An optional string metadata field you can use to store any data you wish. | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The unique ID of the Conversation Service this conversation belongs to. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**FriendlyName** | Pointer to **NullableString** | The human-readable name of this conversation. | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the participants, messages and webhooks of this conversation. | [optional] 
**MessagingServiceSid** | Pointer to **NullableString** | The unique ID of the Messaging Service this conversation belongs to. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**State** | Pointer to **NullableString** | Current state of this conversation. | [optional] 
**Timers** | Pointer to **map[string]interface{}** | Timer date values for this conversation. | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this conversation. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConversation

`func NewConversationsV1ServiceServiceConversation() *ConversationsV1ServiceServiceConversation`

NewConversationsV1ServiceServiceConversation instantiates a new ConversationsV1ServiceServiceConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConversationWithDefaults

`func NewConversationsV1ServiceServiceConversationWithDefaults() *ConversationsV1ServiceServiceConversation`

NewConversationsV1ServiceServiceConversationWithDefaults instantiates a new ConversationsV1ServiceServiceConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceConversation) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceConversation) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceConversation) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceConversation) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceConversation) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceConversation) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ConversationsV1ServiceServiceConversation) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConversationsV1ServiceServiceConversation) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConversationsV1ServiceServiceConversation) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConversationsV1ServiceServiceConversation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ConversationsV1ServiceServiceConversation) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ConversationsV1ServiceServiceConversation) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversation) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConversation) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversation) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConversation) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConversation) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConversation) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceConversation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceConversation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceConversation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceConversation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceConversation) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceConversation) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceConversation) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceConversation) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceConversation) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceConversation) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceConversation) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceConversation) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ConversationsV1ServiceServiceConversation) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ConversationsV1ServiceServiceConversation) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ConversationsV1ServiceServiceConversation) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ConversationsV1ServiceServiceConversation) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ConversationsV1ServiceServiceConversation) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ConversationsV1ServiceServiceConversation) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *ConversationsV1ServiceServiceConversation) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ConversationsV1ServiceServiceConversation) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ConversationsV1ServiceServiceConversation) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ConversationsV1ServiceServiceConversation) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ConversationsV1ServiceServiceConversation) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ConversationsV1ServiceServiceConversation) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMessagingServiceSid

`func (o *ConversationsV1ServiceServiceConversation) GetMessagingServiceSid() string`

GetMessagingServiceSid returns the MessagingServiceSid field if non-nil, zero value otherwise.

### GetMessagingServiceSidOk

`func (o *ConversationsV1ServiceServiceConversation) GetMessagingServiceSidOk() (*string, bool)`

GetMessagingServiceSidOk returns a tuple with the MessagingServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServiceSid

`func (o *ConversationsV1ServiceServiceConversation) SetMessagingServiceSid(v string)`

SetMessagingServiceSid sets MessagingServiceSid field to given value.

### HasMessagingServiceSid

`func (o *ConversationsV1ServiceServiceConversation) HasMessagingServiceSid() bool`

HasMessagingServiceSid returns a boolean if a field has been set.

### SetMessagingServiceSidNil

`func (o *ConversationsV1ServiceServiceConversation) SetMessagingServiceSidNil(b bool)`

 SetMessagingServiceSidNil sets the value for MessagingServiceSid to be an explicit nil

### UnsetMessagingServiceSid
`func (o *ConversationsV1ServiceServiceConversation) UnsetMessagingServiceSid()`

UnsetMessagingServiceSid ensures that no value is present for MessagingServiceSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceConversation) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceConversation) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceConversation) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceConversation) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceConversation) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceConversation) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetState

`func (o *ConversationsV1ServiceServiceConversation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConversationsV1ServiceServiceConversation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConversationsV1ServiceServiceConversation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConversationsV1ServiceServiceConversation) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ConversationsV1ServiceServiceConversation) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ConversationsV1ServiceServiceConversation) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTimers

`func (o *ConversationsV1ServiceServiceConversation) GetTimers() map[string]interface{}`

GetTimers returns the Timers field if non-nil, zero value otherwise.

### GetTimersOk

`func (o *ConversationsV1ServiceServiceConversation) GetTimersOk() (*map[string]interface{}, bool)`

GetTimersOk returns a tuple with the Timers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimers

`func (o *ConversationsV1ServiceServiceConversation) SetTimers(v map[string]interface{})`

SetTimers sets Timers field to given value.

### HasTimers

`func (o *ConversationsV1ServiceServiceConversation) HasTimers() bool`

HasTimers returns a boolean if a field has been set.

### SetTimersNil

`func (o *ConversationsV1ServiceServiceConversation) SetTimersNil(b bool)`

 SetTimersNil sets the value for Timers to be an explicit nil

### UnsetTimers
`func (o *ConversationsV1ServiceServiceConversation) UnsetTimers()`

UnsetTimers ensures that no value is present for Timers, not even an explicit nil
### GetUniqueName

`func (o *ConversationsV1ServiceServiceConversation) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ConversationsV1ServiceServiceConversation) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ConversationsV1ServiceServiceConversation) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ConversationsV1ServiceServiceConversation) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *ConversationsV1ServiceServiceConversation) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *ConversationsV1ServiceServiceConversation) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConversation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConversation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConversation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConversation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConversation) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConversation) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


