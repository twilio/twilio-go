# PricingV1MessagingMessagingCountryInstance

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**InboundSmsPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceInboundSmsPrices.md) | The list of InboundPrice records | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code | [optional] 
**OutboundSmsPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices.md) | The list of OutboundSMSPrice records | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV1MessagingMessagingCountryInstance

`func NewPricingV1MessagingMessagingCountryInstance() *PricingV1MessagingMessagingCountryInstance`

NewPricingV1MessagingMessagingCountryInstance instantiates a new PricingV1MessagingMessagingCountryInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV1MessagingMessagingCountryInstanceWithDefaults

`func NewPricingV1MessagingMessagingCountryInstanceWithDefaults() *PricingV1MessagingMessagingCountryInstance`

NewPricingV1MessagingMessagingCountryInstanceWithDefaults instantiates a new PricingV1MessagingMessagingCountryInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV1MessagingMessagingCountryInstance) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV1MessagingMessagingCountryInstance) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV1MessagingMessagingCountryInstance) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetInboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) GetInboundSmsPrices() []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices`

GetInboundSmsPrices returns the InboundSmsPrices field if non-nil, zero value otherwise.

### GetInboundSmsPricesOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetInboundSmsPricesOk() (*[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices, bool)`

GetInboundSmsPricesOk returns a tuple with the InboundSmsPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) SetInboundSmsPrices(v []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices)`

SetInboundSmsPrices sets InboundSmsPrices field to given value.

### HasInboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) HasInboundSmsPrices() bool`

HasInboundSmsPrices returns a boolean if a field has been set.

### SetInboundSmsPricesNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetInboundSmsPricesNil(b bool)`

 SetInboundSmsPricesNil sets the value for InboundSmsPrices to be an explicit nil

### UnsetInboundSmsPrices
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetInboundSmsPrices()`

UnsetInboundSmsPrices ensures that no value is present for InboundSmsPrices, not even an explicit nil
### GetIsoCountry

`func (o *PricingV1MessagingMessagingCountryInstance) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV1MessagingMessagingCountryInstance) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV1MessagingMessagingCountryInstance) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetOutboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) GetOutboundSmsPrices() []PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices`

GetOutboundSmsPrices returns the OutboundSmsPrices field if non-nil, zero value otherwise.

### GetOutboundSmsPricesOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetOutboundSmsPricesOk() (*[]PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices, bool)`

GetOutboundSmsPricesOk returns a tuple with the OutboundSmsPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) SetOutboundSmsPrices(v []PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices)`

SetOutboundSmsPrices sets OutboundSmsPrices field to given value.

### HasOutboundSmsPrices

`func (o *PricingV1MessagingMessagingCountryInstance) HasOutboundSmsPrices() bool`

HasOutboundSmsPrices returns a boolean if a field has been set.

### SetOutboundSmsPricesNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetOutboundSmsPricesNil(b bool)`

 SetOutboundSmsPricesNil sets the value for OutboundSmsPrices to be an explicit nil

### UnsetOutboundSmsPrices
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetOutboundSmsPrices()`

UnsetOutboundSmsPrices ensures that no value is present for OutboundSmsPrices, not even an explicit nil
### GetPriceUnit

`func (o *PricingV1MessagingMessagingCountryInstance) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV1MessagingMessagingCountryInstance) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV1MessagingMessagingCountryInstance) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV1MessagingMessagingCountryInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV1MessagingMessagingCountryInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV1MessagingMessagingCountryInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV1MessagingMessagingCountryInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV1MessagingMessagingCountryInstance) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV1MessagingMessagingCountryInstance) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


