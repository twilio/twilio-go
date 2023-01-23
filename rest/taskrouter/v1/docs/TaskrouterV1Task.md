# TaskrouterV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Task resource. |
**Age** | Pointer to **int** | The number of seconds since the Task was created. |
**AssignmentStatus** | Pointer to [**string**](TaskEnumStatus.md) |  |
**Attributes** | Pointer to **string** | The JSON string with custom attributes of the work. **Note** If this property has been assigned a value, it will only be displayed in FETCH action that returns a single resource. Otherwise, it will be null. |
**Addons** | Pointer to **string** | An object that contains the [addon](https://www.twilio.com/docs/taskrouter/marketplace) data for all installed addons. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**TaskQueueEnteredDate** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Task entered the TaskQueue, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Priority** | Pointer to **int** | The current priority score of the Task as assigned to a Worker by the workflow. Tasks with higher priority values will be assigned before Tasks with lower values. |
**Reason** | Pointer to **string** | The reason the Task was canceled or completed, if applicable. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Task resource. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue. |
**TaskQueueFriendlyName** | Pointer to **string** | The friendly name of the TaskQueue. |
**TaskChannelSid** | Pointer to **string** | The SID of the TaskChannel. |
**TaskChannelUniqueName** | Pointer to **string** | The unique name of the TaskChannel. |
**Timeout** | Pointer to **int** | The amount of time in seconds that the Task can live before being assigned. |
**WorkflowSid** | Pointer to **string** | The SID of the Workflow that is controlling the Task. |
**WorkflowFriendlyName** | Pointer to **string** | The friendly name of the Workflow that is controlling the Task. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Task. |
**Url** | Pointer to **string** | The absolute URL of the Task resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


