# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActivity**](DefaultApi.md#CreateActivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities | 
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks | 
[**CreateTaskChannel**](DefaultApi.md#CreateTaskChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
[**CreateTaskQueue**](DefaultApi.md#CreateTaskQueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
[**CreateWorker**](DefaultApi.md#CreateWorker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers | 
[**CreateWorkflow**](DefaultApi.md#CreateWorkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows | 
[**CreateWorkspace**](DefaultApi.md#CreateWorkspace) | **Post** /v1/Workspaces | 
[**DeleteActivity**](DefaultApi.md#DeleteActivity) | **Delete** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
[**DeleteTaskChannel**](DefaultApi.md#DeleteTaskChannel) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
[**DeleteTaskQueue**](DefaultApi.md#DeleteTaskQueue) | **Delete** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
[**DeleteWorker**](DefaultApi.md#DeleteWorker) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
[**DeleteWorkflow**](DefaultApi.md#DeleteWorkflow) | **Delete** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
[**DeleteWorkspace**](DefaultApi.md#DeleteWorkspace) | **Delete** /v1/Workspaces/{Sid} | 
[**FetchActivity**](DefaultApi.md#FetchActivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
[**FetchEvent**](DefaultApi.md#FetchEvent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events/{Sid} | 
[**FetchTask**](DefaultApi.md#FetchTask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
[**FetchTaskChannel**](DefaultApi.md#FetchTaskChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
[**FetchTaskQueue**](DefaultApi.md#FetchTaskQueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
[**FetchTaskQueueCumulativeStatistics**](DefaultApi.md#FetchTaskQueueCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/CumulativeStatistics | 
[**FetchTaskQueueRealTimeStatistics**](DefaultApi.md#FetchTaskQueueRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/RealTimeStatistics | 
[**FetchTaskQueueStatistics**](DefaultApi.md#FetchTaskQueueStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{TaskQueueSid}/Statistics | 
[**FetchTaskReservation**](DefaultApi.md#FetchTaskReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 
[**FetchWorker**](DefaultApi.md#FetchWorker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
[**FetchWorkerChannel**](DefaultApi.md#FetchWorkerChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 
[**FetchWorkerInstanceStatistics**](DefaultApi.md#FetchWorkerInstanceStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Statistics | 
[**FetchWorkerReservation**](DefaultApi.md#FetchWorkerReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 
[**FetchWorkerStatistics**](DefaultApi.md#FetchWorkerStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/Statistics | 
[**FetchWorkersCumulativeStatistics**](DefaultApi.md#FetchWorkersCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/CumulativeStatistics | 
[**FetchWorkersRealTimeStatistics**](DefaultApi.md#FetchWorkersRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/RealTimeStatistics | 
[**FetchWorkflow**](DefaultApi.md#FetchWorkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
[**FetchWorkflowCumulativeStatistics**](DefaultApi.md#FetchWorkflowCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/CumulativeStatistics | 
[**FetchWorkflowRealTimeStatistics**](DefaultApi.md#FetchWorkflowRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/RealTimeStatistics | 
[**FetchWorkflowStatistics**](DefaultApi.md#FetchWorkflowStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows/{WorkflowSid}/Statistics | 
[**FetchWorkspace**](DefaultApi.md#FetchWorkspace) | **Get** /v1/Workspaces/{Sid} | 
[**FetchWorkspaceCumulativeStatistics**](DefaultApi.md#FetchWorkspaceCumulativeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/CumulativeStatistics | 
[**FetchWorkspaceRealTimeStatistics**](DefaultApi.md#FetchWorkspaceRealTimeStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/RealTimeStatistics | 
[**FetchWorkspaceStatistics**](DefaultApi.md#FetchWorkspaceStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/Statistics | 
[**ListActivity**](DefaultApi.md#ListActivity) | **Get** /v1/Workspaces/{WorkspaceSid}/Activities | 
[**ListEvent**](DefaultApi.md#ListEvent) | **Get** /v1/Workspaces/{WorkspaceSid}/Events | 
[**ListTask**](DefaultApi.md#ListTask) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks | 
[**ListTaskChannel**](DefaultApi.md#ListTaskChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskChannels | 
[**ListTaskQueue**](DefaultApi.md#ListTaskQueue) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues | 
[**ListTaskQueuesStatistics**](DefaultApi.md#ListTaskQueuesStatistics) | **Get** /v1/Workspaces/{WorkspaceSid}/TaskQueues/Statistics | 
[**ListTaskReservation**](DefaultApi.md#ListTaskReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations | 
[**ListWorker**](DefaultApi.md#ListWorker) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers | 
[**ListWorkerChannel**](DefaultApi.md#ListWorkerChannel) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels | 
[**ListWorkerReservation**](DefaultApi.md#ListWorkerReservation) | **Get** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations | 
[**ListWorkflow**](DefaultApi.md#ListWorkflow) | **Get** /v1/Workspaces/{WorkspaceSid}/Workflows | 
[**ListWorkspace**](DefaultApi.md#ListWorkspace) | **Get** /v1/Workspaces | 
[**UpdateActivity**](DefaultApi.md#UpdateActivity) | **Post** /v1/Workspaces/{WorkspaceSid}/Activities/{Sid} | 
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{Sid} | 
[**UpdateTaskChannel**](DefaultApi.md#UpdateTaskChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskChannels/{Sid} | 
[**UpdateTaskQueue**](DefaultApi.md#UpdateTaskQueue) | **Post** /v1/Workspaces/{WorkspaceSid}/TaskQueues/{Sid} | 
[**UpdateTaskReservation**](DefaultApi.md#UpdateTaskReservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Tasks/{TaskSid}/Reservations/{Sid} | 
[**UpdateWorker**](DefaultApi.md#UpdateWorker) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{Sid} | 
[**UpdateWorkerChannel**](DefaultApi.md#UpdateWorkerChannel) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Channels/{Sid} | 
[**UpdateWorkerReservation**](DefaultApi.md#UpdateWorkerReservation) | **Post** /v1/Workspaces/{WorkspaceSid}/Workers/{WorkerSid}/Reservations/{Sid} | 
[**UpdateWorkflow**](DefaultApi.md#UpdateWorkflow) | **Post** /v1/Workspaces/{WorkspaceSid}/Workflows/{Sid} | 
[**UpdateWorkspace**](DefaultApi.md#UpdateWorkspace) | **Post** /v1/Workspaces/{Sid} | 



## CreateActivity

> TaskrouterV1WorkspaceActivity CreateActivity(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new Activity belongs to. | 
 **optional** | ***CreateActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateActivityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **available** | **optional.Bool**| Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: &#x60;on-call&#x60;, &#x60;break&#x60;, and &#x60;email&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceActivity**](taskrouter.v1.workspace.activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> TaskrouterV1WorkspaceTask CreateTask(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new Task belongs to. | 
 **optional** | ***CreateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **optional.String**| A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow&#39;s &#x60;assignment_callback_url&#x60; when the Task is assigned to a Worker. For example: &#x60;{ \\\&quot;task_type\\\&quot;: \\\&quot;call\\\&quot;, \\\&quot;twilio_call_sid\\\&quot;: \\\&quot;CAxxx\\\&quot;, \\\&quot;customer_ticket_number\\\&quot;: \\\&quot;12345\\\&quot; }&#x60;. | 
 **priority** | **optional.Int32**| The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647). | 
 **taskChannel** | **optional.String**| When MultiTasking is enabled, specify the TaskChannel by passing either its &#x60;unique_name&#x60; or &#x60;sid&#x60;. Default value is &#x60;default&#x60;. | 
 **timeout** | **optional.Int32**| The amount of time in seconds the new task is allowed to live. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the &#x60;task.canceled&#x60; event will fire with description &#x60;Task TTL Exceeded&#x60;. | 
 **workflowSid** | **optional.String**| The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional. | 

### Return type

[**TaskrouterV1WorkspaceTask**](taskrouter.v1.workspace.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskChannel

> TaskrouterV1WorkspaceTaskChannel CreateTaskChannel(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new Task Channel belongs to. | 
 **optional** | ***CreateTaskChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelOptimizedRouting** | **optional.Bool**| Whether the Task Channel should prioritize Workers that have been idle. If &#x60;true&#x60;, Workers that have been idle the longest are prioritized. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the Task Channel, such as &#x60;voice&#x60; or &#x60;sms&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceTaskChannel**](taskrouter.v1.workspace.task_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskQueue

> TaskrouterV1WorkspaceTaskQueue CreateTaskQueue(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new TaskQueue belongs to. | 
 **optional** | ***CreateTaskQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTaskQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentActivitySid** | **optional.String**| The SID of the Activity to assign Workers when a task is assigned to them. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;. | 
 **maxReservedWorkers** | **optional.Int32**| The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1. | 
 **reservationActivitySid** | **optional.String**| The SID of the Activity to assign Workers when a task is reserved for them. | 
 **targetWorkers** | **optional.String**| A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, &#x60;&#39;\\\&quot;language\\\&quot; &#x3D;&#x3D; \\\&quot;spanish\\\&quot;&#39;&#x60;. The default value is &#x60;1&#x3D;&#x3D;1&#x60;. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers). | 
 **taskOrder** | **optional.String**| How Tasks will be assigned to Workers. Set this parameter to &#x60;LIFO&#x60; to assign most recently created Task first or FIFO to assign the oldest Task first. Default is &#x60;FIFO&#x60;. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](taskrouter.v1.workspace.task_queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorker

> TaskrouterV1WorkspaceWorker CreateWorker(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new Worker belongs to. | 
 **optional** | ***CreateWorkerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWorkerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activitySid** | **optional.String**| The SID of a valid Activity that will describe the new Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker&#39;s initial state is the &#x60;default_activity_sid&#x60; configured on the Workspace. | 
 **attributes** | **optional.String**| A valid JSON string that describes the new Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the new Worker. It can be up to 64 characters long. | 

### Return type

[**TaskrouterV1WorkspaceWorker**](taskrouter.v1.workspace.worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflow

> TaskrouterV1WorkspaceWorkflow CreateWorkflow(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace that the new Workflow to create belongs to. | 
 **optional** | ***CreateWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignmentCallbackUrl** | **optional.String**| The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details. | 
 **configuration** | **optional.String**| A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. | 
 **fallbackAssignmentCallbackUrl** | **optional.String**| The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;. | 
 **taskReservationTimeout** | **optional.Int32**| How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkflow**](taskrouter.v1.workspace.workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> TaskrouterV1Workspace CreateWorkspace(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWorkspaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventCallbackUrl** | **optional.String**| The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. | 
 **eventsFilter** | **optional.String**| The list of Workspace events for which to call event_callback_url. For example if &#x60;EventsFilter&#x3D;task.created,task.canceled,worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: &#x60;Customer Support&#x60; or &#x60;2014 Election Campaign&#x60;. | 
 **multiTaskEnabled** | **optional.Bool**| Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. The default is &#x60;false&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking][https://www.twilio.com/docs/taskrouter/multitasking]. | 
 **prioritizeQueueOrder** | **optional.String**| The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: &#x60;LIFO&#x60; or &#x60;FIFO&#x60; and the default is &#x60;FIFO&#x60;. For more information, see [Queue Ordering][https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo]. | 
 **template** | **optional.String**| An available template name. Can be: &#x60;NONE&#x60; or &#x60;FIFO&#x60; and the default is &#x60;NONE&#x60;. Pre-configures the Workspace with the Workflow and Activities specified in the template. &#x60;NONE&#x60; will create a Workspace with only a set of default activities. &#x60;FIFO&#x60; will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter. | 

### Return type

[**TaskrouterV1Workspace**](taskrouter.v1.workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActivity

> DeleteActivity(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Activity resources to delete. | 
**sid** | **string**| The SID of the Activity resource to delete. | 

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


## DeleteTask

> DeleteTask(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task to delete. | 
**sid** | **string**| The SID of the Task resource to delete. | 

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


## DeleteTaskChannel

> DeleteTaskChannel(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task Channel to delete. | 
**sid** | **string**| The SID of the Task Channel resource to delete. | 

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


## DeleteTaskQueue

> DeleteTaskQueue(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to delete. | 
**sid** | **string**| The SID of the TaskQueue resource to delete. | 

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


## DeleteWorker

> DeleteWorker(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Worker to delete. | 
**sid** | **string**| The SID of the Worker resource to delete. | 

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


## DeleteWorkflow

> DeleteWorkflow(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to delete. | 
**sid** | **string**| The SID of the Workflow resource to delete. | 

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


## DeleteWorkspace

> DeleteWorkspace(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Workspace resource to delete. | 

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


## FetchActivity

> TaskrouterV1WorkspaceActivity FetchActivity(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Activity resources to fetch. | 
**sid** | **string**| The SID of the Activity resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceActivity**](taskrouter.v1.workspace.activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvent

> TaskrouterV1WorkspaceEvent FetchEvent(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Event to fetch. | 
**sid** | **string**| The SID of the Event resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceEvent**](taskrouter.v1.workspace.event.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTask

> TaskrouterV1WorkspaceTask FetchTask(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task to fetch. | 
**sid** | **string**| The SID of the Task resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceTask**](taskrouter.v1.workspace.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskChannel

> TaskrouterV1WorkspaceTaskChannel FetchTaskChannel(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task Channel to fetch. | 
**sid** | **string**| The SID of the Task Channel resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceTaskChannel**](taskrouter.v1.workspace.task_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueue

> TaskrouterV1WorkspaceTaskQueue FetchTaskQueue(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to fetch. | 
**sid** | **string**| The SID of the TaskQueue resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](taskrouter.v1.workspace.task_queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueCumulativeStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics FetchTaskQueueCumulativeStatistics(ctx, workspaceSid, taskQueueSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to fetch. | 
**taskQueueSid** | **string**| The SID of the TaskQueue for which to fetch statistics. | 
 **optional** | ***FetchTaskQueueCumulativeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchTaskQueueCumulativeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default is 15 minutes. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics**](taskrouter.v1.workspace.task_queue.task_queue_cumulative_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueRealTimeStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics FetchTaskQueueRealTimeStatistics(ctx, workspaceSid, taskQueueSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to fetch. | 
**taskQueueSid** | **string**| The SID of the TaskQueue for which to fetch statistics. | 
 **optional** | ***FetchTaskQueueRealTimeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchTaskQueueRealTimeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taskChannel** | **optional.String**| The TaskChannel for which to fetch statistics. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics**](taskrouter.v1.workspace.task_queue.task_queue_real_time_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics FetchTaskQueueStatistics(ctx, workspaceSid, taskQueueSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to fetch. | 
**taskQueueSid** | **string**| The SID of the TaskQueue for which to fetch statistics. | 
 **optional** | ***FetchTaskQueueStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchTaskQueueStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default is 15 minutes. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate real-time and cumulative statistics for the specified TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics**](taskrouter.v1.workspace.task_queue.task_queue_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskReservation

> TaskrouterV1WorkspaceTaskTaskReservation FetchTaskReservation(ctx, workspaceSid, taskSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskReservation resource to fetch. | 
**taskSid** | **string**| The SID of the reserved Task resource with the TaskReservation resource to fetch. | 
**sid** | **string**| The SID of the TaskReservation resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceTaskTaskReservation**](taskrouter.v1.workspace.task.task_reservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorker

> TaskrouterV1WorkspaceWorker FetchWorker(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Worker to fetch. | 
**sid** | **string**| The SID of the Worker resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceWorker**](taskrouter.v1.workspace.worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerChannel

> TaskrouterV1WorkspaceWorkerWorkerChannel FetchWorkerChannel(ctx, workspaceSid, workerSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerChannel to fetch. | 
**workerSid** | **string**| The SID of the Worker with the WorkerChannel to fetch. | 
**sid** | **string**| The SID of the WorkerChannel to fetch. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerChannel**](taskrouter.v1.workspace.worker.worker_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerInstanceStatistics

> TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics FetchWorkerInstanceStatistics(ctx, workspaceSid, workerSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerChannel to fetch. | 
**workerSid** | **string**| The SID of the Worker with the WorkerChannel to fetch. | 
 **optional** | ***FetchWorkerInstanceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkerInstanceStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **endDate** | **optional.Time**| Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **taskChannel** | **optional.String**| Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics**](taskrouter.v1.workspace.worker.worker_instance_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerReservation

> TaskrouterV1WorkspaceWorkerWorkerReservation FetchWorkerReservation(ctx, workspaceSid, workerSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerReservation resource to fetch. | 
**workerSid** | **string**| The SID of the reserved Worker resource with the WorkerReservation resource to fetch. | 
**sid** | **string**| The SID of the WorkerReservation resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerReservation**](taskrouter.v1.workspace.worker.worker_reservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerStatistics

> TaskrouterV1WorkspaceWorkerWorkerStatistics FetchWorkerStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Worker to fetch. | 
 **optional** | ***FetchWorkerStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkerStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **taskQueueSid** | **optional.String**| The SID of the TaskQueue for which to fetch Worker statistics. | 
 **taskQueueName** | **optional.String**| The &#x60;friendly_name&#x60; of the TaskQueue for which to fetch Worker statistics. | 
 **friendlyName** | **optional.String**| Only include Workers with &#x60;friendly_name&#x60; values that match this parameter. | 
 **taskChannel** | **optional.String**| Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerStatistics**](taskrouter.v1.workspace.worker.worker_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkersCumulativeStatistics

> TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics FetchWorkersCumulativeStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the resource to fetch. | 
 **optional** | ***FetchWorkersCumulativeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkersCumulativeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics**](taskrouter.v1.workspace.worker.workers_cumulative_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkersRealTimeStatistics

> TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics FetchWorkersRealTimeStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the resource to fetch. | 
 **optional** | ***FetchWorkersRealTimeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkersRealTimeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskChannel** | **optional.String**| Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics**](taskrouter.v1.workspace.worker.workers_real_time_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflow

> TaskrouterV1WorkspaceWorkflow FetchWorkflow(ctx, workspaceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to fetch. | 
**sid** | **string**| The SID of the Workflow resource to fetch. | 

### Return type

[**TaskrouterV1WorkspaceWorkflow**](taskrouter.v1.workspace.workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowCumulativeStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics FetchWorkflowCumulativeStatistics(ctx, workspaceSid, workflowSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the resource to fetch. | 
**workflowSid** | **string**| Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value. | 
 **optional** | ***FetchWorkflowCumulativeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkflowCumulativeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **optional.Time**| Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. | 

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics**](taskrouter.v1.workspace.workflow.workflow_cumulative_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowRealTimeStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics FetchWorkflowRealTimeStatistics(ctx, workspaceSid, workflowSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to fetch. | 
**workflowSid** | **string**| Returns the list of Tasks that are being controlled by the Workflow with the specified SID value. | 
 **optional** | ***FetchWorkflowRealTimeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkflowRealTimeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taskChannel** | **optional.String**| Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics**](taskrouter.v1.workspace.workflow.workflow_real_time_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowStatistics FetchWorkflowStatistics(ctx, workspaceSid, workflowSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to fetch. | 
**workflowSid** | **string**| Returns the list of Tasks that are being controlled by the Workflow with the specified SID value. | 
 **optional** | ***FetchWorkflowStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkflowStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **taskChannel** | **optional.String**| Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. | 

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowStatistics**](taskrouter.v1.workspace.workflow.workflow_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspace

> TaskrouterV1Workspace FetchWorkspace(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Workspace resource to fetch. | 

### Return type

[**TaskrouterV1Workspace**](taskrouter.v1.workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspaceCumulativeStatistics

> TaskrouterV1WorkspaceWorkspaceCumulativeStatistics FetchWorkspaceCumulativeStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace to fetch. | 
 **optional** | ***FetchWorkspaceCumulativeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkspaceCumulativeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.Time**| Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. | 

### Return type

[**TaskrouterV1WorkspaceWorkspaceCumulativeStatistics**](taskrouter.v1.workspace.workspace_cumulative_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspaceRealTimeStatistics

> TaskrouterV1WorkspaceWorkspaceRealTimeStatistics FetchWorkspaceRealTimeStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace to fetch. | 
 **optional** | ***FetchWorkspaceRealTimeStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkspaceRealTimeStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskChannel** | **optional.String**| Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkspaceRealTimeStatistics**](taskrouter.v1.workspace.workspace_real_time_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspaceStatistics

> TaskrouterV1WorkspaceWorkspaceStatistics FetchWorkspaceStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace to fetch. | 
 **optional** | ***FetchWorkspaceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchWorkspaceStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **taskChannel** | **optional.String**| Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. | 

### Return type

[**TaskrouterV1WorkspaceWorkspaceStatistics**](taskrouter.v1.workspace.workspace_statistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivity

> TaskrouterV1WorkspaceActivityReadResponse ListActivity(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Activity resources to read. | 
 **optional** | ***ListActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListActivityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the Activity resources to read. | 
 **available** | **optional.String**| Whether return only Activity resources that are available or unavailable. A value of &#x60;true&#x60; returns only available activities. Values of &#39;1&#39; or &#x60;yes&#x60; also indicate &#x60;true&#x60;. All other values represent &#x60;false&#x60; and return activities that are unavailable. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceActivityReadResponse**](taskrouter_v1_workspace_activityReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> TaskrouterV1WorkspaceEventReadResponse ListEvent(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Events to read. Returns only the Events that pertain to the specified Workspace. | 
 **optional** | ***ListEventOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.Time**| Only include Events that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **eventType** | **optional.String**| The type of Events to read. Returns only Events of the type specified. | 
 **minutes** | **optional.Int32**| The period of events to read in minutes. Returns only Events that occurred since this many minutes in the past. The default is &#x60;15&#x60; minutes. Task Attributes for Events occuring more 43,200 minutes ago will be redacted. | 
 **reservationSid** | **optional.String**| The SID of the Reservation with the Events to read. Returns only Events that pertain to the specified Reservation. | 
 **startDate** | **optional.Time**| Only include Events from on or after this date and time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Task Attributes for Events older than 30 days will be redacted. | 
 **taskQueueSid** | **optional.String**| The SID of the TaskQueue with the Events to read. Returns only the Events that pertain to the specified TaskQueue. | 
 **taskSid** | **optional.String**| The SID of the Task with the Events to read. Returns only the Events that pertain to the specified Task. | 
 **workerSid** | **optional.String**| The SID of the Worker with the Events to read. Returns only the Events that pertain to the specified Worker. | 
 **workflowSid** | **optional.String**| The SID of the Workflow with the Events to read. Returns only the Events that pertain to the specified Workflow. | 
 **taskChannel** | **optional.String**| The TaskChannel with the Events to read. Returns only the Events that pertain to the specified TaskChannel. | 
 **sid** | **optional.String**| The SID of the Event resource to read. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceEventReadResponse**](taskrouter_v1_workspace_eventReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> TaskrouterV1WorkspaceTaskReadResponse ListTask(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Tasks to read. | 
 **optional** | ***ListTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **priority** | **optional.Int32**| The priority value of the Tasks to read. Returns the list of all Tasks in the Workspace with the specified priority. | 
 **assignmentStatus** | [**optional.Interface of []string**](string.md)| The &#x60;assignment_status&#x60; of the Tasks you want to read. Can be: &#x60;pending&#x60;, &#x60;reserved&#x60;, &#x60;assigned&#x60;, &#x60;canceled&#x60;, &#x60;wrapping&#x60;, or &#x60;completed&#x60;. Returns all Tasks in the Workspace with the specified &#x60;assignment_status&#x60;. | 
 **workflowSid** | **optional.String**| The SID of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this SID. | 
 **workflowName** | **optional.String**| The friendly name of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this friendly name. | 
 **taskQueueSid** | **optional.String**| The SID of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this SID. | 
 **taskQueueName** | **optional.String**| The &#x60;friendly_name&#x60; of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this friendly name. | 
 **evaluateTaskAttributes** | **optional.String**| The attributes of the Tasks to read. Returns the Tasks that match the attributes specified in this parameter. | 
 **ordering** | **optional.String**| How to order the returned Task resources. y default, Tasks are sorted by ascending DateCreated. This value is specified as: &#x60;Attribute:Order&#x60;, where &#x60;Attribute&#x60; can be either &#x60;Priority&#x60; or &#x60;DateCreated&#x60; and &#x60;Order&#x60; can be either &#x60;asc&#x60; or &#x60;desc&#x60;. For example, &#x60;Priority:desc&#x60; returns Tasks ordered in descending order of their Priority. Multiple sort orders can be specified in a comma-separated list such as &#x60;Priority:desc,DateCreated:asc&#x60;, which returns the Tasks in descending Priority order and ascending DateCreated Order. | 
 **hasAddons** | **optional.Bool**| Whether to read Tasks with addons. If &#x60;true&#x60;, returns only Tasks with addons. If &#x60;false&#x60;, returns only Tasks without addons. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceTaskReadResponse**](taskrouter_v1_workspace_taskReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskChannel

> TaskrouterV1WorkspaceTaskChannelReadResponse ListTaskChannel(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task Channel to read. | 
 **optional** | ***ListTaskChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceTaskChannelReadResponse**](taskrouter_v1_workspace_task_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskQueue

> TaskrouterV1WorkspaceTaskQueueReadResponse ListTaskQueue(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to read. | 
 **optional** | ***ListTaskQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the TaskQueue resources to read. | 
 **evaluateWorkerAttributes** | **optional.String**| The attributes of the Workers to read. Returns the TaskQueues with Workers that match the attributes specified in this parameter. | 
 **workerSid** | **optional.String**| The SID of the Worker with the TaskQueue resources to read. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueueReadResponse**](taskrouter_v1_workspace_task_queueReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskQueuesStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueuesStatisticsReadResponse ListTaskQueuesStatistics(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueues to read. | 
 **optional** | ***ListTaskQueuesStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskQueuesStatisticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **optional.Time**| Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. | 
 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the TaskQueue statistics to read. | 
 **minutes** | **optional.Int32**| Only calculate statistics since this many minutes in the past. The default is 15 minutes. | 
 **startDate** | **optional.Time**| Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **taskChannel** | **optional.String**| Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 
 **splitByWaitTime** | **optional.String**| A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueuesStatisticsReadResponse**](taskrouter_v1_workspace_task_queue_task_queues_statisticsReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskReservation

> TaskrouterV1WorkspaceTaskTaskReservationReadResponse ListTaskReservation(ctx, workspaceSid, taskSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskReservation resources to read. | 
**taskSid** | **string**| The SID of the reserved Task resource with the TaskReservation resources to read. | 
 **optional** | ***ListTaskReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTaskReservationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reservationStatus** | **optional.String**| Returns the list of reservations for a task with a specified ReservationStatus.  Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, or &#x60;timeout&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceTaskTaskReservationReadResponse**](taskrouter_v1_workspace_task_task_reservationReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorker

> TaskrouterV1WorkspaceWorkerReadResponse ListWorker(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workers to read. | 
 **optional** | ***ListWorkerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityName** | **optional.String**| The &#x60;activity_name&#x60; of the Worker resources to read. | 
 **activitySid** | **optional.String**| The &#x60;activity_sid&#x60; of the Worker resources to read. | 
 **available** | **optional.String**| Whether to return only Worker resources that are available or unavailable. Can be &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; to return Worker resources that are available, and &#x60;false&#x60;, or any value returns the Worker resources that are not available. | 
 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the Worker resources to read. | 
 **targetWorkersExpression** | **optional.String**| Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue. | 
 **taskQueueName** | **optional.String**| The &#x60;friendly_name&#x60; of the TaskQueue that the Workers to read are eligible for. | 
 **taskQueueSid** | **optional.String**| The SID of the TaskQueue that the Workers to read are eligible for. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceWorkerReadResponse**](taskrouter_v1_workspace_workerReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkerChannel

> TaskrouterV1WorkspaceWorkerWorkerChannelReadResponse ListWorkerChannel(ctx, workspaceSid, workerSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerChannels to read. | 
**workerSid** | **string**| The SID of the Worker with the WorkerChannels to read. | 
 **optional** | ***ListWorkerChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkerChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerChannelReadResponse**](taskrouter_v1_workspace_worker_worker_channelReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkerReservation

> TaskrouterV1WorkspaceWorkerWorkerReservationReadResponse ListWorkerReservation(ctx, workspaceSid, workerSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerReservation resources to read. | 
**workerSid** | **string**| The SID of the reserved Worker resource with the WorkerReservation resources to read. | 
 **optional** | ***ListWorkerReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkerReservationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **reservationStatus** | **optional.String**| Returns the list of reservations for a worker with a specified ReservationStatus. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, &#x60;timeout&#x60;, &#x60;canceled&#x60;, or &#x60;rescinded&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerReservationReadResponse**](taskrouter_v1_workspace_worker_worker_reservationReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflow

> TaskrouterV1WorkspaceWorkflowReadResponse ListWorkflow(ctx, workspaceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to read. | 
 **optional** | ***ListWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the Workflow resources to read. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceWorkflowReadResponse**](taskrouter_v1_workspace_workflowReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspace

> TaskrouterV1WorkspaceReadResponse ListWorkspace(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWorkspaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**| The &#x60;friendly_name&#x60; of the Workspace resources to read. For example &#x60;Customer Support&#x60; or &#x60;2014 Election Campaign&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**TaskrouterV1WorkspaceReadResponse**](taskrouter_v1_workspaceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivity

> TaskrouterV1WorkspaceActivity UpdateActivity(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Activity resources to update. | 
**sid** | **string**| The SID of the Activity resource to update. | 
 **optional** | ***UpdateActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateActivityOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: &#x60;on-call&#x60;, &#x60;break&#x60;, and &#x60;email&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceActivity**](taskrouter.v1.workspace.activity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> TaskrouterV1WorkspaceTask UpdateTask(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task to update. | 
**sid** | **string**| The SID of the Task resource to update. | 
 **optional** | ***UpdateTaskOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignmentStatus** | **optional.String**| The new status of the task. Can be: &#x60;canceled&#x60;, to cancel a Task that is currently &#x60;pending&#x60; or &#x60;reserved&#x60;; &#x60;wrapping&#x60;, to move the Task to wrapup state; or &#x60;completed&#x60;, to move a Task to the completed state. | 
 **attributes** | **optional.String**| The JSON string that describes the custom attributes of the task. | 
 **priority** | **optional.Int32**| The Task&#39;s new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647). | 
 **reason** | **optional.String**| The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason. | 
 **taskChannel** | **optional.String**| When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceTask**](taskrouter.v1.workspace.task.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskChannel

> TaskrouterV1WorkspaceTaskChannel UpdateTaskChannel(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Task Channel to update. | 
**sid** | **string**| The SID of the Task Channel resource to update. | 
 **optional** | ***UpdateTaskChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **channelOptimizedRouting** | **optional.Bool**| Whether the TaskChannel should prioritize Workers that have been idle. If &#x60;true&#x60;, Workers that have been idle the longest are prioritized. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long. | 

### Return type

[**TaskrouterV1WorkspaceTaskChannel**](taskrouter.v1.workspace.task_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskQueue

> TaskrouterV1WorkspaceTaskQueue UpdateTaskQueue(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskQueue to update. | 
**sid** | **string**| The SID of the TaskQueue resource to update. | 
 **optional** | ***UpdateTaskQueueOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskQueueOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignmentActivitySid** | **optional.String**| The SID of the Activity to assign Workers when a task is assigned for them. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;. | 
 **maxReservedWorkers** | **optional.Int32**| The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50. | 
 **reservationActivitySid** | **optional.String**| The SID of the Activity to assign Workers when a task is reserved for them. | 
 **targetWorkers** | **optional.String**| A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example &#39;\\\&quot;language\\\&quot; &#x3D;&#x3D; \\\&quot;spanish\\\&quot;&#39; If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below. | 
 **taskOrder** | **optional.String**| How Tasks will be assigned to Workers. Can be: &#x60;FIFO&#x60; or &#x60;LIFO&#x60; and the default is &#x60;FIFO&#x60;. Use &#x60;FIFO&#x60; to assign the oldest task first and &#x60;LIFO&#x60; to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo). | 

### Return type

[**TaskrouterV1WorkspaceTaskQueue**](taskrouter.v1.workspace.task_queue.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskReservation

> TaskrouterV1WorkspaceTaskTaskReservation UpdateTaskReservation(ctx, workspaceSid, taskSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the TaskReservation resources to update. | 
**taskSid** | **string**| The SID of the reserved Task resource with the TaskReservation resources to update. | 
**sid** | **string**| The SID of the TaskReservation resource to update. | 
 **optional** | ***UpdateTaskReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTaskReservationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **beep** | **optional.String**| Whether to play a notification beep when the participant joins or when to play a beep. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;. | 
 **beepOnCustomerEntrance** | **optional.Bool**| Whether to play a notification beep when the customer joins. | 
 **callAccept** | **optional.Bool**| Whether to accept a reservation when executing a Call instruction. | 
 **callFrom** | **optional.String**| The Caller ID of the outbound call when executing a Call instruction. | 
 **callRecord** | **optional.String**| Whether to record both legs of a call when executing a Call instruction or which leg to record. | 
 **callStatusCallbackUrl** | **optional.String**| The URL to call  for the completed call event when executing a Call instruction. | 
 **callTimeout** | **optional.Int32**| Timeout for call when executing a Call instruction. | 
 **callTo** | **optional.String**| The Contact URI of the worker when executing a Call instruction.  Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **callUrl** | **optional.String**| TwiML URI executed on answering the worker&#39;s leg as a result of the Call instruction. | 
 **conferenceRecord** | **optional.String**| Whether to record the conference the participant is joining or when to record the conference. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;. | 
 **conferenceRecordingStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available. | 
 **conferenceRecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored. | 
 **conferenceStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The conference status events that we will send to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;. | 
 **conferenceStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceTrim** | **optional.String**| How to trim the leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;. | 
 **dequeueFrom** | **optional.String**| The Caller ID of the call to the worker when executing a Dequeue instruction. | 
 **dequeuePostWorkActivitySid** | **optional.String**| The SID of the Activity resource to start after executing a Dequeue instruction. | 
 **dequeueRecord** | **optional.String**| Whether to record both legs of a call when executing a Dequeue instruction or which leg to record. | 
 **dequeueStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The Call progress events sent via webhooks as a result of a Dequeue instruction. | 
 **dequeueStatusCallbackUrl** | **optional.String**| The Callback URL for completed call event when executing a Dequeue instruction. | 
 **dequeueTimeout** | **optional.Int32**| Timeout for call when executing a Dequeue instruction. | 
 **dequeueTo** | **optional.String**| The Contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **earlyMedia** | **optional.Bool**| Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is &#x60;true&#x60;. | 
 **endConferenceOnCustomerExit** | **optional.Bool**| Whether to end the conference when the customer leaves. | 
 **endConferenceOnExit** | **optional.Bool**| Whether to end the conference when the agent leaves. | 
 **from** | **optional.String**| The Caller ID of the call to the worker when executing a Conference instruction. | 
 **instruction** | **optional.String**| The assignment instruction for reservation. | 
 **maxParticipants** | **optional.Int32**| The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;. | 
 **muted** | **optional.Bool**| Whether the agent is muted in the conference. The default is &#x60;false&#x60;. | 
 **postWorkActivitySid** | **optional.String**| The new worker activity SID after executing a Conference instruction. | 
 **record** | **optional.Bool**| Whether to record the participant and their conferences, including the time between conferences. The default is &#x60;false&#x60;. | 
 **recordingChannels** | **optional.String**| The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. | 
 **recordingStatusCallback** | **optional.String**| The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes. | 
 **recordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **redirectAccept** | **optional.Bool**| Whether the reservation should be accepted when executing a Redirect instruction. | 
 **redirectCallSid** | **optional.String**| The Call SID of the call parked in the queue when executing a Redirect instruction. | 
 **redirectUrl** | **optional.String**| TwiML URI to redirect the call to when executing the Redirect instruction. | 
 **region** | **optional.String**| The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;. | 
 **reservationStatus** | **optional.String**| The new status of the reservation. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, or &#x60;timeout&#x60;. | 
 **sipAuthPassword** | **optional.String**| The SIP password for authentication. | 
 **sipAuthUsername** | **optional.String**| The SIP username used for authentication. | 
 **startConferenceOnEnter** | **optional.Bool**| Whether to start the conference when the participant joins, if it has not already started. The default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference. | 
 **statusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackEvent** | [**optional.Interface of []string**](string.md)| The call progress events that we will send to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, or &#x60;completed&#x60;. | 
 **statusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. | 
 **supervisor** | **optional.String**| The Supervisor SID/URI when executing the Supervise instruction. | 
 **supervisorMode** | **optional.String**| The Supervisor mode when executing the Supervise instruction. | 
 **timeout** | **optional.Int32**| Timeout for call when executing a Conference instruction. | 
 **to** | **optional.String**| The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **waitMethod** | **optional.String**| The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **waitUrl** | **optional.String**| The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 
 **workerActivitySid** | **optional.String**| The new worker activity SID if rejecting a reservation. | 

### Return type

[**TaskrouterV1WorkspaceTaskTaskReservation**](taskrouter.v1.workspace.task.task_reservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorker

> TaskrouterV1WorkspaceWorker UpdateWorker(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Worker to update. | 
**sid** | **string**| The SID of the Worker resource to update. | 
 **optional** | ***UpdateWorkerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activitySid** | **optional.String**| The SID of a valid Activity that will describe the Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. | 
 **attributes** | **optional.String**| The JSON string that describes the Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Worker. It can be up to 64 characters long. | 
 **rejectPendingReservations** | **optional.Bool**| Whether to reject pending reservations. | 

### Return type

[**TaskrouterV1WorkspaceWorker**](taskrouter.v1.workspace.worker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerChannel

> TaskrouterV1WorkspaceWorkerWorkerChannel UpdateWorkerChannel(ctx, workspaceSid, workerSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerChannel to update. | 
**workerSid** | **string**| The SID of the Worker with the WorkerChannel to update. | 
**sid** | **string**| The SID of the WorkerChannel to update. | 
 **optional** | ***UpdateWorkerChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkerChannelOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **available** | **optional.Bool**| Whether the WorkerChannel is available. Set to &#x60;false&#x60; to prevent the Worker from receiving any new Tasks of this TaskChannel type. | 
 **capacity** | **optional.Int32**| The total number of Tasks that the Worker should handle for the TaskChannel type. TaskRouter creates reservations for Tasks of this TaskChannel type up to the specified capacity. If the capacity is 0, no new reservations will be created. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerChannel**](taskrouter.v1.workspace.worker.worker_channel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerReservation

> TaskrouterV1WorkspaceWorkerWorkerReservation UpdateWorkerReservation(ctx, workspaceSid, workerSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the WorkerReservation resources to update. | 
**workerSid** | **string**| The SID of the reserved Worker resource with the WorkerReservation resources to update. | 
**sid** | **string**| The SID of the WorkerReservation resource to update. | 
 **optional** | ***UpdateWorkerReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkerReservationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **beep** | **optional.String**| Whether to play a notification beep when the participant joins or when to play a beep. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;. | 
 **beepOnCustomerEntrance** | **optional.Bool**| Whether to play a notification beep when the customer joins. | 
 **callAccept** | **optional.Bool**| Whether to accept a reservation when executing a Call instruction. | 
 **callFrom** | **optional.String**| The Caller ID of the outbound call when executing a Call instruction. | 
 **callRecord** | **optional.String**| Whether to record both legs of a call when executing a Call instruction. | 
 **callStatusCallbackUrl** | **optional.String**| The URL to call for the completed call event when executing a Call instruction. | 
 **callTimeout** | **optional.Int32**| The timeout for a call when executing a Call instruction. | 
 **callTo** | **optional.String**| The contact URI of the worker when executing a Call instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **callUrl** | **optional.String**| TwiML URI executed on answering the worker&#39;s leg as a result of the Call instruction. | 
 **conferenceRecord** | **optional.String**| Whether to record the conference the participant is joining or when to record the conference. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;. | 
 **conferenceRecordingStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available. | 
 **conferenceRecordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceStatusCallback** | **optional.String**| The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored. | 
 **conferenceStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The conference status events that we will send to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;. | 
 **conferenceStatusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **conferenceTrim** | **optional.String**| Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;. | 
 **dequeueFrom** | **optional.String**| The caller ID of the call to the worker when executing a Dequeue instruction. | 
 **dequeuePostWorkActivitySid** | **optional.String**| The SID of the Activity resource to start after executing a Dequeue instruction. | 
 **dequeueRecord** | **optional.String**| Whether to record both legs of a call when executing a Dequeue instruction or which leg to record. | 
 **dequeueStatusCallbackEvent** | [**optional.Interface of []string**](string.md)| The call progress events sent via webhooks as a result of a Dequeue instruction. | 
 **dequeueStatusCallbackUrl** | **optional.String**| The callback URL for completed call event when executing a Dequeue instruction. | 
 **dequeueTimeout** | **optional.Int32**| The timeout for call when executing a Dequeue instruction. | 
 **dequeueTo** | **optional.String**| The contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **earlyMedia** | **optional.Bool**| Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is &#x60;true&#x60;. | 
 **endConferenceOnCustomerExit** | **optional.Bool**| Whether to end the conference when the customer leaves. | 
 **endConferenceOnExit** | **optional.Bool**| Whether to end the conference when the agent leaves. | 
 **from** | **optional.String**| The caller ID of the call to the worker when executing a Conference instruction. | 
 **instruction** | **optional.String**| The assignment instruction for the reservation. | 
 **maxParticipants** | **optional.Int32**| The maximum number of participants allowed in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;. | 
 **muted** | **optional.Bool**| Whether the agent is muted in the conference. Defaults to &#x60;false&#x60;. | 
 **postWorkActivitySid** | **optional.String**| The new worker activity SID after executing a Conference instruction. | 
 **record** | **optional.Bool**| Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;. | 
 **recordingChannels** | **optional.String**| The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;. | 
 **recordingStatusCallback** | **optional.String**| The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes. | 
 **recordingStatusCallbackMethod** | **optional.String**| The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;. | 
 **redirectAccept** | **optional.Bool**| Whether the reservation should be accepted when executing a Redirect instruction. | 
 **redirectCallSid** | **optional.String**| The Call SID of the call parked in the queue when executing a Redirect instruction. | 
 **redirectUrl** | **optional.String**| TwiML URI to redirect the call to when executing the Redirect instruction. | 
 **region** | **optional.String**| The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;. | 
 **reservationStatus** | **optional.String**| The new status of the reservation. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, &#x60;timeout&#x60;, &#x60;canceled&#x60;, or &#x60;rescinded&#x60;. | 
 **sipAuthPassword** | **optional.String**| The SIP password for authentication. | 
 **sipAuthUsername** | **optional.String**| The SIP username used for authentication. | 
 **startConferenceOnEnter** | **optional.Bool**| Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference. | 
 **statusCallback** | **optional.String**| The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application. | 
 **statusCallbackEvent** | [**optional.Interface of []string**](string.md)| The call progress events that we will send to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, or &#x60;completed&#x60;. | 
 **statusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;. | 
 **timeout** | **optional.Int32**| The timeout for a call when executing a Conference instruction. | 
 **to** | **optional.String**| The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. | 
 **waitMethod** | **optional.String**| The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file. | 
 **waitUrl** | **optional.String**| The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). | 
 **workerActivitySid** | **optional.String**| The new worker activity SID if rejecting a reservation. | 

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerReservation**](taskrouter.v1.workspace.worker.worker_reservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> TaskrouterV1WorkspaceWorkflow UpdateWorkflow(ctx, workspaceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceSid** | **string**| The SID of the Workspace with the Workflow to update. | 
**sid** | **string**| The SID of the Workflow resource to update. | 
 **optional** | ***UpdateWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkflowOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignmentCallbackUrl** | **optional.String**| The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details. | 
 **configuration** | **optional.String**| A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. | 
 **fallbackAssignmentCallbackUrl** | **optional.String**| The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;. | 
 **reEvaluateTasks** | **optional.String**| Whether or not to re-evaluate Tasks. The default is &#x60;false&#x60;, which means Tasks in the Workflow will not be processed through the assignment loop again. | 
 **taskReservationTimeout** | **optional.Int32**| How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;. | 

### Return type

[**TaskrouterV1WorkspaceWorkflow**](taskrouter.v1.workspace.workflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspace

> TaskrouterV1Workspace UpdateWorkspace(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Workspace resource to update. | 
 **optional** | ***UpdateWorkspaceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWorkspaceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultActivitySid** | **optional.String**| The SID of the Activity that will be used when new Workers are created in the Workspace. | 
 **eventCallbackUrl** | **optional.String**| The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. | 
 **eventsFilter** | **optional.String**| The list of Workspace events for which to call event_callback_url. For example if &#x60;EventsFilter&#x3D;task.created,task.canceled,worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Workspace resource. For example: &#x60;Sales Call Center&#x60; or &#x60;Customer Support Team&#x60;. | 
 **multiTaskEnabled** | **optional.Bool**| Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. The default is &#x60;false&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking][https://www.twilio.com/docs/taskrouter/multitasking]. | 
 **prioritizeQueueOrder** | **optional.String**| The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: &#x60;LIFO&#x60; or &#x60;FIFO&#x60; and the default is &#x60;FIFO&#x60;. For more information, see [Queue Ordering][https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo]. | 
 **timeoutActivitySid** | **optional.String**| The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response. | 

### Return type

[**TaskrouterV1Workspace**](taskrouter.v1.workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

