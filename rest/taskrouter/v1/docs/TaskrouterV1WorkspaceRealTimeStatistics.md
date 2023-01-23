# TaskrouterV1WorkspaceRealTimeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ActivityStatistics** | **[]interface{}** | The number of current Workers by Activity |[optional] 
**LongestTaskWaitingAge** | **int** | The age of the longest waiting Task |[optional] 
**LongestTaskWaitingSid** | **string** | The SID of the longest waiting Task |[optional] 
**TasksByPriority** | Pointer to **interface{}** | The number of Tasks by priority |
**TasksByStatus** | Pointer to **interface{}** | The number of Tasks by their current status |
**TotalTasks** | **int** | The total number of Tasks |[optional] 
**TotalWorkers** | **int** | The total number of Workers in the Workspace |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace |[optional] 
**Url** | **string** | The absolute URL of the Workspace statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


