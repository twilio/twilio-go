# VoiceAnnotationApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAnnotation**](VoiceAnnotationApi.md#FetchAnnotation) | **Get** /v1/Voice/{CallSid}/Annotation | 
[**UpdateAnnotation**](VoiceAnnotationApi.md#UpdateAnnotation) | **Post** /v1/Voice/{CallSid}/Annotation | 



## FetchAnnotation

> InsightsV1Annotation FetchAnnotation(ctx, CallSid)



Fetch a specific Annotation.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The unique SID identifier of the Call.

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



Create/Update the annotation for the call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | The unique string that Twilio created to identify this Call resource. It always starts with a CA.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAnnotationParams struct


Name | Type | Description
------------- | ------------- | -------------
**AnsweredBy** | **string** | 
**ConnectivityIssue** | **string** | 
**QualityIssues** | **string** | Specify if the call had any subjective quality issues. Possible values, one or more of:  no_quality_issue, low_volume, choppy_robotic, echo, dtmf, latency, owa, static_noise. Use comma separated values to indicate multiple quality issues for the same call
**Spam** | **bool** | Specify if the call was a spam call. Use this to provide feedback on whether calls placed from your account were marked as spam, or if inbound calls received by your account were unwanted spam. Is of type Boolean: true, false. Use true if the call was a spam call.
**CallScore** | **int** | Specify the call score. This is of type integer. Use a range of 1-5 to indicate the call experience score, with the following mapping as a reference for rating the call [5: Excellent, 4: Good, 3 : Fair, 2 : Poor, 1: Bad].
**Comment** | **string** | Specify any comments pertaining to the call. This of type string with a max limit of 100 characters. Twilio does not treat this field as PII, so don’t put any PII in here.
**Incident** | **string** | Associate this call with an incident or support ticket. This is of type string with a max limit of 100 characters. Twilio does not treat this field as PII, so don’t put any PII in here.

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

