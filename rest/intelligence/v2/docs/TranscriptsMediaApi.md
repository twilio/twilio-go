# TranscriptsMediaApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMedia**](TranscriptsMediaApi.md#FetchMedia) | **Get** /v2/Transcripts/{Sid}/Media | Get download URLs for media if possible



## FetchMedia

> IntelligenceV2Media FetchMedia(ctx, Sidoptional)

Get download URLs for media if possible

Get download URLs for media if possible

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique SID identifier of the Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**Redacted** | **bool** | Grant access to PII Redacted/Unredacted Media. If redaction is enabled, the default is `true` to access redacted media.

### Return type

[**IntelligenceV2Media**](IntelligenceV2Media.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

