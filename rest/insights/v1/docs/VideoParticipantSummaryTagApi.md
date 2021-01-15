# \VideoParticipantSummaryTagApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVideoParticipantSummary**](VideoParticipantSummaryTagApi.md#FetchVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants/{ParticipantSid} | 
[**ListVideoParticipantSummary**](VideoParticipantSummaryTagApi.md#ListVideoParticipantSummary) | **Get** /v1/Video/Rooms/{RoomSid}/Participants | 



## FetchVideoParticipantSummary

> InsightsV1VideoRoomSummaryVideoParticipantSummary FetchVideoParticipantSummary(ctx, roomSid, participantSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSid** | **string**|  | 
**participantSid** | **string**|  | 

### Return type

[**InsightsV1VideoRoomSummaryVideoParticipantSummary**](insights.v1.video_room_summary.video_participant_summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVideoParticipantSummary

> InsightsV1VideoRoomSummaryVideoParticipantSummaryReadResponse ListVideoParticipantSummary(ctx, roomSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSid** | **string**|  | 
 **optional** | ***ListVideoParticipantSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVideoParticipantSummaryOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InsightsV1VideoRoomSummaryVideoParticipantSummaryReadResponse**](insights_v1_video_room_summary_video_participant_summaryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

