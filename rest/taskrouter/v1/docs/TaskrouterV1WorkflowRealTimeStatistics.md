# TaskrouterV1WorkflowRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource. |
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task. |[optional] [default to 0]
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task. |
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority. For example: `{\"0\": \"10\", \"99\": \"5\"}` shows 10 Tasks at priority 0 and 5 at priority 99. |
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status. For example: `{\"pending\": \"1\", \"reserved\": \"3\", \"assigned\": \"2\", \"completed\": \"5\"}`. |
**TotalTasks** | **int** | The total number of Tasks. |[optional] [default to 0]
**WorkflowSid** | Pointer to **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |
**Url** | Pointer to **string** | The absolute URL of the Workflow statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


