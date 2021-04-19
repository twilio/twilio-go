# PricingV2VoiceVoiceNumber

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**DestinationNumber** | Pointer to **NullableString** | The destination phone number, in E.164 format | [optional] 
**InboundCallPrice** | Pointer to [**NullablePricingV2VoiceVoiceNumberInboundCallPrice**](pricing_v2_voice_voice_number_inbound_call_price.md) |  | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code | [optional] 
**OriginationNumber** | Pointer to **NullableString** | The origination phone number, in E.164 format | [optional] 
**OutboundCallPrices** | Pointer to [**[]PricingV2VoiceVoiceNumberOutboundCallPrices**](PricingV2VoiceVoiceNumberOutboundCallPrices.md) | The list of OutboundCallPriceWithOrigin records | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV2VoiceVoiceNumber

`func NewPricingV2VoiceVoiceNumber() *PricingV2VoiceVoiceNumber`

NewPricingV2VoiceVoiceNumber instantiates a new PricingV2VoiceVoiceNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV2VoiceVoiceNumberWithDefaults

`func NewPricingV2VoiceVoiceNumberWithDefaults() *PricingV2VoiceVoiceNumber`

NewPricingV2VoiceVoiceNumberWithDefaults instantiates a new PricingV2VoiceVoiceNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV2VoiceVoiceNumber) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV2VoiceVoiceNumber) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV2VoiceVoiceNumber) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV2VoiceVoiceNumber) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV2VoiceVoiceNumber) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV2VoiceVoiceNumber) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetDestinationNumber

`func (o *PricingV2VoiceVoiceNumber) GetDestinationNumber() string`

GetDestinationNumber returns the DestinationNumber field if non-nil, zero value otherwise.

### GetDestinationNumberOk

`func (o *PricingV2VoiceVoiceNumber) GetDestinationNumberOk() (*string, bool)`

GetDestinationNumberOk returns a tuple with the DestinationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNumber

`func (o *PricingV2VoiceVoiceNumber) SetDestinationNumber(v string)`

SetDestinationNumber sets DestinationNumber field to given value.

### HasDestinationNumber

`func (o *PricingV2VoiceVoiceNumber) HasDestinationNumber() bool`

HasDestinationNumber returns a boolean if a field has been set.

### SetDestinationNumberNil

`func (o *PricingV2VoiceVoiceNumber) SetDestinationNumberNil(b bool)`

 SetDestinationNumberNil sets the value for DestinationNumber to be an explicit nil

### UnsetDestinationNumber
`func (o *PricingV2VoiceVoiceNumber) UnsetDestinationNumber()`

UnsetDestinationNumber ensures that no value is present for DestinationNumber, not even an explicit nil
### GetInboundCallPrice

`func (o *PricingV2VoiceVoiceNumber) GetInboundCallPrice() PricingV2VoiceVoiceNumberInboundCallPrice`

GetInboundCallPrice returns the InboundCallPrice field if non-nil, zero value otherwise.

### GetInboundCallPriceOk

`func (o *PricingV2VoiceVoiceNumber) GetInboundCallPriceOk() (*PricingV2VoiceVoiceNumberInboundCallPrice, bool)`

GetInboundCallPriceOk returns a tuple with the InboundCallPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallPrice

`func (o *PricingV2VoiceVoiceNumber) SetInboundCallPrice(v PricingV2VoiceVoiceNumberInboundCallPrice)`

SetInboundCallPrice sets InboundCallPrice field to given value.

### HasInboundCallPrice

`func (o *PricingV2VoiceVoiceNumber) HasInboundCallPrice() bool`

HasInboundCallPrice returns a boolean if a field has been set.

### SetInboundCallPriceNil

`func (o *PricingV2VoiceVoiceNumber) SetInboundCallPriceNil(b bool)`

 SetInboundCallPriceNil sets the value for InboundCallPrice to be an explicit nil

### UnsetInboundCallPrice
`func (o *PricingV2VoiceVoiceNumber) UnsetInboundCallPrice()`

UnsetInboundCallPrice ensures that no value is present for InboundCallPrice, not even an explicit nil
### GetIsoCountry

`func (o *PricingV2VoiceVoiceNumber) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV2VoiceVoiceNumber) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV2VoiceVoiceNumber) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV2VoiceVoiceNumber) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV2VoiceVoiceNumber) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV2VoiceVoiceNumber) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetOriginationNumber

`func (o *PricingV2VoiceVoiceNumber) GetOriginationNumber() string`

GetOriginationNumber returns the OriginationNumber field if non-nil, zero value otherwise.

### GetOriginationNumberOk

`func (o *PricingV2VoiceVoiceNumber) GetOriginationNumberOk() (*string, bool)`

GetOriginationNumberOk returns a tuple with the OriginationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationNumber

`func (o *PricingV2VoiceVoiceNumber) SetOriginationNumber(v string)`

SetOriginationNumber sets OriginationNumber field to given value.

### HasOriginationNumber

`func (o *PricingV2VoiceVoiceNumber) HasOriginationNumber() bool`

HasOriginationNumber returns a boolean if a field has been set.

### SetOriginationNumberNil

`func (o *PricingV2VoiceVoiceNumber) SetOriginationNumberNil(b bool)`

 SetOriginationNumberNil sets the value for OriginationNumber to be an explicit nil

### UnsetOriginationNumber
`func (o *PricingV2VoiceVoiceNumber) UnsetOriginationNumber()`

UnsetOriginationNumber ensures that no value is present for OriginationNumber, not even an explicit nil
### GetOutboundCallPrices

`func (o *PricingV2VoiceVoiceNumber) GetOutboundCallPrices() []PricingV2VoiceVoiceNumberOutboundCallPrices`

GetOutboundCallPrices returns the OutboundCallPrices field if non-nil, zero value otherwise.

### GetOutboundCallPricesOk

`func (o *PricingV2VoiceVoiceNumber) GetOutboundCallPricesOk() (*[]PricingV2VoiceVoiceNumberOutboundCallPrices, bool)`

GetOutboundCallPricesOk returns a tuple with the OutboundCallPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundCallPrices

`func (o *PricingV2VoiceVoiceNumber) SetOutboundCallPrices(v []PricingV2VoiceVoiceNumberOutboundCallPrices)`

SetOutboundCallPrices sets OutboundCallPrices field to given value.

### HasOutboundCallPrices

`func (o *PricingV2VoiceVoiceNumber) HasOutboundCallPrices() bool`

HasOutboundCallPrices returns a boolean if a field has been set.

### SetOutboundCallPricesNil

`func (o *PricingV2VoiceVoiceNumber) SetOutboundCallPricesNil(b bool)`

 SetOutboundCallPricesNil sets the value for OutboundCallPrices to be an explicit nil

### UnsetOutboundCallPrices
`func (o *PricingV2VoiceVoiceNumber) UnsetOutboundCallPrices()`

UnsetOutboundCallPrices ensures that no value is present for OutboundCallPrices, not even an explicit nil
### GetPriceUnit

`func (o *PricingV2VoiceVoiceNumber) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV2VoiceVoiceNumber) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV2VoiceVoiceNumber) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV2VoiceVoiceNumber) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV2VoiceVoiceNumber) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV2VoiceVoiceNumber) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV2VoiceVoiceNumber) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV2VoiceVoiceNumber) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV2VoiceVoiceNumber) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV2VoiceVoiceNumber) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV2VoiceVoiceNumber) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV2VoiceVoiceNumber) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


