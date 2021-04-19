# IpMessagingV2ServiceChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**DateUpdated** | Pointer to **NullableTime** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**MembersCount** | Pointer to **NullableInt32** |  | [optional] 
**MessagesCount** | Pointer to **NullableInt32** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UniqueName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpMessagingV2ServiceChannel

`func NewIpMessagingV2ServiceChannel() *IpMessagingV2ServiceChannel`

NewIpMessagingV2ServiceChannel instantiates a new IpMessagingV2ServiceChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV2ServiceChannelWithDefaults

`func NewIpMessagingV2ServiceChannelWithDefaults() *IpMessagingV2ServiceChannel`

NewIpMessagingV2ServiceChannelWithDefaults instantiates a new IpMessagingV2ServiceChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV2ServiceChannel) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV2ServiceChannel) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV2ServiceChannel) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV2ServiceChannel) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV2ServiceChannel) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV2ServiceChannel) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *IpMessagingV2ServiceChannel) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IpMessagingV2ServiceChannel) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IpMessagingV2ServiceChannel) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IpMessagingV2ServiceChannel) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *IpMessagingV2ServiceChannel) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *IpMessagingV2ServiceChannel) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCreatedBy

`func (o *IpMessagingV2ServiceChannel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IpMessagingV2ServiceChannel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IpMessagingV2ServiceChannel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IpMessagingV2ServiceChannel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *IpMessagingV2ServiceChannel) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *IpMessagingV2ServiceChannel) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *IpMessagingV2ServiceChannel) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IpMessagingV2ServiceChannel) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IpMessagingV2ServiceChannel) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IpMessagingV2ServiceChannel) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *IpMessagingV2ServiceChannel) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *IpMessagingV2ServiceChannel) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *IpMessagingV2ServiceChannel) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *IpMessagingV2ServiceChannel) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *IpMessagingV2ServiceChannel) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *IpMessagingV2ServiceChannel) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *IpMessagingV2ServiceChannel) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *IpMessagingV2ServiceChannel) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *IpMessagingV2ServiceChannel) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *IpMessagingV2ServiceChannel) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *IpMessagingV2ServiceChannel) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *IpMessagingV2ServiceChannel) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *IpMessagingV2ServiceChannel) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *IpMessagingV2ServiceChannel) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetLinks

`func (o *IpMessagingV2ServiceChannel) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV2ServiceChannel) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV2ServiceChannel) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV2ServiceChannel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV2ServiceChannel) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV2ServiceChannel) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMembersCount

`func (o *IpMessagingV2ServiceChannel) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *IpMessagingV2ServiceChannel) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *IpMessagingV2ServiceChannel) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *IpMessagingV2ServiceChannel) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### SetMembersCountNil

`func (o *IpMessagingV2ServiceChannel) SetMembersCountNil(b bool)`

 SetMembersCountNil sets the value for MembersCount to be an explicit nil

### UnsetMembersCount
`func (o *IpMessagingV2ServiceChannel) UnsetMembersCount()`

UnsetMembersCount ensures that no value is present for MembersCount, not even an explicit nil
### GetMessagesCount

`func (o *IpMessagingV2ServiceChannel) GetMessagesCount() int32`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *IpMessagingV2ServiceChannel) GetMessagesCountOk() (*int32, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *IpMessagingV2ServiceChannel) SetMessagesCount(v int32)`

SetMessagesCount sets MessagesCount field to given value.

### HasMessagesCount

`func (o *IpMessagingV2ServiceChannel) HasMessagesCount() bool`

HasMessagesCount returns a boolean if a field has been set.

### SetMessagesCountNil

`func (o *IpMessagingV2ServiceChannel) SetMessagesCountNil(b bool)`

 SetMessagesCountNil sets the value for MessagesCount to be an explicit nil

### UnsetMessagesCount
`func (o *IpMessagingV2ServiceChannel) UnsetMessagesCount()`

UnsetMessagesCount ensures that no value is present for MessagesCount, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV2ServiceChannel) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV2ServiceChannel) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV2ServiceChannel) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV2ServiceChannel) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV2ServiceChannel) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV2ServiceChannel) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *IpMessagingV2ServiceChannel) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpMessagingV2ServiceChannel) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpMessagingV2ServiceChannel) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpMessagingV2ServiceChannel) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IpMessagingV2ServiceChannel) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IpMessagingV2ServiceChannel) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetType

`func (o *IpMessagingV2ServiceChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpMessagingV2ServiceChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpMessagingV2ServiceChannel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpMessagingV2ServiceChannel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *IpMessagingV2ServiceChannel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *IpMessagingV2ServiceChannel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUniqueName

`func (o *IpMessagingV2ServiceChannel) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *IpMessagingV2ServiceChannel) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *IpMessagingV2ServiceChannel) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *IpMessagingV2ServiceChannel) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *IpMessagingV2ServiceChannel) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *IpMessagingV2ServiceChannel) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *IpMessagingV2ServiceChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV2ServiceChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV2ServiceChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV2ServiceChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV2ServiceChannel) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV2ServiceChannel) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


