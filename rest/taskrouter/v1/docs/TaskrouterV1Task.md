# TaskrouterV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Age** | Pointer to **int** | The number of seconds since the Task was created |
**AssignmentStatus** | Pointer to [**string**](TaskEnumStatus.md) |  |
**Attributes** | Pointer to **string** | The JSON string with custom attributes of the work |
**Addons** | Pointer to **string** | An object that contains the addon data for all installed addons |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**TaskQueueEnteredDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Task entered the TaskQueue. |
**Priority** | Pointer to **int** | Retrieve the list of all Tasks in the Workspace with the specified priority |
**Reason** | Pointer to **string** | The reason the Task was canceled or completed |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue |
**TaskQueueFriendlyName** | Pointer to **string** | The friendly name of the TaskQueue |
**TaskChannelSid** | Pointer to **string** | The SID of the TaskChannel |
**TaskChannelUniqueName** | Pointer to **string** | The unique name of the TaskChannel |
**Timeout** | Pointer to **int** | The amount of time in seconds that the Task can live before being assigned |
**WorkflowSid** | Pointer to **string** | The SID of the Workflow that is controlling the Task |
**WorkflowFriendlyName** | Pointer to **string** | The friendly name of the Workflow that is controlling the Task |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Task |
**Url** | Pointer to **string** | The absolute URL of the Task resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


