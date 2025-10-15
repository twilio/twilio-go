# LookupBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | **string** | Unique identifier used to match request with response |[optional] 
**TwilioErrorCode** | **int** | Twilio error code in case that the request to downstream fails |[optional] 
**CallingCountryCode** | **string** |  |[optional] 
**CountryCode** | **string** |  |[optional] 
**PhoneNumber** | **string** |  |[optional] 
**NationalFormat** | **string** |  |[optional] 
**Valid** | **bool** |  |[optional] 
**ValidationErrors** | **[]string** |  |[optional] 
**CallerName** | [**CallerName**](CallerName.md) |  |[optional] 
**SimSwap** | [**SimSwap**](SimSwap.md) |  |[optional] 
**CallForwarding** | [**CallForwarding**](CallForwarding.md) |  |[optional] 
**LineTypeIntelligence** | [**LineTypeIntelligence**](LineTypeIntelligence.md) |  |[optional] 
**LineStatus** | [**LineStatus**](LineStatus.md) |  |[optional] 
**IdentityMatch** | [**IdentityMatch**](IdentityMatch.md) |  |[optional] 
**ReassignedNumber** | [**ReassignedNumber**](ReassignedNumber.md) |  |[optional] 
**SmsPumpingRisk** | [**SmsPumpingRisk**](SmsPumpingRisk.md) |  |[optional] 
**PhoneNumberQualityScore** | Pointer to **interface{}** |  |
**PreFill** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


