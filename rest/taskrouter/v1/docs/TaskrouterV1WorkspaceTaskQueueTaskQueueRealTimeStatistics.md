# TaskrouterV1WorkspaceTaskQueueTaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActivityStatistics** | Pointer to **[]map[string]interface{}** | The number of current Workers by Activity |
**LongestRelativeTaskAgeInQueue** | Pointer to **int** | The relative age in the TaskQueue for the longest waiting Task. |
**LongestRelativeTaskSidInQueue** | Pointer to **string** | The SID of the Task waiting in the TaskQueue the longest. |
**LongestTaskWaitingAge** | Pointer to **int** | The age of the longest waiting Task |
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated |
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status |
**TotalAvailableWorkers** | Pointer to **int** | The total number of Workers available for Tasks in the TaskQueue |
**TotalEligibleWorkers** | Pointer to **int** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state |
**TotalTasks** | Pointer to **int** | The total number of Tasks |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


