# VoiceV1IpRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CidrPrefixLength** | Pointer to **NullableInt32** | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. | [optional] 
**DateCreated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**IpAddress** | Pointer to **NullableString** | An IP address in dotted decimal notation, IPv4 only. | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewVoiceV1IpRecord

`func NewVoiceV1IpRecord() *VoiceV1IpRecord`

NewVoiceV1IpRecord instantiates a new VoiceV1IpRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoiceV1IpRecordWithDefaults

`func NewVoiceV1IpRecordWithDefaults() *VoiceV1IpRecord`

NewVoiceV1IpRecordWithDefaults instantiates a new VoiceV1IpRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *VoiceV1IpRecord) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *VoiceV1IpRecord) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *VoiceV1IpRecord) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *VoiceV1IpRecord) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *VoiceV1IpRecord) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *VoiceV1IpRecord) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCidrPrefixLength

`func (o *VoiceV1IpRecord) GetCidrPrefixLength() int32`

GetCidrPrefixLength returns the CidrPrefixLength field if non-nil, zero value otherwise.

### GetCidrPrefixLengthOk

`func (o *VoiceV1IpRecord) GetCidrPrefixLengthOk() (*int32, bool)`

GetCidrPrefixLengthOk returns a tuple with the CidrPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrPrefixLength

`func (o *VoiceV1IpRecord) SetCidrPrefixLength(v int32)`

SetCidrPrefixLength sets CidrPrefixLength field to given value.

### HasCidrPrefixLength

`func (o *VoiceV1IpRecord) HasCidrPrefixLength() bool`

HasCidrPrefixLength returns a boolean if a field has been set.

### SetCidrPrefixLengthNil

`func (o *VoiceV1IpRecord) SetCidrPrefixLengthNil(b bool)`

 SetCidrPrefixLengthNil sets the value for CidrPrefixLength to be an explicit nil

### UnsetCidrPrefixLength
`func (o *VoiceV1IpRecord) UnsetCidrPrefixLength()`

UnsetCidrPrefixLength ensures that no value is present for CidrPrefixLength, not even an explicit nil
### GetDateCreated

`func (o *VoiceV1IpRecord) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VoiceV1IpRecord) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VoiceV1IpRecord) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VoiceV1IpRecord) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *VoiceV1IpRecord) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *VoiceV1IpRecord) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *VoiceV1IpRecord) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *VoiceV1IpRecord) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *VoiceV1IpRecord) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *VoiceV1IpRecord) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *VoiceV1IpRecord) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *VoiceV1IpRecord) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *VoiceV1IpRecord) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *VoiceV1IpRecord) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *VoiceV1IpRecord) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *VoiceV1IpRecord) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *VoiceV1IpRecord) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *VoiceV1IpRecord) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIpAddress

`func (o *VoiceV1IpRecord) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VoiceV1IpRecord) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VoiceV1IpRecord) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VoiceV1IpRecord) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *VoiceV1IpRecord) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *VoiceV1IpRecord) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetSid

`func (o *VoiceV1IpRecord) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *VoiceV1IpRecord) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *VoiceV1IpRecord) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *VoiceV1IpRecord) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *VoiceV1IpRecord) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *VoiceV1IpRecord) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *VoiceV1IpRecord) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VoiceV1IpRecord) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VoiceV1IpRecord) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VoiceV1IpRecord) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *VoiceV1IpRecord) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *VoiceV1IpRecord) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


