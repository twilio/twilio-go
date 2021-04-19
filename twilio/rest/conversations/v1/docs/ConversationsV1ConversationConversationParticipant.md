# ConversationsV1ConversationConversationParticipant

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this participant. | [optional] 
**Attributes** | Pointer to **NullableString** | An optional string metadata field you can use to store any data you wish. | [optional] 
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

### NewConversationsV1ConversationConversationParticipant

`func NewConversationsV1ConversationConversationParticipant() *ConversationsV1ConversationConversationParticipant`

NewConversationsV1ConversationConversationParticipant instantiates a new ConversationsV1ConversationConversationParticipant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ConversationConversationParticipantWithDefaults

`func NewConversationsV1ConversationConversationParticipantWithDefaults() *ConversationsV1ConversationConversationParticipant`

NewConversationsV1ConversationConversationParticipantWithDefaults instantiates a new ConversationsV1ConversationConversationParticipant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ConversationConversationParticipant) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ConversationConversationParticipant) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ConversationConversationParticipant) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ConversationConversationParticipant) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ConversationConversationParticipant) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ConversationConversationParticipant) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *ConversationsV1ConversationConversationParticipant) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ConversationsV1ConversationConversationParticipant) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ConversationsV1ConversationConversationParticipant) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ConversationsV1ConversationConversationParticipant) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *ConversationsV1ConversationConversationParticipant) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *ConversationsV1ConversationConversationParticipant) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ConversationConversationParticipant) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ConversationConversationParticipant) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ConversationConversationParticipant) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ConversationConversationParticipant) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ConversationConversationParticipant) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ConversationConversationParticipant) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ConversationConversationParticipant) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ConversationConversationParticipant) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ConversationConversationParticipant) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ConversationConversationParticipant) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ConversationConversationParticipant) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ConversationConversationParticipant) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ConversationConversationParticipant) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ConversationConversationParticipant) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ConversationConversationParticipant) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ConversationConversationParticipant) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ConversationConversationParticipant) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ConversationConversationParticipant) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIdentity

`func (o *ConversationsV1ConversationConversationParticipant) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ConversationsV1ConversationConversationParticipant) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ConversationsV1ConversationConversationParticipant) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ConversationsV1ConversationConversationParticipant) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ConversationsV1ConversationConversationParticipant) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ConversationsV1ConversationConversationParticipant) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLastReadMessageIndex

`func (o *ConversationsV1ConversationConversationParticipant) GetLastReadMessageIndex() int32`

GetLastReadMessageIndex returns the LastReadMessageIndex field if non-nil, zero value otherwise.

### GetLastReadMessageIndexOk

`func (o *ConversationsV1ConversationConversationParticipant) GetLastReadMessageIndexOk() (*int32, bool)`

GetLastReadMessageIndexOk returns a tuple with the LastReadMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadMessageIndex

`func (o *ConversationsV1ConversationConversationParticipant) SetLastReadMessageIndex(v int32)`

SetLastReadMessageIndex sets LastReadMessageIndex field to given value.

### HasLastReadMessageIndex

`func (o *ConversationsV1ConversationConversationParticipant) HasLastReadMessageIndex() bool`

HasLastReadMessageIndex returns a boolean if a field has been set.

### SetLastReadMessageIndexNil

`func (o *ConversationsV1ConversationConversationParticipant) SetLastReadMessageIndexNil(b bool)`

 SetLastReadMessageIndexNil sets the value for LastReadMessageIndex to be an explicit nil

### UnsetLastReadMessageIndex
`func (o *ConversationsV1ConversationConversationParticipant) UnsetLastReadMessageIndex()`

UnsetLastReadMessageIndex ensures that no value is present for LastReadMessageIndex, not even an explicit nil
### GetLastReadTimestamp

`func (o *ConversationsV1ConversationConversationParticipant) GetLastReadTimestamp() string`

GetLastReadTimestamp returns the LastReadTimestamp field if non-nil, zero value otherwise.

### GetLastReadTimestampOk

`func (o *ConversationsV1ConversationConversationParticipant) GetLastReadTimestampOk() (*string, bool)`

GetLastReadTimestampOk returns a tuple with the LastReadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReadTimestamp

`func (o *ConversationsV1ConversationConversationParticipant) SetLastReadTimestamp(v string)`

SetLastReadTimestamp sets LastReadTimestamp field to given value.

### HasLastReadTimestamp

`func (o *ConversationsV1ConversationConversationParticipant) HasLastReadTimestamp() bool`

HasLastReadTimestamp returns a boolean if a field has been set.

### SetLastReadTimestampNil

`func (o *ConversationsV1ConversationConversationParticipant) SetLastReadTimestampNil(b bool)`

 SetLastReadTimestampNil sets the value for LastReadTimestamp to be an explicit nil

### UnsetLastReadTimestamp
`func (o *ConversationsV1ConversationConversationParticipant) UnsetLastReadTimestamp()`

UnsetLastReadTimestamp ensures that no value is present for LastReadTimestamp, not even an explicit nil
### GetMessagingBinding

`func (o *ConversationsV1ConversationConversationParticipant) GetMessagingBinding() map[string]interface{}`

GetMessagingBinding returns the MessagingBinding field if non-nil, zero value otherwise.

### GetMessagingBindingOk

`func (o *ConversationsV1ConversationConversationParticipant) GetMessagingBindingOk() (*map[string]interface{}, bool)`

GetMessagingBindingOk returns a tuple with the MessagingBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingBinding

`func (o *ConversationsV1ConversationConversationParticipant) SetMessagingBinding(v map[string]interface{})`

SetMessagingBinding sets MessagingBinding field to given value.

### HasMessagingBinding

`func (o *ConversationsV1ConversationConversationParticipant) HasMessagingBinding() bool`

HasMessagingBinding returns a boolean if a field has been set.

### SetMessagingBindingNil

`func (o *ConversationsV1ConversationConversationParticipant) SetMessagingBindingNil(b bool)`

 SetMessagingBindingNil sets the value for MessagingBinding to be an explicit nil

### UnsetMessagingBinding
`func (o *ConversationsV1ConversationConversationParticipant) UnsetMessagingBinding()`

UnsetMessagingBinding ensures that no value is present for MessagingBinding, not even an explicit nil
### GetRoleSid

`func (o *ConversationsV1ConversationConversationParticipant) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ConversationsV1ConversationConversationParticipant) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ConversationsV1ConversationConversationParticipant) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ConversationsV1ConversationConversationParticipant) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ConversationsV1ConversationConversationParticipant) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ConversationsV1ConversationConversationParticipant) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ConversationConversationParticipant) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ConversationConversationParticipant) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ConversationConversationParticipant) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ConversationConversationParticipant) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ConversationConversationParticipant) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ConversationConversationParticipant) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ConversationConversationParticipant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ConversationConversationParticipant) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ConversationConversationParticipant) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ConversationConversationParticipant) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ConversationConversationParticipant) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ConversationConversationParticipant) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


