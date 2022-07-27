# TaskrouterV1TaskQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AssignmentActivitySid** | Pointer to **string** | The SID of the Activity to assign Workers when a task is assigned for them |
**AssignmentActivityName** | Pointer to **string** | The name of the Activity to assign Workers when a task is assigned for them |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**MaxReservedWorkers** | Pointer to **int** | The maximum number of Workers to reserve |
**ReservationActivitySid** | Pointer to **string** | The SID of the Activity to assign Workers once a task is reserved for them |
**ReservationActivityName** | Pointer to **string** | The name of the Activity to assign Workers once a task is reserved for them |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TargetWorkers** | Pointer to **string** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue |
**TaskOrder** | Pointer to [**string**](TaskQueueEnumTaskOrder.md) |  |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


