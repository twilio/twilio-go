# \MetricsTagApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetric**](MetricsTagApi.md#ListMetric) | **Get** /v1/Voice/{CallSid}/Metrics | 



## ListMetric

> InsightsV1CallMetricReadResponse ListMetric(ctx, CallSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string**|  | 
 **optional** | ***ListMetricOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMetricOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Edge** | **optional.String**|  | 
 **Direction** | **optional.String**|  | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InsightsV1CallMetricReadResponse**](insights_v1_call_metricReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

