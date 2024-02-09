# InsightsInstancesAIReportInsightsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInsightsConversationalAiReportInsights**](InsightsInstancesAIReportInsightsApi.md#FetchInsightsConversationalAiReportInsights) | **Get** /v1/Insights/Instances/{InstanceSid}/AI/ReportInsights | 



## FetchInsightsConversationalAiReportInsights

> FlexV1InsightsConversationalAiReportInsights FetchInsightsConversationalAiReportInsights(ctx, InstanceSidoptional)



Fetch Instance Based Conversational AI Report Insights

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstanceSid** | **string** | The Instance SID of the instance for which report insights will be fetched

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsConversationalAiReportInsightsParams struct


Name | Type | Description
------------- | ------------- | -------------
**MaxRows** | **int** | Maximum number of rows to return
**ReportId** | **string** | The type of report insights required to fetch.Like gauge,channel-metrics,queue-metrics
**Granularity** | **string** | The time period for which report insights is needed
**IncludeDate** | **time.Time** | A reference date that should be included in the returned period

### Return type

[**FlexV1InsightsConversationalAiReportInsights**](FlexV1InsightsConversationalAiReportInsights.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

