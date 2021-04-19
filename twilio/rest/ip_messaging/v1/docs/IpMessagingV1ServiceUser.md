# IpMessagingV1ServiceUser

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**DateUpdated** | Pointer to **NullableTime** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**IsNotifiable** | Pointer to **NullableBool** |  | [optional] 
**IsOnline** | Pointer to **NullableBool** |  | [optional] 
**JoinedChannelsCount** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**RoleSid** | Pointer to **NullableString** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpMessagingV1ServiceUser

`func NewIpMessagingV1ServiceUser() *IpMessagingV1ServiceUser`

NewIpMessagingV1ServiceUser instantiates a new IpMessagingV1ServiceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV1ServiceUserWithDefaults

`func NewIpMessagingV1ServiceUserWithDefaults() *IpMessagingV1ServiceUser`

NewIpMessagingV1ServiceUserWithDefaults instantiates a new IpMessagingV1ServiceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV1ServiceUser) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV1ServiceUser) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV1ServiceUser) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV1ServiceUser) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV1ServiceUser) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV1ServiceUser) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetAttributes

`func (o *IpMessagingV1ServiceUser) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IpMessagingV1ServiceUser) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IpMessagingV1ServiceUser) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IpMessagingV1ServiceUser) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *IpMessagingV1ServiceUser) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *IpMessagingV1ServiceUser) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetDateCreated

`func (o *IpMessagingV1ServiceUser) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IpMessagingV1ServiceUser) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IpMessagingV1ServiceUser) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IpMessagingV1ServiceUser) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *IpMessagingV1ServiceUser) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *IpMessagingV1ServiceUser) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *IpMessagingV1ServiceUser) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *IpMessagingV1ServiceUser) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *IpMessagingV1ServiceUser) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *IpMessagingV1ServiceUser) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *IpMessagingV1ServiceUser) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *IpMessagingV1ServiceUser) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *IpMessagingV1ServiceUser) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *IpMessagingV1ServiceUser) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *IpMessagingV1ServiceUser) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *IpMessagingV1ServiceUser) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *IpMessagingV1ServiceUser) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *IpMessagingV1ServiceUser) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIdentity

`func (o *IpMessagingV1ServiceUser) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IpMessagingV1ServiceUser) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IpMessagingV1ServiceUser) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IpMessagingV1ServiceUser) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *IpMessagingV1ServiceUser) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *IpMessagingV1ServiceUser) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetIsNotifiable

`func (o *IpMessagingV1ServiceUser) GetIsNotifiable() bool`

GetIsNotifiable returns the IsNotifiable field if non-nil, zero value otherwise.

### GetIsNotifiableOk

`func (o *IpMessagingV1ServiceUser) GetIsNotifiableOk() (*bool, bool)`

GetIsNotifiableOk returns a tuple with the IsNotifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotifiable

`func (o *IpMessagingV1ServiceUser) SetIsNotifiable(v bool)`

SetIsNotifiable sets IsNotifiable field to given value.

### HasIsNotifiable

`func (o *IpMessagingV1ServiceUser) HasIsNotifiable() bool`

HasIsNotifiable returns a boolean if a field has been set.

### SetIsNotifiableNil

`func (o *IpMessagingV1ServiceUser) SetIsNotifiableNil(b bool)`

 SetIsNotifiableNil sets the value for IsNotifiable to be an explicit nil

### UnsetIsNotifiable
`func (o *IpMessagingV1ServiceUser) UnsetIsNotifiable()`

UnsetIsNotifiable ensures that no value is present for IsNotifiable, not even an explicit nil
### GetIsOnline

`func (o *IpMessagingV1ServiceUser) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *IpMessagingV1ServiceUser) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *IpMessagingV1ServiceUser) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *IpMessagingV1ServiceUser) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.

### SetIsOnlineNil

`func (o *IpMessagingV1ServiceUser) SetIsOnlineNil(b bool)`

 SetIsOnlineNil sets the value for IsOnline to be an explicit nil

### UnsetIsOnline
`func (o *IpMessagingV1ServiceUser) UnsetIsOnline()`

UnsetIsOnline ensures that no value is present for IsOnline, not even an explicit nil
### GetJoinedChannelsCount

`func (o *IpMessagingV1ServiceUser) GetJoinedChannelsCount() int32`

GetJoinedChannelsCount returns the JoinedChannelsCount field if non-nil, zero value otherwise.

### GetJoinedChannelsCountOk

`func (o *IpMessagingV1ServiceUser) GetJoinedChannelsCountOk() (*int32, bool)`

GetJoinedChannelsCountOk returns a tuple with the JoinedChannelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedChannelsCount

`func (o *IpMessagingV1ServiceUser) SetJoinedChannelsCount(v int32)`

SetJoinedChannelsCount sets JoinedChannelsCount field to given value.

### HasJoinedChannelsCount

`func (o *IpMessagingV1ServiceUser) HasJoinedChannelsCount() bool`

HasJoinedChannelsCount returns a boolean if a field has been set.

### SetJoinedChannelsCountNil

`func (o *IpMessagingV1ServiceUser) SetJoinedChannelsCountNil(b bool)`

 SetJoinedChannelsCountNil sets the value for JoinedChannelsCount to be an explicit nil

### UnsetJoinedChannelsCount
`func (o *IpMessagingV1ServiceUser) UnsetJoinedChannelsCount()`

UnsetJoinedChannelsCount ensures that no value is present for JoinedChannelsCount, not even an explicit nil
### GetLinks

`func (o *IpMessagingV1ServiceUser) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IpMessagingV1ServiceUser) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IpMessagingV1ServiceUser) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IpMessagingV1ServiceUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *IpMessagingV1ServiceUser) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *IpMessagingV1ServiceUser) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRoleSid

`func (o *IpMessagingV1ServiceUser) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *IpMessagingV1ServiceUser) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *IpMessagingV1ServiceUser) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *IpMessagingV1ServiceUser) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *IpMessagingV1ServiceUser) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *IpMessagingV1ServiceUser) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV1ServiceUser) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV1ServiceUser) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV1ServiceUser) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV1ServiceUser) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV1ServiceUser) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV1ServiceUser) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *IpMessagingV1ServiceUser) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpMessagingV1ServiceUser) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpMessagingV1ServiceUser) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpMessagingV1ServiceUser) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IpMessagingV1ServiceUser) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IpMessagingV1ServiceUser) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *IpMessagingV1ServiceUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV1ServiceUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV1ServiceUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV1ServiceUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV1ServiceUser) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV1ServiceUser) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


