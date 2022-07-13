# WorkspacesCumulativeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkspaceCumulativeStatistics**](WorkspacesCumulativeStatisticsApi.md#FetchWorkspaceCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/CumulativeStatistics | 



## FetchWorkspaceCumulativeStatistics

> TaskrouterV1WorkspaceCumulativeStatistics FetchWorkspaceCumulativeStatistics(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceCumulativeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**EndDate** | **time.Time** | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
**Minutes** | **int** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
**StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
**TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
**SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. TaskRouter will calculate statistics on up to 10,000 Tasks for any given threshold.

### Return type

[**TaskrouterV1WorkspaceCumulativeStatistics**](TaskrouterV1WorkspaceCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

