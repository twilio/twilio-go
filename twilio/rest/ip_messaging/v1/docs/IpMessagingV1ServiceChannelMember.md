# IpMessagingV1ServiceChannelMember

## Properties

Name | Type | Description
------------ | ------------- | -------------
**AccountSid** | Pointer to **NullableString** |  | [optional] 
**ChannelSid** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**DateUpdated** | Pointer to **NullableTime** |  | [optional] 
**Identity** | Pointer to **NullableString** |  | [optional] 
**LastConsumedMessageIndex** | Pointer to **NullableInt32** |  | [optional] 
**LastConsumptionTimestamp** | Pointer to **NullableTime** |  | [optional] 
**RoleSid** | Pointer to **NullableString** |  | [optional] 
**ServiceSid** | Pointer to **NullableString** |  | [optional] 
**Sid** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewIpMessagingV1ServiceChannelMember

`func NewIpMessagingV1ServiceChannelMember() *IpMessagingV1ServiceChannelMember`

NewIpMessagingV1ServiceChannelMember instantiates a new IpMessagingV1ServiceChannelMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpMessagingV1ServiceChannelMemberWithDefaults

`func NewIpMessagingV1ServiceChannelMemberWithDefaults() *IpMessagingV1ServiceChannelMember`

NewIpMessagingV1ServiceChannelMemberWithDefaults instantiates a new IpMessagingV1ServiceChannelMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *IpMessagingV1ServiceChannelMember) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *IpMessagingV1ServiceChannelMember) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *IpMessagingV1ServiceChannelMember) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *IpMessagingV1ServiceChannelMember) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *IpMessagingV1ServiceChannelMember) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *IpMessagingV1ServiceChannelMember) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetChannelSid

`func (o *IpMessagingV1ServiceChannelMember) GetChannelSid() string`

GetChannelSid returns the ChannelSid field if non-nil, zero value otherwise.

### GetChannelSidOk

`func (o *IpMessagingV1ServiceChannelMember) GetChannelSidOk() (*string, bool)`

GetChannelSidOk returns a tuple with the ChannelSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSid

`func (o *IpMessagingV1ServiceChannelMember) SetChannelSid(v string)`

SetChannelSid sets ChannelSid field to given value.

### HasChannelSid

`func (o *IpMessagingV1ServiceChannelMember) HasChannelSid() bool`

HasChannelSid returns a boolean if a field has been set.

### SetChannelSidNil

`func (o *IpMessagingV1ServiceChannelMember) SetChannelSidNil(b bool)`

 SetChannelSidNil sets the value for ChannelSid to be an explicit nil

### UnsetChannelSid
`func (o *IpMessagingV1ServiceChannelMember) UnsetChannelSid()`

UnsetChannelSid ensures that no value is present for ChannelSid, not even an explicit nil
### GetDateCreated

`func (o *IpMessagingV1ServiceChannelMember) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *IpMessagingV1ServiceChannelMember) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *IpMessagingV1ServiceChannelMember) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *IpMessagingV1ServiceChannelMember) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *IpMessagingV1ServiceChannelMember) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *IpMessagingV1ServiceChannelMember) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *IpMessagingV1ServiceChannelMember) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *IpMessagingV1ServiceChannelMember) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *IpMessagingV1ServiceChannelMember) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *IpMessagingV1ServiceChannelMember) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *IpMessagingV1ServiceChannelMember) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *IpMessagingV1ServiceChannelMember) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetIdentity

`func (o *IpMessagingV1ServiceChannelMember) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *IpMessagingV1ServiceChannelMember) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *IpMessagingV1ServiceChannelMember) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *IpMessagingV1ServiceChannelMember) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### SetIdentityNil

`func (o *IpMessagingV1ServiceChannelMember) SetIdentityNil(b bool)`

 SetIdentityNil sets the value for Identity to be an explicit nil

### UnsetIdentity
`func (o *IpMessagingV1ServiceChannelMember) UnsetIdentity()`

UnsetIdentity ensures that no value is present for Identity, not even an explicit nil
### GetLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceChannelMember) GetLastConsumedMessageIndex() int32`

GetLastConsumedMessageIndex returns the LastConsumedMessageIndex field if non-nil, zero value otherwise.

### GetLastConsumedMessageIndexOk

`func (o *IpMessagingV1ServiceChannelMember) GetLastConsumedMessageIndexOk() (*int32, bool)`

GetLastConsumedMessageIndexOk returns a tuple with the LastConsumedMessageIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceChannelMember) SetLastConsumedMessageIndex(v int32)`

SetLastConsumedMessageIndex sets LastConsumedMessageIndex field to given value.

### HasLastConsumedMessageIndex

`func (o *IpMessagingV1ServiceChannelMember) HasLastConsumedMessageIndex() bool`

HasLastConsumedMessageIndex returns a boolean if a field has been set.

### SetLastConsumedMessageIndexNil

`func (o *IpMessagingV1ServiceChannelMember) SetLastConsumedMessageIndexNil(b bool)`

 SetLastConsumedMessageIndexNil sets the value for LastConsumedMessageIndex to be an explicit nil

### UnsetLastConsumedMessageIndex
`func (o *IpMessagingV1ServiceChannelMember) UnsetLastConsumedMessageIndex()`

UnsetLastConsumedMessageIndex ensures that no value is present for LastConsumedMessageIndex, not even an explicit nil
### GetLastConsumptionTimestamp

`func (o *IpMessagingV1ServiceChannelMember) GetLastConsumptionTimestamp() time.Time`

GetLastConsumptionTimestamp returns the LastConsumptionTimestamp field if non-nil, zero value otherwise.

### GetLastConsumptionTimestampOk

`func (o *IpMessagingV1ServiceChannelMember) GetLastConsumptionTimestampOk() (*time.Time, bool)`

GetLastConsumptionTimestampOk returns a tuple with the LastConsumptionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumptionTimestamp

`func (o *IpMessagingV1ServiceChannelMember) SetLastConsumptionTimestamp(v time.Time)`

SetLastConsumptionTimestamp sets LastConsumptionTimestamp field to given value.

### HasLastConsumptionTimestamp

`func (o *IpMessagingV1ServiceChannelMember) HasLastConsumptionTimestamp() bool`

HasLastConsumptionTimestamp returns a boolean if a field has been set.

### SetLastConsumptionTimestampNil

`func (o *IpMessagingV1ServiceChannelMember) SetLastConsumptionTimestampNil(b bool)`

 SetLastConsumptionTimestampNil sets the value for LastConsumptionTimestamp to be an explicit nil

### UnsetLastConsumptionTimestamp
`func (o *IpMessagingV1ServiceChannelMember) UnsetLastConsumptionTimestamp()`

UnsetLastConsumptionTimestamp ensures that no value is present for LastConsumptionTimestamp, not even an explicit nil
### GetRoleSid

`func (o *IpMessagingV1ServiceChannelMember) GetRoleSid() string`

GetRoleSid returns the RoleSid field if non-nil, zero value otherwise.

### GetRoleSidOk

`func (o *IpMessagingV1ServiceChannelMember) GetRoleSidOk() (*string, bool)`

GetRoleSidOk returns a tuple with the RoleSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleSid

`func (o *IpMessagingV1ServiceChannelMember) SetRoleSid(v string)`

SetRoleSid sets RoleSid field to given value.

### HasRoleSid

`func (o *IpMessagingV1ServiceChannelMember) HasRoleSid() bool`

HasRoleSid returns a boolean if a field has been set.

### SetRoleSidNil

`func (o *IpMessagingV1ServiceChannelMember) SetRoleSidNil(b bool)`

 SetRoleSidNil sets the value for RoleSid to be an explicit nil

### UnsetRoleSid
`func (o *IpMessagingV1ServiceChannelMember) UnsetRoleSid()`

UnsetRoleSid ensures that no value is present for RoleSid, not even an explicit nil
### GetServiceSid

`func (o *IpMessagingV1ServiceChannelMember) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *IpMessagingV1ServiceChannelMember) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *IpMessagingV1ServiceChannelMember) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *IpMessagingV1ServiceChannelMember) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *IpMessagingV1ServiceChannelMember) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *IpMessagingV1ServiceChannelMember) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *IpMessagingV1ServiceChannelMember) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpMessagingV1ServiceChannelMember) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpMessagingV1ServiceChannelMember) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpMessagingV1ServiceChannelMember) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *IpMessagingV1ServiceChannelMember) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *IpMessagingV1ServiceChannelMember) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *IpMessagingV1ServiceChannelMember) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IpMessagingV1ServiceChannelMember) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IpMessagingV1ServiceChannelMember) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IpMessagingV1ServiceChannelMember) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *IpMessagingV1ServiceChannelMember) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IpMessagingV1ServiceChannelMember) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


