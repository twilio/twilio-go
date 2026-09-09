# AccountReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallDeliverabilityScore** | **float32** | The call deliverability score measures the network effectiveness in delivering calls by scoring calls reach the intended recipient. The score is a value between 0 and 100, where 100 indicates that all calls were successfully delivered.  |[optional] 
**CallAnswerScore** | **float32** | The call answer score measures customers behavior to the delivered calls. The score is a value between 0 and 100, where 100 indicates that all calls were successfully answered.  |[optional] 
**TotalCalls** | **int** | Total number of calls made during the report period. |[optional] 
**CallDirection** | Pointer to [**AccountReportCallDirection**](AccountReportCallDirection.md) |  |
**CallState** | Pointer to [**AccountReportCallState**](AccountReportCallState.md) |  |
**CallType** | Pointer to [**AccountReportCallType**](AccountReportCallType.md) |  |
**Aloc** | **float32** | Average length of call in seconds. |[optional] 
**TwilioEdgeLocation** | **map[string]int** | Number of calls made in each Twilio Edge location. Refer to [Public Edge Locations](https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations) for more detail. |[optional] 
**CallerCountryCode** | **map[string]int** | Number of calls originating from each country (ISO alpha-2). |[optional] 
**CalleeCountryCode** | **map[string]int** | Number of calls terminating in each country (ISO alpha-2). |[optional] 
**AverageQueueTimeMs** | **float32** | Average queue time in milliseconds. |[optional] 
**SilentCallsPercentage** | **float32** | Percentage of silent calls. |[optional] 
**NetworkIssues** | Pointer to [**AccountReportNetworkIssues**](AccountReportNetworkIssues.md) |  |
**KYT** | Pointer to [**AccountReportKYT**](AccountReportKYT.md) |  |
**AnsweringMachineDetection** | Pointer to [**AccountReportAnsweringMachineDetection**](AccountReportAnsweringMachineDetection.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


