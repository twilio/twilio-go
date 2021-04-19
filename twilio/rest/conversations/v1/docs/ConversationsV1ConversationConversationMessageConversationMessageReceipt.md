# ConversationsV1ConversationConversationMessageConversationMessageReceipt

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique ID of the Account responsible for this participant. | [optional] 
**ChannelMessageSid** | Pointer to **NullableString** | A messaging channel-specific identifier for the message delivered to participant | [optional] 
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

### NewConversationsV1ConversationConversationMessageConversationMessageReceipt

`func NewConversationsV1ConversationConversationMessageConversationMessageReceipt() *ConversationsV1ConversationConversationMessageConversationMessageReceipt`

NewConversationsV1ConversationConversationMessageConversationMessageReceipt instantiates a new ConversationsV1ConversationConversationMessageConversationMessageReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationsV1ConversationConversationMessageConversationMessageReceiptWithDefaults

`func NewConversationsV1ConversationConversationMessageConversationMessageReceiptWithDefaults() *ConversationsV1ConversationConversationMessageConversationMessageReceipt`

NewConversationsV1ConversationConversationMessageConversationMessageReceiptWithDefaults instantiates a new ConversationsV1ConversationConversationMessageConversationMessageReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetChannelMessageSid() string`

GetChannelMessageSid returns the ChannelMessageSid field if non-nil, zero value otherwise.

### GetChannelMessageSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetChannelMessageSidOk() (*string, bool)`

GetChannelMessageSidOk returns a tuple with the ChannelMessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetChannelMessageSid(v string)`

SetChannelMessageSid sets ChannelMessageSid field to given value.

### HasChannelMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasChannelMessageSid() bool`

HasChannelMessageSid returns a boolean if a field has been set.

### SetChannelMessageSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetChannelMessageSidNil(b bool)`

 SetChannelMessageSidNil sets the value for ChannelMessageSid to be an explicit nil

### UnsetChannelMessageSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetChannelMessageSid()`

UnsetChannelMessageSid ensures that no value is present for ChannelMessageSid, not even an explicit nil
### GetConversationSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetConversationSid() string`

GetConversationSid returns the ConversationSid field if non-nil, zero value otherwise.

### GetConversationSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetConversationSidOk() (*string, bool)`

GetConversationSidOk returns a tuple with the ConversationSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetConversationSid(v string)`

SetConversationSid sets ConversationSid field to given value.

### HasConversationSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasConversationSid() bool`

HasConversationSid returns a boolean if a field has been set.

### SetConversationSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetConversationSidNil(b bool)`

 SetConversationSidNil sets the value for ConversationSid to be an explicit nil

### UnsetConversationSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetConversationSid()`

UnsetConversationSid ensures that no value is present for ConversationSid, not even an explicit nil
### GetDateCreated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetErrorCode

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetMessageSid() string`

GetMessageSid returns the MessageSid field if non-nil, zero value otherwise.

### GetMessageSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetMessageSidOk() (*string, bool)`

GetMessageSidOk returns a tuple with the MessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetMessageSid(v string)`

SetMessageSid sets MessageSid field to given value.

### HasMessageSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasMessageSid() bool`

HasMessageSid returns a boolean if a field has been set.

### SetMessageSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetMessageSidNil(b bool)`

 SetMessageSidNil sets the value for MessageSid to be an explicit nil

### UnsetMessageSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetMessageSid()`

UnsetMessageSid ensures that no value is present for MessageSid, not even an explicit nil
### GetParticipantSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetParticipantSid() string`

GetParticipantSid returns the ParticipantSid field if non-nil, zero value otherwise.

### GetParticipantSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetParticipantSidOk() (*string, bool)`

GetParticipantSidOk returns a tuple with the ParticipantSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetParticipantSid(v string)`

SetParticipantSid sets ParticipantSid field to given value.

### HasParticipantSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasParticipantSid() bool`

HasParticipantSid returns a boolean if a field has been set.

### SetParticipantSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetParticipantSidNil(b bool)`

 SetParticipantSidNil sets the value for ParticipantSid to be an explicit nil

### UnsetParticipantSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetParticipantSid()`

UnsetParticipantSid ensures that no value is present for ParticipantSid, not even an explicit nil
### GetSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStatus

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ConversationsV1ConversationConversationMessageConversationMessageReceipt) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


