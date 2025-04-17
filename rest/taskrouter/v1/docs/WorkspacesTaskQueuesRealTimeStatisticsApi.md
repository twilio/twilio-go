# WorkspacesTaskQueuesRealTimeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskQueueBulkRealTimeStatistics**](WorkspacesTaskQueuesRealTimeStatisticsApi.md#CreateTaskQueueBulkRealTimeStatistics) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues/RealTimeStatistics | Fetch a Task Queue Real Time Statistics in bulk for the array of TaskQueue SIDs, support upto 50 in a request.
[**FetchTaskQueueRealTimeStatistics**](WorkspacesTaskQueuesRealTimeStatisticsApi.md#FetchTaskQueueRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/RealTimeStatistics | 



## CreateTaskQueueBulkRealTimeStatistics

> TaskrouterV1TaskQueueBulkRealTimeStatistics CreateTaskQueueBulkRealTimeStatistics(ctx, WorkspaceSidoptional)

Fetch a Task Queue Real Time Statistics in bulk for the array of TaskQueue SIDs, support upto 50 in a request.

Fetch a Task Queue Real Time Statistics in bulk for the array of TaskQueue SIDs, support upto 50 in a request.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The unique SID identifier of the Workspace.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskQueueBulkRealTimeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Body** | **map[string]interface{}** | 

### Return type

[**TaskrouterV1TaskQueueBulkRealTimeStatistics**](TaskrouterV1TaskQueueBulkRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueRealTimeStatistics

> TaskrouterV1TaskQueueRealTimeStatistics FetchTaskQueueRealTimeStatistics(ctx, WorkspaceSidTaskQueueSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueRealTimeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**TaskChannel** | **string** | The TaskChannel for which to fetch statistics. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.

### Return type

[**TaskrouterV1TaskQueueRealTimeStatistics**](TaskrouterV1TaskQueueRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

