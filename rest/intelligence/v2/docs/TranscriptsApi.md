# TranscriptsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscript**](TranscriptsApi.md#CreateTranscript) | **Post** /v2/Transcripts | Create a new Transcript for the service
[**DeleteTranscript**](TranscriptsApi.md#DeleteTranscript) | **Delete** /v2/Transcripts/{Sid} | Delete a specific Transcript.
[**FetchTranscript**](TranscriptsApi.md#FetchTranscript) | **Get** /v2/Transcripts/{Sid} | Fetch a specific Transcript.
[**ListTranscript**](TranscriptsApi.md#ListTranscript) | **Get** /v2/Transcripts | Retrieve a list of Transcripts for a given service.



## CreateTranscript

> IntelligenceV2Transcript CreateTranscript(ctx, optional)

Create a new Transcript for the service

Create a new Transcript for the service

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Channel** | [**interface{}**](interface{}.md) | JSON object describing Media Channel including Source and Participants
**CustomerKey** | **string** | Used to store client provided metadata. Maximum of 64 double-byte UTF8 characters.
**MediaStartTime** | **time.Time** | The date that this Transcript's media was started, given in ISO 8601 format.

### Return type

[**IntelligenceV2Transcript**](IntelligenceV2Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranscript

> DeleteTranscript(ctx, Sid)

Delete a specific Transcript.

Delete a specific Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Transcript.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTranscript

> IntelligenceV2Transcript FetchTranscript(ctx, Sid)

Fetch a specific Transcript.

Fetch a specific Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2Transcript**](IntelligenceV2Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscript

> []IntelligenceV2Transcript ListTranscript(ctx, optional)

Retrieve a list of Transcripts for a given service.

Retrieve a list of Transcripts for a given service.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------
**ServiceSid** | **string** | The unique SID identifier of the Service.
**BeforeStartTime** | **string** | Filter by before StartTime.
**AfterStartTime** | **string** | Filter by after StartTime.
**BeforeDateCreated** | **string** | Filter by before DateCreated.
**AfterDateCreated** | **string** | Filter by after DateCreated.
**Status** | **string** | Filter by status.
**LanguageCode** | **string** | Filter by Language Code.
**SourceSid** | **string** | Filter by SourceSid.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2Transcript**](IntelligenceV2Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

