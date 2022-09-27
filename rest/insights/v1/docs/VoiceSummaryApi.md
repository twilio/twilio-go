# VoiceSummaryApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchSummary**](VoiceSummaryApi.md#FetchSummary) | **Get** /v1/Voice/{CallSid}/Summary | 



## FetchSummary

> InsightsV1Summary FetchSummary(ctx, CallSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**ProcessingState** | **string** | 

### Return type

[**InsightsV1Summary**](InsightsV1Summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

