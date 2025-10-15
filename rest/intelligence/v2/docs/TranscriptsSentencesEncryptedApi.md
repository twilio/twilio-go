# TranscriptsSentencesEncryptedApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEncryptedSentences**](TranscriptsSentencesEncryptedApi.md#FetchEncryptedSentences) | **Get** /v2/Transcripts/{TranscriptSid}/Sentences/Encrypted | Fetch Public Key Encrypted Sentences by TranscriptSid



## FetchEncryptedSentences

> IntelligenceV2EncryptedSentences FetchEncryptedSentences(ctx, TranscriptSidoptional)

Fetch Public Key Encrypted Sentences by TranscriptSid

Fetch Public Key Encrypted Sentences by TranscriptSid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptSid** | **string** | The unique SID identifier of the Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchEncryptedSentencesParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII Redacted/Unredacted Sentences. If redaction is enabled, the default is `true` to access redacted sentences.

### Return type

[**IntelligenceV2EncryptedSentences**](IntelligenceV2EncryptedSentences.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

