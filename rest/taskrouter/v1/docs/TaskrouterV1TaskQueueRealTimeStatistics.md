# TaskrouterV1TaskQueueRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource. |
**ActivityStatistics** | Pointer to **[]interface{}** | The number of current Workers by Activity. |
**LongestTaskWaitingAge** | Pointer to **int** | The age of the longest waiting Task. |
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task. |
**LongestRelativeTaskAgeInQueue** | Pointer to **int** | The relative age in the TaskQueue for the longest waiting Task. Calculation is based on the time when the Task entered the TaskQueue. |
**LongestRelativeTaskSidInQueue** | Pointer to **string** | The Task SID of the Task waiting in the TaskQueue the longest. Calculation is based on the time when the Task entered the TaskQueue. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated. |
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99. |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`. |
**TotalAvailableWorkers** | Pointer to **int** | The total number of Workers available for Tasks in the TaskQueue. |
**TotalEligibleWorkers** | Pointer to **int** | The total number of Workers eligible for Tasks in the TaskQueue, independent of their Activity state. |
**TotalTasks** | Pointer to **int** | The total number of Tasks. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue. |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


