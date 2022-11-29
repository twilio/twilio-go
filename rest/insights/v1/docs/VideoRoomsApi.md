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

> []InsightsV1VideoRoomSummary ListVideoRoomSummary(ctx, optional)



Get a list of Programmable Video Rooms.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoRoomSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**RoomType** | [**[]VideoRoomSummaryEnumRoomType**](VideoRoomSummaryEnumRoomType.md) | Type of room. Can be `go`, `peer_to_peer`, `group`, or `group_small`.
**Codec** | [**[]VideoRoomSummaryEnumCodec**](VideoRoomSummaryEnumCodec.md) | Codecs used by participants in the room. Can be `VP8`, `H264`, or `VP9`.
**RoomName** | **string** | Room friendly name.
**CreatedAfter** | **time.Time** | Only read rooms that started on or after this ISO 8601 timestamp.
**CreatedBefore** | **time.Time** | Only read rooms that started before this ISO 8601 timestamp.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1VideoRoomSummary**](InsightsV1VideoRoomSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

