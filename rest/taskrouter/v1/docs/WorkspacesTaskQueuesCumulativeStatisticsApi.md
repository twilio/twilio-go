# WorkspacesTaskQueuesCumulativeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskQueueCumulativeStatistics**](WorkspacesTaskQueuesCumulativeStatisticsApi.md#FetchTaskQueueCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/CumulativeStatistics | 



## FetchTaskQueueCumulativeStatistics

> TaskrouterV1TaskQueueCumulativeStatistics FetchTaskQueueCumulativeStatistics(ctx, WorkspaceSidTaskQueueSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueCumulativeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.
**SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. TaskRouter will calculate statistics on up to 10,000 Tasks/Reservations for any given threshold.

### Return type

[**TaskrouterV1TaskQueueCumulativeStatistics**](TaskrouterV1TaskQueueCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

