# WorkspacesTaskQueuesStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskQueueStatistics**](WorkspacesTaskQueuesStatisticsApi.md#FetchTaskQueueStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/Statistics | 
[**ListTaskQueuesStatistics**](WorkspacesTaskQueuesStatisticsApi.md#ListTaskQueuesStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/Statistics | 



## FetchTaskQueueStatistics

> TaskrouterV1TaskQueueStatistics FetchTaskQueueStatistics(ctx, WorkspaceSidTaskQueueSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**TaskChannel** | **string** | Only calculate real-time and cumulative statistics for the specified TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
**SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.

### Return type

[**TaskrouterV1TaskQueueStatistics**](TaskrouterV1TaskQueueStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskQueuesStatistics

> []TaskrouterV1TaskQueuesStatistics ListTaskQueuesStatistics(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueues to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskQueuesStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**FriendlyName** | **string** | The `friendly_name` of the TaskQueue statistics to read.
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
**SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1TaskQueuesStatistics**](TaskrouterV1TaskQueuesStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

