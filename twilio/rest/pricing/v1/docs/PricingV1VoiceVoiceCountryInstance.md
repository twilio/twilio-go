# PricingV1VoiceVoiceCountryInstance

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**InboundCallPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceInboundSmsPrices.md) | The list of InboundCallPrice records | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code | [optional] 
**OutboundPrefixPrices** | Pointer to [**[]PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices**](PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices.md) | The list of OutboundPrefixPrice records | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV1VoiceVoiceCountryInstance

`func NewPricingV1VoiceVoiceCountryInstance() *PricingV1VoiceVoiceCountryInstance`

NewPricingV1VoiceVoiceCountryInstance instantiates a new PricingV1VoiceVoiceCountryInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV1VoiceVoiceCountryInstanceWithDefaults

`func NewPricingV1VoiceVoiceCountryInstanceWithDefaults() *PricingV1VoiceVoiceCountryInstance`

NewPricingV1VoiceVoiceCountryInstanceWithDefaults instantiates a new PricingV1VoiceVoiceCountryInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV1VoiceVoiceCountryInstance) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV1VoiceVoiceCountryInstance) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV1VoiceVoiceCountryInstance) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetInboundCallPrices

`func (o *PricingV1VoiceVoiceCountryInstance) GetInboundCallPrices() []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices`

GetInboundCallPrices returns the InboundCallPrices field if non-nil, zero value otherwise.

### GetInboundCallPricesOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetInboundCallPricesOk() (*[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices, bool)`

GetInboundCallPricesOk returns a tuple with the InboundCallPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallPrices

`func (o *PricingV1VoiceVoiceCountryInstance) SetInboundCallPrices(v []PricingV1MessagingMessagingCountryInstanceInboundSmsPrices)`

SetInboundCallPrices sets InboundCallPrices field to given value.

### HasInboundCallPrices

`func (o *PricingV1VoiceVoiceCountryInstance) HasInboundCallPrices() bool`

HasInboundCallPrices returns a boolean if a field has been set.

### SetInboundCallPricesNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetInboundCallPricesNil(b bool)`

 SetInboundCallPricesNil sets the value for InboundCallPrices to be an explicit nil

### UnsetInboundCallPrices
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetInboundCallPrices()`

UnsetInboundCallPrices ensures that no value is present for InboundCallPrices, not even an explicit nil
### GetIsoCountry

`func (o *PricingV1VoiceVoiceCountryInstance) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV1VoiceVoiceCountryInstance) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV1VoiceVoiceCountryInstance) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetOutboundPrefixPrices

`func (o *PricingV1VoiceVoiceCountryInstance) GetOutboundPrefixPrices() []PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices`

GetOutboundPrefixPrices returns the OutboundPrefixPrices field if non-nil, zero value otherwise.

### GetOutboundPrefixPricesOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetOutboundPrefixPricesOk() (*[]PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices, bool)`

GetOutboundPrefixPricesOk returns a tuple with the OutboundPrefixPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPrefixPrices

`func (o *PricingV1VoiceVoiceCountryInstance) SetOutboundPrefixPrices(v []PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices)`

SetOutboundPrefixPrices sets OutboundPrefixPrices field to given value.

### HasOutboundPrefixPrices

`func (o *PricingV1VoiceVoiceCountryInstance) HasOutboundPrefixPrices() bool`

HasOutboundPrefixPrices returns a boolean if a field has been set.

### SetOutboundPrefixPricesNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetOutboundPrefixPricesNil(b bool)`

 SetOutboundPrefixPricesNil sets the value for OutboundPrefixPrices to be an explicit nil

### UnsetOutboundPrefixPrices
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetOutboundPrefixPrices()`

UnsetOutboundPrefixPrices ensures that no value is present for OutboundPrefixPrices, not even an explicit nil
### GetPriceUnit

`func (o *PricingV1VoiceVoiceCountryInstance) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV1VoiceVoiceCountryInstance) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV1VoiceVoiceCountryInstance) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV1VoiceVoiceCountryInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV1VoiceVoiceCountryInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV1VoiceVoiceCountryInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV1VoiceVoiceCountryInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV1VoiceVoiceCountryInstance) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV1VoiceVoiceCountryInstance) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


