# TranscriptionsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV3Transcriptions**](TranscriptionsApi.md#CreateV3Transcriptions) | **Post** /v3/Transcriptions | Creates a new transcription from either a sourceId or a mediaUrl. Either sourceId or mediaUrl must be provided, but not both.
[**FetchTranscription**](TranscriptionsApi.md#FetchTranscription) | **Get** /v3/Transcriptions/{transcriptionId} | Fetch metadata about a specific transcription
[**ListV3Transcriptions**](TranscriptionsApi.md#ListV3Transcriptions) | **Get** /v3/Transcriptions | Retrieves a paginated, filterable list of transcriptions for the account, sorted by creation date (newest first).



## CreateV3Transcriptions

> VoiceV3LongRunningOperationResponse CreateV3Transcriptions(ctx, optional)

Creates a new transcription from either a sourceId or a mediaUrl. Either sourceId or mediaUrl must be provided, but not both.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateV3TranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**IdempotencyKey** | **string** | A unique key to ensure idempotency. We recommend using UUID v7. Requests with the same key within the idempotency window return the original response.
**CreateV3TranscriptionsRequest** | [**CreateV3TranscriptionsRequest**](CreateV3TranscriptionsRequest.md) | 

### Return type

[**VoiceV3LongRunningOperationResponse**](VoiceV3LongRunningOperation202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTranscription

> VoiceV3LongRunningOperationResponse FetchTranscription(ctx, TranscriptionId)

Fetch metadata about a specific transcription

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TranscriptionId** | **string** | The unique identifier of the transcription to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV3LongRunningOperationResponse**](VoiceV3LongRunningOperationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListV3Transcriptions

> []VoiceV3Transcription ListV3Transcriptions(ctx, optional)

Retrieves a paginated, filterable list of transcriptions for the account, sorted by creation date (newest first).

Returns the account's transcriptions, newest first. Filters combine with AND, and paging uses the opaque cursors in meta rather than page numbers: pass meta.nextToken to advance and meta.previousToken to go back. A token is only valid for the filter set that produced it, so changing sourceId or the date range means starting a new walk. Items are the transcription resource itself. Create and fetch-by-id return the long running operation envelope instead, because creation is asynchronous.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListV3TranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreatedAfter** | **time.Time** | Only include transcriptions created at or after this time (inclusive)
**CreatedBefore** | **time.Time** | Only include transcriptions created strictly before this time (exclusive)
**LanguageCode** | **string** | Only include transcriptions whose resolved language matches this value exactly. The comparison is case sensitive, so use the stored form, for example en-US.
**SourceId** | **string** | Only include transcriptions for this source audio. Must be a Recording SID in lowercase hex; anything else is rejected with a 400.
**Status** | **string** | Only include transcriptions in this status
**PageSize** | **int** | Number of results per page. This endpoint caps at 100, which is lower than the shared pagination component's ceiling and matches what the service enforces.
**PageToken** | **string** | Opaque cursor for retrieving the next or previous page of results
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV3Transcription**](VoiceV3Transcription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

