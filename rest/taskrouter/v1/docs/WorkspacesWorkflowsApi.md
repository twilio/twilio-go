# WorkspacesWorkflowsApi

All URIs are relative to *https://taskrouter.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflow**](WorkspacesWorkflowsApi.md#CreateWorkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows | 
[**DeleteWorkflow**](WorkspacesWorkflowsApi.md#DeleteWorkflow) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
[**FetchWorkflow**](WorkspacesWorkflowsApi.md#FetchWorkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
[**ListWorkflow**](WorkspacesWorkflowsApi.md#ListWorkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows | 
[**UpdateWorkflow**](WorkspacesWorkflowsApi.md#UpdateWorkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 



## CreateWorkflow

> TaskrouterV1Workflow CreateWorkflow(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Workflow to create belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateWorkflowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
**Configuration** | **string** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
**AssignmentCallbackUrl** | **string** | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
**FallbackAssignmentCallbackUrl** | **string** | The URL that we should call when a call to the `assignment_callback_url` fails.
**TaskReservationTimeout** | **int** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.

### Return type

[**TaskrouterV1Workflow**](TaskrouterV1Workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to delete.
**Sid** | **string** | The SID of the Workflow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWorkflowParams struct


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


## FetchWorkflow

> TaskrouterV1Workflow FetchWorkflow(ctx, WorkspaceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to fetch.
**Sid** | **string** | The SID of the Workflow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TaskrouterV1Workflow**](TaskrouterV1Workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflow

> []TaskrouterV1Workflow ListWorkflow(ctx, WorkspaceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkflowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | The `friendly_name` of the Workflow resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TaskrouterV1Workflow**](TaskrouterV1Workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> TaskrouterV1Workflow UpdateWorkflow(ctx, WorkspaceSidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to update.
**Sid** | **string** | The SID of the Workflow resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkflowParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`.
**AssignmentCallbackUrl** | **string** | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
**FallbackAssignmentCallbackUrl** | **string** | The URL that we should call when a call to the `assignment_callback_url` fails.
**Configuration** | **string** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
**TaskReservationTimeout** | **int** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`.
**ReEvaluateTasks** | **string** | Whether or not to re-evaluate Tasks. The default is `false`, which means Tasks in the Workflow will not be processed through the assignment loop again.

### Return type

[**TaskrouterV1Workflow**](TaskrouterV1Workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

