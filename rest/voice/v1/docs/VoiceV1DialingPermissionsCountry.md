# VoiceV1DialingPermissionsCountry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCode** | Pointer to **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). |
**Name** | Pointer to **string** | The name of the country. |
**Continent** | Pointer to **string** | The name of the continent in which the country is located. |
**CountryCodes** | Pointer to **[]string** | The E.164 assigned [country codes(s)](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html) |
**LowRiskNumbersEnabled** | Pointer to **bool** | Whether dialing to low-risk numbers is enabled. |
**HighRiskSpecialNumbersEnabled** | Pointer to **bool** | Whether dialing to high-risk special services numbers is enabled. These prefixes include number ranges allocated by the country and include premium numbers, special services, shared cost, and others |
**HighRiskTollfraudNumbersEnabled** | Pointer to **bool** | Whether dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers is enabled. These prefixes include narrow number ranges that have a high-risk of international revenue sharing fraud (IRSF) attacks, also known as [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud). These prefixes are collected from anti-fraud databases and verified by analyzing calls on our network. These prefixes are not available for download and are updated frequently |
**Url** | Pointer to **string** | The absolute URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | A list of URLs related to this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


