# PricingV2VoiceNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationNumber** | Pointer to **string** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**OriginationNumber** | Pointer to **string** | The origination phone number in [[E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**Country** | Pointer to **string** | The name of the country. |
**IsoCountry** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) |
**OutboundCallPrices** | Pointer to [**[]PricingV2VoiceVoiceNumberOutboundCallPrices**](PricingV2VoiceVoiceNumberOutboundCallPrices.md) | The list of [OutboundCallPriceWithOrigin](https://www.twilio.com/docs/voice/pricing#outbound-call-price-with-origin) records. |
**InboundCallPrice** | Pointer to [**PricingV2VoiceVoiceNumberInboundCallPrice**](PricingV2VoiceVoiceNumberInboundCallPrice.md) |  |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, specified in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


