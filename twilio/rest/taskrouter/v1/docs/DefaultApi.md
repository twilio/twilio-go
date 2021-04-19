# \DefaultApi

All URIs are relative to *https://taskrouter.twilio.com*

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

> TaskrouterV1WorkspaceActivity CreateActivity(ctx, WorkspaceSid).Available(Available).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new Activity belongs to.
    Available := true // bool | Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of `true`, `1`, or `yes` specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateActivity(context.Background(), WorkspaceSid).Available(Available).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActivity`: TaskrouterV1WorkspaceActivity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Activity belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateActivityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Available** | **bool** | Whether the Worker should be eligible to receive a Task when it occupies the Activity. A value of &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; specifies the Activity is available. All other values specify that it is not. The value cannot be changed after the Activity is created.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: &#x60;on-call&#x60;, &#x60;break&#x60;, and &#x60;email&#x60;.

### Return type

[**TaskrouterV1WorkspaceActivity**](TaskrouterV1WorkspaceActivity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> TaskrouterV1WorkspaceTask CreateTask(ctx, WorkspaceSid).Attributes(Attributes).Priority(Priority).TaskChannel(TaskChannel).Timeout(Timeout).WorkflowSid(WorkflowSid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new Task belongs to.
    Attributes := "Attributes_example" // string | A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow's `assignment_callback_url` when the Task is assigned to a Worker. For example: `{ \\\"task_type\\\": \\\"call\\\", \\\"twilio_call_sid\\\": \\\"CAxxx\\\", \\\"customer_ticket_number\\\": \\\"12345\\\" }`. (optional)
    Priority := int32(56) // int32 | The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647). (optional)
    TaskChannel := "TaskChannel_example" // string | When MultiTasking is enabled, specify the TaskChannel by passing either its `unique_name` or `sid`. Default value is `default`. (optional)
    Timeout := int32(56) // int32 | The amount of time in seconds the new task can live before being assigned. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the `task.canceled` event will fire with description `Task TTL Exceeded`. (optional)
    WorkflowSid := "WorkflowSid_example" // string | The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTask(context.Background(), WorkspaceSid).Attributes(Attributes).Priority(Priority).TaskChannel(TaskChannel).Timeout(Timeout).WorkflowSid(WorkflowSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTask`: TaskrouterV1WorkspaceTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Task belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Attributes** | **string** | A URL-encoded JSON string with the attributes of the new task. This value is passed to the Workflow&#39;s &#x60;assignment_callback_url&#x60; when the Task is assigned to a Worker. For example: &#x60;{ \\\&quot;task_type\\\&quot;: \\\&quot;call\\\&quot;, \\\&quot;twilio_call_sid\\\&quot;: \\\&quot;CAxxx\\\&quot;, \\\&quot;customer_ticket_number\\\&quot;: \\\&quot;12345\\\&quot; }&#x60;.
 **Priority** | **int32** | The priority to assign the new task and override the default. When supplied, the new Task will have this priority unless it matches a Workflow Target with a Priority set. When not supplied, the new Task will have the priority of the matching Workflow Target. Value can be 0 to 2^31^ (2,147,483,647).
 **TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel by passing either its &#x60;unique_name&#x60; or &#x60;sid&#x60;. Default value is &#x60;default&#x60;.
 **Timeout** | **int32** | The amount of time in seconds the new task can live before being assigned. Can be up to a maximum of 2 weeks (1,209,600 seconds). The default value is 24 hours (86,400 seconds). On timeout, the &#x60;task.canceled&#x60; event will fire with description &#x60;Task TTL Exceeded&#x60;.
 **WorkflowSid** | **string** | The SID of the Workflow that you would like to handle routing for the new Task. If there is only one Workflow defined for the Workspace that you are posting the new task to, this parameter is optional.

### Return type

[**TaskrouterV1WorkspaceTask**](TaskrouterV1WorkspaceTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskChannel

> TaskrouterV1WorkspaceTaskChannel CreateTaskChannel(ctx, WorkspaceSid).ChannelOptimizedRouting(ChannelOptimizedRouting).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new Task Channel belongs to.
    ChannelOptimizedRouting := true // bool | Whether the Task Channel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long. (optional)
    UniqueName := "UniqueName_example" // string | An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTaskChannel(context.Background(), WorkspaceSid).ChannelOptimizedRouting(ChannelOptimizedRouting).FriendlyName(FriendlyName).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTaskChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTaskChannel`: TaskrouterV1WorkspaceTaskChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTaskChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Task Channel belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ChannelOptimizedRouting** | **bool** | Whether the Task Channel should prioritize Workers that have been idle. If &#x60;true&#x60;, Workers that have been idle the longest are prioritized.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.
 **UniqueName** | **string** | An application-defined string that uniquely identifies the Task Channel, such as &#x60;voice&#x60; or &#x60;sms&#x60;.

### Return type

[**TaskrouterV1WorkspaceTaskChannel**](TaskrouterV1WorkspaceTaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskQueue

> TaskrouterV1WorkspaceTaskQueue CreateTaskQueue(ctx, WorkspaceSid).AssignmentActivitySid(AssignmentActivitySid).FriendlyName(FriendlyName).MaxReservedWorkers(MaxReservedWorkers).ReservationActivitySid(ReservationActivitySid).TargetWorkers(TargetWorkers).TaskOrder(TaskOrder).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new TaskQueue belongs to.
    AssignmentActivitySid := "AssignmentActivitySid_example" // string | The SID of the Activity to assign Workers when a task is assigned to them. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`. (optional)
    MaxReservedWorkers := int32(56) // int32 | The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1. (optional)
    ReservationActivitySid := "ReservationActivitySid_example" // string | The SID of the Activity to assign Workers when a task is reserved for them. (optional)
    TargetWorkers := "TargetWorkers_example" // string | A string that describes the Worker selection criteria for any Tasks that enter the TaskQueue. For example, `'\\\"language\\\" == \\\"spanish\\\"'`. The default value is `1==1`. If this value is empty, Tasks will wait in the TaskQueue until they are deleted or moved to another TaskQueue. For more information about Worker selection, see [Describing Worker selection criteria](https://www.twilio.com/docs/taskrouter/api/taskqueues#target-workers). (optional)
    TaskOrder := "TaskOrder_example" // string | How Tasks will be assigned to Workers. Set this parameter to `LIFO` to assign most recently created Task first or FIFO to assign the oldest Task first. Default is `FIFO`. [Click here](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo) to learn more. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTaskQueue(context.Background(), WorkspaceSid).AssignmentActivitySid(AssignmentActivitySid).FriendlyName(FriendlyName).MaxReservedWorkers(MaxReservedWorkers).ReservationActivitySid(ReservationActivitySid).TargetWorkers(TargetWorkers).TaskOrder(TaskOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTaskQueue`: TaskrouterV1WorkspaceTaskQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTaskQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new TaskQueue belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateTaskQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned to them.
 **FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;.
 **MaxReservedWorkers** | **int32** | The maximum number of Workers to reserve for the assignment of a Task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1.
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


## CreateWorker

> TaskrouterV1WorkspaceWorker CreateWorker(ctx, WorkspaceSid).ActivitySid(ActivitySid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new Worker belongs to.
    ActivitySid := "ActivitySid_example" // string | The SID of a valid Activity that will describe the new Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker's initial state is the `default_activity_sid` configured on the Workspace. (optional)
    Attributes := "Attributes_example" // string | A valid JSON string that describes the new Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the new Worker. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWorker(context.Background(), WorkspaceSid).ActivitySid(ActivitySid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorker`: TaskrouterV1WorkspaceWorker
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Worker belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateWorkerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ActivitySid** | **string** | The SID of a valid Activity that will describe the new Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. If not provided, the new Worker&#39;s initial state is the &#x60;default_activity_sid&#x60; configured on the Workspace.
 **Attributes** | **string** | A valid JSON string that describes the new Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}.
 **FriendlyName** | **string** | A descriptive string that you create to describe the new Worker. It can be up to 64 characters long.

### Return type

[**TaskrouterV1WorkspaceWorker**](TaskrouterV1WorkspaceWorker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflow

> TaskrouterV1WorkspaceWorkflow CreateWorkflow(ctx, WorkspaceSid).AssignmentCallbackUrl(AssignmentCallbackUrl).Configuration(Configuration).FallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl).FriendlyName(FriendlyName).TaskReservationTimeout(TaskReservationTimeout).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace that the new Workflow to create belongs to.
    AssignmentCallbackUrl := "AssignmentCallbackUrl_example" // string | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details. (optional)
    Configuration := "Configuration_example" // string | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. (optional)
    FallbackAssignmentCallbackUrl := "FallbackAssignmentCallbackUrl_example" // string | The URL that we should call when a call to the `assignment_callback_url` fails. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`. (optional)
    TaskReservationTimeout := int32(56) // int32 | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWorkflow(context.Background(), WorkspaceSid).AssignmentCallbackUrl(AssignmentCallbackUrl).Configuration(Configuration).FallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl).FriendlyName(FriendlyName).TaskReservationTimeout(TaskReservationTimeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflow`: TaskrouterV1WorkspaceWorkflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace that the new Workflow to create belongs to.

### Other Parameters

Other parameters are passed through a pointer to a CreateWorkflowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AssignmentCallbackUrl** | **string** | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
 **Configuration** | **string** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
 **FallbackAssignmentCallbackUrl** | **string** | The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;.
 **TaskReservationTimeout** | **int32** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkflow**](TaskrouterV1WorkspaceWorkflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> TaskrouterV1Workspace CreateWorkspace(ctx).EventCallbackUrl(EventCallbackUrl).EventsFilter(EventsFilter).FriendlyName(FriendlyName).MultiTaskEnabled(MultiTaskEnabled).PrioritizeQueueOrder(PrioritizeQueueOrder).Template(Template).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    EventCallbackUrl := "EventCallbackUrl_example" // string | The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. (optional)
    EventsFilter := "EventsFilter_example" // string | The list of Workspace events for which to call event_callback_url. For example, if `EventsFilter=task.created, task.canceled, worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: `Customer Support` or `2014 Election Campaign`. (optional)
    MultiTaskEnabled := true // bool | Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. The default is `false`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking). (optional)
    PrioritizeQueueOrder := "PrioritizeQueueOrder_example" // string | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: `LIFO` or `FIFO` and the default is `FIFO`. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo). (optional)
    Template := "Template_example" // string | An available template name. Can be: `NONE` or `FIFO` and the default is `NONE`. Pre-configures the Workspace with the Workflow and Activities specified in the template. `NONE` will create a Workspace with only a set of default activities. `FIFO` will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWorkspace(context.Background()).EventCallbackUrl(EventCallbackUrl).EventsFilter(EventsFilter).FriendlyName(FriendlyName).MultiTaskEnabled(MultiTaskEnabled).PrioritizeQueueOrder(PrioritizeQueueOrder).Template(Template).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: TaskrouterV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateWorkspaceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **EventCallbackUrl** | **string** | The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information.
 **EventsFilter** | **string** | The list of Workspace events for which to call event_callback_url. For example, if &#x60;EventsFilter&#x3D;task.created, task.canceled, worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: &#x60;Customer Support&#x60; or &#x60;2014 Election Campaign&#x60;.
 **MultiTaskEnabled** | **bool** | Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. The default is &#x60;false&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
 **PrioritizeQueueOrder** | **string** | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: &#x60;LIFO&#x60; or &#x60;FIFO&#x60; and the default is &#x60;FIFO&#x60;. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
 **Template** | **string** | An available template name. Can be: &#x60;NONE&#x60; or &#x60;FIFO&#x60; and the default is &#x60;NONE&#x60;. Pre-configures the Workspace with the Workflow and Activities specified in the template. &#x60;NONE&#x60; will create a Workspace with only a set of default activities. &#x60;FIFO&#x60; will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter.

### Return type

[**TaskrouterV1Workspace**](TaskrouterV1Workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteActivity

> DeleteActivity(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Activity resources to delete.
    Sid := "Sid_example" // string | The SID of the Activity resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteActivity(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to delete.
**Sid** | **string** | The SID of the Activity resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteActivityParams struct via the builder pattern


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


## DeleteTask

> DeleteTask(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task to delete.
    Sid := "Sid_example" // string | The SID of the Task resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTask(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to delete.
**Sid** | **string** | The SID of the Task resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskParams struct via the builder pattern


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


## DeleteTaskChannel

> DeleteTaskChannel(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task Channel to delete.
    Sid := "Sid_example" // string | The SID of the Task Channel resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTaskChannel(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTaskChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to delete.
**Sid** | **string** | The SID of the Task Channel resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskChannelParams struct via the builder pattern


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


## DeleteTaskQueue

> DeleteTaskQueue(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to delete.
    Sid := "Sid_example" // string | The SID of the TaskQueue resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTaskQueue(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to delete.
**Sid** | **string** | The SID of the TaskQueue resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTaskQueueParams struct via the builder pattern


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


## DeleteWorker

> DeleteWorker(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Worker to delete.
    Sid := "Sid_example" // string | The SID of the Worker resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWorker(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to delete.
**Sid** | **string** | The SID of the Worker resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWorkerParams struct via the builder pattern


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


## DeleteWorkflow

> DeleteWorkflow(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to delete.
    Sid := "Sid_example" // string | The SID of the Workflow resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWorkflow(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to delete.
**Sid** | **string** | The SID of the Workflow resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWorkflowParams struct via the builder pattern


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


## DeleteWorkspace

> DeleteWorkspace(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Workspace resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWorkspace(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Workspace resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWorkspaceParams struct via the builder pattern


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


## FetchActivity

> TaskrouterV1WorkspaceActivity FetchActivity(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Activity resources to fetch.
    Sid := "Sid_example" // string | The SID of the Activity resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchActivity(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchActivity`: TaskrouterV1WorkspaceActivity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to fetch.
**Sid** | **string** | The SID of the Activity resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchActivityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceActivity**](TaskrouterV1WorkspaceActivity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvent

> TaskrouterV1WorkspaceEvent FetchEvent(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Event to fetch.
    Sid := "Sid_example" // string | The SID of the Event resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEvent(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEvent`: TaskrouterV1WorkspaceEvent
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Event to fetch.
**Sid** | **string** | The SID of the Event resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceEvent**](TaskrouterV1WorkspaceEvent.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTask

> TaskrouterV1WorkspaceTask FetchTask(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task to fetch.
    Sid := "Sid_example" // string | The SID of the Task resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTask(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTask`: TaskrouterV1WorkspaceTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to fetch.
**Sid** | **string** | The SID of the Task resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceTask**](TaskrouterV1WorkspaceTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskChannel

> TaskrouterV1WorkspaceTaskChannel FetchTaskChannel(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task Channel to fetch.
    Sid := "Sid_example" // string | The SID of the Task Channel resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskChannel(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskChannel`: TaskrouterV1WorkspaceTaskChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to fetch.
**Sid** | **string** | The SID of the Task Channel resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceTaskChannel**](TaskrouterV1WorkspaceTaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueue

> TaskrouterV1WorkspaceTaskQueue FetchTaskQueue(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to fetch.
    Sid := "Sid_example" // string | The SID of the TaskQueue resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskQueue(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskQueue`: TaskrouterV1WorkspaceTaskQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**Sid** | **string** | The SID of the TaskQueue resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueParams struct via the builder pattern


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


## FetchTaskQueueCumulativeStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics FetchTaskQueueCumulativeStatistics(ctx, WorkspaceSid, TaskQueueSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to fetch.
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue for which to fetch statistics.
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default is 15 minutes. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskQueueCumulativeStatistics(context.Background(), WorkspaceSid, TaskQueueSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskQueueCumulativeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskQueueCumulativeStatistics`: TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskQueueCumulativeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueCumulativeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics**](TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueRealTimeStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics FetchTaskQueueRealTimeStatistics(ctx, WorkspaceSid, TaskQueueSid).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to fetch.
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue for which to fetch statistics.
    TaskChannel := "TaskChannel_example" // string | The TaskChannel for which to fetch statistics. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskQueueRealTimeStatistics(context.Background(), WorkspaceSid, TaskQueueSid).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskQueueRealTimeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskQueueRealTimeStatistics`: TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskQueueRealTimeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueRealTimeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **TaskChannel** | **string** | The TaskChannel for which to fetch statistics. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics**](TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskQueueStatistics

> TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics FetchTaskQueueStatistics(ctx, WorkspaceSid, TaskQueueSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to fetch.
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue for which to fetch statistics.
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default is 15 minutes. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate real-time and cumulative statistics for the specified TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskQueueStatistics(context.Background(), WorkspaceSid, TaskQueueSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskQueueStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskQueueStatistics`: TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskQueueStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to fetch.
**TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch statistics.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskQueueStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate real-time and cumulative statistics for the specified TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.

### Return type

[**TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics**](TaskrouterV1WorkspaceTaskQueueTaskQueueStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTaskReservation

> TaskrouterV1WorkspaceTaskTaskReservation FetchTaskReservation(ctx, WorkspaceSid, TaskSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskReservation resource to fetch.
    TaskSid := "TaskSid_example" // string | The SID of the reserved Task resource with the TaskReservation resource to fetch.
    Sid := "Sid_example" // string | The SID of the TaskReservation resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTaskReservation(context.Background(), WorkspaceSid, TaskSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTaskReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTaskReservation`: TaskrouterV1WorkspaceTaskTaskReservation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTaskReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resource to fetch.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resource to fetch.
**Sid** | **string** | The SID of the TaskReservation resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTaskReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**TaskrouterV1WorkspaceTaskTaskReservation**](TaskrouterV1WorkspaceTaskTaskReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorker

> TaskrouterV1WorkspaceWorker FetchWorker(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Worker to fetch.
    Sid := "Sid_example" // string | The SID of the Worker resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorker(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorker`: TaskrouterV1WorkspaceWorker
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to fetch.
**Sid** | **string** | The SID of the Worker resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceWorker**](TaskrouterV1WorkspaceWorker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerChannel

> TaskrouterV1WorkspaceWorkerWorkerChannel FetchWorkerChannel(ctx, WorkspaceSid, WorkerSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerChannel to fetch.
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the WorkerChannel to fetch.
    Sid := "Sid_example" // string | The SID of the WorkerChannel to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkerChannel(context.Background(), WorkspaceSid, WorkerSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkerChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkerChannel`: TaskrouterV1WorkspaceWorkerWorkerChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkerChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to fetch.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to fetch.
**Sid** | **string** | The SID of the WorkerChannel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**TaskrouterV1WorkspaceWorkerWorkerChannel**](TaskrouterV1WorkspaceWorkerWorkerChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerInstanceStatistics

> TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics FetchWorkerInstanceStatistics(ctx, WorkspaceSid, WorkerSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerChannel to fetch.
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the WorkerChannel to fetch.
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    EndDate := time.Now() // time.Time | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkerInstanceStatistics(context.Background(), WorkspaceSid, WorkerSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkerInstanceStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkerInstanceStatistics`: TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkerInstanceStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to fetch.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerInstanceStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **EndDate** | **time.Time** | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics**](TaskrouterV1WorkspaceWorkerWorkerInstanceStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerReservation

> TaskrouterV1WorkspaceWorkerWorkerReservation FetchWorkerReservation(ctx, WorkspaceSid, WorkerSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerReservation resource to fetch.
    WorkerSid := "WorkerSid_example" // string | The SID of the reserved Worker resource with the WorkerReservation resource to fetch.
    Sid := "Sid_example" // string | The SID of the WorkerReservation resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkerReservation(context.Background(), WorkspaceSid, WorkerSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkerReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkerReservation`: TaskrouterV1WorkspaceWorkerWorkerReservation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkerReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resource to fetch.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resource to fetch.
**Sid** | **string** | The SID of the WorkerReservation resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




### Return type

[**TaskrouterV1WorkspaceWorkerWorkerReservation**](TaskrouterV1WorkspaceWorkerWorkerReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkerStatistics

> TaskrouterV1WorkspaceWorkerWorkerStatistics FetchWorkerStatistics(ctx, WorkspaceSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskQueueSid(TaskQueueSid).TaskQueueName(TaskQueueName).FriendlyName(FriendlyName).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Worker to fetch.
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue for which to fetch Worker statistics. (optional)
    TaskQueueName := "TaskQueueName_example" // string | The `friendly_name` of the TaskQueue for which to fetch Worker statistics. (optional)
    FriendlyName := "FriendlyName_example" // string | Only include Workers with `friendly_name` values that match this parameter. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkerStatistics(context.Background(), WorkspaceSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskQueueSid(TaskQueueSid).TaskQueueName(TaskQueueName).FriendlyName(FriendlyName).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkerStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkerStatistics`: TaskrouterV1WorkspaceWorkerWorkerStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkerStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkerStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **TaskQueueSid** | **string** | The SID of the TaskQueue for which to fetch Worker statistics.
 **TaskQueueName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue for which to fetch Worker statistics.
 **FriendlyName** | **string** | Only include Workers with &#x60;friendly_name&#x60; values that match this parameter.
 **TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerStatistics**](TaskrouterV1WorkspaceWorkerWorkerStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkersCumulativeStatistics

> TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics FetchWorkersCumulativeStatistics(ctx, WorkspaceSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the resource to fetch.
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkersCumulativeStatistics(context.Background(), WorkspaceSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkersCumulativeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkersCumulativeStatistics`: TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkersCumulativeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkersCumulativeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics**](TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkersRealTimeStatistics

> TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics FetchWorkersRealTimeStatistics(ctx, WorkspaceSid).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the resource to fetch.
    TaskChannel := "TaskChannel_example" // string | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkersRealTimeStatistics(context.Background(), WorkspaceSid).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkersRealTimeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkersRealTimeStatistics`: TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkersRealTimeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkersRealTimeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **TaskChannel** | **string** | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics**](TaskrouterV1WorkspaceWorkerWorkersRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflow

> TaskrouterV1WorkspaceWorkflow FetchWorkflow(ctx, WorkspaceSid, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to fetch.
    Sid := "Sid_example" // string | The SID of the Workflow resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkflow(context.Background(), WorkspaceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkflow`: TaskrouterV1WorkspaceWorkflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to fetch.
**Sid** | **string** | The SID of the Workflow resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TaskrouterV1WorkspaceWorkflow**](TaskrouterV1WorkspaceWorkflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowCumulativeStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics FetchWorkflowCumulativeStatistics(ctx, WorkspaceSid, WorkflowSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the resource to fetch.
    WorkflowSid := "WorkflowSid_example" // string | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value.
    EndDate := time.Now() // time.Time | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, `5,30` would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkflowCumulativeStatistics(context.Background(), WorkspaceSid, WorkflowSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkflowCumulativeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkflowCumulativeStatistics`: TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkflowCumulativeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the resource to fetch.
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowCumulativeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **EndDate** | **time.Time** | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA.

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics**](TaskrouterV1WorkspaceWorkflowWorkflowCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowRealTimeStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics FetchWorkflowRealTimeStatistics(ctx, WorkspaceSid, WorkflowSid).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to fetch.
    WorkflowSid := "WorkflowSid_example" // string | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.
    TaskChannel := "TaskChannel_example" // string | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkflowRealTimeStatistics(context.Background(), WorkspaceSid, WorkflowSid).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkflowRealTimeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkflowRealTimeStatistics`: TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkflowRealTimeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to fetch.
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowRealTimeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **TaskChannel** | **string** | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics**](TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkflowStatistics

> TaskrouterV1WorkspaceWorkflowWorkflowStatistics FetchWorkflowStatistics(ctx, WorkspaceSid, WorkflowSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to fetch.
    WorkflowSid := "WorkflowSid_example" // string | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, `5,30` would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkflowStatistics(context.Background(), WorkspaceSid, WorkflowSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkflowStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkflowStatistics`: TaskrouterV1WorkspaceWorkflowWorkflowStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkflowStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to fetch.
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkflowStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **TaskChannel** | **string** | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA.

### Return type

[**TaskrouterV1WorkspaceWorkflowWorkflowStatistics**](TaskrouterV1WorkspaceWorkflowWorkflowStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspace

> TaskrouterV1Workspace FetchWorkspace(ctx, Sid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Workspace resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkspace(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkspace`: TaskrouterV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Workspace resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**TaskrouterV1Workspace**](TaskrouterV1Workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspaceCumulativeStatistics

> TaskrouterV1WorkspaceWorkspaceCumulativeStatistics FetchWorkspaceCumulativeStatistics(ctx, WorkspaceSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace to fetch.
    EndDate := time.Now() // time.Time | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, `5,30` would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkspaceCumulativeStatistics(context.Background(), WorkspaceSid).EndDate(EndDate).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkspaceCumulativeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkspaceCumulativeStatistics`: TaskrouterV1WorkspaceWorkspaceCumulativeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkspaceCumulativeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceCumulativeStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EndDate** | **time.Time** | Only include usage that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate cumulative statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA.

### Return type

[**TaskrouterV1WorkspaceWorkspaceCumulativeStatistics**](TaskrouterV1WorkspaceWorkspaceCumulativeStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWorkspaceRealTimeStatistics

> TaskrouterV1WorkspaceWorkspaceRealTimeStatistics FetchWorkspaceRealTimeStatistics(ctx, WorkspaceSid).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace to fetch.
    TaskChannel := "TaskChannel_example" // string | Only calculate real-time statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkspaceRealTimeStatistics(context.Background(), WorkspaceSid).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkspaceRealTimeStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkspaceRealTimeStatistics`: TaskrouterV1WorkspaceWorkspaceRealTimeStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkspaceRealTimeStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceRealTimeStatisticsParams struct via the builder pattern


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


## FetchWorkspaceStatistics

> TaskrouterV1WorkspaceWorkspaceStatistics FetchWorkspaceStatistics(ctx, WorkspaceSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace to fetch.
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, `5,30` would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWorkspaceStatistics(context.Background(), WorkspaceSid).Minutes(Minutes).StartDate(StartDate).EndDate(EndDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWorkspaceStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWorkspaceStatistics`: TaskrouterV1WorkspaceWorkspaceStatistics
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWorkspaceStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWorkspaceStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default 15 minutes. This is helpful for displaying statistics for the last 15 minutes, 240 minutes (4 hours), and 480 minutes (8 hours) to see trends.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. For example, &#x60;5,30&#x60; would show splits of Tasks that were canceled or accepted before and after 5 seconds and before and after 30 seconds. This can be used to show short abandoned Tasks or Tasks that failed to meet an SLA.

### Return type

[**TaskrouterV1WorkspaceWorkspaceStatistics**](TaskrouterV1WorkspaceWorkspaceStatistics.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListActivity

> ListActivityResponse ListActivity(ctx, WorkspaceSid).FriendlyName(FriendlyName).Available(Available).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Activity resources to read.
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the Activity resources to read. (optional)
    Available := "Available_example" // string | Whether return only Activity resources that are available or unavailable. A value of `true` returns only available activities. Values of '1' or `yes` also indicate `true`. All other values represent `false` and return activities that are unavailable. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListActivity(context.Background(), WorkspaceSid).FriendlyName(FriendlyName).Available(Available).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActivity`: ListActivityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListActivityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Activity resources to read.
 **Available** | **string** | Whether return only Activity resources that are available or unavailable. A value of &#x60;true&#x60; returns only available activities. Values of &#39;1&#39; or &#x60;yes&#x60; also indicate &#x60;true&#x60;. All other values represent &#x60;false&#x60; and return activities that are unavailable.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListActivityResponse**](ListActivityResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvent

> ListEventResponse ListEvent(ctx, WorkspaceSid).EndDate(EndDate).EventType(EventType).Minutes(Minutes).ReservationSid(ReservationSid).StartDate(StartDate).TaskQueueSid(TaskQueueSid).TaskSid(TaskSid).WorkerSid(WorkerSid).WorkflowSid(WorkflowSid).TaskChannel(TaskChannel).Sid(Sid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Events to read. Returns only the Events that pertain to the specified Workspace.
    EndDate := time.Now() // time.Time | Only include Events that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    EventType := "EventType_example" // string | The type of Events to read. Returns only Events of the type specified. (optional)
    Minutes := int32(56) // int32 | The period of events to read in minutes. Returns only Events that occurred since this many minutes in the past. The default is `15` minutes. Task Attributes for Events occuring more 43,200 minutes ago will be redacted. (optional)
    ReservationSid := "ReservationSid_example" // string | The SID of the Reservation with the Events to read. Returns only Events that pertain to the specified Reservation. (optional)
    StartDate := time.Now() // time.Time | Only include Events from on or after this date and time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Task Attributes for Events older than 30 days will be redacted. (optional)
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue with the Events to read. Returns only the Events that pertain to the specified TaskQueue. (optional)
    TaskSid := "TaskSid_example" // string | The SID of the Task with the Events to read. Returns only the Events that pertain to the specified Task. (optional)
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the Events to read. Returns only the Events that pertain to the specified Worker. (optional)
    WorkflowSid := "WorkflowSid_example" // string | The SID of the Workflow with the Events to read. Returns only the Events that pertain to the specified Workflow. (optional)
    TaskChannel := "TaskChannel_example" // string | The TaskChannel with the Events to read. Returns only the Events that pertain to the specified TaskChannel. (optional)
    Sid := "Sid_example" // string | The SID of the Event resource to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEvent(context.Background(), WorkspaceSid).EndDate(EndDate).EventType(EventType).Minutes(Minutes).ReservationSid(ReservationSid).StartDate(StartDate).TaskQueueSid(TaskQueueSid).TaskSid(TaskSid).WorkerSid(WorkerSid).WorkflowSid(WorkflowSid).TaskChannel(TaskChannel).Sid(Sid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvent`: ListEventResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Events to read. Returns only the Events that pertain to the specified Workspace.

### Other Parameters

Other parameters are passed through a pointer to a ListEventParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EndDate** | **time.Time** | Only include Events that occurred on or before this date, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **EventType** | **string** | The type of Events to read. Returns only Events of the type specified.
 **Minutes** | **int32** | The period of events to read in minutes. Returns only Events that occurred since this many minutes in the past. The default is &#x60;15&#x60; minutes. Task Attributes for Events occuring more 43,200 minutes ago will be redacted.
 **ReservationSid** | **string** | The SID of the Reservation with the Events to read. Returns only Events that pertain to the specified Reservation.
 **StartDate** | **time.Time** | Only include Events from on or after this date and time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Task Attributes for Events older than 30 days will be redacted.
 **TaskQueueSid** | **string** | The SID of the TaskQueue with the Events to read. Returns only the Events that pertain to the specified TaskQueue.
 **TaskSid** | **string** | The SID of the Task with the Events to read. Returns only the Events that pertain to the specified Task.
 **WorkerSid** | **string** | The SID of the Worker with the Events to read. Returns only the Events that pertain to the specified Worker.
 **WorkflowSid** | **string** | The SID of the Workflow with the Events to read. Returns only the Events that pertain to the specified Workflow.
 **TaskChannel** | **string** | The TaskChannel with the Events to read. Returns only the Events that pertain to the specified TaskChannel.
 **Sid** | **string** | The SID of the Event resource to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEventResponse**](ListEventResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTask

> ListTaskResponse ListTask(ctx, WorkspaceSid).Priority(Priority).AssignmentStatus(AssignmentStatus).WorkflowSid(WorkflowSid).WorkflowName(WorkflowName).TaskQueueSid(TaskQueueSid).TaskQueueName(TaskQueueName).EvaluateTaskAttributes(EvaluateTaskAttributes).Ordering(Ordering).HasAddons(HasAddons).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Tasks to read.
    Priority := int32(56) // int32 | The priority value of the Tasks to read. Returns the list of all Tasks in the Workspace with the specified priority. (optional)
    AssignmentStatus := []string{"Inner_example"} // []string | The `assignment_status` of the Tasks you want to read. Can be: `pending`, `reserved`, `assigned`, `canceled`, `wrapping`, or `completed`. Returns all Tasks in the Workspace with the specified `assignment_status`. (optional)
    WorkflowSid := "WorkflowSid_example" // string | The SID of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this SID. (optional)
    WorkflowName := "WorkflowName_example" // string | The friendly name of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this friendly name. (optional)
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this SID. (optional)
    TaskQueueName := "TaskQueueName_example" // string | The `friendly_name` of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this friendly name. (optional)
    EvaluateTaskAttributes := "EvaluateTaskAttributes_example" // string | The attributes of the Tasks to read. Returns the Tasks that match the attributes specified in this parameter. (optional)
    Ordering := "Ordering_example" // string | How to order the returned Task resources. y default, Tasks are sorted by ascending DateCreated. This value is specified as: `Attribute:Order`, where `Attribute` can be either `Priority` or `DateCreated` and `Order` can be either `asc` or `desc`. For example, `Priority:desc` returns Tasks ordered in descending order of their Priority. Multiple sort orders can be specified in a comma-separated list such as `Priority:desc,DateCreated:asc`, which returns the Tasks in descending Priority order and ascending DateCreated Order. (optional)
    HasAddons := true // bool | Whether to read Tasks with addons. If `true`, returns only Tasks with addons. If `false`, returns only Tasks without addons. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTask(context.Background(), WorkspaceSid).Priority(Priority).AssignmentStatus(AssignmentStatus).WorkflowSid(WorkflowSid).WorkflowName(WorkflowName).TaskQueueSid(TaskQueueSid).TaskQueueName(TaskQueueName).EvaluateTaskAttributes(EvaluateTaskAttributes).Ordering(Ordering).HasAddons(HasAddons).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTask`: ListTaskResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Tasks to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Priority** | **int32** | The priority value of the Tasks to read. Returns the list of all Tasks in the Workspace with the specified priority.
 **AssignmentStatus** | **[]string** | The &#x60;assignment_status&#x60; of the Tasks you want to read. Can be: &#x60;pending&#x60;, &#x60;reserved&#x60;, &#x60;assigned&#x60;, &#x60;canceled&#x60;, &#x60;wrapping&#x60;, or &#x60;completed&#x60;. Returns all Tasks in the Workspace with the specified &#x60;assignment_status&#x60;.
 **WorkflowSid** | **string** | The SID of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this SID.
 **WorkflowName** | **string** | The friendly name of the Workflow with the Tasks to read. Returns the Tasks controlled by the Workflow identified by this friendly name.
 **TaskQueueSid** | **string** | The SID of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this SID.
 **TaskQueueName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue with the Tasks to read. Returns the Tasks waiting in the TaskQueue identified by this friendly name.
 **EvaluateTaskAttributes** | **string** | The attributes of the Tasks to read. Returns the Tasks that match the attributes specified in this parameter.
 **Ordering** | **string** | How to order the returned Task resources. y default, Tasks are sorted by ascending DateCreated. This value is specified as: &#x60;Attribute:Order&#x60;, where &#x60;Attribute&#x60; can be either &#x60;Priority&#x60; or &#x60;DateCreated&#x60; and &#x60;Order&#x60; can be either &#x60;asc&#x60; or &#x60;desc&#x60;. For example, &#x60;Priority:desc&#x60; returns Tasks ordered in descending order of their Priority. Multiple sort orders can be specified in a comma-separated list such as &#x60;Priority:desc,DateCreated:asc&#x60;, which returns the Tasks in descending Priority order and ascending DateCreated Order.
 **HasAddons** | **bool** | Whether to read Tasks with addons. If &#x60;true&#x60;, returns only Tasks with addons. If &#x60;false&#x60;, returns only Tasks without addons.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTaskResponse**](ListTaskResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskChannel

> ListTaskChannelResponse ListTaskChannel(ctx, WorkspaceSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task Channel to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTaskChannel(context.Background(), WorkspaceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTaskChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskChannel`: ListTaskChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTaskChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTaskChannelResponse**](ListTaskChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskQueue

> ListTaskQueueResponse ListTaskQueue(ctx, WorkspaceSid).FriendlyName(FriendlyName).EvaluateWorkerAttributes(EvaluateWorkerAttributes).WorkerSid(WorkerSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to read.
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the TaskQueue resources to read. (optional)
    EvaluateWorkerAttributes := "EvaluateWorkerAttributes_example" // string | The attributes of the Workers to read. Returns the TaskQueues with Workers that match the attributes specified in this parameter. (optional)
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the TaskQueue resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTaskQueue(context.Background(), WorkspaceSid).FriendlyName(FriendlyName).EvaluateWorkerAttributes(EvaluateWorkerAttributes).WorkerSid(WorkerSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskQueue`: ListTaskQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTaskQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue resources to read.
 **EvaluateWorkerAttributes** | **string** | The attributes of the Workers to read. Returns the TaskQueues with Workers that match the attributes specified in this parameter.
 **WorkerSid** | **string** | The SID of the Worker with the TaskQueue resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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


## ListTaskQueuesStatistics

> ListTaskQueuesStatisticsResponse ListTaskQueuesStatistics(ctx, WorkspaceSid).EndDate(EndDate).FriendlyName(FriendlyName).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueues to read.
    EndDate := time.Now() // time.Time | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time. (optional)
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the TaskQueue statistics to read. (optional)
    Minutes := int32(56) // int32 | Only calculate statistics since this many minutes in the past. The default is 15 minutes. (optional)
    StartDate := time.Now() // time.Time | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. (optional)
    TaskChannel := "TaskChannel_example" // string | Only calculate statistics on this TaskChannel. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)
    SplitByWaitTime := "SplitByWaitTime_example" // string | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTaskQueuesStatistics(context.Background(), WorkspaceSid).EndDate(EndDate).FriendlyName(FriendlyName).Minutes(Minutes).StartDate(StartDate).TaskChannel(TaskChannel).SplitByWaitTime(SplitByWaitTime).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTaskQueuesStatistics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskQueuesStatistics`: ListTaskQueuesStatisticsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTaskQueuesStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueues to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskQueuesStatisticsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EndDate** | **time.Time** | Only calculate statistics from this date and time and earlier, specified in GMT as an [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time.
 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue statistics to read.
 **Minutes** | **int32** | Only calculate statistics since this many minutes in the past. The default is 15 minutes.
 **StartDate** | **time.Time** | Only calculate statistics from this date and time and later, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
 **TaskChannel** | **string** | Only calculate statistics on this TaskChannel. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.
 **SplitByWaitTime** | **string** | A comma separated list of values that describes the thresholds, in seconds, to calculate statistics on. For each threshold specified, the number of Tasks canceled and reservations accepted above and below the specified thresholds in seconds are computed.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTaskQueuesStatisticsResponse**](ListTaskQueuesStatisticsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaskReservation

> ListTaskReservationResponse ListTaskReservation(ctx, WorkspaceSid, TaskSid).ReservationStatus(ReservationStatus).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskReservation resources to read.
    TaskSid := "TaskSid_example" // string | The SID of the reserved Task resource with the TaskReservation resources to read.
    ReservationStatus := "ReservationStatus_example" // string | Returns the list of reservations for a task with a specified ReservationStatus.  Can be: `pending`, `accepted`, `rejected`, or `timeout`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListTaskReservation(context.Background(), WorkspaceSid, TaskSid).ReservationStatus(ReservationStatus).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTaskReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaskReservation`: ListTaskReservationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTaskReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resources to read.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListTaskReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ReservationStatus** | **string** | Returns the list of reservations for a task with a specified ReservationStatus.  Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, or &#x60;timeout&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTaskReservationResponse**](ListTaskReservationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorker

> ListWorkerResponse ListWorker(ctx, WorkspaceSid).ActivityName(ActivityName).ActivitySid(ActivitySid).Available(Available).FriendlyName(FriendlyName).TargetWorkersExpression(TargetWorkersExpression).TaskQueueName(TaskQueueName).TaskQueueSid(TaskQueueSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workers to read.
    ActivityName := "ActivityName_example" // string | The `activity_name` of the Worker resources to read. (optional)
    ActivitySid := "ActivitySid_example" // string | The `activity_sid` of the Worker resources to read. (optional)
    Available := "Available_example" // string | Whether to return only Worker resources that are available or unavailable. Can be `true`, `1`, or `yes` to return Worker resources that are available, and `false`, or any value returns the Worker resources that are not available. (optional)
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the Worker resources to read. (optional)
    TargetWorkersExpression := "TargetWorkersExpression_example" // string | Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue. (optional)
    TaskQueueName := "TaskQueueName_example" // string | The `friendly_name` of the TaskQueue that the Workers to read are eligible for. (optional)
    TaskQueueSid := "TaskQueueSid_example" // string | The SID of the TaskQueue that the Workers to read are eligible for. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorker(context.Background(), WorkspaceSid).ActivityName(ActivityName).ActivitySid(ActivitySid).Available(Available).FriendlyName(FriendlyName).TargetWorkersExpression(TargetWorkersExpression).TaskQueueName(TaskQueueName).TaskQueueSid(TaskQueueSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorker`: ListWorkerResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workers to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ActivityName** | **string** | The &#x60;activity_name&#x60; of the Worker resources to read.
 **ActivitySid** | **string** | The &#x60;activity_sid&#x60; of the Worker resources to read.
 **Available** | **string** | Whether to return only Worker resources that are available or unavailable. Can be &#x60;true&#x60;, &#x60;1&#x60;, or &#x60;yes&#x60; to return Worker resources that are available, and &#x60;false&#x60;, or any value returns the Worker resources that are not available.
 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Worker resources to read.
 **TargetWorkersExpression** | **string** | Filter by Workers that would match an expression on a TaskQueue. This is helpful for debugging which Workers would match a potential queue.
 **TaskQueueName** | **string** | The &#x60;friendly_name&#x60; of the TaskQueue that the Workers to read are eligible for.
 **TaskQueueSid** | **string** | The SID of the TaskQueue that the Workers to read are eligible for.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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


## ListWorkerChannel

> ListWorkerChannelResponse ListWorkerChannel(ctx, WorkspaceSid, WorkerSid).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerChannels to read.
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the WorkerChannels to read.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorkerChannel(context.Background(), WorkspaceSid, WorkerSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkerChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkerChannel`: ListWorkerChannelResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkerChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannels to read.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannels to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWorkerChannelResponse**](ListWorkerChannelResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkerReservation

> ListWorkerReservationResponse ListWorkerReservation(ctx, WorkspaceSid, WorkerSid).ReservationStatus(ReservationStatus).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerReservation resources to read.
    WorkerSid := "WorkerSid_example" // string | The SID of the reserved Worker resource with the WorkerReservation resources to read.
    ReservationStatus := "ReservationStatus_example" // string | Returns the list of reservations for a worker with a specified ReservationStatus. Can be: `pending`, `accepted`, `rejected`, `timeout`, `canceled`, or `rescinded`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorkerReservation(context.Background(), WorkspaceSid, WorkerSid).ReservationStatus(ReservationStatus).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkerReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkerReservation`: ListWorkerReservationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkerReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resources to read.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resources to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkerReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ReservationStatus** | **string** | Returns the list of reservations for a worker with a specified ReservationStatus. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, &#x60;timeout&#x60;, &#x60;canceled&#x60;, or &#x60;rescinded&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWorkerReservationResponse**](ListWorkerReservationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflow

> ListWorkflowResponse ListWorkflow(ctx, WorkspaceSid).FriendlyName(FriendlyName).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to read.
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the Workflow resources to read. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorkflow(context.Background(), WorkspaceSid).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflow`: ListWorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to read.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkflowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Workflow resources to read.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWorkflowResponse**](ListWorkflowResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspace

> ListWorkspaceResponse ListWorkspace(ctx).FriendlyName(FriendlyName).PageSize(PageSize).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    FriendlyName := "FriendlyName_example" // string | The `friendly_name` of the Workspace resources to read. For example `Customer Support` or `2014 Election Campaign`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorkspace(context.Background()).FriendlyName(FriendlyName).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspace`: ListWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkspace`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListWorkspaceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | The &#x60;friendly_name&#x60; of the Workspace resources to read. For example &#x60;Customer Support&#x60; or &#x60;2014 Election Campaign&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListWorkspaceResponse**](ListWorkspaceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivity

> TaskrouterV1WorkspaceActivity UpdateActivity(ctx, WorkspaceSid, Sid).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Activity resources to update.
    Sid := "Sid_example" // string | The SID of the Activity resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: `on-call`, `break`, and `email`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateActivity(context.Background(), WorkspaceSid, Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateActivity`: TaskrouterV1WorkspaceActivity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Activity resources to update.
**Sid** | **string** | The SID of the Activity resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateActivityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FriendlyName** | **string** | A descriptive string that you create to describe the Activity resource. It can be up to 64 characters long. These names are used to calculate and expose statistics about Workers, and provide visibility into the state of each Worker. Examples of friendly names include: &#x60;on-call&#x60;, &#x60;break&#x60;, and &#x60;email&#x60;.

### Return type

[**TaskrouterV1WorkspaceActivity**](TaskrouterV1WorkspaceActivity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> TaskrouterV1WorkspaceTask UpdateTask(ctx, WorkspaceSid, Sid).IfMatch(IfMatch).AssignmentStatus(AssignmentStatus).Attributes(Attributes).Priority(Priority).Reason(Reason).TaskChannel(TaskChannel).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task to update.
    Sid := "Sid_example" // string | The SID of the Task resource to update.
    IfMatch := "IfMatch_example" // string | If provided, applies this mutation if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match). (optional)
    AssignmentStatus := "AssignmentStatus_example" // string | The new status of the task. Can be: `canceled`, to cancel a Task that is currently `pending` or `reserved`; `wrapping`, to move the Task to wrapup state; or `completed`, to move a Task to the completed state. (optional)
    Attributes := "Attributes_example" // string | The JSON string that describes the custom attributes of the task. (optional)
    Priority := int32(56) // int32 | The Task's new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647). (optional)
    Reason := "Reason_example" // string | The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason. (optional)
    TaskChannel := "TaskChannel_example" // string | When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel's SID or its `unique_name`, such as `voice`, `sms`, or `default`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTask(context.Background(), WorkspaceSid, Sid).IfMatch(IfMatch).AssignmentStatus(AssignmentStatus).Attributes(Attributes).Priority(Priority).Reason(Reason).TaskChannel(TaskChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: TaskrouterV1WorkspaceTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task to update.
**Sid** | **string** | The SID of the Task resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **IfMatch** | **string** | If provided, applies this mutation if (and only if) the [ETag](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag) header of the Task matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
 **AssignmentStatus** | **string** | The new status of the task. Can be: &#x60;canceled&#x60;, to cancel a Task that is currently &#x60;pending&#x60; or &#x60;reserved&#x60;; &#x60;wrapping&#x60;, to move the Task to wrapup state; or &#x60;completed&#x60;, to move a Task to the completed state.
 **Attributes** | **string** | The JSON string that describes the custom attributes of the task.
 **Priority** | **int32** | The Task&#39;s new priority value. When supplied, the Task takes on the specified priority unless it matches a Workflow Target with a Priority set. Value can be 0 to 2^31^ (2,147,483,647).
 **Reason** | **string** | The reason that the Task was canceled or completed. This parameter is required only if the Task is canceled or completed. Setting this value queues the task for deletion and logs the reason.
 **TaskChannel** | **string** | When MultiTasking is enabled, specify the TaskChannel with the task to update. Can be the TaskChannel&#39;s SID or its &#x60;unique_name&#x60;, such as &#x60;voice&#x60;, &#x60;sms&#x60;, or &#x60;default&#x60;.

### Return type

[**TaskrouterV1WorkspaceTask**](TaskrouterV1WorkspaceTask.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskChannel

> TaskrouterV1WorkspaceTaskChannel UpdateTaskChannel(ctx, WorkspaceSid, Sid).ChannelOptimizedRouting(ChannelOptimizedRouting).FriendlyName(FriendlyName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Task Channel to update.
    Sid := "Sid_example" // string | The SID of the Task Channel resource to update.
    ChannelOptimizedRouting := true // bool | Whether the TaskChannel should prioritize Workers that have been idle. If `true`, Workers that have been idle the longest are prioritized. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTaskChannel(context.Background(), WorkspaceSid, Sid).ChannelOptimizedRouting(ChannelOptimizedRouting).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTaskChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskChannel`: TaskrouterV1WorkspaceTaskChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTaskChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Task Channel to update.
**Sid** | **string** | The SID of the Task Channel resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ChannelOptimizedRouting** | **bool** | Whether the TaskChannel should prioritize Workers that have been idle. If &#x60;true&#x60;, Workers that have been idle the longest are prioritized.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Task Channel. It can be up to 64 characters long.

### Return type

[**TaskrouterV1WorkspaceTaskChannel**](TaskrouterV1WorkspaceTaskChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaskQueue

> TaskrouterV1WorkspaceTaskQueue UpdateTaskQueue(ctx, WorkspaceSid, Sid).AssignmentActivitySid(AssignmentActivitySid).FriendlyName(FriendlyName).MaxReservedWorkers(MaxReservedWorkers).ReservationActivitySid(ReservationActivitySid).TargetWorkers(TargetWorkers).TaskOrder(TaskOrder).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskQueue to update.
    Sid := "Sid_example" // string | The SID of the TaskQueue resource to update.
    AssignmentActivitySid := "AssignmentActivitySid_example" // string | The SID of the Activity to assign Workers when a task is assigned for them. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the TaskQueue. For example `Support-Tier 1`, `Sales`, or `Escalation`. (optional)
    MaxReservedWorkers := int32(56) // int32 | The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50. (optional)
    ReservationActivitySid := "ReservationActivitySid_example" // string | The SID of the Activity to assign Workers when a task is reserved for them. (optional)
    TargetWorkers := "TargetWorkers_example" // string | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example '\\\"language\\\" == \\\"spanish\\\"' If no TargetWorkers parameter is provided, Tasks will wait in the queue until they are either deleted or moved to another queue. Additional examples on how to describing Worker selection criteria below. (optional)
    TaskOrder := "TaskOrder_example" // string | How Tasks will be assigned to Workers. Can be: `FIFO` or `LIFO` and the default is `FIFO`. Use `FIFO` to assign the oldest task first and `LIFO` to assign the most recent task first. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTaskQueue(context.Background(), WorkspaceSid, Sid).AssignmentActivitySid(AssignmentActivitySid).FriendlyName(FriendlyName).MaxReservedWorkers(MaxReservedWorkers).ReservationActivitySid(ReservationActivitySid).TargetWorkers(TargetWorkers).TaskOrder(TaskOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTaskQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskQueue`: TaskrouterV1WorkspaceTaskQueue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTaskQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskQueue to update.
**Sid** | **string** | The SID of the TaskQueue resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskQueueParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned for them.
 **FriendlyName** | **string** | A descriptive string that you create to describe the TaskQueue. For example &#x60;Support-Tier 1&#x60;, &#x60;Sales&#x60;, or &#x60;Escalation&#x60;.
 **MaxReservedWorkers** | **int32** | The maximum number of Workers to create reservations for the assignment of a task while in the queue. Maximum of 50.
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


## UpdateTaskReservation

> TaskrouterV1WorkspaceTaskTaskReservation UpdateTaskReservation(ctx, WorkspaceSid, TaskSid, Sid).Beep(Beep).BeepOnCustomerEntrance(BeepOnCustomerEntrance).CallAccept(CallAccept).CallFrom(CallFrom).CallRecord(CallRecord).CallStatusCallbackUrl(CallStatusCallbackUrl).CallTimeout(CallTimeout).CallTo(CallTo).CallUrl(CallUrl).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).DequeueFrom(DequeueFrom).DequeuePostWorkActivitySid(DequeuePostWorkActivitySid).DequeueRecord(DequeueRecord).DequeueStatusCallbackEvent(DequeueStatusCallbackEvent).DequeueStatusCallbackUrl(DequeueStatusCallbackUrl).DequeueTimeout(DequeueTimeout).DequeueTo(DequeueTo).EarlyMedia(EarlyMedia).EndConferenceOnCustomerExit(EndConferenceOnCustomerExit).EndConferenceOnExit(EndConferenceOnExit).From(From).Instruction(Instruction).MaxParticipants(MaxParticipants).Muted(Muted).PostWorkActivitySid(PostWorkActivitySid).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RedirectAccept(RedirectAccept).RedirectCallSid(RedirectCallSid).RedirectUrl(RedirectUrl).Region(Region).ReservationStatus(ReservationStatus).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Supervisor(Supervisor).SupervisorMode(SupervisorMode).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).WorkerActivitySid(WorkerActivitySid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the TaskReservation resources to update.
    TaskSid := "TaskSid_example" // string | The SID of the reserved Task resource with the TaskReservation resources to update.
    Sid := "Sid_example" // string | The SID of the TaskReservation resource to update.
    Beep := "Beep_example" // string | Whether to play a notification beep when the participant joins or when to play a beep. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`. (optional)
    BeepOnCustomerEntrance := true // bool | Whether to play a notification beep when the customer joins. (optional)
    CallAccept := true // bool | Whether to accept a reservation when executing a Call instruction. (optional)
    CallFrom := "CallFrom_example" // string | The Caller ID of the outbound call when executing a Call instruction. (optional)
    CallRecord := "CallRecord_example" // string | Whether to record both legs of a call when executing a Call instruction or which leg to record. (optional)
    CallStatusCallbackUrl := "CallStatusCallbackUrl_example" // string | The URL to call  for the completed call event when executing a Call instruction. (optional)
    CallTimeout := int32(56) // int32 | Timeout for call when executing a Call instruction. (optional)
    CallTo := "CallTo_example" // string | The Contact URI of the worker when executing a Call instruction.  Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    CallUrl := "CallUrl_example" // string | TwiML URI executed on answering the worker's leg as a result of the Call instruction. (optional)
    ConferenceRecord := "ConferenceRecord_example" // string | Whether to record the conference the participant is joining or when to record the conference. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`. (optional)
    ConferenceRecordingStatusCallback := "ConferenceRecordingStatusCallback_example" // string | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available. (optional)
    ConferenceRecordingStatusCallbackMethod := "ConferenceRecordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceStatusCallback := "ConferenceStatusCallback_example" // string | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored. (optional)
    ConferenceStatusCallbackEvent := []string{"Inner_example"} // []string | The conference status events that we will send to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`. (optional)
    ConferenceStatusCallbackMethod := "ConferenceStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceTrim := "ConferenceTrim_example" // string | How to trim the leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`. (optional)
    DequeueFrom := "DequeueFrom_example" // string | The Caller ID of the call to the worker when executing a Dequeue instruction. (optional)
    DequeuePostWorkActivitySid := "DequeuePostWorkActivitySid_example" // string | The SID of the Activity resource to start after executing a Dequeue instruction. (optional)
    DequeueRecord := "DequeueRecord_example" // string | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record. (optional)
    DequeueStatusCallbackEvent := []string{"Inner_example"} // []string | The Call progress events sent via webhooks as a result of a Dequeue instruction. (optional)
    DequeueStatusCallbackUrl := "DequeueStatusCallbackUrl_example" // string | The Callback URL for completed call event when executing a Dequeue instruction. (optional)
    DequeueTimeout := int32(56) // int32 | Timeout for call when executing a Dequeue instruction. (optional)
    DequeueTo := "DequeueTo_example" // string | The Contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    EarlyMedia := true // bool | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is `true`. (optional)
    EndConferenceOnCustomerExit := true // bool | Whether to end the conference when the customer leaves. (optional)
    EndConferenceOnExit := true // bool | Whether to end the conference when the agent leaves. (optional)
    From := "From_example" // string | The Caller ID of the call to the worker when executing a Conference instruction. (optional)
    Instruction := "Instruction_example" // string | The assignment instruction for reservation. (optional)
    MaxParticipants := int32(56) // int32 | The maximum number of participants in the conference. Can be a positive integer from `2` to `250`. The default value is `250`. (optional)
    Muted := true // bool | Whether the agent is muted in the conference. The default is `false`. (optional)
    PostWorkActivitySid := "PostWorkActivitySid_example" // string | The new worker activity SID after executing a Conference instruction. (optional)
    Record := true // bool | Whether to record the participant and their conferences, including the time between conferences. The default is `false`. (optional)
    RecordingChannels := "RecordingChannels_example" // string | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`. (optional)
    RecordingStatusCallback := "RecordingStatusCallback_example" // string | The URL that we should call using the `recording_status_callback_method` when the recording status changes. (optional)
    RecordingStatusCallbackMethod := "RecordingStatusCallbackMethod_example" // string | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    RedirectAccept := true // bool | Whether the reservation should be accepted when executing a Redirect instruction. (optional)
    RedirectCallSid := "RedirectCallSid_example" // string | The Call SID of the call parked in the queue when executing a Redirect instruction. (optional)
    RedirectUrl := "RedirectUrl_example" // string | TwiML URI to redirect the call to when executing the Redirect instruction. (optional)
    Region := "Region_example" // string | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`. (optional)
    ReservationStatus := "ReservationStatus_example" // string | The new status of the reservation. Can be: `pending`, `accepted`, `rejected`, or `timeout`. (optional)
    SipAuthPassword := "SipAuthPassword_example" // string | The SIP password for authentication. (optional)
    SipAuthUsername := "SipAuthUsername_example" // string | The SIP username used for authentication. (optional)
    StartConferenceOnEnter := true // bool | Whether to start the conference when the participant joins, if it has not already started. The default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackEvent := []string{"Inner_example"} // []string | The call progress events that we will send to `status_callback`. Can be: `initiated`, `ringing`, `answered`, or `completed`. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`. (optional)
    Supervisor := "Supervisor_example" // string | The Supervisor SID/URI when executing the Supervise instruction. (optional)
    SupervisorMode := "SupervisorMode_example" // string | The Supervisor mode when executing the Supervise instruction. (optional)
    Timeout := int32(56) // int32 | Timeout for call when executing a Conference instruction. (optional)
    To := "To_example" // string | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    WaitMethod := "WaitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    WaitUrl := "WaitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)
    WorkerActivitySid := "WorkerActivitySid_example" // string | The new worker activity SID if rejecting a reservation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTaskReservation(context.Background(), WorkspaceSid, TaskSid, Sid).Beep(Beep).BeepOnCustomerEntrance(BeepOnCustomerEntrance).CallAccept(CallAccept).CallFrom(CallFrom).CallRecord(CallRecord).CallStatusCallbackUrl(CallStatusCallbackUrl).CallTimeout(CallTimeout).CallTo(CallTo).CallUrl(CallUrl).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).DequeueFrom(DequeueFrom).DequeuePostWorkActivitySid(DequeuePostWorkActivitySid).DequeueRecord(DequeueRecord).DequeueStatusCallbackEvent(DequeueStatusCallbackEvent).DequeueStatusCallbackUrl(DequeueStatusCallbackUrl).DequeueTimeout(DequeueTimeout).DequeueTo(DequeueTo).EarlyMedia(EarlyMedia).EndConferenceOnCustomerExit(EndConferenceOnCustomerExit).EndConferenceOnExit(EndConferenceOnExit).From(From).Instruction(Instruction).MaxParticipants(MaxParticipants).Muted(Muted).PostWorkActivitySid(PostWorkActivitySid).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RedirectAccept(RedirectAccept).RedirectCallSid(RedirectCallSid).RedirectUrl(RedirectUrl).Region(Region).ReservationStatus(ReservationStatus).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Supervisor(Supervisor).SupervisorMode(SupervisorMode).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).WorkerActivitySid(WorkerActivitySid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTaskReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTaskReservation`: TaskrouterV1WorkspaceTaskTaskReservation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTaskReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the TaskReservation resources to update.
**TaskSid** | **string** | The SID of the reserved Task resource with the TaskReservation resources to update.
**Sid** | **string** | The SID of the TaskReservation resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTaskReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Beep** | **string** | Whether to play a notification beep when the participant joins or when to play a beep. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
 **BeepOnCustomerEntrance** | **bool** | Whether to play a notification beep when the customer joins.
 **CallAccept** | **bool** | Whether to accept a reservation when executing a Call instruction.
 **CallFrom** | **string** | The Caller ID of the outbound call when executing a Call instruction.
 **CallRecord** | **string** | Whether to record both legs of a call when executing a Call instruction or which leg to record.
 **CallStatusCallbackUrl** | **string** | The URL to call  for the completed call event when executing a Call instruction.
 **CallTimeout** | **int32** | Timeout for call when executing a Call instruction.
 **CallTo** | **string** | The Contact URI of the worker when executing a Call instruction.  Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **CallUrl** | **string** | TwiML URI executed on answering the worker&#39;s leg as a result of the Call instruction.
 **ConferenceRecord** | **string** | Whether to record the conference the participant is joining or when to record the conference. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
 **ConferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
 **ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
 **ConferenceStatusCallbackEvent** | **[]string** | The conference status events that we will send to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;.
 **ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceTrim** | **string** | How to trim the leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
 **DequeueFrom** | **string** | The Caller ID of the call to the worker when executing a Dequeue instruction.
 **DequeuePostWorkActivitySid** | **string** | The SID of the Activity resource to start after executing a Dequeue instruction.
 **DequeueRecord** | **string** | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
 **DequeueStatusCallbackEvent** | **[]string** | The Call progress events sent via webhooks as a result of a Dequeue instruction.
 **DequeueStatusCallbackUrl** | **string** | The Callback URL for completed call event when executing a Dequeue instruction.
 **DequeueTimeout** | **int32** | Timeout for call when executing a Dequeue instruction.
 **DequeueTo** | **string** | The Contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is &#x60;true&#x60;.
 **EndConferenceOnCustomerExit** | **bool** | Whether to end the conference when the customer leaves.
 **EndConferenceOnExit** | **bool** | Whether to end the conference when the agent leaves.
 **From** | **string** | The Caller ID of the call to the worker when executing a Conference instruction.
 **Instruction** | **string** | The assignment instruction for reservation.
 **MaxParticipants** | **int32** | The maximum number of participants in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
 **Muted** | **bool** | Whether the agent is muted in the conference. The default is &#x60;false&#x60;.
 **PostWorkActivitySid** | **string** | The new worker activity SID after executing a Conference instruction.
 **Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. The default is &#x60;false&#x60;.
 **RecordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
 **RecordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
 **RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **RedirectAccept** | **bool** | Whether the reservation should be accepted when executing a Redirect instruction.
 **RedirectCallSid** | **string** | The Call SID of the call parked in the queue when executing a Redirect instruction.
 **RedirectUrl** | **string** | TwiML URI to redirect the call to when executing the Redirect instruction.
 **Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
 **ReservationStatus** | **string** | The new status of the reservation. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, or &#x60;timeout&#x60;.
 **SipAuthPassword** | **string** | The SIP password for authentication.
 **SipAuthUsername** | **string** | The SIP username used for authentication.
 **StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. The default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackEvent** | **[]string** | The call progress events that we will send to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, or &#x60;completed&#x60;.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
 **Supervisor** | **string** | The Supervisor SID/URI when executing the Supervise instruction.
 **SupervisorMode** | **string** | The Supervisor mode when executing the Supervise instruction.
 **Timeout** | **int32** | Timeout for call when executing a Conference instruction.
 **To** | **string** | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
 **WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
 **WorkerActivitySid** | **string** | The new worker activity SID if rejecting a reservation.

### Return type

[**TaskrouterV1WorkspaceTaskTaskReservation**](TaskrouterV1WorkspaceTaskTaskReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorker

> TaskrouterV1WorkspaceWorker UpdateWorker(ctx, WorkspaceSid, Sid).ActivitySid(ActivitySid).Attributes(Attributes).FriendlyName(FriendlyName).RejectPendingReservations(RejectPendingReservations).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Worker to update.
    Sid := "Sid_example" // string | The SID of the Worker resource to update.
    ActivitySid := "ActivitySid_example" // string | The SID of a valid Activity that will describe the Worker's initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information. (optional)
    Attributes := "Attributes_example" // string | The JSON string that describes the Worker. For example: `{ \\\"email\\\": \\\"Bob@example.com\\\", \\\"phone\\\": \\\"+5095551234\\\" }`. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker. Defaults to {}. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Worker. It can be up to 64 characters long. (optional)
    RejectPendingReservations := true // bool | Whether to reject pending reservations. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorker(context.Background(), WorkspaceSid, Sid).ActivitySid(ActivitySid).Attributes(Attributes).FriendlyName(FriendlyName).RejectPendingReservations(RejectPendingReservations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorker`: TaskrouterV1WorkspaceWorker
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Worker to update.
**Sid** | **string** | The SID of the Worker resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **ActivitySid** | **string** | The SID of a valid Activity that will describe the Worker&#39;s initial state. See [Activities](https://www.twilio.com/docs/taskrouter/api/activity) for more information.
 **Attributes** | **string** | The JSON string that describes the Worker. For example: &#x60;{ \\\&quot;email\\\&quot;: \\\&quot;Bob@example.com\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;+5095551234\\\&quot; }&#x60;. This data is passed to the &#x60;assignment_callback_url&#x60; when TaskRouter assigns a Task to the Worker. Defaults to {}.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Worker. It can be up to 64 characters long.
 **RejectPendingReservations** | **bool** | Whether to reject pending reservations.

### Return type

[**TaskrouterV1WorkspaceWorker**](TaskrouterV1WorkspaceWorker.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerChannel

> TaskrouterV1WorkspaceWorkerWorkerChannel UpdateWorkerChannel(ctx, WorkspaceSid, WorkerSid, Sid).Available(Available).Capacity(Capacity).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerChannel to update.
    WorkerSid := "WorkerSid_example" // string | The SID of the Worker with the WorkerChannel to update.
    Sid := "Sid_example" // string | The SID of the WorkerChannel to update.
    Available := true // bool | Whether the WorkerChannel is available. Set to `false` to prevent the Worker from receiving any new Tasks of this TaskChannel type. (optional)
    Capacity := int32(56) // int32 | The total number of Tasks that the Worker should handle for the TaskChannel type. TaskRouter creates reservations for Tasks of this TaskChannel type up to the specified capacity. If the capacity is 0, no new reservations will be created. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorkerChannel(context.Background(), WorkspaceSid, WorkerSid, Sid).Available(Available).Capacity(Capacity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorkerChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkerChannel`: TaskrouterV1WorkspaceWorkerWorkerChannel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorkerChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerChannel to update.
**WorkerSid** | **string** | The SID of the Worker with the WorkerChannel to update.
**Sid** | **string** | The SID of the WorkerChannel to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerChannelParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Available** | **bool** | Whether the WorkerChannel is available. Set to &#x60;false&#x60; to prevent the Worker from receiving any new Tasks of this TaskChannel type.
 **Capacity** | **int32** | The total number of Tasks that the Worker should handle for the TaskChannel type. TaskRouter creates reservations for Tasks of this TaskChannel type up to the specified capacity. If the capacity is 0, no new reservations will be created.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerChannel**](TaskrouterV1WorkspaceWorkerWorkerChannel.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkerReservation

> TaskrouterV1WorkspaceWorkerWorkerReservation UpdateWorkerReservation(ctx, WorkspaceSid, WorkerSid, Sid).Beep(Beep).BeepOnCustomerEntrance(BeepOnCustomerEntrance).CallAccept(CallAccept).CallFrom(CallFrom).CallRecord(CallRecord).CallStatusCallbackUrl(CallStatusCallbackUrl).CallTimeout(CallTimeout).CallTo(CallTo).CallUrl(CallUrl).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).DequeueFrom(DequeueFrom).DequeuePostWorkActivitySid(DequeuePostWorkActivitySid).DequeueRecord(DequeueRecord).DequeueStatusCallbackEvent(DequeueStatusCallbackEvent).DequeueStatusCallbackUrl(DequeueStatusCallbackUrl).DequeueTimeout(DequeueTimeout).DequeueTo(DequeueTo).EarlyMedia(EarlyMedia).EndConferenceOnCustomerExit(EndConferenceOnCustomerExit).EndConferenceOnExit(EndConferenceOnExit).From(From).Instruction(Instruction).MaxParticipants(MaxParticipants).Muted(Muted).PostWorkActivitySid(PostWorkActivitySid).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RedirectAccept(RedirectAccept).RedirectCallSid(RedirectCallSid).RedirectUrl(RedirectUrl).Region(Region).ReservationStatus(ReservationStatus).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).WorkerActivitySid(WorkerActivitySid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the WorkerReservation resources to update.
    WorkerSid := "WorkerSid_example" // string | The SID of the reserved Worker resource with the WorkerReservation resources to update.
    Sid := "Sid_example" // string | The SID of the WorkerReservation resource to update.
    Beep := "Beep_example" // string | Whether to play a notification beep when the participant joins or when to play a beep. Can be: `true`, `false`, `onEnter`, or `onExit`. The default value is `true`. (optional)
    BeepOnCustomerEntrance := true // bool | Whether to play a notification beep when the customer joins. (optional)
    CallAccept := true // bool | Whether to accept a reservation when executing a Call instruction. (optional)
    CallFrom := "CallFrom_example" // string | The Caller ID of the outbound call when executing a Call instruction. (optional)
    CallRecord := "CallRecord_example" // string | Whether to record both legs of a call when executing a Call instruction. (optional)
    CallStatusCallbackUrl := "CallStatusCallbackUrl_example" // string | The URL to call for the completed call event when executing a Call instruction. (optional)
    CallTimeout := int32(56) // int32 | The timeout for a call when executing a Call instruction. (optional)
    CallTo := "CallTo_example" // string | The contact URI of the worker when executing a Call instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    CallUrl := "CallUrl_example" // string | TwiML URI executed on answering the worker's leg as a result of the Call instruction. (optional)
    ConferenceRecord := "ConferenceRecord_example" // string | Whether to record the conference the participant is joining or when to record the conference. Can be: `true`, `false`, `record-from-start`, and `do-not-record`. The default value is `false`. (optional)
    ConferenceRecordingStatusCallback := "ConferenceRecordingStatusCallback_example" // string | The URL we should call using the `conference_recording_status_callback_method` when the conference recording is available. (optional)
    ConferenceRecordingStatusCallbackMethod := "ConferenceRecordingStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceStatusCallback := "ConferenceStatusCallback_example" // string | The URL we should call using the `conference_status_callback_method` when the conference events in `conference_status_callback_event` occur. Only the value set by the first participant to join the conference is used. Subsequent `conference_status_callback` values are ignored. (optional)
    ConferenceStatusCallbackEvent := []string{"Inner_example"} // []string | The conference status events that we will send to `conference_status_callback`. Can be: `start`, `end`, `join`, `leave`, `mute`, `hold`, `speaker`. (optional)
    ConferenceStatusCallbackMethod := "ConferenceStatusCallbackMethod_example" // string | The HTTP method we should use to call `conference_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    ConferenceTrim := "ConferenceTrim_example" // string | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: `trim-silence` or `do-not-trim` and defaults to `trim-silence`. (optional)
    DequeueFrom := "DequeueFrom_example" // string | The caller ID of the call to the worker when executing a Dequeue instruction. (optional)
    DequeuePostWorkActivitySid := "DequeuePostWorkActivitySid_example" // string | The SID of the Activity resource to start after executing a Dequeue instruction. (optional)
    DequeueRecord := "DequeueRecord_example" // string | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record. (optional)
    DequeueStatusCallbackEvent := []string{"Inner_example"} // []string | The call progress events sent via webhooks as a result of a Dequeue instruction. (optional)
    DequeueStatusCallbackUrl := "DequeueStatusCallbackUrl_example" // string | The callback URL for completed call event when executing a Dequeue instruction. (optional)
    DequeueTimeout := int32(56) // int32 | The timeout for call when executing a Dequeue instruction. (optional)
    DequeueTo := "DequeueTo_example" // string | The contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    EarlyMedia := true // bool | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is `true`. (optional)
    EndConferenceOnCustomerExit := true // bool | Whether to end the conference when the customer leaves. (optional)
    EndConferenceOnExit := true // bool | Whether to end the conference when the agent leaves. (optional)
    From := "From_example" // string | The caller ID of the call to the worker when executing a Conference instruction. (optional)
    Instruction := "Instruction_example" // string | The assignment instruction for the reservation. (optional)
    MaxParticipants := int32(56) // int32 | The maximum number of participants allowed in the conference. Can be a positive integer from `2` to `250`. The default value is `250`. (optional)
    Muted := true // bool | Whether the agent is muted in the conference. Defaults to `false`. (optional)
    PostWorkActivitySid := "PostWorkActivitySid_example" // string | The new worker activity SID after executing a Conference instruction. (optional)
    Record := true // bool | Whether to record the participant and their conferences, including the time between conferences. Can be `true` or `false` and the default is `false`. (optional)
    RecordingChannels := "RecordingChannels_example" // string | The recording channels for the final recording. Can be: `mono` or `dual` and the default is `mono`. (optional)
    RecordingStatusCallback := "RecordingStatusCallback_example" // string | The URL that we should call using the `recording_status_callback_method` when the recording status changes. (optional)
    RecordingStatusCallbackMethod := "RecordingStatusCallbackMethod_example" // string | The HTTP method we should use when we call `recording_status_callback`. Can be: `GET` or `POST` and defaults to `POST`. (optional)
    RedirectAccept := true // bool | Whether the reservation should be accepted when executing a Redirect instruction. (optional)
    RedirectCallSid := "RedirectCallSid_example" // string | The Call SID of the call parked in the queue when executing a Redirect instruction. (optional)
    RedirectUrl := "RedirectUrl_example" // string | TwiML URI to redirect the call to when executing the Redirect instruction. (optional)
    Region := "Region_example" // string | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:`us1`, `ie1`, `de1`, `sg1`, `br1`, `au1`, or `jp1`. (optional)
    ReservationStatus := "ReservationStatus_example" // string | The new status of the reservation. Can be: `pending`, `accepted`, `rejected`, `timeout`, `canceled`, or `rescinded`. (optional)
    SipAuthPassword := "SipAuthPassword_example" // string | The SIP password for authentication. (optional)
    SipAuthUsername := "SipAuthUsername_example" // string | The SIP username used for authentication. (optional)
    StartConferenceOnEnter := true // bool | Whether to start the conference when the participant joins, if it has not already started. Can be: `true` or `false` and the default is `true`. If `false` and the conference has not started, the participant is muted and hears background music until another participant starts the conference. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call using the `status_callback_method` to send status information to your application. (optional)
    StatusCallbackEvent := []string{"Inner_example"} // []string | The call progress events that we will send to `status_callback`. Can be: `initiated`, `ringing`, `answered`, or `completed`. (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback`. Can be: `POST` or `GET` and the default is `POST`. (optional)
    Timeout := int32(56) // int32 | The timeout for a call when executing a Conference instruction. (optional)
    To := "To_example" // string | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination. (optional)
    WaitMethod := "WaitMethod_example" // string | The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file. (optional)
    WaitUrl := "WaitUrl_example" // string | The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic). (optional)
    WorkerActivitySid := "WorkerActivitySid_example" // string | The new worker activity SID if rejecting a reservation. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorkerReservation(context.Background(), WorkspaceSid, WorkerSid, Sid).Beep(Beep).BeepOnCustomerEntrance(BeepOnCustomerEntrance).CallAccept(CallAccept).CallFrom(CallFrom).CallRecord(CallRecord).CallStatusCallbackUrl(CallStatusCallbackUrl).CallTimeout(CallTimeout).CallTo(CallTo).CallUrl(CallUrl).ConferenceRecord(ConferenceRecord).ConferenceRecordingStatusCallback(ConferenceRecordingStatusCallback).ConferenceRecordingStatusCallbackMethod(ConferenceRecordingStatusCallbackMethod).ConferenceStatusCallback(ConferenceStatusCallback).ConferenceStatusCallbackEvent(ConferenceStatusCallbackEvent).ConferenceStatusCallbackMethod(ConferenceStatusCallbackMethod).ConferenceTrim(ConferenceTrim).DequeueFrom(DequeueFrom).DequeuePostWorkActivitySid(DequeuePostWorkActivitySid).DequeueRecord(DequeueRecord).DequeueStatusCallbackEvent(DequeueStatusCallbackEvent).DequeueStatusCallbackUrl(DequeueStatusCallbackUrl).DequeueTimeout(DequeueTimeout).DequeueTo(DequeueTo).EarlyMedia(EarlyMedia).EndConferenceOnCustomerExit(EndConferenceOnCustomerExit).EndConferenceOnExit(EndConferenceOnExit).From(From).Instruction(Instruction).MaxParticipants(MaxParticipants).Muted(Muted).PostWorkActivitySid(PostWorkActivitySid).Record(Record).RecordingChannels(RecordingChannels).RecordingStatusCallback(RecordingStatusCallback).RecordingStatusCallbackMethod(RecordingStatusCallbackMethod).RedirectAccept(RedirectAccept).RedirectCallSid(RedirectCallSid).RedirectUrl(RedirectUrl).Region(Region).ReservationStatus(ReservationStatus).SipAuthPassword(SipAuthPassword).SipAuthUsername(SipAuthUsername).StartConferenceOnEnter(StartConferenceOnEnter).StatusCallback(StatusCallback).StatusCallbackEvent(StatusCallbackEvent).StatusCallbackMethod(StatusCallbackMethod).Timeout(Timeout).To(To).WaitMethod(WaitMethod).WaitUrl(WaitUrl).WorkerActivitySid(WorkerActivitySid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorkerReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkerReservation`: TaskrouterV1WorkspaceWorkerWorkerReservation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorkerReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the WorkerReservation resources to update.
**WorkerSid** | **string** | The SID of the reserved Worker resource with the WorkerReservation resources to update.
**Sid** | **string** | The SID of the WorkerReservation resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkerReservationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Beep** | **string** | Whether to play a notification beep when the participant joins or when to play a beep. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;onEnter&#x60;, or &#x60;onExit&#x60;. The default value is &#x60;true&#x60;.
 **BeepOnCustomerEntrance** | **bool** | Whether to play a notification beep when the customer joins.
 **CallAccept** | **bool** | Whether to accept a reservation when executing a Call instruction.
 **CallFrom** | **string** | The Caller ID of the outbound call when executing a Call instruction.
 **CallRecord** | **string** | Whether to record both legs of a call when executing a Call instruction.
 **CallStatusCallbackUrl** | **string** | The URL to call for the completed call event when executing a Call instruction.
 **CallTimeout** | **int32** | The timeout for a call when executing a Call instruction.
 **CallTo** | **string** | The contact URI of the worker when executing a Call instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **CallUrl** | **string** | TwiML URI executed on answering the worker&#39;s leg as a result of the Call instruction.
 **ConferenceRecord** | **string** | Whether to record the conference the participant is joining or when to record the conference. Can be: &#x60;true&#x60;, &#x60;false&#x60;, &#x60;record-from-start&#x60;, and &#x60;do-not-record&#x60;. The default value is &#x60;false&#x60;.
 **ConferenceRecordingStatusCallback** | **string** | The URL we should call using the &#x60;conference_recording_status_callback_method&#x60; when the conference recording is available.
 **ConferenceRecordingStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceStatusCallback** | **string** | The URL we should call using the &#x60;conference_status_callback_method&#x60; when the conference events in &#x60;conference_status_callback_event&#x60; occur. Only the value set by the first participant to join the conference is used. Subsequent &#x60;conference_status_callback&#x60; values are ignored.
 **ConferenceStatusCallbackEvent** | **[]string** | The conference status events that we will send to &#x60;conference_status_callback&#x60;. Can be: &#x60;start&#x60;, &#x60;end&#x60;, &#x60;join&#x60;, &#x60;leave&#x60;, &#x60;mute&#x60;, &#x60;hold&#x60;, &#x60;speaker&#x60;.
 **ConferenceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;conference_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **ConferenceTrim** | **string** | Whether to trim leading and trailing silence from your recorded conference audio files. Can be: &#x60;trim-silence&#x60; or &#x60;do-not-trim&#x60; and defaults to &#x60;trim-silence&#x60;.
 **DequeueFrom** | **string** | The caller ID of the call to the worker when executing a Dequeue instruction.
 **DequeuePostWorkActivitySid** | **string** | The SID of the Activity resource to start after executing a Dequeue instruction.
 **DequeueRecord** | **string** | Whether to record both legs of a call when executing a Dequeue instruction or which leg to record.
 **DequeueStatusCallbackEvent** | **[]string** | The call progress events sent via webhooks as a result of a Dequeue instruction.
 **DequeueStatusCallbackUrl** | **string** | The callback URL for completed call event when executing a Dequeue instruction.
 **DequeueTimeout** | **int32** | The timeout for call when executing a Dequeue instruction.
 **DequeueTo** | **string** | The contact URI of the worker when executing a Dequeue instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **EarlyMedia** | **bool** | Whether to allow an agent to hear the state of the outbound call, including ringing or disconnect messages. The default is &#x60;true&#x60;.
 **EndConferenceOnCustomerExit** | **bool** | Whether to end the conference when the customer leaves.
 **EndConferenceOnExit** | **bool** | Whether to end the conference when the agent leaves.
 **From** | **string** | The caller ID of the call to the worker when executing a Conference instruction.
 **Instruction** | **string** | The assignment instruction for the reservation.
 **MaxParticipants** | **int32** | The maximum number of participants allowed in the conference. Can be a positive integer from &#x60;2&#x60; to &#x60;250&#x60;. The default value is &#x60;250&#x60;.
 **Muted** | **bool** | Whether the agent is muted in the conference. Defaults to &#x60;false&#x60;.
 **PostWorkActivitySid** | **string** | The new worker activity SID after executing a Conference instruction.
 **Record** | **bool** | Whether to record the participant and their conferences, including the time between conferences. Can be &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;false&#x60;.
 **RecordingChannels** | **string** | The recording channels for the final recording. Can be: &#x60;mono&#x60; or &#x60;dual&#x60; and the default is &#x60;mono&#x60;.
 **RecordingStatusCallback** | **string** | The URL that we should call using the &#x60;recording_status_callback_method&#x60; when the recording status changes.
 **RecordingStatusCallbackMethod** | **string** | The HTTP method we should use when we call &#x60;recording_status_callback&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and defaults to &#x60;POST&#x60;.
 **RedirectAccept** | **bool** | Whether the reservation should be accepted when executing a Redirect instruction.
 **RedirectCallSid** | **string** | The Call SID of the call parked in the queue when executing a Redirect instruction.
 **RedirectUrl** | **string** | TwiML URI to redirect the call to when executing the Redirect instruction.
 **Region** | **string** | The [region](https://support.twilio.com/hc/en-us/articles/223132167-How-global-low-latency-routing-and-region-selection-work-for-conferences-and-Client-calls) where we should mix the recorded audio. Can be:&#x60;us1&#x60;, &#x60;ie1&#x60;, &#x60;de1&#x60;, &#x60;sg1&#x60;, &#x60;br1&#x60;, &#x60;au1&#x60;, or &#x60;jp1&#x60;.
 **ReservationStatus** | **string** | The new status of the reservation. Can be: &#x60;pending&#x60;, &#x60;accepted&#x60;, &#x60;rejected&#x60;, &#x60;timeout&#x60;, &#x60;canceled&#x60;, or &#x60;rescinded&#x60;.
 **SipAuthPassword** | **string** | The SIP password for authentication.
 **SipAuthUsername** | **string** | The SIP username used for authentication.
 **StartConferenceOnEnter** | **bool** | Whether to start the conference when the participant joins, if it has not already started. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If &#x60;false&#x60; and the conference has not started, the participant is muted and hears background music until another participant starts the conference.
 **StatusCallback** | **string** | The URL we should call using the &#x60;status_callback_method&#x60; to send status information to your application.
 **StatusCallbackEvent** | **[]string** | The call progress events that we will send to &#x60;status_callback&#x60;. Can be: &#x60;initiated&#x60;, &#x60;ringing&#x60;, &#x60;answered&#x60;, or &#x60;completed&#x60;.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60; and the default is &#x60;POST&#x60;.
 **Timeout** | **int32** | The timeout for a call when executing a Conference instruction.
 **To** | **string** | The Contact URI of the worker when executing a Conference instruction. Can be the URI of the Twilio Client, the SIP URI for Programmable SIP, or the [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number, depending on the destination.
 **WaitMethod** | **string** | The HTTP method we should use to call &#x60;wait_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;. When using a static audio file, this should be &#x60;GET&#x60; so that we can cache the file.
 **WaitUrl** | **string** | The URL we should call using the &#x60;wait_method&#x60; for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
 **WorkerActivitySid** | **string** | The new worker activity SID if rejecting a reservation.

### Return type

[**TaskrouterV1WorkspaceWorkerWorkerReservation**](TaskrouterV1WorkspaceWorkerWorkerReservation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> TaskrouterV1WorkspaceWorkflow UpdateWorkflow(ctx, WorkspaceSid, Sid).AssignmentCallbackUrl(AssignmentCallbackUrl).Configuration(Configuration).FallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl).FriendlyName(FriendlyName).ReEvaluateTasks(ReEvaluateTasks).TaskReservationTimeout(TaskReservationTimeout).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    WorkspaceSid := "WorkspaceSid_example" // string | The SID of the Workspace with the Workflow to update.
    Sid := "Sid_example" // string | The SID of the Workflow resource to update.
    AssignmentCallbackUrl := "AssignmentCallbackUrl_example" // string | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details. (optional)
    Configuration := "Configuration_example" // string | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information. (optional)
    FallbackAssignmentCallbackUrl := "FallbackAssignmentCallbackUrl_example" // string | The URL that we should call when a call to the `assignment_callback_url` fails. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Workflow resource. For example, `Inbound Call Workflow` or `2014 Outbound Campaign`. (optional)
    ReEvaluateTasks := "ReEvaluateTasks_example" // string | Whether or not to re-evaluate Tasks. The default is `false`, which means Tasks in the Workflow will not be processed through the assignment loop again. (optional)
    TaskReservationTimeout := int32(56) // int32 | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to `86,400` (24 hours) and the default is `120`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorkflow(context.Background(), WorkspaceSid, Sid).AssignmentCallbackUrl(AssignmentCallbackUrl).Configuration(Configuration).FallbackAssignmentCallbackUrl(FallbackAssignmentCallbackUrl).FriendlyName(FriendlyName).ReEvaluateTasks(ReEvaluateTasks).TaskReservationTimeout(TaskReservationTimeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflow`: TaskrouterV1WorkspaceWorkflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**WorkspaceSid** | **string** | The SID of the Workspace with the Workflow to update.
**Sid** | **string** | The SID of the Workflow resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkflowParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **AssignmentCallbackUrl** | **string** | The URL from your application that will process task assignment events. See [Handling Task Assignment Callback](https://www.twilio.com/docs/taskrouter/handle-assignment-callbacks) for more details.
 **Configuration** | **string** | A JSON string that contains the rules to apply to the Workflow. See [Configuring Workflows](https://www.twilio.com/docs/taskrouter/workflow-configuration) for more information.
 **FallbackAssignmentCallbackUrl** | **string** | The URL that we should call when a call to the &#x60;assignment_callback_url&#x60; fails.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Workflow resource. For example, &#x60;Inbound Call Workflow&#x60; or &#x60;2014 Outbound Campaign&#x60;.
 **ReEvaluateTasks** | **string** | Whether or not to re-evaluate Tasks. The default is &#x60;false&#x60;, which means Tasks in the Workflow will not be processed through the assignment loop again.
 **TaskReservationTimeout** | **int32** | How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker. Can be up to &#x60;86,400&#x60; (24 hours) and the default is &#x60;120&#x60;.

### Return type

[**TaskrouterV1WorkspaceWorkflow**](TaskrouterV1WorkspaceWorkflow.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspace

> TaskrouterV1Workspace UpdateWorkspace(ctx, Sid).DefaultActivitySid(DefaultActivitySid).EventCallbackUrl(EventCallbackUrl).EventsFilter(EventsFilter).FriendlyName(FriendlyName).MultiTaskEnabled(MultiTaskEnabled).PrioritizeQueueOrder(PrioritizeQueueOrder).TimeoutActivitySid(TimeoutActivitySid).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The SID of the Workspace resource to update.
    DefaultActivitySid := "DefaultActivitySid_example" // string | The SID of the Activity that will be used when new Workers are created in the Workspace. (optional)
    EventCallbackUrl := "EventCallbackUrl_example" // string | The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. (optional)
    EventsFilter := "EventsFilter_example" // string | The list of Workspace events for which to call event_callback_url. For example if `EventsFilter=task.created,task.canceled,worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the Workspace resource. For example: `Sales Call Center` or `Customer Support Team`. (optional)
    MultiTaskEnabled := true // bool | Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. The default is `false`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking). (optional)
    PrioritizeQueueOrder := "PrioritizeQueueOrder_example" // string | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: `LIFO` or `FIFO` and the default is `FIFO`. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo). (optional)
    TimeoutActivitySid := "TimeoutActivitySid_example" // string | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWorkspace(context.Background(), Sid).DefaultActivitySid(DefaultActivitySid).EventCallbackUrl(EventCallbackUrl).EventsFilter(EventsFilter).FriendlyName(FriendlyName).MultiTaskEnabled(MultiTaskEnabled).PrioritizeQueueOrder(PrioritizeQueueOrder).TimeoutActivitySid(TimeoutActivitySid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkspace`: TaskrouterV1Workspace
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Workspace resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWorkspaceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **DefaultActivitySid** | **string** | The SID of the Activity that will be used when new Workers are created in the Workspace.
 **EventCallbackUrl** | **string** | The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information.
 **EventsFilter** | **string** | The list of Workspace events for which to call event_callback_url. For example if &#x60;EventsFilter&#x3D;task.created,task.canceled,worker.activity.update&#x60;, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
 **FriendlyName** | **string** | A descriptive string that you create to describe the Workspace resource. For example: &#x60;Sales Call Center&#x60; or &#x60;Customer Support Team&#x60;.
 **MultiTaskEnabled** | **bool** | Whether to enable multi-tasking. Can be: &#x60;true&#x60; to enable multi-tasking, or &#x60;false&#x60; to disable it. The default is &#x60;false&#x60;. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (&#x60;true&#x60;), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. Otherwise, each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
 **PrioritizeQueueOrder** | **string** | The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues. Can be: &#x60;LIFO&#x60; or &#x60;FIFO&#x60; and the default is &#x60;FIFO&#x60;. For more information, see [Queue Ordering](https://www.twilio.com/docs/taskrouter/queue-ordering-last-first-out-lifo).
 **TimeoutActivitySid** | **string** | The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response.

### Return type

[**TaskrouterV1Workspace**](TaskrouterV1Workspace.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

