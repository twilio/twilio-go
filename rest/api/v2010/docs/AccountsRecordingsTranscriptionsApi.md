# AccountsRecordingsTranscriptionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecordingTranscription**](AccountsRecordingsTranscriptionsApi.md#DeleteRecordingTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**FetchRecordingTranscription**](AccountsRecordingsTranscriptionsApi.md#FetchRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions/{Sid}.json | 
[**ListRecordingTranscription**](AccountsRecordingsTranscriptionsApi.md#ListRecordingTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Recordings/{RecordingSid}/Transcriptions.json | 



## DeleteRecordingTranscription

> DeleteRecordingTranscription(ctx, RecordingSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to delete.

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


## FetchRecordingTranscription

> ApiV2010RecordingTranscription FetchRecordingTranscription(ctx, RecordingSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcription to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.

### Return type

[**ApiV2010RecordingTranscription**](ApiV2010RecordingTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecordingTranscription

> []ApiV2010RecordingTranscription ListRecordingTranscription(ctx, RecordingSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RecordingSid** | **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) that created the transcriptions to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010RecordingTranscription**](ApiV2010RecordingTranscription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

