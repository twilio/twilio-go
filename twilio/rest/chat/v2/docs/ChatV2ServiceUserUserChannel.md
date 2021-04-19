# ChatV2ServiceUserUserChannel

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ChannelSid** | Pointer to **NullableString** | The SID of the Channel the resource belongs to | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** | The index of the last Message in the Channel the Member has read | [optional] 
**Links** | Pointer to **map[string]interface{}** | Absolute URLs to access the Members, Messages , Invites and, if it exists, the last Message for the Channel | [optional] 
**MemberSid** | Pointer to **NullableString** | The SID of the User as a Member in the Channel | [optional] 
**NotificationLevel** | Pointer to **NullableString** | The push notification level of the User for the Channel | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Status** | Pointer to **NullableString** | The status of the User on the Channel | [optional] 
**UnreadMessagesCount** | Pointer to **NullableInt32** | The number of unread Messages in the Channel for the User | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 
**UserSid** | Pointer to **NullableString** | The SID of the User the User Channel belongs to | [optional] 

## Methods

### NewChatV2ServiceUserUserChannel

`func NewChatV2ServiceUserUserChannel() *ChatV2ServiceUserUserChannel`

NewChatV2ServiceUserUserChannel instantiates a new ChatV2ServiceUserUserChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceUserUserChannelWithDefaults

`func NewChatV2ServiceUserUserChannelWithDefaults() *ChatV2ServiceUserUserChannel`

NewChatV2ServiceUserUserChannelWithDefaults instantiates a new ChatV2ServiceUserUserChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2ServiceUserUserChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2ServiceUserUserChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2ServiceUserUserChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2ServiceUserUserChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2ServiceUserUserChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2ServiceUserUserChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *ChatV2ServiceUserUserChannel) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV2ServiceUserUserChannel) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV2ServiceUserUserChannel) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV2ServiceUserUserChannel) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV2ServiceUserUserChannel) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV2ServiceUserUserChannel) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *ChatV2ServiceUserUserChannel) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *ChatV2ServiceUserUserChannel) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *ChatV2ServiceUserUserChannel) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *ChatV2ServiceUserUserChannel) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *ChatV2ServiceUserUserChannel) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *ChatV2ServiceUserUserChannel) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLinks

`func (o *ChatV2ServiceUserUserChannel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ChatV2ServiceUserUserChannel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ChatV2ServiceUserUserChannel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ChatV2ServiceUserUserChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ChatV2ServiceUserUserChannel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ChatV2ServiceUserUserChannel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMemberSid

`func (o *ChatV2ServiceUserUserChannel) GetMemberSid() string`

GetMemberSid returns the MemberSid field if non-nil, zero value otherwise.

### GetMemberSidOk

`func (o *ChatV2ServiceUserUserChannel) GetMemberSidOk() (*string, bool)`

GetMemberSidOk returns a tuple with the MemberSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSid

`func (o *ChatV2ServiceUserUserChannel) SetMemberSid(v string)`

SetMemberSid sets MemberSid field to given value.

### HasMemberSid

`func (o *ChatV2ServiceUserUserChannel) HasMemberSid() bool`

HasMemberSid returns a boolean if a field has been set.

### SetMemberSidNil

`func (o *ChatV2ServiceUserUserChannel) SetMemberSidNil(b bool)`

 SetMemberSidNil sets the value for MemberSid to be an explicit nil

### UnsetMemberSid
`func (o *ChatV2ServiceUserUserChannel) UnsetMemberSid()`

UnsetMemberSid ensures that no value is present for MemberSid, not even an explicit nil
### GetNotificationLevel

`func (o *ChatV2ServiceUserUserChannel) GetNotificationLevel() string`

GetNotificationLevel returns the NotificationLevel field if non-nil, zero value otherwise.

### GetNotificationLevelOk

`func (o *ChatV2ServiceUserUserChannel) GetNotificationLevelOk() (*string, bool)`

GetNotificationLevelOk returns a tuple with the NotificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLevel

`func (o *ChatV2ServiceUserUserChannel) SetNotificationLevel(v string)`

SetNotificationLevel sets NotificationLevel field to given value.

### HasNotificationLevel

`func (o *ChatV2ServiceUserUserChannel) HasNotificationLevel() bool`

HasNotificationLevel returns a boolean if a field has been set.

### SetNotificationLevelNil

`func (o *ChatV2ServiceUserUserChannel) SetNotificationLevelNil(b bool)`

 SetNotificationLevelNil sets the value for NotificationLevel to be an explicit nil

### UnsetNotificationLevel
`func (o *ChatV2ServiceUserUserChannel) UnsetNotificationLevel()`

UnsetNotificationLevel ensures that no value is present for NotificationLevel, not even an explicit nil
### GetServiceSid

`func (o *ChatV2ServiceUserUserChannel) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV2ServiceUserUserChannel) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV2ServiceUserUserChannel) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV2ServiceUserUserChannel) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV2ServiceUserUserChannel) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV2ServiceUserUserChannel) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetStatus

`func (o *ChatV2ServiceUserUserChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChatV2ServiceUserUserChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChatV2ServiceUserUserChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChatV2ServiceUserUserChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ChatV2ServiceUserUserChannel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ChatV2ServiceUserUserChannel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnreadMessagesCount

`func (o *ChatV2ServiceUserUserChannel) GetUnreadMessagesCount() int32`

GetUnreadMessagesCount returns the UnreadMessagesCount field if non-nil, zero value otherwise.

### GetUnreadMessagesCountOk

`func (o *ChatV2ServiceUserUserChannel) GetUnreadMessagesCountOk() (*int32, bool)`

GetUnreadMessagesCountOk returns a tuple with the UnreadMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessagesCount

`func (o *ChatV2ServiceUserUserChannel) SetUnreadMessagesCount(v int32)`

SetUnreadMessagesCount sets UnreadMessagesCount field to given value.

### HasUnreadMessagesCount

`func (o *ChatV2ServiceUserUserChannel) HasUnreadMessagesCount() bool`

HasUnreadMessagesCount returns a boolean if a field has been set.

### SetUnreadMessagesCountNil

`func (o *ChatV2ServiceUserUserChannel) SetUnreadMessagesCountNil(b bool)`

 SetUnreadMessagesCountNil sets the value for UnreadMessagesCount to be an explicit nil

### UnsetUnreadMessagesCount
`func (o *ChatV2ServiceUserUserChannel) UnsetUnreadMessagesCount()`

UnsetUnreadMessagesCount ensures that no value is present for UnreadMessagesCount, not even an explicit nil
### GetUrl

`func (o *ChatV2ServiceUserUserChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2ServiceUserUserChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2ServiceUserUserChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2ServiceUserUserChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2ServiceUserUserChannel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2ServiceUserUserChannel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUserSid

`func (o *ChatV2ServiceUserUserChannel) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *ChatV2ServiceUserUserChannel) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *ChatV2ServiceUserUserChannel) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *ChatV2ServiceUserUserChannel) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *ChatV2ServiceUserUserChannel) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *ChatV2ServiceUserUserChannel) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


