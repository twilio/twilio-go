# MediaRecordingsApi

All URIs are relative to *https://media.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaRecording**](MediaRecordingsApi.md#DeleteMediaRecording) | **Delete** /v1/MediaRecordings/{Sid} | 
[**FetchMediaRecording**](MediaRecordingsApi.md#FetchMediaRecording) | **Get** /v1/MediaRecordings/{Sid} | 
[**ListMediaRecording**](MediaRecordingsApi.md#ListMediaRecording) | **Get** /v1/MediaRecordings | 



## DeleteMediaRecording

> DeleteMediaRecording(ctx, Sid)



Deletes a MediaRecording resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the MediaRecording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMediaRecordingParams struct


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


## FetchMediaRecording

> MediaV1MediaRecording FetchMediaRecording(ctx, Sid)



Returns a single MediaRecording resource identified by a SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the MediaRecording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMediaRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MediaV1MediaRecording**](MediaV1MediaRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMediaRecording

> []MediaV1MediaRecording ListMediaRecording(ctx, optional)



Returns a list of MediaRecordings.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMediaRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Order** | **string** | The sort order of the list by `date_created`. Can be: `asc` (ascending) or `desc` (descending) with `desc` as the default.
**Status** | **string** | Status to filter by, with possible values `processing`, `completed`, `deleted`, or `failed`.
**ProcessorSid** | **string** | SID of a MediaProcessor to filter by.
**SourceSid** | **string** | SID of a MediaRecording source to filter by.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MediaV1MediaRecording**](MediaV1MediaRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

