# InsightsV2InboundPhoneNumberReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | **string** | Inbound phone number handle represented in the report. |[optional] 
**TotalCalls** | **int** | Total number of calls made with the given handle during the report period. |[optional] 
**CallAnswerScore** | **float32** | The call answer score measures customers behavior to the delivered calls. The score is a value between 0 and 100, where 100 indicates that all calls were successfully answered.  |[optional] 
**CallStatePercentage** | [**InsightsV2InboundPhoneNumberReportCallStatePercentage**](InsightsV2InboundPhoneNumberReportCallStatePercentage.md) |  |[optional] 
**SilentCallsPercentage** | **float32** | Percentage of inbound calls with silence tags over total outbound calls. A silent tag is indicative of a connectivity issue or muted audio. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


