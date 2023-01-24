# ApiV2010AvailablePhoneNumberCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country. |
**Country** | Pointer to **string** | The name of the country. |
**Uri** | Pointer to **string** | The URI of the Country resource, relative to `https://api.twilio.com`. |
**Beta** | Pointer to **bool** | Whether all phone numbers available in the country are new to the Twilio platform. `true` if they are and `false` if all numbers are not in the Twilio Phone Number Beta program. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related AvailablePhoneNumber resources identified by their URIs relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


