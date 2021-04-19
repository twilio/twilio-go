# IpMessagingV2ServiceUserUserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**ChannelSid** | Pointer to **NullableString** |  | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**MemberSid** | Pointer to **NullableString** |  | [optional] 
**NotificationLevel** | Pointer to **NullableString** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UnreadMessagesCount** | Pointer to **NullableInt32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**UserSid** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpMessagingV2ServiceUserUserChannel

`func NewIpMessagingV2ServiceUserUserChannel() *IpMessagingV2ServiceUserUserChannel`

NewIpMessagingV2ServiceUserUserChannel instantiates a new IpMessagingV2ServiceUserUserChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV2ServiceUserUserChannelWithDefaults

`func NewIpMessagingV2ServiceUserUserChannelWithDefaults() *IpMessagingV2ServiceUserUserChannel`

NewIpMessagingV2ServiceUserUserChannelWithDefaults instantiates a new IpMessagingV2ServiceUserUserChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV2ServiceUserUserChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV2ServiceUserUserChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV2ServiceUserUserChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *IpMessagingV2ServiceUserUserChannel) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *IpMessagingV2ServiceUserUserChannel) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *IpMessagingV2ServiceUserUserChannel) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *IpMessagingV2ServiceUserUserChannel) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *IpMessagingV2ServiceUserUserChannel) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *IpMessagingV2ServiceUserUserChannel) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLinks

`func (o *IpMessagingV2ServiceUserUserChannel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV2ServiceUserUserChannel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV2ServiceUserUserChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMemberSid

`func (o *IpMessagingV2ServiceUserUserChannel) GetMemberSid() string`

GetMemberSid returns the MemberSid field if non-nil, zero value otherwise.

### GetMemberSidOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetMemberSidOk() (*string, bool)`

GetMemberSidOk returns a tuple with the MemberSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSid

`func (o *IpMessagingV2ServiceUserUserChannel) SetMemberSid(v string)`

SetMemberSid sets MemberSid field to given value.

### HasMemberSid

`func (o *IpMessagingV2ServiceUserUserChannel) HasMemberSid() bool`

HasMemberSid returns a boolean if a field has been set.

### SetMemberSidNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetMemberSidNil(b bool)`

 SetMemberSidNil sets the value for MemberSid to be an explicit nil

### UnsetMemberSid
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetMemberSid()`

UnsetMemberSid ensures that no value is present for MemberSid, not even an explicit nil
### GetNotificationLevel

`func (o *IpMessagingV2ServiceUserUserChannel) GetNotificationLevel() string`

GetNotificationLevel returns the NotificationLevel field if non-nil, zero value otherwise.

### GetNotificationLevelOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetNotificationLevelOk() (*string, bool)`

GetNotificationLevelOk returns a tuple with the NotificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLevel

`func (o *IpMessagingV2ServiceUserUserChannel) SetNotificationLevel(v string)`

SetNotificationLevel sets NotificationLevel field to given value.

### HasNotificationLevel

`func (o *IpMessagingV2ServiceUserUserChannel) HasNotificationLevel() bool`

HasNotificationLevel returns a boolean if a field has been set.

### SetNotificationLevelNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetNotificationLevelNil(b bool)`

 SetNotificationLevelNil sets the value for NotificationLevel to be an explicit nil

### UnsetNotificationLevel
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetNotificationLevel()`

UnsetNotificationLevel ensures that no value is present for NotificationLevel, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV2ServiceUserUserChannel) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV2ServiceUserUserChannel) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV2ServiceUserUserChannel) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetStatus

`func (o *IpMessagingV2ServiceUserUserChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpMessagingV2ServiceUserUserChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpMessagingV2ServiceUserUserChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnreadMessagesCount

`func (o *IpMessagingV2ServiceUserUserChannel) GetUnreadMessagesCount() int32`

GetUnreadMessagesCount returns the UnreadMessagesCount field if non-nil, zero value otherwise.

### GetUnreadMessagesCountOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetUnreadMessagesCountOk() (*int32, bool)`

GetUnreadMessagesCountOk returns a tuple with the UnreadMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessagesCount

`func (o *IpMessagingV2ServiceUserUserChannel) SetUnreadMessagesCount(v int32)`

SetUnreadMessagesCount sets UnreadMessagesCount field to given value.

### HasUnreadMessagesCount

`func (o *IpMessagingV2ServiceUserUserChannel) HasUnreadMessagesCount() bool`

HasUnreadMessagesCount returns a boolean if a field has been set.

### SetUnreadMessagesCountNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetUnreadMessagesCountNil(b bool)`

 SetUnreadMessagesCountNil sets the value for UnreadMessagesCount to be an explicit nil

### UnsetUnreadMessagesCount
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetUnreadMessagesCount()`

UnsetUnreadMessagesCount ensures that no value is present for UnreadMessagesCount, not even an explicit nil
### GetUrl

`func (o *IpMessagingV2ServiceUserUserChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV2ServiceUserUserChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV2ServiceUserUserChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetUserSid

`func (o *IpMessagingV2ServiceUserUserChannel) GetUserSid() string`

GetUserSid returns the UserSid field if non-nil, zero value otherwise.

### GetUserSidOk

`func (o *IpMessagingV2ServiceUserUserChannel) GetUserSidOk() (*string, bool)`

GetUserSidOk returns a tuple with the UserSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSid

`func (o *IpMessagingV2ServiceUserUserChannel) SetUserSid(v string)`

SetUserSid sets UserSid field to given value.

### HasUserSid

`func (o *IpMessagingV2ServiceUserUserChannel) HasUserSid() bool`

HasUserSid returns a boolean if a field has been set.

### SetUserSidNil

`func (o *IpMessagingV2ServiceUserUserChannel) SetUserSidNil(b bool)`

 SetUserSidNil sets the value for UserSid to be an explicit nil

### UnsetUserSid
`func (o *IpMessagingV2ServiceUserUserChannel) UnsetUserSid()`

UnsetUserSid ensures that no value is present for UserSid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


