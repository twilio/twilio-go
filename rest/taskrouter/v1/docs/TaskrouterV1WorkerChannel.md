# TaskrouterV1WorkerChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AssignedTasks** | **int** | The total number of Tasks assigned to Worker for the TaskChannel type |[optional] 
**Available** | **bool** | Whether the Worker should receive Tasks of the TaskChannel type |[optional] 
**AvailableCapacityPercentage** | **int** | The current available capacity between 0 to 100 for the TaskChannel |[optional] 
**ConfiguredCapacity** | **int** | The current configured capacity for the WorkerChannel |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TaskChannelSid** | **string** | The SID of the TaskChannel |[optional] 
**TaskChannelUniqueName** | **string** | The unique name of the TaskChannel, such as 'voice' or 'sms' |[optional] 
**WorkerSid** | **string** | The SID of the Worker that contains the WorkerChannel |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the WorkerChannel |[optional] 
**Url** | **string** | The absolute URL of the WorkerChannel resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


