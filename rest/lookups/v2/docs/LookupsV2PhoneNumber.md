# LookupsV2PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCountryCode** | Pointer to **string** | International dialing prefix of the phone number defined in the E.164 standard. |
**CountryCode** | Pointer to **string** | The phone number's [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**NationalFormat** | Pointer to **string** | The phone number in [national format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers). |
**Valid** | Pointer to **bool** | Boolean which indicates if the phone number is in a valid range that can be freely assigned by a carrier to a user. |
**ValidationErrors** | Pointer to [**[]string**](PhoneNumberEnumValidationError.md) | Contains reasons why a phone number is invalid. Possible values: TOO_SHORT, TOO_LONG, INVALID_BUT_POSSIBLE, INVALID_COUNTRY_CODE, INVALID_LENGTH, NOT_A_NUMBER. |
**CallerName** | Pointer to **interface{}** | An object that contains caller name information based on [CNAM](https://support.twilio.com/hc/en-us/articles/360051670533-Getting-Started-with-CNAM-Caller-ID). |
**SimSwap** | Pointer to **interface{}** | An object that contains information on the last date the subscriber identity module (SIM) was changed for a mobile phone number. |
**CallForwarding** | Pointer to **interface{}** | An object that contains information on the unconditional call forwarding status of mobile phone number. |
**LiveActivity** | Pointer to **interface{}** | An object that contains live activity information for a mobile phone number. |
**LineTypeIntelligence** | Pointer to **interface{}** | An object that contains line type information including the carrier name, mobile country code, and mobile network code. |
**IdentityMatch** | Pointer to **interface{}** | An object that contains identity match information. The result of comparing user-provided information including name, address, date of birth, national ID, against authoritative phone-based data sources |
**SmsPumpingRisk** | Pointer to **interface{}** | An object that contains information on if a phone number has been currently or previously blocked by Verify Fraud Guard for receiving malicious SMS pumping traffic as well as other signals associated with risky carriers and low conversion rates. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


