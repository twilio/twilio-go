# FlexV1InsightsConversationalAi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSid** | Pointer to **string** | Sid of Flex Service Instance |
**ReportId** | Pointer to **string** | The type of report required to fetch.Like gauge,channel-metrics,queue-metrics |
**Granularity** | Pointer to [**string**](InsightsConversationalAiEnumGranularity.md) |  |
**PeriodStart** | Pointer to [**time.Time**](time.Time.md) | The start date from which report data is included |
**PeriodEnd** | Pointer to [**time.Time**](time.Time.md) | The end date till report data is included |
**Updated** | Pointer to [**time.Time**](time.Time.md) | Updated time of the report |
**TotalPages** | Pointer to **int** | Represents total number of pages fetched report has |
**Page** | Pointer to **int** | Page offset required for pagination |
**Rows** | Pointer to **[]interface{}** | List of report breakdown  |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


