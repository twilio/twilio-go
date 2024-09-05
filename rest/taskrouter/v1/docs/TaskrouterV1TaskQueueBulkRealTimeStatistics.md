# TaskrouterV1TaskQueueBulkRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue. |
**TaskQueueData** | Pointer to **[]interface{}** | The real-time statistics for each requested TaskQueue SID. `task_queue_data` returns the following attributes:  `task_queue_sid`: The SID of the TaskQueue from which these statistics were calculated.  `total_available_workers`: The total number of Workers available for Tasks in the TaskQueue.  `total_eligible_workers`: The total number of Workers eligible for Tasks in the TaskQueue, regardless of their Activity state.  `total_tasks`: The total number of Tasks.  `longest_task_waiting_age`: The age of the longest waiting Task.  `longest_task_waiting_sid`: The SID of the longest waiting Task.  `tasks_by_status`: The number of Tasks grouped by their current status.  `tasks_by_priority`: The number of Tasks grouped by priority.  `activity_statistics`: The number of current Workers grouped by Activity.  |
**TaskQueueResponseCount** | **int** | The number of TaskQueue statistics received in task_queue_data. |[optional] [default to 0]
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


