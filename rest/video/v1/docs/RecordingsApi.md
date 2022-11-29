# RecordingsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecording**](RecordingsApi.md#DeleteRecording) | **Delete** /v1/Recordings/{Sid} | 
[**FetchRecording**](RecordingsApi.md#FetchRecording) | **Get** /v1/Recordings/{Sid} | 
[**ListRecording**](RecordingsApi.md#ListRecording) | **Get** /v1/Recordings | 



## DeleteRecording

> DeleteRecording(ctx, Sid)



Delete a Recording resource identified by a Recording SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Recording resource to delete.

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

> VideoV1Recording FetchRecording(ctx, Sid)



Returns a single Recording resource identified by a Recording SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Recording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1Recording**](VideoV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecording

> []VideoV1Recording ListRecording(ctx, optional)



List of all Track recordings.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Read only the recordings that have this status. Can be: `processing`, `completed`, or `deleted`.
**SourceSid** | **string** | Read only the recordings that have this `source_sid`.
**GroupingSid** | **[]string** | Read only recordings with this `grouping_sid`, which may include a `participant_sid` and/or a `room_sid`.
**DateCreatedAfter** | **time.Time** | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone.
**DateCreatedBefore** | **time.Time** | Read only recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time with time zone, given as `YYYY-MM-DDThh:mm:ss+|-hh:mm` or `YYYY-MM-DDThh:mm:ssZ`.
**MediaType** | **string** | Read only recordings that have this media type. Can be either `audio` or `video`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1Recording**](VideoV1Recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

