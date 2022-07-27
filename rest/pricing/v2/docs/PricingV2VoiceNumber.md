# PricingV2VoiceNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationNumber** | Pointer to **string** | The destination phone number, in E.164 format |
**OriginationNumber** | Pointer to **string** | The origination phone number, in E.164 format |
**Country** | Pointer to **string** | The name of the country |
**IsoCountry** | Pointer to **string** | The ISO country code |
**OutboundCallPrices** | Pointer to [**[]PricingV2VoiceVoiceNumberOutboundCallPrices**](PricingV2VoiceVoiceNumberOutboundCallPrices.md) | The list of OutboundCallPriceWithOrigin records |
**InboundCallPrice** | Pointer to [**PricingV2VoiceVoiceNumberInboundCallPrice**](PricingV2VoiceVoiceNumberInboundCallPrice.md) |  |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy) |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


