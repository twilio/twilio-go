# ServicesTranscriptsApi

All URIs are relative to *https://ai.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranscript**](ServicesTranscriptsApi.md#CreateTranscript) | **Post** /v1/Services/{ServiceSid}/Transcripts | 
[**DeleteTranscript**](ServicesTranscriptsApi.md#DeleteTranscript) | **Delete** /v1/Services/{ServiceSid}/Transcripts/{Sid} | 
[**FetchTranscript**](ServicesTranscriptsApi.md#FetchTranscript) | **Get** /v1/Services/{ServiceSid}/Transcripts/{Sid} | 
[**ListTranscript**](ServicesTranscriptsApi.md#ListTranscript) | **Get** /v1/Services/{ServiceSid}/Transcripts | 



## CreateTranscript

> AiV1Transcript CreateTranscript(ctx, ServiceSidoptional)



Create a new Transcript for the service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------
**DataLogging** | **bool** | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**ModelSid** | **string** | The unique SID identifier of the Model to use.
**MediaUrl** | **string** | The URL of the media to transcribe.
**RecordingSid** | **string** | The unique SID identifier of the Recording to use.
**Participants** | [**interface{}**](interface{}.md) | The array containing the transcript's participants. Participant fields: id, channel, type,role, full_name, email and image_url
**CallDirection** | **string** | 

### Return type

[**AiV1Transcript**](AiV1Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranscript

> DeleteTranscript(ctx, ServiceSidSid)



Delete a specific Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
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

> AiV1Transcript FetchTranscript(ctx, ServiceSidSid)



Fetch a specific Transcript.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | A 34 character string that uniquely identifies this Transcript.

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AiV1Transcript**](AiV1Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscript

> []AiV1Transcript ListTranscript(ctx, ServiceSidoptional)



Retrieve a list of Transcripts for a given service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListTranscriptParams struct


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

[**[]AiV1Transcript**](AiV1Transcript.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

