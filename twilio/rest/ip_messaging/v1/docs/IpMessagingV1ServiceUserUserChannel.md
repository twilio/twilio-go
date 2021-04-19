# IpMessagingV1ServiceUserUserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**ChannelSid** | Pointer to **NullableString** |  | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**MemberSid** | Pointer to **NullableString** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**UnreadMessagesCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewIpMessagingV1ServiceUserUserChannel

`func NewIpMessagingV1ServiceUserUserChannel() *IpMessagingV1ServiceUserUserChannel`

NewIpMessagingV1ServiceUserUserChannel instantiates a new IpMessagingV1ServiceUserUserChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV1ServiceUserUserChannelWithDefaults

`func NewIpMessagingV1ServiceUserUserChannelWithDefaults() *IpMessagingV1ServiceUserUserChannel`

NewIpMessagingV1ServiceUserUserChannelWithDefaults instantiates a new IpMessagingV1ServiceUserUserChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV1ServiceUserUserChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV1ServiceUserUserChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV1ServiceUserUserChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *IpMessagingV1ServiceUserUserChannel) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *IpMessagingV1ServiceUserUserChannel) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *IpMessagingV1ServiceUserUserChannel) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceUserUserChannel) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceUserUserChannel) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceUserUserChannel) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLinks

`func (o *IpMessagingV1ServiceUserUserChannel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV1ServiceUserUserChannel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV1ServiceUserUserChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMemberSid

`func (o *IpMessagingV1ServiceUserUserChannel) GetMemberSid() string`

GetMemberSid returns the MemberSid field if non-nil, zero value otherwise.

### GetMemberSidOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetMemberSidOk() (*string, bool)`

GetMemberSidOk returns a tuple with the MemberSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSid

`func (o *IpMessagingV1ServiceUserUserChannel) SetMemberSid(v string)`

SetMemberSid sets MemberSid field to given value.

### HasMemberSid

`func (o *IpMessagingV1ServiceUserUserChannel) HasMemberSid() bool`

HasMemberSid returns a boolean if a field has been set.

### SetMemberSidNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetMemberSidNil(b bool)`

 SetMemberSidNil sets the value for MemberSid to be an explicit nil

### UnsetMemberSid
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetMemberSid()`

UnsetMemberSid ensures that no value is present for MemberSid, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV1ServiceUserUserChannel) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV1ServiceUserUserChannel) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV1ServiceUserUserChannel) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetStatus

`func (o *IpMessagingV1ServiceUserUserChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IpMessagingV1ServiceUserUserChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IpMessagingV1ServiceUserUserChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnreadMessagesCount

`func (o *IpMessagingV1ServiceUserUserChannel) GetUnreadMessagesCount() int32`

GetUnreadMessagesCount returns the UnreadMessagesCount field if non-nil, zero value otherwise.

### GetUnreadMessagesCountOk

`func (o *IpMessagingV1ServiceUserUserChannel) GetUnreadMessagesCountOk() (*int32, bool)`

GetUnreadMessagesCountOk returns a tuple with the UnreadMessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadMessagesCount

`func (o *IpMessagingV1ServiceUserUserChannel) SetUnreadMessagesCount(v int32)`

SetUnreadMessagesCount sets UnreadMessagesCount field to given value.

### HasUnreadMessagesCount

`func (o *IpMessagingV1ServiceUserUserChannel) HasUnreadMessagesCount() bool`

HasUnreadMessagesCount returns a boolean if a field has been set.

### SetUnreadMessagesCountNil

`func (o *IpMessagingV1ServiceUserUserChannel) SetUnreadMessagesCountNil(b bool)`

 SetUnreadMessagesCountNil sets the value for UnreadMessagesCount to be an explicit nil

### UnsetUnreadMessagesCount
`func (o *IpMessagingV1ServiceUserUserChannel) UnsetUnreadMessagesCount()`

UnsetUnreadMessagesCount ensures that no value is present for UnreadMessagesCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


