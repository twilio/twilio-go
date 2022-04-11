# RoomsParticipantsSubscribeRulesApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRoomParticipantSubscribeRule**](RoomsParticipantsSubscribeRulesApi.md#FetchRoomParticipantSubscribeRule) | **Get** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | 
[**UpdateRoomParticipantSubscribeRule**](RoomsParticipantsSubscribeRulesApi.md#UpdateRoomParticipantSubscribeRule) | **Post** /v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribeRules | 



## FetchRoomParticipantSubscribeRule

> VideoV1RoomParticipantSubscribeRule FetchRoomParticipantSubscribeRule(ctx, RoomSidParticipantSid)



Returns a list of Subscribe Rules for the Participant.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the subscribe rules to fetch apply.
**ParticipantSid** | **string** | The SID of the Participant resource with the subscribe rules to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomParticipantSubscribeRuleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomParticipantSubscribeRule**](VideoV1RoomParticipantSubscribeRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomParticipantSubscribeRule

> VideoV1RoomParticipantSubscribeRule UpdateRoomParticipantSubscribeRule(ctx, RoomSidParticipantSidoptional)



Update the Subscribe Rules for the Participant

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the subscribe rules to update apply.
**ParticipantSid** | **string** | The SID of the Participant resource to update the Subscribe Rules.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomParticipantSubscribeRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Rules** | [**interface{}**](interface{}.md) | A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information.

### Return type

[**VideoV1RoomParticipantSubscribeRule**](VideoV1RoomParticipantSubscribeRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

