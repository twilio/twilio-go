# PricingV2TrunkingCountryInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The name of the country. |
**IsoCountry** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**TerminatingPrefixPrices** | Pointer to [**[]PricingV2TrunkingCountryInstanceTerminatingPrefixPrices**](PricingV2TrunkingCountryInstanceTerminatingPrefixPrices.md) | The list of [TerminatingPrefixPrice](https://www.twilio.com/docs/voice/pricing#outbound-prefix-price-with-origin) records. |
**OriginatingCallPrices** | Pointer to [**[]PricingV2TrunkingCountryInstanceOriginatingCallPrices**](PricingV2TrunkingCountryInstanceOriginatingCallPrices.md) | The list of [OriginatingCallPrice](https://www.twilio.com/docs/voice/pricing#inbound-call-price) records. |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, specified in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


