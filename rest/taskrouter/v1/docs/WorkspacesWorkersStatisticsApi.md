# WorkspacesWorkersStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkerInstanceStatistics**](WorkspacesWorkersStatisticsApi.md#FetchWorkerInstanceStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Statistics | 
[**FetchWorkerStatistics**](WorkspacesWorkersStatisticsApi.md#FetchWorkerStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/Statistics | 



## FetchWorkerInstanceStatistics

> TaskrouterV1WorkerInstanceStatistics FetchWorkerInstanceStatistics(ctx, WorkspaceSidWorkerSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to fetch.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerInstanceStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**EndDate** | **time.Time** | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkerInstanceStatistics**](TaskrouterV1WorkerInstanceStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerStatistics

> TaskrouterV1WorkerStatistics FetchWorkerStatistics(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch Worker statistics.
**TaskQueueName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue for which to fetch Worker statistics.
**FriendlyName** | **string** | Only include Workers with &#x60;friendly_name&#x60; values that match this parameter.
**TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkerStatistics**](TaskrouterV1WorkerStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

