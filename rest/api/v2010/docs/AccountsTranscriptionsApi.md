# AccountsTranscriptionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTranscription**](AccountsTranscriptionsApi.md#DeleteTranscription) | **Delete** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**FetchTranscription**](AccountsTranscriptionsApi.md#FetchTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions/{Sid}.json | 
[**ListTranscription**](AccountsTranscriptionsApi.md#ListTranscription) | **Get** /2010-04-01/Accounts/{AccountSid}/Transcriptions.json | 



## DeleteTranscription

> DeleteTranscription(ctx, Sidoptional)



Delete a transcription from the account used to make the request

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTranscriptionParams struct


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


## FetchTranscription

> ApiV2010Transcription FetchTranscription(ctx, Sidoptional)



Fetch an instance of a Transcription

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Transcription resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource to fetch.

### Return type

[**ApiV2010Transcription**](ApiV2010Transcription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTranscription

> []ApiV2010Transcription ListTranscription(ctx, optional)



Retrieve a list of transcriptions belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTranscriptionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Transcription**](ApiV2010Transcription.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

