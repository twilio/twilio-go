# ConversationsV1ServiceServiceConversationServiceConversationParticipant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this participant. | [optional] 
**Attributes** | Pointer to **NullableString** | An optional string metadata field you can use to store any data you wish. | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the resource is associated with. | [optional] 
**ConversationSid** | Pointer to **NullableString** | The unique ID of the Conversation for this participant. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**Identity** | Pointer to **NullableString** | A unique string identifier for the conversation participant as Conversation User. | [optional] 
**LastReadMessageIndex** | Pointer to **NullableInt32** | Index of last “read” message in the Conversation for the Participant. | [optional] 
**LastReadTimestamp** | Pointer to **NullableString** | Timestamp of last “read” message in the Conversation for the Participant. | [optional] 
**MessagingBinding** | Pointer to **map[string]interface{}** | Information about how this participant exchanges messages with the conversation. | [optional] 
**RoleSid** | Pointer to **NullableString** | The SID of a conversation-level Role to assign to the participant | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this participant. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConversationServiceConversationParticipant

`func NewConversationsV1ServiceServiceConversationServiceConversationParticipant() *ConversationsV1ServiceServiceConversationServiceConversationParticipant`

NewConversationsV1ServiceServiceConversationServiceConversationParticipant instantiates a new ConversationsV1ServiceServiceConversationServiceConversationParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConversationServiceConversationParticipantWithDefaults

`func NewConversationsV1ServiceServiceConversationServiceConversationParticipantWithDefaults() *ConversationsV1ServiceServiceConversationServiceConversationParticipant`

NewConversationsV1ServiceServiceConversationServiceConversationParticipantWithDefaults instantiates a new ConversationsV1ServiceServiceConversationServiceConversationParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIdentity

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLastReadMessageIndex

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetLastReadMessageIndex() int32`

GetLastReadMessageIndex returns the LastReadMessageIndex field if non-nil, zero value otherwise.

### GetLastReadMessageIndexOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetLastReadMessageIndexOk() (*int32, bool)`

GetLastReadMessageIndexOk returns a tuple with the LastReadMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadMessageIndex

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetLastReadMessageIndex(v int32)`

SetLastReadMessageIndex sets LastReadMessageIndex field to given value.

### HasLastReadMessageIndex

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasLastReadMessageIndex() bool`

HasLastReadMessageIndex returns a boolean if a field has been set.

### SetLastReadMessageIndexNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetLastReadMessageIndexNil(b bool)`

 SetLastReadMessageIndexNil sets the value for LastReadMessageIndex to be an explicit nil

### UnsetLastReadMessageIndex
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetLastReadMessageIndex()`

UnsetLastReadMessageIndex ensures that no value is present for LastReadMessageIndex, not even an explicit nil
### GetLastReadTimestamp

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetLastReadTimestamp() string`

GetLastReadTimestamp returns the LastReadTimestamp field if non-nil, zero value otherwise.

### GetLastReadTimestampOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetLastReadTimestampOk() (*string, bool)`

GetLastReadTimestampOk returns a tuple with the LastReadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadTimestamp

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetLastReadTimestamp(v string)`

SetLastReadTimestamp sets LastReadTimestamp field to given value.

### HasLastReadTimestamp

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasLastReadTimestamp() bool`

HasLastReadTimestamp returns a boolean if a field has been set.

### SetLastReadTimestampNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetLastReadTimestampNil(b bool)`

 SetLastReadTimestampNil sets the value for LastReadTimestamp to be an explicit nil

### UnsetLastReadTimestamp
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetLastReadTimestamp()`

UnsetLastReadTimestamp ensures that no value is present for LastReadTimestamp, not even an explicit nil
### GetMessagingBinding

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetMessagingBinding() map[string]interface{}`

GetMessagingBinding returns the MessagingBinding field if non-nil, zero value otherwise.

### GetMessagingBindingOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetMessagingBindingOk() (*map[string]interface{}, bool)`

GetMessagingBindingOk returns a tuple with the MessagingBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingBinding

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetMessagingBinding(v map[string]interface{})`

SetMessagingBinding sets MessagingBinding field to given value.

### HasMessagingBinding

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasMessagingBinding() bool`

HasMessagingBinding returns a boolean if a field has been set.

### SetMessagingBindingNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetMessagingBindingNil(b bool)`

 SetMessagingBindingNil sets the value for MessagingBinding to be an explicit nil

### UnsetMessagingBinding
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetMessagingBinding()`

UnsetMessagingBinding ensures that no value is present for MessagingBinding, not even an explicit nil
### GetRoleSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConversationServiceConversationParticipant) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


