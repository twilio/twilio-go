# RoomsRecordingRulesApi

All URIs are relative to *https://video.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchRoomRecordingRule**](RoomsRecordingRulesApi.md#FetchRoomRecordingRule) | **Get** /v1/Rooms/{RoomSid}/RecordingRules | 
[**UpdateRoomRecordingRule**](RoomsRecordingRulesApi.md#UpdateRoomRecordingRule) | **Post** /v1/Rooms/{RoomSid}/RecordingRules | 



## FetchRoomRecordingRule

> VideoV1RoomRecordingRule FetchRoomRecordingRule(ctx, RoomSid)



Returns a list of Recording Rules for the Room.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the recording rules to fetch apply.

### Other Parameters

Other parameters are passed through a pointer to a FetchRoomRecordingRuleParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VideoV1RoomRecordingRule**](VideoV1RoomRecordingRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoomRecordingRule

> VideoV1RoomRecordingRule UpdateRoomRecordingRule(ctx, RoomSidoptional)



Update the Recording Rules for the Room

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RoomSid** | **string** | The SID of the Room resource where the recording rules to update apply.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRoomRecordingRuleParams struct


Name | Type | Description
------------- | ------------- | -------------
**Rules** | [**interface{}**](interface{}.md) | A JSON-encoded array of recording rules.

### Return type

[**VideoV1RoomRecordingRule**](VideoV1RoomRecordingRule.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

