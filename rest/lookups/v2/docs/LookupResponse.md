# LookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallingCountryCode** | Pointer to **string** | International dialing prefix of the phone number defined in the E.164 standard. |
**CountryCode** | Pointer to **string** | The phone number's [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**PhoneNumber** | Pointer to **string** | The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number. |
**NationalFormat** | Pointer to **string** | The phone number in [national format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers). |
**Valid** | **bool** | Boolean which indicates if the phone number is in a valid range that can be freely assigned by a carrier to a user. |[optional] 
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Contains reasons why a phone number is invalid. Possible values: TOO_SHORT, TOO_LONG, INVALID_BUT_POSSIBLE, INVALID_COUNTRY_CODE, INVALID_LENGTH, NOT_A_NUMBER. |[optional] 
**CallerName** | [**CallerNameInfo**](CallerNameInfo.md) |  |[optional] 
**SimSwap** | [**SimSwapInfo**](SimSwapInfo.md) |  |[optional] 
**CallForwarding** | [**CallForwardingInfo**](CallForwardingInfo.md) |  |[optional] 
**LineTypeIntelligence** | [**LineTypeIntelligenceInfo**](LineTypeIntelligenceInfo.md) |  |[optional] 
**LineStatus** | [**LineStatusInfo**](LineStatusInfo.md) |  |[optional] 
**IdentityMatch** | [**IdentityMatchInfo**](IdentityMatchInfo.md) |  |[optional] 
**ReassignedNumber** | [**ReassignedNumberInfo**](ReassignedNumberInfo.md) |  |[optional] 
**SmsPumpingRisk** | [**SmsPumpingRiskInfo**](SmsPumpingRiskInfo.md) |  |[optional] 
**PhoneNumberQualityScore** | Pointer to **interface{}** | An object that contains information of a mobile phone number quality score. Quality score will return a risk score about the phone number. |
**PreFill** | Pointer to **interface{}** | An object that contains pre fill information. pre_fill will return PII information associated with the phone number like first name, last name, address line, country code, state and postal code.  |
**Url** | **string** | The absolute URL of the resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


