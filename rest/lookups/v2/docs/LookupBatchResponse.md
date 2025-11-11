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
**CallerName** | [**CallerNameInfo**](CallerNameInfo.md) |  |[optional] 
**SimSwap** | [**SimSwapInfo**](SimSwapInfo.md) |  |[optional] 
**CallForwarding** | [**CallForwardingInfo**](CallForwardingInfo.md) |  |[optional] 
**LineTypeIntelligence** | [**LineTypeIntelligenceInfo**](LineTypeIntelligenceInfo.md) |  |[optional] 
**LineStatus** | [**LineStatusInfo**](LineStatusInfo.md) |  |[optional] 
**IdentityMatch** | [**IdentityMatchInfo**](IdentityMatchInfo.md) |  |[optional] 
**ReassignedNumber** | [**ReassignedNumberInfo**](ReassignedNumberInfo.md) |  |[optional] 
**SmsPumpingRisk** | [**SmsPumpingRiskInfo**](SmsPumpingRiskInfo.md) |  |[optional] 
**PhoneNumberQualityScore** | Pointer to **interface{}** |  |
**PreFill** | Pointer to **interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


