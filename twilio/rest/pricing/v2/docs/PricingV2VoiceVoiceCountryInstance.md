# PricingV2VoiceVoiceCountryInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**InboundCallPrices** | Pointer to [**[]PricingV2VoiceVoiceCountryInstanceInboundCallPrices**](PricingV2VoiceVoiceCountryInstanceInboundCallPrices.md) | The list of InboundCallPrice records | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code | [optional] 
**OutboundPrefixPrices** | Pointer to [**[]PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices**](PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices.md) | The list of OutboundPrefixPriceWithOrigin records | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV2VoiceVoiceCountryInstance

`func NewPricingV2VoiceVoiceCountryInstance() *PricingV2VoiceVoiceCountryInstance`

NewPricingV2VoiceVoiceCountryInstance instantiates a new PricingV2VoiceVoiceCountryInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV2VoiceVoiceCountryInstanceWithDefaults

`func NewPricingV2VoiceVoiceCountryInstanceWithDefaults() *PricingV2VoiceVoiceCountryInstance`

NewPricingV2VoiceVoiceCountryInstanceWithDefaults instantiates a new PricingV2VoiceVoiceCountryInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV2VoiceVoiceCountryInstance) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV2VoiceVoiceCountryInstance) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV2VoiceVoiceCountryInstance) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetInboundCallPrices

`func (o *PricingV2VoiceVoiceCountryInstance) GetInboundCallPrices() []PricingV2VoiceVoiceCountryInstanceInboundCallPrices`

GetInboundCallPrices returns the InboundCallPrices field if non-nil, zero value otherwise.

### GetInboundCallPricesOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetInboundCallPricesOk() (*[]PricingV2VoiceVoiceCountryInstanceInboundCallPrices, bool)`

GetInboundCallPricesOk returns a tuple with the InboundCallPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallPrices

`func (o *PricingV2VoiceVoiceCountryInstance) SetInboundCallPrices(v []PricingV2VoiceVoiceCountryInstanceInboundCallPrices)`

SetInboundCallPrices sets InboundCallPrices field to given value.

### HasInboundCallPrices

`func (o *PricingV2VoiceVoiceCountryInstance) HasInboundCallPrices() bool`

HasInboundCallPrices returns a boolean if a field has been set.

### SetInboundCallPricesNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetInboundCallPricesNil(b bool)`

 SetInboundCallPricesNil sets the value for InboundCallPrices to be an explicit nil

### UnsetInboundCallPrices
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetInboundCallPrices()`

UnsetInboundCallPrices ensures that no value is present for InboundCallPrices, not even an explicit nil
### GetIsoCountry

`func (o *PricingV2VoiceVoiceCountryInstance) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV2VoiceVoiceCountryInstance) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV2VoiceVoiceCountryInstance) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetOutboundPrefixPrices

`func (o *PricingV2VoiceVoiceCountryInstance) GetOutboundPrefixPrices() []PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices`

GetOutboundPrefixPrices returns the OutboundPrefixPrices field if non-nil, zero value otherwise.

### GetOutboundPrefixPricesOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetOutboundPrefixPricesOk() (*[]PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices, bool)`

GetOutboundPrefixPricesOk returns a tuple with the OutboundPrefixPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundPrefixPrices

`func (o *PricingV2VoiceVoiceCountryInstance) SetOutboundPrefixPrices(v []PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices)`

SetOutboundPrefixPrices sets OutboundPrefixPrices field to given value.

### HasOutboundPrefixPrices

`func (o *PricingV2VoiceVoiceCountryInstance) HasOutboundPrefixPrices() bool`

HasOutboundPrefixPrices returns a boolean if a field has been set.

### SetOutboundPrefixPricesNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetOutboundPrefixPricesNil(b bool)`

 SetOutboundPrefixPricesNil sets the value for OutboundPrefixPrices to be an explicit nil

### UnsetOutboundPrefixPrices
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetOutboundPrefixPrices()`

UnsetOutboundPrefixPrices ensures that no value is present for OutboundPrefixPrices, not even an explicit nil
### GetPriceUnit

`func (o *PricingV2VoiceVoiceCountryInstance) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV2VoiceVoiceCountryInstance) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV2VoiceVoiceCountryInstance) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV2VoiceVoiceCountryInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV2VoiceVoiceCountryInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV2VoiceVoiceCountryInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV2VoiceVoiceCountryInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV2VoiceVoiceCountryInstance) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV2VoiceVoiceCountryInstance) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


