# TaskrouterV1WorkspaceRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workspace resource. |
**ActivityStatistics** | Pointer to **[]interface{}** | The number of current Workers by Activity. |
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task. |[optional] [default to 0]
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task. |
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99. |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`. |
**TotalTasks** | **int** | The total number of Tasks. |[optional] [default to 0]
**TotalWorkers** | **int** | The total number of Workers in the Workspace. |[optional] [default to 0]
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace. |
**Url** | Pointer to **string** | The absolute URL of the Workspace statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


