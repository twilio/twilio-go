# RecordingsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecording**](RecordingsApi.md#DeleteRecording) | **Delete** /v1/Recordings/{Sid} | 
[**FetchRecording**](RecordingsApi.md#FetchRecording) | **Get** /v1/Recordings/{Sid} | 
[**ListRecording**](RecordingsApi.md#ListRecording) | **Get** /v1/Recordings | 



## DeleteRecording

> DeleteRecording(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteRecordingParams struct


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


## FetchRecording

> VoiceV1Recording FetchRecording(ctx, Sid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1Recording**](VoiceV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecording

> []VoiceV1Recording ListRecording(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateCreated** | **time.Time** | 
**DateCreatedBefore** | **time.Time** | 
**DateCreatedAfter** | **time.Time** | 
**CallSid** | **string** | 
**ConferenceSid** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1Recording**](VoiceV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

