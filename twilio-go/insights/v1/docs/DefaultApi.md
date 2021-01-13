# \DefaultApi

All URIs are relative to *http://localhost*

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

> InsightsV1Call FetchCall(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**|  | 

### Return type

[**InsightsV1Call**](insights.v1.call.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSummary

> InsightsV1CallSummary FetchSummary(ctx, callSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string**|  | 
 **optional** | ***FetchSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **processingState** | **optional.String**|  | 

### Return type

[**InsightsV1CallSummary**](insights.v1.call.summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## FetchVideoRoomSummary

> InsightsV1VideoRoomSummary FetchVideoRoomSummary(ctx, roomSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomSid** | **string**|  | 

### Return type

[**InsightsV1VideoRoomSummary**](insights.v1.video_room_summary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> InsightsV1CallEventReadResponse ListEvent(ctx, callSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string**|  | 
 **optional** | ***ListEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edge** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InsightsV1CallEventReadResponse**](insights_v1_call_eventReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetric

> InsightsV1CallMetricReadResponse ListMetric(ctx, callSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callSid** | **string**|  | 
 **optional** | ***ListMetricOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMetricOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edge** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InsightsV1CallMetricReadResponse**](insights_v1_call_metricReadResponse.md)

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


## ListVideoRoomSummary

> InsightsV1VideoRoomSummaryReadResponse ListVideoRoomSummary(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVideoRoomSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVideoRoomSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roomType** | [**optional.Interface of []string**](string.md)|  | 
 **codec** | [**optional.Interface of []string**](string.md)|  | 
 **roomName** | **optional.String**|  | 
 **createdAfter** | **optional.Time**|  | 
 **createdBefore** | **optional.Time**|  | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InsightsV1VideoRoomSummaryReadResponse**](insights_v1_video_room_summaryReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

