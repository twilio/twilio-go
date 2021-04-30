# DefaultApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCall**](DefaultApi.md#FetchCall) | **Get** /v1/Voice/{Sid} | 
[**FetchSummary**](DefaultApi.md#FetchSummary) | **Get** /v1/Voice/{CallSid}/Summary | 
[**FetchVideoParticipantSummary**](DefaultApi.md#FetchVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid} | 
[**FetchVideoRoomSummary**](DefaultApi.md#FetchVideoRoomSummary) | **Get** /v1/Video/Rooms/{RoomSid} | 
[**ListEvent**](DefaultApi.md#ListEvent) | **Get** /v1/Voice/{CallSid}/Events | 
[**ListMetric**](DefaultApi.md#ListMetric) | **Get** /v1/Voice/{CallSid}/Metrics | 
[**ListVideoParticipantSummary**](DefaultApi.md#ListVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants | 
[**ListVideoRoomSummary**](DefaultApi.md#ListVideoRoomSummary) | **Get** /v1/Video/Rooms | 



## FetchCall

> InsightsV1Call FetchCall(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCallParams struct

Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1Call**](InsightsV1Call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSummary

> InsightsV1CallSummary FetchSummary(ctx, CallSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSummaryParams struct

Name | Type | Description
------------- | ------------- | -------------
**ProcessingState** | **string** | 

### Return type

[**InsightsV1CallSummary**](InsightsV1CallSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVideoParticipantSummary

> InsightsV1VideoRoomSummaryVideoParticipantSummary FetchVideoParticipantSummary(ctx, RoomSidParticipantSid)



Get Video Log Analyzer data for a Room Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.
**ParticipantSid** | **string** | The SID of the Participant resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchVideoParticipantSummaryParams struct

Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1VideoRoomSummaryVideoParticipantSummary**](InsightsV1VideoRoomSummaryVideoParticipantSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListEvent

> ListEventResponse ListEvent(ctx, CallSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct

Name | Type | Description
------------- | ------------- | -------------
**Edge** | **string** | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventResponse**](ListEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetric

> ListMetricResponse ListMetric(ctx, CallSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CallSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a ListMetricParams struct

Name | Type | Description
------------- | ------------- | -------------
**Edge** | **string** | 
**Direction** | **string** | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMetricResponse**](ListMetricResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoParticipantSummary

> ListVideoParticipantSummaryResponse ListVideoParticipantSummary(ctx, RoomSidoptional)



Get a list of room participants.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource.

### Other Parameters

Other parameters are passed through a pointer to a ListVideoParticipantSummaryParams struct

Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVideoParticipantSummaryResponse**](ListVideoParticipantSummaryResponse.md)

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
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

