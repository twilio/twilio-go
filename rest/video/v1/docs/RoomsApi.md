# RoomsApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoom**](RoomsApi.md#CreateRoom) | **Post** /v1/Rooms | 
[**FetchRoom**](RoomsApi.md#FetchRoom) | **Get** /v1/Rooms/{Sid} | 
[**ListRoom**](RoomsApi.md#ListRoom) | **Get** /v1/Rooms | 
[**UpdateRoom**](RoomsApi.md#UpdateRoom) | **Post** /v1/Rooms/{Sid} | 



## CreateRoom

> VideoV1Room CreateRoom(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRoomParams struct


Name | Type | Description
------------- | ------------- | -------------
**AudioOnly** | **bool** | When set to true, indicates that the participants in the room will only publish audio. No video tracks will be allowed. Group rooms only.
**EmptyRoomTimeout** | **int** | Configures how long (in minutes) a room will remain active after last participant leaves. Valid values range from 1 to 60 minutes (no fractions).
**EnableTurn** | **bool** | Deprecated, now always considered to be true.
**MaxParticipantDuration** | **int** | The maximum number of seconds a Participant can be connected to the room. The maximum possible value is 86400 seconds (24 hours). The default is 14400 seconds (4 hours).
**MaxParticipants** | **int** | The maximum number of concurrent Participants allowed in the room. Peer-to-peer rooms can have up to 10 Participants. Small Group rooms can have up to 4 Participants. Group rooms can have up to 50 Participants.
**MediaRegion** | **string** | The region for the media server in Group Rooms.  Can be: one of the [available Media Regions](https://www.twilio.com/docs/video/ip-address-whitelisting#group-rooms-media-servers). ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.***
**RecordParticipantsOnConnect** | **bool** | Whether to start recording when Participants connect. ***This feature is not available in &#x60;peer-to-peer&#x60; rooms.***
**RecordingRules** | [**map[string]interface{}**](map[string]interface{}.md) | A collection of Recording Rules that describe how to include or exclude matching tracks for recording
**StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application on every room event. See [Status Callbacks](https://www.twilio.com/docs/video/api/status-callbacks) for more info.
**StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be &#x60;POST&#x60; or &#x60;GET&#x60;.
**Type** | **string** | The type of room. Can be: &#x60;go&#x60;, &#x60;peer-to-peer&#x60;, &#x60;group-small&#x60;, or &#x60;group&#x60;. The default value is &#x60;group&#x60;.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used as a &#x60;room_sid&#x60; in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource, assuming it does not contain any [reserved characters](https://tools.ietf.org/html/rfc3986#section-2.2) that would need to be URL encoded. This value is unique for &#x60;in-progress&#x60; rooms. SDK clients can use this name to connect to the room. REST API clients can use this name in place of the Room SID to interact with the room as long as the room is &#x60;in-progress&#x60;.
**UnusedRoomTimeout** | **int** | Configures how long (in minutes) a room will remain active if no one joins. Valid values range from 1 to 60 minutes (no fractions).
**VideoCodecs** | **[]string** | An array of the video codecs that are supported when publishing a track in the room.  Can be: &#x60;VP8&#x60; and &#x60;H264&#x60;.  ***This feature is not available in &#x60;peer-to-peer&#x60; rooms***

### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRoom

> VideoV1Room FetchRoom(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Room resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoom

> []VideoV1Room ListRoom(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | Read only the rooms with this status. Can be: &#x60;in-progress&#x60; (default) or &#x60;completed&#x60;
**UniqueName** | **string** | Read only rooms with the this &#x60;unique_name&#x60;.
**DateCreatedAfter** | **time.Time** | Read only rooms that started on or after this date, given as &#x60;YYYY-MM-DD&#x60;.
**DateCreatedBefore** | **time.Time** | Read only rooms that started before this date, given as &#x60;YYYY-MM-DD&#x60;.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoom

> VideoV1Room UpdateRoom(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Room resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParams struct


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The new status of the resource. Set to &#x60;completed&#x60; to end the room.

### Return type

[**VideoV1Room**](VideoV1Room.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

