# ApiV2010AvailablePhoneNumberTollFree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** | A formatted version of the phone number. |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**Lata** | Pointer to **string** | The [LATA](https://en.wikipedia.org/wiki/Local_access_and_transport_area) of this phone number. Available for only phone numbers from the US and Canada. |
**Locality** | Pointer to **string** | The locality or city of this phone number's location. |
**RateCenter** | Pointer to **string** | The [rate center](https://en.wikipedia.org/wiki/Telephone_exchange) of this phone number. Available for only phone numbers from the US and Canada. |
**Latitude** | Pointer to **float32** | The latitude of this phone number's location. Available for only phone numbers from the US and Canada. |
**Longitude** | Pointer to **float32** | The longitude of this phone number's location. Available for only phone numbers from the US and Canada. |
**Region** | Pointer to **string** | The two-letter state or province abbreviation of this phone number's location. Available for only phone numbers from the US and Canada. |
**PostalCode** | Pointer to **string** | The postal or ZIP code of this phone number's location. Available for only phone numbers from the US and Canada. |
**IsoCountry** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of this phone number. |
**AddressRequirements** | Pointer to **string** | The type of [Address](https://www.twilio.com/docs/usage/api/address) resource the phone number requires. Can be: `none`, `any`, `local`, or `foreign`. `none` means no address is required. `any` means an address is required, but it can be anywhere in the world. `local` means an address in the phone number's country is required. `foreign` means an address outside of the phone number's country is required. |
**Beta** | Pointer to **bool** | Whether the phone number is new to the Twilio platform. Can be: `true` or `false`. |
**Capabilities** | Pointer to [**ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities**](ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


