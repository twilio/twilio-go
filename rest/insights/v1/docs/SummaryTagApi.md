# \SummaryTagApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSummary**](SummaryTagApi.md#FetchSummary) | **Get** /v1/Voice/{CallSid}/Summary | 



## FetchSummary

> InsightsV1CallSummary FetchSummary(ctx, CallSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string**|  | 
 **optional** | ***FetchSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSummaryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ProcessingState** | **optional.String**|  | 

### Return type

[**InsightsV1CallSummary**](insights.v1.call.summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

