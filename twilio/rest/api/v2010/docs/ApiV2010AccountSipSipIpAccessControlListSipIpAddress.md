# ApiV2010AccountSipSipIpAccessControlListSipIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The unique id of the Account that is responsible for this resource. | [optional] 
**CidrPrefixLength** | Pointer to **NullableInt32** | An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used. | [optional] 
**DateCreated** | Pointer to **NullableString** | The date that this resource was created, given as GMT in RFC 2822 format. | [optional] 
**DateUpdated** | Pointer to **NullableString** | The date that this resource was last updated, given as GMT in RFC 2822 format. | [optional] 
**FriendlyName** | Pointer to **NullableString** | A human readable descriptive text for this resource, up to 64 characters long. | [optional] 
**IpAccessControlListSid** | Pointer to **NullableString** | The unique id of the IpAccessControlList resource that includes this resource. | [optional] 
**IpAddress** | Pointer to **NullableString** | An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today. | [optional] 
**Sid** | Pointer to **NullableString** | A 34 character string that uniquely identifies this resource. | [optional] 
**Uri** | Pointer to **NullableString** | The URI for this resource, relative to https://api.twilio.com | [optional] 

## Methods

### NewApiV2010AccountSipSipIpAccessControlListSipIpAddress

`func NewApiV2010AccountSipSipIpAccessControlListSipIpAddress() *ApiV2010AccountSipSipIpAccessControlListSipIpAddress`

NewApiV2010AccountSipSipIpAccessControlListSipIpAddress instantiates a new ApiV2010AccountSipSipIpAccessControlListSipIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountSipSipIpAccessControlListSipIpAddressWithDefaults

`func NewApiV2010AccountSipSipIpAccessControlListSipIpAddressWithDefaults() *ApiV2010AccountSipSipIpAccessControlListSipIpAddress`

NewApiV2010AccountSipSipIpAccessControlListSipIpAddressWithDefaults instantiates a new ApiV2010AccountSipSipIpAccessControlListSipIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCidrPrefixLength

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetCidrPrefixLength() int32`

GetCidrPrefixLength returns the CidrPrefixLength field if non-nil, zero value otherwise.

### GetCidrPrefixLengthOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetCidrPrefixLengthOk() (*int32, bool)`

GetCidrPrefixLengthOk returns a tuple with the CidrPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrPrefixLength

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetCidrPrefixLength(v int32)`

SetCidrPrefixLength sets CidrPrefixLength field to given value.

### HasCidrPrefixLength

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasCidrPrefixLength() bool`

HasCidrPrefixLength returns a boolean if a field has been set.

### SetCidrPrefixLengthNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetCidrPrefixLengthNil(b bool)`

 SetCidrPrefixLengthNil sets the value for CidrPrefixLength to be an explicit nil

### UnsetCidrPrefixLength
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetCidrPrefixLength()`

UnsetCidrPrefixLength ensures that no value is present for CidrPrefixLength, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIpAccessControlListSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetIpAccessControlListSid() string`

GetIpAccessControlListSid returns the IpAccessControlListSid field if non-nil, zero value otherwise.

### GetIpAccessControlListSidOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetIpAccessControlListSidOk() (*string, bool)`

GetIpAccessControlListSidOk returns a tuple with the IpAccessControlListSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAccessControlListSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetIpAccessControlListSid(v string)`

SetIpAccessControlListSid sets IpAccessControlListSid field to given value.

### HasIpAccessControlListSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasIpAccessControlListSid() bool`

HasIpAccessControlListSid returns a boolean if a field has been set.

### SetIpAccessControlListSidNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetIpAccessControlListSidNil(b bool)`

 SetIpAccessControlListSidNil sets the value for IpAccessControlListSid to be an explicit nil

### UnsetIpAccessControlListSid
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetIpAccessControlListSid()`

UnsetIpAccessControlListSid ensures that no value is present for IpAccessControlListSid, not even an explicit nil
### GetIpAddress

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountSipSipIpAccessControlListSipIpAddress) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


