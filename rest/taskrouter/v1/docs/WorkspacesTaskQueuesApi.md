# WorkspacesTaskQueuesApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTaskQueue**](WorkspacesTaskQueuesApi.md#CreateTaskQueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
[**DeleteTaskQueue**](WorkspacesTaskQueuesApi.md#DeleteTaskQueue) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
[**FetchTaskQueue**](WorkspacesTaskQueuesApi.md#FetchTaskQueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
[**ListTaskQueue**](WorkspacesTaskQueuesApi.md#ListTaskQueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
[**UpdateTaskQueue**](WorkspacesTaskQueuesApi.md#UpdateTaskQueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 



## CreateTaskQueue

> TaskrouterV1WorkspaceTaskQueue CreateTaskQueue(ctx, WorkspaceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new TaskQueue belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskQueueParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned to them.
**FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;.
**MaxReservedWorkers** | **int** | The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
**ReservationActivitySid** | **string** | The SID of the Activity to assign Workers when a task is reserved for them.
**TargetWorkers** | **string** | A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, &#x60;&#39;\\\&quot;language\\\&quot; &#x3D;&#x3D; \\\&quot;spanish\\\&quot;&#39;&#x60;. The default value is &#x60;1&#x3D;&#x3D;1&#x60;. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers).
**TaskOrder** | **string** | How Tasks will be assigned to Workers. Set this parameter to &#x60;LIFO&#x60; to assign most recently created Task first or FIFO to assign the oldest Task first. Default is &#x60;FIFO&#x60;. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more.

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](TaskrouterV1WorkspaceTaskQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTaskQueue

> DeleteTaskQueue(ctx, WorkspaceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to delete.
**Sid** | **string** | The SID of the TaskQueue resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskQueueParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchTaskQueue

> TaskrouterV1WorkspaceTaskQueue FetchTaskQueue(ctx, WorkspaceSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**Sid** | **string** | The SID of the TaskQueue resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](TaskrouterV1WorkspaceTaskQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskQueue

> ListTaskQueueResponse ListTaskQueue(ctx, WorkspaceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskQueueParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue resources to read.
**EvaluateWorkerAttributes** | **string** | The attributes of the Workers to read. Returns the TaskQueues with Workers that match the attributes specified in this parameter.
**WorkerSid** | **string** | The SID of the Worker with the TaskQueue resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTaskQueueResponse**](ListTaskQueueResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskQueue

> TaskrouterV1WorkspaceTaskQueue UpdateTaskQueue(ctx, WorkspaceSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to update.
**Sid** | **string** | The SID of the TaskQueue resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskQueueParams struct


Name | Type | Description
------------- | ------------- | -------------
**AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned for them.
**FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;.
**MaxReservedWorkers** | **int** | The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50.
**ReservationActivitySid** | **string** | The SID of the Activity to assign Workers when a task is reserved for them.
**TargetWorkers** | **string** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example &#39;\\\&quot;language\\\&quot; &#x3D;&#x3D; \\\&quot;spanish\\\&quot;&#39; If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below.
**TaskOrder** | **string** | How Tasks will be assigned to Workers. Can be: &#x60;FIFO&#x60; or &#x60;LIFO&#x60; and the default is &#x60;FIFO&#x60;. Use &#x60;FIFO&#x60; to assign the oldest task first and &#x60;LIFO&#x60; to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](TaskrouterV1WorkspaceTaskQueue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

