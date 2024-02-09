# FlexV1InsightsConversationalAiReportInsights

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSid** | Pointer to **string** | The Instance SID of the instance for which report insights is fetched |
**ReportId** | Pointer to **string** | The type of report insights required to fetch.Like gauge,channel-metrics,queue-metrics |
**PeriodStart** | Pointer to [**time.Time**](time.Time.md) | The start date from which report insights data is included |
**PeriodEnd** | Pointer to [**time.Time**](time.Time.md) | The end date till report insights data is included |
**Updated** | Pointer to [**time.Time**](time.Time.md) | Updated time of the report insights |
**Insights** | Pointer to **[]interface{}** | List of report insights breakdown |
**Url** | Pointer to **string** | The URL of this resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


