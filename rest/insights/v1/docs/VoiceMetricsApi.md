# VoiceMetricsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetric**](VoiceMetricsApi.md#ListMetric) | **Get** /v1/Voice/{CallSid}/Metrics | 



## ListMetric

> []InsightsV1Metric ListMetric(ctx, CallSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMetricParams struct


Name | Type | Description
------------- | ------------- | -------------
**Edge** | **string** | 
**Direction** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1Metric**](InsightsV1Metric.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

