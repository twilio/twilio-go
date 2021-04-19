# ChatV1ServiceChannelMember

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**ChannelSid** | Pointer to **NullableString** | The unique ID of the Channel for the member | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT when the resource was last updated | [optional] 
**Identity** | Pointer to **NullableString** | The string that identifies the resource&#39;s User | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** | The index of the last Message that the Member has read within the Channel | [optional] 
**LastConsumptionTimestamp** | Pointer to **NullableTime** | The ISO 8601 based timestamp string that represents the date-time of the last Message read event for the Member within the Channel | [optional] 
**RoleSid** | Pointer to **NullableString** | The SID of the Role assigned to the member | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the Service that the resource is associated with | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Member resource | [optional] 

## Methods

### NewChatV1ServiceChannelMember

`func NewChatV1ServiceChannelMember() *ChatV1ServiceChannelMember`

NewChatV1ServiceChannelMember instantiates a new ChatV1ServiceChannelMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatV1ServiceChannelMemberWithDefaults

`func NewChatV1ServiceChannelMemberWithDefaults() *ChatV1ServiceChannelMember`

NewChatV1ServiceChannelMemberWithDefaults instantiates a new ChatV1ServiceChannelMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ChatV1ServiceChannelMember) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ChatV1ServiceChannelMember) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ChatV1ServiceChannelMember) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ChatV1ServiceChannelMember) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ChatV1ServiceChannelMember) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ChatV1ServiceChannelMember) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *ChatV1ServiceChannelMember) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *ChatV1ServiceChannelMember) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *ChatV1ServiceChannelMember) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *ChatV1ServiceChannelMember) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *ChatV1ServiceChannelMember) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *ChatV1ServiceChannelMember) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetDateCreated

`func (o *ChatV1ServiceChannelMember) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ChatV1ServiceChannelMember) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ChatV1ServiceChannelMember) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ChatV1ServiceChannelMember) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ChatV1ServiceChannelMember) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ChatV1ServiceChannelMember) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ChatV1ServiceChannelMember) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ChatV1ServiceChannelMember) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ChatV1ServiceChannelMember) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ChatV1ServiceChannelMember) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ChatV1ServiceChannelMember) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ChatV1ServiceChannelMember) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIdentity

`func (o *ChatV1ServiceChannelMember) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *ChatV1ServiceChannelMember) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *ChatV1ServiceChannelMember) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *ChatV1ServiceChannelMember) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *ChatV1ServiceChannelMember) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *ChatV1ServiceChannelMember) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *ChatV1ServiceChannelMember) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *ChatV1ServiceChannelMember) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *ChatV1ServiceChannelMember) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *ChatV1ServiceChannelMember) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *ChatV1ServiceChannelMember) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *ChatV1ServiceChannelMember) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLastConsumptionTimestamp

`func (o *ChatV1ServiceChannelMember) GetLastConsumptionTimestamp() time.Time`

GetLastConsumptionTimestamp returns the LastConsumptionTimestamp field if non-nil, zero value otherwise.

### GetLastConsumptionTimestampOk

`func (o *ChatV1ServiceChannelMember) GetLastConsumptionTimestampOk() (*time.Time, bool)`

GetLastConsumptionTimestampOk returns a tuple with the LastConsumptionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumptionTimestamp

`func (o *ChatV1ServiceChannelMember) SetLastConsumptionTimestamp(v time.Time)`

SetLastConsumptionTimestamp sets LastConsumptionTimestamp field to given value.

### HasLastConsumptionTimestamp

`func (o *ChatV1ServiceChannelMember) HasLastConsumptionTimestamp() bool`

HasLastConsumptionTimestamp returns a boolean if a field has been set.

### SetLastConsumptionTimestampNil

`func (o *ChatV1ServiceChannelMember) SetLastConsumptionTimestampNil(b bool)`

 SetLastConsumptionTimestampNil sets the value for LastConsumptionTimestamp to be an explicit nil

### UnsetLastConsumptionTimestamp
`func (o *ChatV1ServiceChannelMember) UnsetLastConsumptionTimestamp()`

UnsetLastConsumptionTimestamp ensures that no value is present for LastConsumptionTimestamp, not even an explicit nil
### GetRoleSid

`func (o *ChatV1ServiceChannelMember) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *ChatV1ServiceChannelMember) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *ChatV1ServiceChannelMember) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *ChatV1ServiceChannelMember) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *ChatV1ServiceChannelMember) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *ChatV1ServiceChannelMember) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetServiceSid

`func (o *ChatV1ServiceChannelMember) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ChatV1ServiceChannelMember) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ChatV1ServiceChannelMember) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ChatV1ServiceChannelMember) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ChatV1ServiceChannelMember) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ChatV1ServiceChannelMember) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ChatV1ServiceChannelMember) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ChatV1ServiceChannelMember) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ChatV1ServiceChannelMember) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ChatV1ServiceChannelMember) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ChatV1ServiceChannelMember) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ChatV1ServiceChannelMember) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ChatV1ServiceChannelMember) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChatV1ServiceChannelMember) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChatV1ServiceChannelMember) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChatV1ServiceChannelMember) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ChatV1ServiceChannelMember) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ChatV1ServiceChannelMember) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


