# RoomsTranscriptionsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoomTranscriptions**](RoomsTranscriptionsApi.md#CreateRoomTranscriptions) | **Post** /v1/Rooms/{RoomSid}/Transcriptions | 
[**FetchRoomTranscriptions**](RoomsTranscriptionsApi.md#FetchRoomTranscriptions) | **Get** /v1/Rooms/{RoomSid}/Transcriptions/{Ttid} | 
[**ListRoomTranscriptions**](RoomsTranscriptionsApi.md#ListRoomTranscriptions) | **Get** /v1/Rooms/{RoomSid}/Transcriptions | 
[**UpdateRoomTranscriptions**](RoomsTranscriptionsApi.md#UpdateRoomTranscriptions) | **Post** /v1/Rooms/{RoomSid}/Transcriptions/{Ttid} | 



## CreateRoomTranscriptions

> VideoV1RoomTranscriptions CreateRoomTranscriptions(ctx, RoomSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room new transcriptions resource to be created.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoomTranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Configuration** | [**map[string]interface{}**](map[string]interface{}.md) | A collection of properties that describe transcription behaviour.

### Return type

[**VideoV1RoomTranscriptions**](VideoV1RoomTranscriptions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoomTranscriptions

> VideoV1RoomTranscriptions FetchRoomTranscriptions(ctx, RoomSidTtid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the transcriptions resource to fetch.
**Ttid** | **string** | The Twilio type id of the transcriptions resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomTranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomTranscriptions**](VideoV1RoomTranscriptions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomTranscriptions

> []VideoV1RoomTranscriptions ListRoomTranscriptions(ctx, RoomSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the transcriptions resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomTranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1RoomTranscriptions**](VideoV1RoomTranscriptions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomTranscriptions

> VideoV1RoomTranscriptions UpdateRoomTranscriptions(ctx, RoomSidTtidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the transcriptions resource to update.
**Ttid** | **string** | The Twilio type id of the transcriptions resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomTranscriptionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | [**string**](string.md) | 

### Return type

[**VideoV1RoomTranscriptions**](VideoV1RoomTranscriptions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

