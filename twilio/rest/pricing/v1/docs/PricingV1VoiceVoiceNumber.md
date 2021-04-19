# PricingV1VoiceVoiceNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** | The name of the country | [optional] 
**InboundCallPrice** | Pointer to [**NullablePricingV1VoiceVoiceNumberInboundCallPrice**](pricing_v1_voice_voice_number_inbound_call_price.md) |  | [optional] 
**IsoCountry** | Pointer to **NullableString** | The ISO country code | [optional] 
**Number** | Pointer to **NullableString** | The phone number | [optional] 
**OutboundCallPrice** | Pointer to [**NullablePricingV1VoiceVoiceNumberOutboundCallPrice**](pricing_v1_voice_voice_number_outbound_call_price.md) |  | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the resource | [optional] 

## Methods

### NewPricingV1VoiceVoiceNumber

`func NewPricingV1VoiceVoiceNumber() *PricingV1VoiceVoiceNumber`

NewPricingV1VoiceVoiceNumber instantiates a new PricingV1VoiceVoiceNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingV1VoiceVoiceNumberWithDefaults

`func NewPricingV1VoiceVoiceNumberWithDefaults() *PricingV1VoiceVoiceNumber`

NewPricingV1VoiceVoiceNumberWithDefaults instantiates a new PricingV1VoiceVoiceNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PricingV1VoiceVoiceNumber) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PricingV1VoiceVoiceNumber) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PricingV1VoiceVoiceNumber) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PricingV1VoiceVoiceNumber) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *PricingV1VoiceVoiceNumber) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *PricingV1VoiceVoiceNumber) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetInboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) GetInboundCallPrice() PricingV1VoiceVoiceNumberInboundCallPrice`

GetInboundCallPrice returns the InboundCallPrice field if non-nil, zero value otherwise.

### GetInboundCallPriceOk

`func (o *PricingV1VoiceVoiceNumber) GetInboundCallPriceOk() (*PricingV1VoiceVoiceNumberInboundCallPrice, bool)`

GetInboundCallPriceOk returns a tuple with the InboundCallPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) SetInboundCallPrice(v PricingV1VoiceVoiceNumberInboundCallPrice)`

SetInboundCallPrice sets InboundCallPrice field to given value.

### HasInboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) HasInboundCallPrice() bool`

HasInboundCallPrice returns a boolean if a field has been set.

### SetInboundCallPriceNil

`func (o *PricingV1VoiceVoiceNumber) SetInboundCallPriceNil(b bool)`

 SetInboundCallPriceNil sets the value for InboundCallPrice to be an explicit nil

### UnsetInboundCallPrice
`func (o *PricingV1VoiceVoiceNumber) UnsetInboundCallPrice()`

UnsetInboundCallPrice ensures that no value is present for InboundCallPrice, not even an explicit nil
### GetIsoCountry

`func (o *PricingV1VoiceVoiceNumber) GetIsoCountry() string`

GetIsoCountry returns the IsoCountry field if non-nil, zero value otherwise.

### GetIsoCountryOk

`func (o *PricingV1VoiceVoiceNumber) GetIsoCountryOk() (*string, bool)`

GetIsoCountryOk returns a tuple with the IsoCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCountry

`func (o *PricingV1VoiceVoiceNumber) SetIsoCountry(v string)`

SetIsoCountry sets IsoCountry field to given value.

### HasIsoCountry

`func (o *PricingV1VoiceVoiceNumber) HasIsoCountry() bool`

HasIsoCountry returns a boolean if a field has been set.

### SetIsoCountryNil

`func (o *PricingV1VoiceVoiceNumber) SetIsoCountryNil(b bool)`

 SetIsoCountryNil sets the value for IsoCountry to be an explicit nil

### UnsetIsoCountry
`func (o *PricingV1VoiceVoiceNumber) UnsetIsoCountry()`

UnsetIsoCountry ensures that no value is present for IsoCountry, not even an explicit nil
### GetNumber

`func (o *PricingV1VoiceVoiceNumber) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PricingV1VoiceVoiceNumber) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PricingV1VoiceVoiceNumber) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PricingV1VoiceVoiceNumber) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *PricingV1VoiceVoiceNumber) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *PricingV1VoiceVoiceNumber) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetOutboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) GetOutboundCallPrice() PricingV1VoiceVoiceNumberOutboundCallPrice`

GetOutboundCallPrice returns the OutboundCallPrice field if non-nil, zero value otherwise.

### GetOutboundCallPriceOk

`func (o *PricingV1VoiceVoiceNumber) GetOutboundCallPriceOk() (*PricingV1VoiceVoiceNumberOutboundCallPrice, bool)`

GetOutboundCallPriceOk returns a tuple with the OutboundCallPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) SetOutboundCallPrice(v PricingV1VoiceVoiceNumberOutboundCallPrice)`

SetOutboundCallPrice sets OutboundCallPrice field to given value.

### HasOutboundCallPrice

`func (o *PricingV1VoiceVoiceNumber) HasOutboundCallPrice() bool`

HasOutboundCallPrice returns a boolean if a field has been set.

### SetOutboundCallPriceNil

`func (o *PricingV1VoiceVoiceNumber) SetOutboundCallPriceNil(b bool)`

 SetOutboundCallPriceNil sets the value for OutboundCallPrice to be an explicit nil

### UnsetOutboundCallPrice
`func (o *PricingV1VoiceVoiceNumber) UnsetOutboundCallPrice()`

UnsetOutboundCallPrice ensures that no value is present for OutboundCallPrice, not even an explicit nil
### GetPriceUnit

`func (o *PricingV1VoiceVoiceNumber) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *PricingV1VoiceVoiceNumber) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *PricingV1VoiceVoiceNumber) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *PricingV1VoiceVoiceNumber) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *PricingV1VoiceVoiceNumber) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *PricingV1VoiceVoiceNumber) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetUrl

`func (o *PricingV1VoiceVoiceNumber) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PricingV1VoiceVoiceNumber) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PricingV1VoiceVoiceNumber) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PricingV1VoiceVoiceNumber) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PricingV1VoiceVoiceNumber) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PricingV1VoiceVoiceNumber) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


