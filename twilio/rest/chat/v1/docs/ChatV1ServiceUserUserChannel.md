# ChatV1ServiceUserUserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ChannelSid** | Pointer to **NullableString** | The SID of the Channel the resource belongs to | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** | The index of the last Message in the Channel the Member has read | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel | [optional] 
**MemberSid** | Pointer to **NullableString** | The SID of the User as a Member in the Channel | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Status** | Pointer to **NullableString** | The status of the User on the Channel | [optional] 
**UnreadMessagesCount** | Pointer to **NullableInt32** | The number of unread Messages in the Channel for the User | [optional] 

## Methods

### NewChatV1ServiceUserUserChannel

`func NewChatV1ServiceUserUserChannel() *ChatV1ServiceUserUserChannel`

NewChatV1ServiceUserUserChannel instantiates a new ChatV1ServiceUserUserChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV1ServiceUserUserChannelWithDefaults

`func NewChatV1ServiceUserUserChannelWithDefaults() *ChatV1ServiceUserUserChannel`

NewChatV1ServiceUserUserChannelWithDefaults instantiates a new ChatV1ServiceUserUserChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV1ServiceUserUserChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV1ServiceUserUserChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV1ServiceUserUserChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV1ServiceUserUserChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV1ServiceUserUserChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV1ServiceUserUserChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *ChatV1ServiceUserUserChannel) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV1ServiceUserUserChannel) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV1ServiceUserUserChannel) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV1ServiceUserUserChannel) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV1ServiceUserUserChannel) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV1ServiceUserUserChannel) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *ChatV1ServiceUserUserChannel) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *ChatV1ServiceUserUserChannel) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *ChatV1ServiceUserUserChannel) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *ChatV1ServiceUserUserChannel) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *ChatV1ServiceUserUserChannel) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *ChatV1ServiceUserUserChannel) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLinks

`func (o *ChatV1ServiceUserUserChannel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChatV1ServiceUserUserChannel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChatV1ServiceUserUserChannel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChatV1ServiceUserUserChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ChatV1ServiceUserUserChannel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ChatV1ServiceUserUserChannel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMemberSid

`func (o *ChatV1ServiceUserUserChannel) GetMemberSid() string`

GetMemberSid returns the MemberSid field if non-nil, zero value otherwise.

### GetMemberSidOk

`func (o *ChatV1ServiceUserUserChannel) GetMemberSidOk() (*string, bool)`

GetMemberSidOk returns a tuple with the MemberSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSid

`func (o *ChatV1ServiceUserUserChannel) SetMemberSid(v string)`

SetMemberSid sets MemberSid field to given value.

### HasMemberSid

`func (o *ChatV1ServiceUserUserChannel) HasMemberSid() bool`

HasMemberSid returns a boolean if a field has been set.

### SetMemberSidNil

`func (o *ChatV1ServiceUserUserChannel) SetMemberSidNil(b bool)`

 SetMemberSidNil sets the value for MemberSid to be an explicit nil

### UnsetMemberSid
`func (o *ChatV1ServiceUserUserChannel) UnsetMemberSid()`

UnsetMemberSid ensures that no value is present for MemberSid, not even an explicit nil
### GetServiceSid

`func (o *ChatV1ServiceUserUserChannel) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV1ServiceUserUserChannel) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV1ServiceUserUserChannel) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV1ServiceUserUserChannel) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV1ServiceUserUserChannel) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV1ServiceUserUserChannel) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetStatus

`func (o *ChatV1ServiceUserUserChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatV1ServiceUserUserChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatV1ServiceUserUserChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatV1ServiceUserUserChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ChatV1ServiceUserUserChannel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ChatV1ServiceUserUserChannel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnreadMessagesCount

`func (o *ChatV1ServiceUserUserChannel) GetUnreadMessagesCount() int32`

GetUnreadMessagesCount returns the UnreadMessagesCount field if non-nil, zero value otherwise.

### GetUnreadMessagesCountOk

`func (o *ChatV1ServiceUserUserChannel) GetUnreadMessagesCountOk() (*int32, bool)`

GetUnreadMessagesCountOk returns a tuple with the UnreadMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessagesCount

`func (o *ChatV1ServiceUserUserChannel) SetUnreadMessagesCount(v int32)`

SetUnreadMessagesCount sets UnreadMessagesCount field to given value.

### HasUnreadMessagesCount

`func (o *ChatV1ServiceUserUserChannel) HasUnreadMessagesCount() bool`

HasUnreadMessagesCount returns a boolean if a field has been set.

### SetUnreadMessagesCountNil

`func (o *ChatV1ServiceUserUserChannel) SetUnreadMessagesCountNil(b bool)`

 SetUnreadMessagesCountNil sets the value for UnreadMessagesCount to be an explicit nil

### UnsetUnreadMessagesCount
`func (o *ChatV1ServiceUserUserChannel) UnsetUnreadMessagesCount()`

UnsetUnreadMessagesCount ensures that no value is present for UnreadMessagesCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


