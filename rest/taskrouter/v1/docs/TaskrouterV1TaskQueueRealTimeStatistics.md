# TaskrouterV1TaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource. |
**ActivityStatistics** | Pointer to **[]interface{}** | The number of current Workers by Activity. |
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task. |[optional] [default to 0]
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task. |
**LongestRelativeTaskAgeInQueue** | **int** | The relative age in the TaskQueue for the longest waiting Task. Calculation is based on the time when the Task entered the TaskQueue. |[optional] [default to 0]
**LongestRelativeTaskSidInQueue** | Pointer to **string** | The Task SID of the Task waiting in the TaskQueue the longest. Calculation is based on the time when the Task entered the TaskQueue. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated. |
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99. |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`. |
**TotalAvailableWorkers** | **int** | The total number of Workers in the TaskQueue with an `available` status. Workers with an `available` status may already have active interactions or may have none. |[optional] [default to 0]
**TotalEligibleWorkers** | **int** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state. |[optional] [default to 0]
**TotalTasks** | **int** | The total number of Tasks. |[optional] [default to 0]
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue. |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


