# PricingV1MessagingCountryInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The name of the country. |
**IsoCountry** | Pointer to **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**OutboundSmsPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceOutboundSmsPrices.md) | The list of [OutboundSMSPrice](https://www.twilio.com/docs/sms/api/pricing#outbound-sms-price) records that represent the price to send a message for each MCC/MNC applicable in this country. |
**InboundSmsPrices** | Pointer to [**[]PricingV1MessagingMessagingCountryInstanceInboundSmsPrices**](PricingV1MessagingMessagingCountryInstanceInboundSmsPrices.md) | The list of [InboundPrice](https://www.twilio.com/docs/sms/api/pricing#inbound-price) records that describe the price to receive an inbound SMS to the different Twilio phone number types supported in this country |
**PriceUnit** | Pointer to **string** | The currency in which prices are measured, specified in [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


