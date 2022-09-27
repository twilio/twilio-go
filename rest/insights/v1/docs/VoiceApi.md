# VoiceApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCall**](VoiceApi.md#FetchCall) | **Get** /v1/Voice/{Sid} | 



## FetchCall

> InsightsV1Call FetchCall(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCallParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1Call**](InsightsV1Call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

