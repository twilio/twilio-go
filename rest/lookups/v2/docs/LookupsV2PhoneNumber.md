# LookupsV2PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCountryCode** | Pointer to **string** | International dialing prefix |
**CountryCode** | Pointer to **string** | Phone number's ISO country code |
**PhoneNumber** | Pointer to **string** | Phone number in E.164 format |
**NationalFormat** | Pointer to **string** | Phone number in national format |
**Valid** | Pointer to **bool** | Boolean which indicates if the phone number is valid |
**ValidationErrors** | Pointer to [**[]string**](PhoneNumberEnumValidationError.md) | Contains reasons why a phone number is invalid |
**CallerName** | Pointer to **interface{}** | An object that contains caller name information |
**SimSwap** | Pointer to **interface{}** | An object that contains SIM swap information |
**CallForwarding** | Pointer to **interface{}** | An object that contains call forwarding status information |
**LiveActivity** | Pointer to **interface{}** | An object that contains live activity information |
**LineTypeIntelligence** | Pointer to **interface{}** | An object that contains line type information |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


