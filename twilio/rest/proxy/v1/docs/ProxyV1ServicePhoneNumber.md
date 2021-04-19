# ProxyV1ServicePhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**Capabilities** | Pointer to [**NullableProxyV1ServicePhoneNumberCapabilities**](proxy_v1_service_phone_number_capabilities.md) |  | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**InUse** | Pointer to **NullableInt32** | The number of open session assigned to the number. | [optional] 
**IsReserved** | Pointer to **NullableBool** | Reserve the phone number for manual assignment to participants only | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO Country Code | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The phone number in E.164 format | [optional] 
**ServiceSid** | Pointer to **NullableString** | The SID of the PhoneNumber resource&#39;s parent Service resource | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the PhoneNumber resource | [optional] 

## Methods

### NewProxyV1ServicePhoneNumber

`func NewProxyV1ServicePhoneNumber() *ProxyV1ServicePhoneNumber`

NewProxyV1ServicePhoneNumber instantiates a new ProxyV1ServicePhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServicePhoneNumberWithDefaults

`func NewProxyV1ServicePhoneNumberWithDefaults() *ProxyV1ServicePhoneNumber`

NewProxyV1ServicePhoneNumberWithDefaults instantiates a new ProxyV1ServicePhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1ServicePhoneNumber) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1ServicePhoneNumber) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1ServicePhoneNumber) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1ServicePhoneNumber) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1ServicePhoneNumber) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1ServicePhoneNumber) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCapabilities

`func (o *ProxyV1ServicePhoneNumber) GetCapabilities() ProxyV1ServicePhoneNumberCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ProxyV1ServicePhoneNumber) GetCapabilitiesOk() (*ProxyV1ServicePhoneNumberCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ProxyV1ServicePhoneNumber) SetCapabilities(v ProxyV1ServicePhoneNumberCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ProxyV1ServicePhoneNumber) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ProxyV1ServicePhoneNumber) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ProxyV1ServicePhoneNumber) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1ServicePhoneNumber) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1ServicePhoneNumber) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1ServicePhoneNumber) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1ServicePhoneNumber) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1ServicePhoneNumber) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1ServicePhoneNumber) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1ServicePhoneNumber) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1ServicePhoneNumber) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1ServicePhoneNumber) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1ServicePhoneNumber) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1ServicePhoneNumber) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1ServicePhoneNumber) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetFriendlyName

`func (o *ProxyV1ServicePhoneNumber) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ProxyV1ServicePhoneNumber) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ProxyV1ServicePhoneNumber) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ProxyV1ServicePhoneNumber) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ProxyV1ServicePhoneNumber) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ProxyV1ServicePhoneNumber) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetInUse

`func (o *ProxyV1ServicePhoneNumber) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *ProxyV1ServicePhoneNumber) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *ProxyV1ServicePhoneNumber) SetInUse(v int32)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *ProxyV1ServicePhoneNumber) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### SetInUseNil

`func (o *ProxyV1ServicePhoneNumber) SetInUseNil(b bool)`

 SetInUseNil sets the value for InUse to be an explicit nil

### UnsetInUse
`func (o *ProxyV1ServicePhoneNumber) UnsetInUse()`

UnsetInUse ensures that no value is present for InUse, not even an explicit nil
### GetIsReserved

`func (o *ProxyV1ServicePhoneNumber) GetIsReserved() bool`

GetIsReserved returns the IsReserved field if non-nil, zero value otherwise.

### GetIsReservedOk

`func (o *ProxyV1ServicePhoneNumber) GetIsReservedOk() (*bool, bool)`

GetIsReservedOk returns a tuple with the IsReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReserved

`func (o *ProxyV1ServicePhoneNumber) SetIsReserved(v bool)`

SetIsReserved sets IsReserved field to given value.

### HasIsReserved

`func (o *ProxyV1ServicePhoneNumber) HasIsReserved() bool`

HasIsReserved returns a boolean if a field has been set.

### SetIsReservedNil

`func (o *ProxyV1ServicePhoneNumber) SetIsReservedNil(b bool)`

 SetIsReservedNil sets the value for IsReserved to be an explicit nil

### UnsetIsReserved
`func (o *ProxyV1ServicePhoneNumber) UnsetIsReserved()`

UnsetIsReserved ensures that no value is present for IsReserved, not even an explicit nil
### GetIsoCountry

`func (o *ProxyV1ServicePhoneNumber) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *ProxyV1ServicePhoneNumber) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *ProxyV1ServicePhoneNumber) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *ProxyV1ServicePhoneNumber) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *ProxyV1ServicePhoneNumber) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *ProxyV1ServicePhoneNumber) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetPhoneNumber

`func (o *ProxyV1ServicePhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ProxyV1ServicePhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ProxyV1ServicePhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ProxyV1ServicePhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *ProxyV1ServicePhoneNumber) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *ProxyV1ServicePhoneNumber) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetServiceSid

`func (o *ProxyV1ServicePhoneNumber) GetServiceSid() string`

GetServiceSid returns the ServiceSid field if non-nil, zero value otherwise.

### GetServiceSidOk

`func (o *ProxyV1ServicePhoneNumber) GetServiceSidOk() (*string, bool)`

GetServiceSidOk returns a tuple with the ServiceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSid

`func (o *ProxyV1ServicePhoneNumber) SetServiceSid(v string)`

SetServiceSid sets ServiceSid field to given value.

### HasServiceSid

`func (o *ProxyV1ServicePhoneNumber) HasServiceSid() bool`

HasServiceSid returns a boolean if a field has been set.

### SetServiceSidNil

`func (o *ProxyV1ServicePhoneNumber) SetServiceSidNil(b bool)`

 SetServiceSidNil sets the value for ServiceSid to be an explicit nil

### UnsetServiceSid
`func (o *ProxyV1ServicePhoneNumber) UnsetServiceSid()`

UnsetServiceSid ensures that no value is present for ServiceSid, not even an explicit nil
### GetSid

`func (o *ProxyV1ServicePhoneNumber) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1ServicePhoneNumber) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1ServicePhoneNumber) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1ServicePhoneNumber) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1ServicePhoneNumber) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1ServicePhoneNumber) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUrl

`func (o *ProxyV1ServicePhoneNumber) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1ServicePhoneNumber) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1ServicePhoneNumber) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1ServicePhoneNumber) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1ServicePhoneNumber) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1ServicePhoneNumber) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


