# LookupsV1PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerName** | Pointer to **interface{}** | The name of the phone number's owner. If `null`, that information was not available. |
**CountryCode** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) for the phone number. |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**NationalFormat** | Pointer to **string** | The phone number, in national format. |
**Carrier** | Pointer to **interface{}** | The telecom company that provides the phone number. |
**AddOns** | Pointer to **interface{}** | A JSON string with the results of the Add-ons you specified in the `add_ons` parameters. For the format of the object, see [Using Add-ons](https://www.twilio.com/docs/add-ons). |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


