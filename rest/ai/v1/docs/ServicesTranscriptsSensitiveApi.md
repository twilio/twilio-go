# ServicesTranscriptsSensitiveApi

All URIs are relative to *https://ai.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTranscriptSensitive**](ServicesTranscriptsSensitiveApi.md#FetchTranscriptSensitive) | **Get** /v1/Services/{ServiceSid}/Transcripts/{Sid}/Sensitive | 
[**ListTranscriptSensitive**](ServicesTranscriptsSensitiveApi.md#ListTranscriptSensitive) | **Get** /v1/Services/{ServiceSid}/Transcripts/Sensitive | 



## FetchTranscriptSensitive

> AiV1TranscriptSensitive FetchTranscriptSensitive(ctx, ServiceSidSid)



Fetch a specific Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | A 34 character string that uniquely identifies this Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptSensitiveParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AiV1TranscriptSensitive**](AiV1TranscriptSensitive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscriptSensitive

> []AiV1TranscriptSensitive ListTranscriptSensitive(ctx, ServiceSidoptional)



Retrieve a list of Transcripts for a given service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListTranscriptSensitiveParams struct


Name | Type | Description
------------- | ------------- | -------------
**BeforeStartTime** | **string** | Filter by before StartTime.
**AfterStartTime** | **string** | Filter by after StartTime.
**BeforeDateCreated** | **string** | Filter by before DateCreated.
**AfterDateCreated** | **string** | Filter by after DateCreated.
**Status** | **string** | Filter by status.
**LanguageLocale** | **string** | Filter by LanguageLocale.
**CallSid** | **string** | Filter by CallSid.
**RecordingSid** | **string** | Filter by RecordingSid.
**ConfidenceLessThan** | **float32** | Filter by confidence (less than or equal to).
**ConfidenceGreaterThan** | **float32** | Filter by confidence (greater than or equal to).
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AiV1TranscriptSensitive**](AiV1TranscriptSensitive.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

