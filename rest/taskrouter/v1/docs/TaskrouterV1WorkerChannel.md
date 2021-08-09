# TaskrouterV1WorkerChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AssignedTasks** | Pointer to **int** | The total number of Tasks assigned to Worker for the TaskChannel type |
**Available** | Pointer to **bool** | Whether the Worker should receive Tasks of the TaskChannel type |
**AvailableCapacityPercentage** | Pointer to **int** | The current available capacity between 0 to 100 for the TaskChannel |
**ConfiguredCapacity** | Pointer to **int** | The current configured capacity for the WorkerChannel |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TaskChannelSid** | Pointer to **string** | The SID of the TaskChannel |
**TaskChannelUniqueName** | Pointer to **string** | The unique name of the TaskChannel, such as 'voice' or 'sms' |
**Url** | Pointer to **string** | The absolute URL of the WorkerChannel resource |
**WorkerSid** | Pointer to **string** | The SID of the Worker that contains the WorkerChannel |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the WorkerChannel |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


