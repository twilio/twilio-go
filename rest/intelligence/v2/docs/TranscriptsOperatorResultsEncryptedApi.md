# TranscriptsOperatorResultsEncryptedApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEncryptedOperatorResults**](TranscriptsOperatorResultsEncryptedApi.md#FetchEncryptedOperatorResults) | **Get** /v2/Transcripts/{TranscriptSid}/OperatorResults/Encrypted | Fetch Public Key Encrypted Operator Results by TranscriptSid



## FetchEncryptedOperatorResults

> IntelligenceV2EncryptedOperatorResults FetchEncryptedOperatorResults(ctx, TranscriptSidoptional)

Fetch Public Key Encrypted Operator Results by TranscriptSid

Fetch Public Key Encrypted Operator Results by TranscriptSid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptSid** | **string** | The unique SID identifier of the Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchEncryptedOperatorResultsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII Redacted/Unredacted Operator Results. If redaction is enabled, the default is `true` to access redacted operator results.

### Return type

[**IntelligenceV2EncryptedOperatorResults**](IntelligenceV2EncryptedOperatorResults.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

