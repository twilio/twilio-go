# ApiV2010AccountAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that is responsible for the resource | [optional] 
**City** | Pointer to **NullableString** | The city in which the address is located | [optional] 
**CustomerName** | Pointer to **NullableString** | The name associated with the address | [optional] 
**DateCreated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableString** | The RFC 2822 date and time in GMT that the resource was last updated | [optional] 
**EmergencyEnabled** | Pointer to **NullableBool** | Whether emergency calling has been enabled on this number | [optional] 
**FriendlyName** | Pointer to **NullableString** | The string that you assigned to describe the resource | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code of the address | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code of the address | [optional] 
**Region** | Pointer to **NullableString** | The state or region of the address | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**Street** | Pointer to **NullableString** | The number and street address of the address | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**Validated** | Pointer to **NullableBool** | Whether the address has been validated to comply with local regulation | [optional] 
**Verified** | Pointer to **NullableBool** | Whether the address has been verified to comply with regulation | [optional] 

## Methods

### NewApiV2010AccountAddress

`func NewApiV2010AccountAddress() *ApiV2010AccountAddress`

NewApiV2010AccountAddress instantiates a new ApiV2010AccountAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountAddressWithDefaults

`func NewApiV2010AccountAddressWithDefaults() *ApiV2010AccountAddress`

NewApiV2010AccountAddressWithDefaults instantiates a new ApiV2010AccountAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountAddress) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountAddress) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountAddress) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountAddress) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountAddress) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountAddress) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCity

`func (o *ApiV2010AccountAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ApiV2010AccountAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ApiV2010AccountAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ApiV2010AccountAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ApiV2010AccountAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ApiV2010AccountAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCustomerName

`func (o *ApiV2010AccountAddress) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *ApiV2010AccountAddress) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *ApiV2010AccountAddress) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *ApiV2010AccountAddress) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### SetCustomerNameNil

`func (o *ApiV2010AccountAddress) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *ApiV2010AccountAddress) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetDateCreated

`func (o *ApiV2010AccountAddress) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ApiV2010AccountAddress) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ApiV2010AccountAddress) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ApiV2010AccountAddress) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ApiV2010AccountAddress) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ApiV2010AccountAddress) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ApiV2010AccountAddress) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ApiV2010AccountAddress) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ApiV2010AccountAddress) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ApiV2010AccountAddress) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ApiV2010AccountAddress) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ApiV2010AccountAddress) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetEmergencyEnabled

`func (o *ApiV2010AccountAddress) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *ApiV2010AccountAddress) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *ApiV2010AccountAddress) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.

### HasEmergencyEnabled

`func (o *ApiV2010AccountAddress) HasEmergencyEnabled() bool`

HasEmergencyEnabled returns a boolean if a field has been set.

### SetEmergencyEnabledNil

`func (o *ApiV2010AccountAddress) SetEmergencyEnabledNil(b bool)`

 SetEmergencyEnabledNil sets the value for EmergencyEnabled to be an explicit nil

### UnsetEmergencyEnabled
`func (o *ApiV2010AccountAddress) UnsetEmergencyEnabled()`

UnsetEmergencyEnabled ensures that no value is present for EmergencyEnabled, not even an explicit nil
### GetFriendlyName

`func (o *ApiV2010AccountAddress) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ApiV2010AccountAddress) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ApiV2010AccountAddress) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ApiV2010AccountAddress) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ApiV2010AccountAddress) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ApiV2010AccountAddress) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetIsoCountry

`func (o *ApiV2010AccountAddress) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *ApiV2010AccountAddress) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *ApiV2010AccountAddress) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *ApiV2010AccountAddress) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *ApiV2010AccountAddress) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *ApiV2010AccountAddress) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetPostalCode

`func (o *ApiV2010AccountAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ApiV2010AccountAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ApiV2010AccountAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ApiV2010AccountAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ApiV2010AccountAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ApiV2010AccountAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetRegion

`func (o *ApiV2010AccountAddress) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiV2010AccountAddress) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiV2010AccountAddress) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiV2010AccountAddress) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ApiV2010AccountAddress) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ApiV2010AccountAddress) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSid

`func (o *ApiV2010AccountAddress) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ApiV2010AccountAddress) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ApiV2010AccountAddress) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ApiV2010AccountAddress) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ApiV2010AccountAddress) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ApiV2010AccountAddress) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetStreet

`func (o *ApiV2010AccountAddress) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ApiV2010AccountAddress) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ApiV2010AccountAddress) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ApiV2010AccountAddress) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ApiV2010AccountAddress) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ApiV2010AccountAddress) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountAddress) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountAddress) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountAddress) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountAddress) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountAddress) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountAddress) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetValidated

`func (o *ApiV2010AccountAddress) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *ApiV2010AccountAddress) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *ApiV2010AccountAddress) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *ApiV2010AccountAddress) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### SetValidatedNil

`func (o *ApiV2010AccountAddress) SetValidatedNil(b bool)`

 SetValidatedNil sets the value for Validated to be an explicit nil

### UnsetValidated
`func (o *ApiV2010AccountAddress) UnsetValidated()`

UnsetValidated ensures that no value is present for Validated, not even an explicit nil
### GetVerified

`func (o *ApiV2010AccountAddress) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ApiV2010AccountAddress) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ApiV2010AccountAddress) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ApiV2010AccountAddress) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### SetVerifiedNil

`func (o *ApiV2010AccountAddress) SetVerifiedNil(b bool)`

 SetVerifiedNil sets the value for Verified to be an explicit nil

### UnsetVerified
`func (o *ApiV2010AccountAddress) UnsetVerified()`

UnsetVerified ensures that no value is present for Verified, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


