# TaskrouterV1TaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActivityStatistics** | Pointer to **[]interface{}** | The number of current Workers by Activity |
**LongestTaskWaitingAge** | Pointer to **int** | The age of the longest waiting Task |
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task |
**LongestRelativeTaskAgeInQueue** | Pointer to **int** | The relative age in the TaskQueue for the longest waiting Task. |
**LongestRelativeTaskSidInQueue** | Pointer to **string** | The SID of the Task waiting in the TaskQueue the longest. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated |
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status |
**TotalAvailableWorkers** | Pointer to **int** | The total number of Workers available for Tasks in the TaskQueue |
**TotalEligibleWorkers** | Pointer to **int** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state |
**TotalTasks** | Pointer to **int** | The total number of Tasks |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


