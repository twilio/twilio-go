# RoomsParticipantsAnonymizeApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRoomParticipantAnonymize**](RoomsParticipantsAnonymizeApi.md#UpdateRoomParticipantAnonymize) | **Post** /v1/Rooms/{RoomSid}/Participants/{Sid}/Anonymize | 



## UpdateRoomParticipantAnonymize

> VideoV1RoomParticipantAnonymize UpdateRoomParticipantAnonymize(ctx, RoomSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the room with the participant to update.
**Sid** | **string** | The SID of the RoomParticipant resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParticipantAnonymizeParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomParticipantAnonymize**](VideoV1RoomParticipantAnonymize.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

