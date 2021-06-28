# WorkspacesRealTimeStatisticsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchWorkspaceRealTimeStatistics**](WorkspacesRealTimeStatisticsApi.md#FetchWorkspaceRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/RealTimeStatistics | 



## FetchWorkspaceRealTimeStatistics

> TaskrouterV1WorkspaceWorkspaceRealTimeStatistics FetchWorkspaceRealTimeStatistics(ctx, WorkspaceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceRealTimeStatisticsParams struct


Name | Type | Description
------------- | ------------- | -------------
**TaskChannel** | **string** | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkspaceRealTimeStatistics**](TaskrouterV1WorkspaceWorkspaceRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

