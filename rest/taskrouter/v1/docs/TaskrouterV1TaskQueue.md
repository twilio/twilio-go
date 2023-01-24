# TaskrouterV1TaskQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource. |
**AssignmentActivitySid** | Pointer to **string** | The SID of the Activity to assign Workers when a task is assigned for them. |
**AssignmentActivityName** | Pointer to **string** | The name of the Activity to assign Workers when a task is assigned for them. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**MaxReservedWorkers** | Pointer to **int** | The maximum number of Workers to reserve for the assignment of a task in the queue. Can be an integer between 1 and 50, inclusive and defaults to 1. |
**ReservationActivitySid** | Pointer to **string** | The SID of the Activity to assign Workers once a task is reserved for them. |
**ReservationActivityName** | Pointer to **string** | The name of the Activity to assign Workers once a task is reserved for them. |
**Sid** | Pointer to **string** | The unique string that we created to identify the TaskQueue resource. |
**TargetWorkers** | Pointer to **string** | A string describing the Worker selection criteria for any Tasks that enter the TaskQueue. For example `'\"language\" == \"spanish\"'` If no TargetWorkers parameter is provided, Tasks will wait in the TaskQueue until they are either deleted or moved to another TaskQueue. Additional examples on how to describing Worker selection criteria below. Defaults to 1==1. |
**TaskOrder** | Pointer to [**string**](TaskQueueEnumTaskOrder.md) |  |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue resource. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


