# WorkspacesTaskQueuesRealTimeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaskQueueRealTimeStatistics**](WorkspacesTaskQueuesRealTimeStatisticsApi.md#FetchTaskQueueRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/RealTimeStatistics | 



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
**TaskChannel** | **string** | The TaskChannel for which to fetch statistics. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

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

