# VideoRoomsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVideoRoomSummary**](VideoRoomsApi.md#FetchVideoRoomSummary) | **Get** /v1/Video/Rooms/{RoomSid} | 
[**ListVideoRoomSummary**](VideoRoomsApi.md#ListVideoRoomSummary) | **Get** /v1/Video/Rooms | 



## FetchVideoRoomSummary

> InsightsV1VideoRoomSummary FetchVideoRoomSummary(ctx, RoomSid)



Get Video Log Analyzer data for a Room.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchVideoRoomSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1VideoRoomSummary**](InsightsV1VideoRoomSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoRoomSummary

> ListVideoRoomSummaryResponse ListVideoRoomSummary(ctx, optional)



Get a list of Programmable Video Rooms.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoRoomSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**RoomType** | **[]string** | Type of room. Can be &#x60;go&#x60;, &#x60;peer_to_peer&#x60;, &#x60;group&#x60;, or &#x60;group_small&#x60;.
**Codec** | **[]string** | Codecs used by participants in the room. Can be &#x60;VP8&#x60;, &#x60;H264&#x60;, or &#x60;VP9&#x60;.
**RoomName** | **string** | Room friendly name.
**CreatedAfter** | **time.Time** | Only read rooms that started on or after this ISO 8601 timestamp.
**CreatedBefore** | **time.Time** | Only read rooms that started before this ISO 8601 timestamp.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVideoRoomSummaryResponse**](ListVideoRoomSummaryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

