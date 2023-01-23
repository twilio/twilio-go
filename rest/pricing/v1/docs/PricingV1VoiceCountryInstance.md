# PricingV1VoiceCountryInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The name of the country. |
**IsoCountry** | Pointer to **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**OutboundPrefixPrices** | Pointer to [**[]PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices**](PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices.md) | The list of OutboundPrefixPrice records, which include a list of the `prefixes`, the `friendly_name`, `base_price`, and the   `current_price` for those prefixes. |
**InboundCallPrices** | Pointer to [**[]PricingV1VoiceVoiceCountryInstanceInboundCallPrices**](PricingV1VoiceVoiceCountryInstanceInboundCallPrices.md) | The list of [InboundCallPrice](https://www.twilio.com/docs/voice/pricing#inbound-call-price) records. |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, specified in [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


