# TranscriptionsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateV3Transcriptions**](TranscriptionsApi.md#CreateV3Transcriptions) | **Post** /v3/Transcriptions | Creates a new transcription from either a sourceId or a mediaUrl. Either sourceId or mediaUrl must be provided, but not both.
[**FetchTranscription**](TranscriptionsApi.md#FetchTranscription) | **Get** /v3/Transcriptions/{transcriptionId} | Fetch metadata about a specific transcription



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

