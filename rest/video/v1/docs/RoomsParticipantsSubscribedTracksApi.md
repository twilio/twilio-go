# RoomsParticipantsSubscribedTracksApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRoomParticipantSubscribedTrack**](RoomsParticipantsSubscribedTracksApi.md#FetchRoomParticipantSubscribedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks/{Sid} | 
[**ListRoomParticipantSubscribedTrack**](RoomsParticipantsSubscribedTracksApi.md#ListRoomParticipantSubscribedTrack) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks | 



## FetchRoomParticipantSubscribedTrack

> VideoV1RoomParticipantSubscribedTrack FetchRoomParticipantSubscribedTrack(ctx, RoomSidParticipantSidSid)



Returns a single Track resource represented by `track_sid`.  Note: This is one resource with the Video API that requires a SID, be Track Name on the subscriber side is not guaranteed to be unique.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room where the Track resource to fetch is subscribed.
**ParticipantSid** | **string** | The SID of the participant that subscribes to the Track resource to fetch.
**Sid** | **string** | The SID of the RoomParticipantSubscribedTrack resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantSubscribedTrackParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomParticipantSubscribedTrack**](VideoV1RoomParticipantSubscribedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoomParticipantSubscribedTrack

> []VideoV1RoomParticipantSubscribedTrack ListRoomParticipantSubscribedTrack(ctx, RoomSidParticipantSidoptional)



Returns a list of tracks that are subscribed for the participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource with the Track resources to read.
**ParticipantSid** | **string** | The SID of the participant that subscribes to the Track resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListRoomParticipantSubscribedTrackParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VideoV1RoomParticipantSubscribedTrack**](VideoV1RoomParticipantSubscribedTrack.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

