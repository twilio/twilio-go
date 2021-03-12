# TaskrouterV1WorkspaceWorkflowWorkflowRealTimeStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**LongestTaskWaitingAge** | Pointer to **int32** | The age of the longest waiting Task |
**LongestTaskWaitingSid** | Pointer to **string** | The SID of the longest waiting Task |
**TasksByPriority** | Pointer to **map[string]interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **map[string]interface{}** | The number of Tasks by their current status |
**TotalTasks** | Pointer to **int32** | The total number of Tasks |
**Url** | Pointer to **string** | The absolute URL of the Workflow statistics resource |
**WorkflowSid** | Pointer to **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


