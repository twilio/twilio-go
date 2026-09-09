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
**CallerName** | Pointer to [**CallerNameInfo**](CallerNameInfo.md) |  |
**SimSwap** | Pointer to [**SimSwapInfo**](SimSwapInfo.md) |  |
**CallForwarding** | Pointer to [**CallForwardingInfo**](CallForwardingInfo.md) |  |
**LineTypeIntelligence** | Pointer to [**LineTypeIntelligenceInfo**](LineTypeIntelligenceInfo.md) |  |
**LineStatus** | Pointer to [**LineStatusInfo**](LineStatusInfo.md) |  |
**IdentityMatch** | Pointer to [**IdentityMatchInfo**](IdentityMatchInfo.md) |  |
**ReassignedNumber** | Pointer to [**ReassignedNumberInfo**](ReassignedNumberInfo.md) |  |
**SmsPumpingRisk** | Pointer to [**SmsPumpingRiskInfo**](SmsPumpingRiskInfo.md) |  |
**PhoneNumberQualityScore** | Pointer to **interface{}** |  |
**PreFill** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


