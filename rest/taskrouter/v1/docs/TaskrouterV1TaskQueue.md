# TaskrouterV1TaskQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AssignmentActivitySid** | **string** | The SID of the Activity to assign Workers when a task is assigned for them |[optional] 
**AssignmentActivityName** | **string** | The name of the Activity to assign Workers when a task is assigned for them |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**MaxReservedWorkers** | **int** | The maximum number of Workers to reserve |[optional] 
**ReservationActivitySid** | **string** | The SID of the Activity to assign Workers once a task is reserved for them |[optional] 
**ReservationActivityName** | **string** | The name of the Activity to assign Workers once a task is reserved for them |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TargetWorkers** | **string** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue |[optional] 
**TaskOrder** | Pointer to [**string**](TaskQueueEnumTaskOrder.md) |  |
**Url** | **string** | The absolute URL of the TaskQueue resource |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the TaskQueue |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


