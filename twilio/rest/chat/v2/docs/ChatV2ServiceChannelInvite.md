# ChatV2ServiceChannelInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ChannelSid** | Pointer to **NullableString** | The SID of the Channel the new resource belongs to | [optional] 
**CreatedBy** | Pointer to **NullableString** | The identity of the User that created the invite | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**RoleSid** | Pointer to **NullableString** | The SID of the Role assigned to the member | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Invite resource | [optional] 

## Methods

### NewChatV2ServiceChannelInvite

`func NewChatV2ServiceChannelInvite() *ChatV2ServiceChannelInvite`

NewChatV2ServiceChannelInvite instantiates a new ChatV2ServiceChannelInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV2ServiceChannelInviteWithDefaults

`func NewChatV2ServiceChannelInviteWithDefaults() *ChatV2ServiceChannelInvite`

NewChatV2ServiceChannelInviteWithDefaults instantiates a new ChatV2ServiceChannelInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV2ServiceChannelInvite) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV2ServiceChannelInvite) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV2ServiceChannelInvite) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV2ServiceChannelInvite) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV2ServiceChannelInvite) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV2ServiceChannelInvite) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *ChatV2ServiceChannelInvite) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV2ServiceChannelInvite) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV2ServiceChannelInvite) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV2ServiceChannelInvite) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV2ServiceChannelInvite) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV2ServiceChannelInvite) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetCreatedBy

`func (o *ChatV2ServiceChannelInvite) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChatV2ServiceChannelInvite) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChatV2ServiceChannelInvite) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ChatV2ServiceChannelInvite) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ChatV2ServiceChannelInvite) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ChatV2ServiceChannelInvite) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *ChatV2ServiceChannelInvite) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV2ServiceChannelInvite) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV2ServiceChannelInvite) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV2ServiceChannelInvite) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV2ServiceChannelInvite) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV2ServiceChannelInvite) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV2ServiceChannelInvite) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV2ServiceChannelInvite) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV2ServiceChannelInvite) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV2ServiceChannelInvite) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV2ServiceChannelInvite) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV2ServiceChannelInvite) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIdentity

`func (o *ChatV2ServiceChannelInvite) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ChatV2ServiceChannelInvite) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ChatV2ServiceChannelInvite) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ChatV2ServiceChannelInvite) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ChatV2ServiceChannelInvite) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ChatV2ServiceChannelInvite) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetRoleSid

`func (o *ChatV2ServiceChannelInvite) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ChatV2ServiceChannelInvite) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ChatV2ServiceChannelInvite) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ChatV2ServiceChannelInvite) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ChatV2ServiceChannelInvite) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ChatV2ServiceChannelInvite) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetServiceSid

`func (o *ChatV2ServiceChannelInvite) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV2ServiceChannelInvite) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV2ServiceChannelInvite) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV2ServiceChannelInvite) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV2ServiceChannelInvite) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV2ServiceChannelInvite) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV2ServiceChannelInvite) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV2ServiceChannelInvite) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV2ServiceChannelInvite) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV2ServiceChannelInvite) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV2ServiceChannelInvite) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV2ServiceChannelInvite) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ChatV2ServiceChannelInvite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV2ServiceChannelInvite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV2ServiceChannelInvite) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV2ServiceChannelInvite) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV2ServiceChannelInvite) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV2ServiceChannelInvite) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


