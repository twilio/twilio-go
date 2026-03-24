# AccountReportKYTOutboundCarrierCalling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UniqueCallingNumbers** | **int** | Number of unique PSTN calling numbers to non-Twilio numbers during the report period. |[optional] 
**UniqueCalledNumbers** | **int** | Number of unique non-Twilio PSTN called numbers during the report period. |[optional] 
**BlockedCallsByCarrier** | [**[]CountyCarrierValue**](CountyCarrierValue.md) | Percentage of blocked calls by carrier per country. |[optional] 
**ShortDurationCallsPercentage** | **float32** | Percentage of completed outbound calls under 10 seconds (PSTN Short call tags); More than 15% is typically low trust measured. |[optional] 
**LongDurationCallsPercentage** | **float32** | Percentage of long duration calls ( >= 60 seconds) |[optional] 
**PotentialRobocallsPercentage** | **float32** | Percentage of completed outbound calls to unassigned or unallocated phone numbers. |[optional] 
**BrandedCalling** | [**BrandedCalling**](BrandedCalling.md) |  |[optional] 
**VoiceIntegrity** | [**VoiceIntegrity**](VoiceIntegrity.md) |  |[optional] 
**StirShaken** | [**StirShaken**](StirShaken.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


