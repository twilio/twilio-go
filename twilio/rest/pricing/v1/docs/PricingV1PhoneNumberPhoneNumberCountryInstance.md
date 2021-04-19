# PricingV1PhoneNumberPhoneNumberCountryInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code  | [optional] 
**PhoneNumberPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceInboundSmsPrices.md) | The list of PhoneNumberPrices records | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV1PhoneNumberPhoneNumberCountryInstance

`func NewPricingV1PhoneNumberPhoneNumberCountryInstance() *PricingV1PhoneNumberPhoneNumberCountryInstance`

NewPricingV1PhoneNumberPhoneNumberCountryInstance instantiates a new PricingV1PhoneNumberPhoneNumberCountryInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV1PhoneNumberPhoneNumberCountryInstanceWithDefaults

`func NewPricingV1PhoneNumberPhoneNumberCountryInstanceWithDefaults() *PricingV1PhoneNumberPhoneNumberCountryInstance`

NewPricingV1PhoneNumberPhoneNumberCountryInstanceWithDefaults instantiates a new PricingV1PhoneNumberPhoneNumberCountryInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetIsoCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetPhoneNumberPrices

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetPhoneNumberPrices() []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices`

GetPhoneNumberPrices returns the PhoneNumberPrices field if non-nil, zero value otherwise.

### GetPhoneNumberPricesOk

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetPhoneNumberPricesOk() (*[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices, bool)`

GetPhoneNumberPricesOk returns a tuple with the PhoneNumberPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberPrices

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetPhoneNumberPrices(v []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices)`

SetPhoneNumberPrices sets PhoneNumberPrices field to given value.

### HasPhoneNumberPrices

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) HasPhoneNumberPrices() bool`

HasPhoneNumberPrices returns a boolean if a field has been set.

### SetPhoneNumberPricesNil

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetPhoneNumberPricesNil(b bool)`

 SetPhoneNumberPricesNil sets the value for PhoneNumberPrices to be an explicit nil

### UnsetPhoneNumberPrices
`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) UnsetPhoneNumberPrices()`

UnsetPhoneNumberPrices ensures that no value is present for PhoneNumberPrices, not even an explicit nil
### GetPriceUnit

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV1PhoneNumberPhoneNumberCountryInstance) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


