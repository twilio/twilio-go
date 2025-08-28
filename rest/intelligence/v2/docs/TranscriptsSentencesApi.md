# TranscriptsSentencesApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSentence**](TranscriptsSentencesApi.md#ListSentence) | **Get** /v2/Transcripts/{TranscriptSid}/Sentences | Get all Transcript Sentences by TranscriptSid



## ListSentence

> []IntelligenceV2Sentence ListSentence(ctx, TranscriptSidoptional)

Get all Transcript Sentences by TranscriptSid

Get all Transcript Sentences by TranscriptSid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptSid** | **string** | The unique SID identifier of the Transcript.

### Other Parameters

Other parameters are passed through a pointer to a ListSentenceParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII Redacted/Unredacted Sentences. If redaction is enabled, the default is `true` to access redacted sentences.
**WordTimestamps** | **bool** | Returns word level timestamps information, if word_timestamps is enabled. The default is `false`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 5000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2Sentence**](IntelligenceV2Sentence.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

