# InsightsInstancesAIReportsApi

All URIs are relative to *https://flex-api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchInsightsConversationalAi**](InsightsInstancesAIReportsApi.md#FetchInsightsConversationalAi) | **Get** /v1/Insights/Instances/{InstanceSid}/AI/Reports | 



## FetchInsightsConversationalAi

> FlexV1InsightsConversationalAi FetchInsightsConversationalAi(ctx, InstanceSidoptional)



Fetch Account Based Conversational AI Reports

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**InstanceSid** | **string** | Sid of Flex Service Instance

### Other Parameters

Other parameters are passed through a pointer to a FetchInsightsConversationalAiParams struct


Name | Type | Description
------------- | ------------- | -------------
**MaxRows** | **int** | Maximum number of rows to return
**ReportId** | **string** | The type of report required to fetch.Like gauge,channel-metrics,queue-metrics
**Granularity** | **string** | The time period for which report is needed
**IncludeDate** | **time.Time** | A reference date that should be included in the returned period

### Return type

[**FlexV1InsightsConversationalAi**](FlexV1InsightsConversationalAi.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

