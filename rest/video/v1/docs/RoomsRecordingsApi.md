# RoomsRecordingsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoomRecording**](RoomsRecordingsApi.md#DeleteRoomRecording) | **Delete** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
[**FetchRoomRecording**](RoomsRecordingsApi.md#FetchRoomRecording) | **Get** /v1/Rooms/{RoomSid}/Recordings/{Sid} | 
[**ListRoomRecording**](RoomsRecordingsApi.md#ListRoomRecording) | **Get** /v1/Rooms/{RoomSid}/Recordings | 



## DeleteRoomRecording

> DeleteRoomRecording(ctx, RoomSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the RoomRecording resource to delete.
**Sid** | **string** | The SID of the RoomRecording resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRoomRecordingParams struct


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


## FetchRoomRecording

> VideoV1RoomRecording FetchRoomRecording(ctx, RoomSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource with the recording to fetch.
**Sid** | **string** | The SID of the RoomRecording resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomRecording**](VideoV1RoomRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomRecording

> []VideoV1RoomRecording ListRoomRecording(ctx, RoomSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the RoomRecording resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Read only the recordings with this status. Can be: `processing`, `completed`, or `deleted`.
**SourceSid** | **string** | Read only the recordings that have this `source_sid`.
**DateCreatedAfter** | **time.Time** | Read only recordings that started on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
**DateCreatedBefore** | **time.Time** | Read only Recordings that started before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) datetime with time zone.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1RoomRecording**](VideoV1RoomRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

