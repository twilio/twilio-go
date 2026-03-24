# InsightsV2OutboundPhoneNumberReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** | The outbound phone number handle. |[optional] 
**TotalCalls** | **int** | Total number of outbound calls made with the given handle during the report period. |[optional] 
**CallAnswerScore** | **float32** | The call answer score measures customers behavior to the delivered calls. The score is a value between 0 and 100, where 100 indicates that all calls were successfully answered.  |[optional] 
**CallsByDeviceType** | **map[string]int** | Number of calls made with each device type. `voip`, `mobile`, `landline`, `unknown`  |[optional] 
**AnswerRateDeviceType** | **map[string]float32** | Answer rate for each device type. `voip`, `mobile`, `landline`, `unknown`  |[optional] 
**CallStatePercentage** | [**InsightsV2OutboundPhoneNumberReportCallStatePercentage**](InsightsV2OutboundPhoneNumberReportCallStatePercentage.md) |  |[optional] 
**BlockedCallsByCarrier** | [**[]CountyCarrierValue**](CountyCarrierValue.md) | Percentage of blocked calls by carrier per country. |[optional] 
**SilentCallsPercentage** | **float32** | Percentage of calls with silence tags over total calls. A silent tag is indicative of a connectivity issue or muted audio. |[optional] 
**ShortDurationCallsPercentage** | **float32** | Percentage of completed outbound calls under 10 seconds (PSTN Short call tags); More than 15% is typically low trust measured. |[optional] 
**LongDurationCallsPercentage** | **float32** | Percentage of long duration calls ( >= 60 seconds) |[optional] 
**PotentialRobocallsPercentage** | **float32** | Percentage of completed outbound calls to unassigned or unallocated phone numbers. |[optional] 
**AnsweringMachineDetection** | [**InsightsV2OutboundPhoneNumberReportAnsweringMachineDetection**](InsightsV2OutboundPhoneNumberReportAnsweringMachineDetection.md) |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


