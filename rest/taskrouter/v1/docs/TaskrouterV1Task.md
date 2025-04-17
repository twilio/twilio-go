# TaskrouterV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Task resource. |
**Age** | **int** | The number of seconds since the Task was created. |[optional] [default to 0]
**AssignmentStatus** | Pointer to [**string**](TaskEnumStatus.md) |  |
**Attributes** | Pointer to **string** | The JSON string with custom attributes of the work. **Note** If this property has been assigned a value, it will only be displayed in FETCH action that returns a single resource. Otherwise, it will be null. |
**Addons** | Pointer to **string** | An object that contains the [Add-on](https://www.twilio.com/docs/add-ons) data for all installed Add-ons. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**TaskQueueEnteredDate** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Task entered the TaskQueue, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Priority** | **int** | The current priority score of the Task as assigned to a Worker by the workflow. Tasks with higher priority values will be assigned before Tasks with lower values. |[optional] [default to 0]
**Reason** | Pointer to **string** | The reason the Task was canceled or completed, if applicable. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Task resource. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue. |
**TaskQueueFriendlyName** | Pointer to **string** | The friendly name of the TaskQueue. |
**TaskChannelSid** | Pointer to **string** | The SID of the TaskChannel. |
**TaskChannelUniqueName** | Pointer to **string** | The unique name of the TaskChannel. |
**Timeout** | **int** | The amount of time in seconds that the Task can live before being assigned. |[optional] [default to 0]
**WorkflowSid** | Pointer to **string** | The SID of the Workflow that is controlling the Task. |
**WorkflowFriendlyName** | Pointer to **string** | The friendly name of the Workflow that is controlling the Task. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Task. |
**Url** | Pointer to **string** | The absolute URL of the Task resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |
**VirtualStartTime** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT indicating the ordering for routing of the Task specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**IgnoreCapacity** | Pointer to **bool** | A boolean that indicates if the Task should respect a Worker's capacity and availability during assignment. This field can only be used when the `RoutingTarget` field is set to a Worker SID. By setting `IgnoreCapacity` to a value of `true`, `1`, or `yes`, the Task will be routed to the Worker without respecting their capacity and availability. Any other value will enforce the Worker's capacity and availability. The default value of `IgnoreCapacity` is `true` when the `RoutingTarget` is set to a Worker SID.  |
**RoutingTarget** | Pointer to **string** | A SID of a Worker, Queue, or Workflow to route a Task to |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


