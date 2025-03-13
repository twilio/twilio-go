# TranscriptsOperatorResultsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOperatorResult**](TranscriptsOperatorResultsApi.md#FetchOperatorResult) | **Get** /v2/Transcripts/{TranscriptSid}/OperatorResults/{OperatorSid} | Fetch a specific Operator Result for the given Transcript.
[**ListOperatorResult**](TranscriptsOperatorResultsApi.md#ListOperatorResult) | **Get** /v2/Transcripts/{TranscriptSid}/OperatorResults | Retrieve a list of Operator Results for the given Transcript.



## FetchOperatorResult

> IntelligenceV2OperatorResult FetchOperatorResult(ctx, TranscriptSidOperatorSidoptional)

Fetch a specific Operator Result for the given Transcript.

Fetch a specific Operator Result for the given Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptSid** | **string** | A 34 character string that uniquely identifies this Transcript.
**OperatorSid** | **string** | A 34 character string that identifies this Language Understanding operator sid.

### Other Parameters

Other parameters are passed through a pointer to a FetchOperatorResultParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII redacted/unredacted Language Understanding operator. If redaction is enabled, the default is True.

### Return type

[**IntelligenceV2OperatorResult**](IntelligenceV2OperatorResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperatorResult

> []IntelligenceV2OperatorResult ListOperatorResult(ctx, TranscriptSidoptional)

Retrieve a list of Operator Results for the given Transcript.

Retrieve a list of Operator Results for the given Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptSid** | **string** | A 34 character string that uniquely identifies this Transcript.

### Other Parameters

Other parameters are passed through a pointer to a ListOperatorResultParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII redacted/unredacted Language Understanding operator. If redaction is enabled, the default is True.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2OperatorResult**](IntelligenceV2OperatorResult.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

