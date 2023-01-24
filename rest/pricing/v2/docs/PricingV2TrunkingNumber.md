# PricingV2TrunkingNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationNumber** | Pointer to **string** | The destination phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**OriginationNumber** | Pointer to **string** | The origination phone number in [[E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**Country** | Pointer to **string** | The name of the country. |
**IsoCountry** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) |
**TerminatingPrefixPrices** | Pointer to [**[]PricingV2TrunkingCountryInstanceTerminatingPrefixPrices**](PricingV2TrunkingCountryInstanceTerminatingPrefixPrices.md) |  |
**OriginatingCallPrice** | Pointer to [**PricingV2TrunkingNumberOriginatingCallPrice**](PricingV2TrunkingNumberOriginatingCallPrice.md) |  |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, specified in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


