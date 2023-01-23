# TaskrouterV1Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Age** | **int** | The number of seconds since the Task was created |[optional] 
**AssignmentStatus** | Pointer to [**string**](TaskEnumStatus.md) |  |
**Attributes** | **string** | The JSON string with custom attributes of the work |[optional] 
**Addons** | **string** | An object that contains the addon data for all installed addons |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**TaskQueueEnteredDate** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the Task entered the TaskQueue. |[optional] 
**Priority** | **int** | Retrieve the list of all Tasks in the Workspace with the specified priority |[optional] 
**Reason** | **string** | The reason the Task was canceled or completed |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TaskQueueSid** | **string** | The SID of the TaskQueue |[optional] 
**TaskQueueFriendlyName** | **string** | The friendly name of the TaskQueue |[optional] 
**TaskChannelSid** | **string** | The SID of the TaskChannel |[optional] 
**TaskChannelUniqueName** | **string** | The unique name of the TaskChannel |[optional] 
**Timeout** | **int** | The amount of time in seconds that the Task can live before being assigned |[optional] 
**WorkflowSid** | **string** | The SID of the Workflow that is controlling the Task |[optional] 
**WorkflowFriendlyName** | **string** | The friendly name of the Workflow that is controlling the Task |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Task |[optional] 
**Url** | **string** | The absolute URL of the Task resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


