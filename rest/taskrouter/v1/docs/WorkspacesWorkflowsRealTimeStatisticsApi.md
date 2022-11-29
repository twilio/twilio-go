# WorkspacesWorkflowsRealTimeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkflowRealTimeStatistics**](WorkspacesWorkflowsRealTimeStatisticsApi.md#FetchWorkflowRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/RealTimeStatistics | 



## FetchWorkflowRealTimeStatistics

> TaskrouterV1WorkflowRealTimeStatistics FetchWorkflowRealTimeStatistics(ctx, WorkspaceSidWorkflowSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to fetch.
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowRealTimeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**TaskChannel** | **string** | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.

### Return type

[**TaskrouterV1WorkflowRealTimeStatistics**](TaskrouterV1WorkflowRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

