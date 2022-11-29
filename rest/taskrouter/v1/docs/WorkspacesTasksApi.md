# WorkspacesTasksApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](WorkspacesTasksApi.md#CreateTask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks | 
[**DeleteTask**](WorkspacesTasksApi.md#DeleteTask) | **Delete** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
[**FetchTask**](WorkspacesTasksApi.md#FetchTask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
[**ListTask**](WorkspacesTasksApi.md#ListTask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks | 
[**UpdateTask**](WorkspacesTasksApi.md#UpdateTask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 



## CreateTask

> TaskrouterV1Task CreateTask(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Task belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskParams struct


Name | Type | Description
------------- | ------------- | -------------
**Timeout** | **int** | The amount of time in seconds the new task can live before being assigned. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the `task.canceled` event will fire with description `Task TTL Exceeded`.
**Priority** | **int** | The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647).
**TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel by passing either its `unique_name` or `sid`. Default value is `default`.
**WorkflowSid** | **string** | The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional.
**Attributes** | **string** | A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow's `assignment_callback_url` when the Task is assigned to a Worker. For example: `{ \\\"task_type\\\": \\\"call\\\", \\\"twilio_call_sid\\\": \\\"CAxxx\\\", \\\"customer_ticket_number\\\": \\\"12345\\\" }`.

### Return type

[**TaskrouterV1Task**](TaskrouterV1Task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> DeleteTask(ctx, WorkspaceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to delete.
**Sid** | **string** | The SID of the Task resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, deletes this Task if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).

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


## FetchTask

> TaskrouterV1Task FetchTask(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to fetch.
**Sid** | **string** | The SID of the Task resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1Task**](TaskrouterV1Task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> []TaskrouterV1Task ListTask(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Tasks to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskParams struct


Name | Type | Description
------------- | ------------- | -------------
**Priority** | **int** | The priority value of the Tasks to read. Returns the list of all Tasks in the Workspace with the specified priority.
**AssignmentStatus** | **[]string** | The `assignment_status` of the Tasks you want to read. Can be: `pending`, `reserved`, `assigned`, `canceled`, `wrapping`, or `completed`. Returns all Tasks in the Workspace with the specified `assignment_status`.
**WorkflowSid** | **string** | The SID of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this SID.
**WorkflowName** | **string** | The friendly name of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this friendly name.
**TaskQueueSid** | **string** | The SID of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this SID.
**TaskQueueName** | **string** | The `friendly_name` of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this friendly name.
**EvaluateTaskAttributes** | **string** | The attributes of the Tasks to read. Returns the Tasks that match the attributes specified in this parameter.
**Ordering** | **string** | How to order the returned Task resources. y default, Tasks are sorted by ascending DateCreated. This value is specified as: `Attribute:Order`, where `Attribute` can be either `Priority` or `DateCreated` and `Order` can be either `asc` or `desc`. For example, `Priority:desc` returns Tasks ordered in descending order of their Priority. Multiple sort orders can be specified in a comma-separated list such as `Priority:desc,DateCreated:asc`, which returns the Tasks in descending Priority order and ascending DateCreated Order.
**HasAddons** | **bool** | Whether to read Tasks with addons. If `true`, returns only Tasks with addons. If `false`, returns only Tasks without addons.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1Task**](TaskrouterV1Task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> TaskrouterV1Task UpdateTask(ctx, WorkspaceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to update.
**Sid** | **string** | The SID of the Task resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | If provided, applies this mutation if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
**Attributes** | **string** | The JSON string that describes the custom attributes of the task.
**AssignmentStatus** | **string** | 
**Reason** | **string** | The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason.
**Priority** | **int** | The Task's new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647).
**TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`.

### Return type

[**TaskrouterV1Task**](TaskrouterV1Task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

