# WorkspacesWorkersApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorker**](WorkspacesWorkersApi.md#CreateWorker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers | 
[**DeleteWorker**](WorkspacesWorkersApi.md#DeleteWorker) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
[**FetchWorker**](WorkspacesWorkersApi.md#FetchWorker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
[**ListWorker**](WorkspacesWorkersApi.md#ListWorker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers | 
[**UpdateWorker**](WorkspacesWorkersApi.md#UpdateWorker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 



## CreateWorker

> TaskrouterV1Worker CreateWorker(ctx, WorkspaceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Worker belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateWorkerParams struct


Name | Type | Description
------------- | ------------- | -------------
**ActivitySid** | **string** | The SID of a valid Activity that will describe the new Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker&#39;s initial state is the &#x60;default_activity_sid&#x60; configured on the Workspace.
**Attributes** | **string** | A valid JSON string that describes the new Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}.
**FriendlyName** | **string** | A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.

### Return type

[**TaskrouterV1Worker**](TaskrouterV1Worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorker

> DeleteWorker(ctx, WorkspaceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to delete.
**Sid** | **string** | The SID of the Worker resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWorkerParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorker

> TaskrouterV1Worker FetchWorker(ctx, WorkspaceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to fetch.
**Sid** | **string** | The SID of the Worker resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1Worker**](TaskrouterV1Worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorker

> ListWorkerResponse ListWorker(ctx, WorkspaceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workers to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerParams struct


Name | Type | Description
------------- | ------------- | -------------
**ActivityName** | **string** | The &#x60;activity_name&#x60; of the Worker resources to read.
**ActivitySid** | **string** | The &#x60;activity_sid&#x60; of the Worker resources to read.
**Available** | **string** | Whether to return only Worker resources that are available or unavailable. Can be &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; to return Worker resources that are available, and &#x60;false&#x60;, or any value returns the Worker resources that are not available.
**FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Worker resources to read.
**TargetWorkersExpression** | **string** | Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue.
**TaskQueueName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue that the Workers to read are eligible for.
**TaskQueueSid** | **string** | The SID of the TaskQueue that the Workers to read are eligible for.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListWorkerResponse**](ListWorkerResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorker

> TaskrouterV1Worker UpdateWorker(ctx, WorkspaceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to update.
**Sid** | **string** | The SID of the Worker resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**ActivitySid** | **string** | The SID of a valid Activity that will describe the Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information.
**Attributes** | **string** | The JSON string that describes the Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}.
**FriendlyName** | **string** | A descriptive string that you create to describe the Worker. It can be up to 64 characters long.
**RejectPendingReservations** | **bool** | Whether to reject the Worker&#39;s pending reservations. This option is only valid if the Worker&#39;s new [Activity](https://www.twilio.com/docs/taskrouter/api/activity) resource has its &#x60;availability&#x60; property set to &#x60;False&#x60;.

### Return type

[**TaskrouterV1Worker**](TaskrouterV1Worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

