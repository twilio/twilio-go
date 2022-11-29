# RoomsParticipantsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRoomParticipant**](RoomsParticipantsApi.md#FetchRoomParticipant) | **Get** /v1/Rooms/{RoomSid}/Participants/{Sid} | 
[**ListRoomParticipant**](RoomsParticipantsApi.md#ListRoomParticipant) | **Get** /v1/Rooms/{RoomSid}/Participants | 
[**UpdateRoomParticipant**](RoomsParticipantsApi.md#UpdateRoomParticipant) | **Post** /v1/Rooms/{RoomSid}/Participants/{Sid} | 



## FetchRoomParticipant

> VideoV1RoomParticipant FetchRoomParticipant(ctx, RoomSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the Participant resource to fetch.
**Sid** | **string** | The SID of the RoomParticipant resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomParticipant**](VideoV1RoomParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipant

> []VideoV1RoomParticipant ListRoomParticipant(ctx, RoomSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the Participant resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Read only the participants with this status. Can be: `connected` or `disconnected`. For `in-progress` Rooms the default Status is `connected`, for `completed` Rooms only `disconnected` Participants are returned.
**Identity** | **string** | Read only the Participants with this [User](https://www.twilio.com/docs/chat/rest/user-resource) `identity` value.
**DateCreatedAfter** | **time.Time** | Read only Participants that started after this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
**DateCreatedBefore** | **time.Time** | Read only Participants that started before this date in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1RoomParticipant**](VideoV1RoomParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipant

> VideoV1RoomParticipant UpdateRoomParticipant(ctx, RoomSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the participant to update.
**Sid** | **string** | The SID of the RoomParticipant resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParticipantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | 

### Return type

[**VideoV1RoomParticipant**](VideoV1RoomParticipant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

