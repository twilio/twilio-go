# VoiceAnnotationApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAnnotation**](VoiceAnnotationApi.md#FetchAnnotation) | **Get** /v1/Voice/{CallSid}/Annotation | 
[**UpdateAnnotation**](VoiceAnnotationApi.md#UpdateAnnotation) | **Post** /v1/Voice/{CallSid}/Annotation | 



## FetchAnnotation

> InsightsV1Annotation FetchAnnotation(ctx, CallSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchAnnotationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1Annotation**](InsightsV1Annotation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAnnotation

> InsightsV1Annotation UpdateAnnotation(ctx, CallSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateAnnotationParams struct


Name | Type | Description
------------- | ------------- | -------------
**AnsweredBy** | **string** | 
**CallScore** | **int** | 
**Comment** | **string** | 
**ConnectivityIssue** | **string** | 
**Incident** | **string** | 
**QualityIssues** | **string** | 
**Spam** | **bool** | 

### Return type

[**InsightsV1Annotation**](InsightsV1Annotation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

