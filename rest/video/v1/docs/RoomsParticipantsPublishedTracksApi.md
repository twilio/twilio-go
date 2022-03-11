# RoomsParticipantsPublishedTracksApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRoomParticipantPublishedTrack**](RoomsParticipantsPublishedTracksApi.md#FetchRoomParticipantPublishedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks/{Sid} | 
[**ListRoomParticipantPublishedTrack**](RoomsParticipantsPublishedTracksApi.md#ListRoomParticipantPublishedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/PublishedTracks | 



## FetchRoomParticipantPublishedTrack

> VideoV1RoomParticipantPublishedTrack FetchRoomParticipantPublishedTrack(ctx, RoomSidParticipantSidSid)



Returns a single Track resource represented by TrackName or SID.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the Track resource to fetch is published.
**ParticipantSid** | **string** | The SID of the Participant resource with the published track to fetch.
**Sid** | **string** | The SID of the RoomParticipantPublishedTrack resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantPublishedTrackParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomParticipantPublishedTrack**](VideoV1RoomParticipantPublishedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipantPublishedTrack

> []VideoV1RoomParticipantPublishedTrack ListRoomParticipantPublishedTrack(ctx, RoomSidParticipantSidoptional)



Returns a list of tracks associated with a given Participant. Only `currently` Published Tracks are in the list resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the Track resources to read are published.
**ParticipantSid** | **string** | The SID of the Participant resource with the published tracks to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantPublishedTrackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1RoomParticipantPublishedTrack**](VideoV1RoomParticipantPublishedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

