# \VideoRoomSummaryTagApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVideoRoomSummary**](VideoRoomSummaryTagApi.md#FetchVideoRoomSummary) | **Get** /v1/Video/Rooms/{RoomSid} | 
[**ListVideoRoomSummary**](VideoRoomSummaryTagApi.md#ListVideoRoomSummary) | **Get** /v1/Video/Rooms | 



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

