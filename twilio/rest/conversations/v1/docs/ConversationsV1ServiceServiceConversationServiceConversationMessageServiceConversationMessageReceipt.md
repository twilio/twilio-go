# ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this participant. | [optional] 
**ChannelMessageSid** | Pointer to **NullableString** | A messaging channel-specific identifier for the message delivered to participant | [optional] 
**ChatServiceSid** | Pointer to **NullableString** | The SID of the Conversation Service that the resource is associated with. | [optional] 
**ConversationSid** | Pointer to **NullableString** | The unique ID of the Conversation for this message. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The date that this resource was created. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The date that this resource was last updated. | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | The message [delivery error code](https://www.twilio.com/docs/sms/api/message-resource#delivery-related-errors) for a &#x60;failed&#x60; status | [optional] 
**MessageSid** | Pointer to **NullableString** | The SID of the message the delivery receipt belongs to | [optional] 
**ParticipantSid** | Pointer to **NullableString** | The unique ID of the participant the delivery receipt belongs to. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Status** | Pointer to **NullableString** | The message delivery status | [optional] 
**Url** | Pointer to **NullableString** | An absolute URL for this delivery receipt. | [optional] 

## Methods

### NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt

`func NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt() *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt`

NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt instantiates a new ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceiptWithDefaults

`func NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceiptWithDefaults() *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt`

NewConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceiptWithDefaults instantiates a new ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetChannelMessageSid() string`

GetChannelMessageSid returns the ChannelMessageSid field if non-nil, zero value otherwise.

### GetChannelMessageSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetChannelMessageSidOk() (*string, bool)`

GetChannelMessageSidOk returns a tuple with the ChannelMessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetChannelMessageSid(v string)`

SetChannelMessageSid sets ChannelMessageSid field to given value.

### HasChannelMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasChannelMessageSid() bool`

HasChannelMessageSid returns a boolean if a field has been set.

### SetChannelMessageSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetChannelMessageSidNil(b bool)`

 SetChannelMessageSidNil sets the value for ChannelMessageSid to be an explicit nil

### UnsetChannelMessageSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetChannelMessageSid()`

UnsetChannelMessageSid ensures that no value is present for ChannelMessageSid, not even an explicit nil
### GetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetChatServiceSid() string`

GetChatServiceSid returns the ChatServiceSid field if non-nil, zero value otherwise.

### GetChatServiceSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetChatServiceSidOk() (*string, bool)`

GetChatServiceSidOk returns a tuple with the ChatServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetChatServiceSid(v string)`

SetChatServiceSid sets ChatServiceSid field to given value.

### HasChatServiceSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasChatServiceSid() bool`

HasChatServiceSid returns a boolean if a field has been set.

### SetChatServiceSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetChatServiceSidNil(b bool)`

 SetChatServiceSidNil sets the value for ChatServiceSid to be an explicit nil

### UnsetChatServiceSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetChatServiceSid()`

UnsetChatServiceSid ensures that no value is present for ChatServiceSid, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetErrorCode

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetMessageSid() string`

GetMessageSid returns the MessageSid field if non-nil, zero value otherwise.

### GetMessageSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetMessageSidOk() (*string, bool)`

GetMessageSidOk returns a tuple with the MessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetMessageSid(v string)`

SetMessageSid sets MessageSid field to given value.

### HasMessageSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasMessageSid() bool`

HasMessageSid returns a boolean if a field has been set.

### SetMessageSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetMessageSidNil(b bool)`

 SetMessageSidNil sets the value for MessageSid to be an explicit nil

### UnsetMessageSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetMessageSid()`

UnsetMessageSid ensures that no value is present for MessageSid, not even an explicit nil
### GetParticipantSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ServiceServiceConversationServiceConversationMessageServiceConversationMessageReceipt) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


