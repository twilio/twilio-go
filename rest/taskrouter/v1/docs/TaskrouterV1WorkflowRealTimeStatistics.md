# TaskrouterV1WorkflowRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task |[optional] 
**LongestTaskWaitingSid** | **string** | The SID of the longest waiting Task |[optional] 
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status |
**TotalTasks** | **int** | The total number of Tasks |[optional] 
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified SID value |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Workflow. |[optional] 
**Url** | **string** | The absolute URL of the Workflow statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


