# TaskrouterV1TaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ActivityStatistics** | **[]interface{}** | The number of current Workers by Activity |[optional] 
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task |[optional] 
**LongestTaskWaitingSid** | **string** | The SID of the longest waiting Task |[optional] 
**LongestRelativeTaskAgeInQueue** | **int** | The relative age in the TaskQueue for the longest waiting Task. |[optional] 
**LongestRelativeTaskSidInQueue** | **string** | The SID of the Task waiting in the TaskQueue the longest. |[optional] 
**TaskQueueSid** | **string** | The SID of the TaskQueue from which these statistics were calculated |[optional] 
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status |
**TotalAvailableWorkers** | **int** | The total number of Workers available for Tasks in the TaskQueue |[optional] 
**TotalEligibleWorkers** | **int** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state |[optional] 
**TotalTasks** | **int** | The total number of Tasks |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the TaskQueue |[optional] 
**Url** | **string** | The absolute URL of the TaskQueue statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


