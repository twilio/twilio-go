# TaskrouterV1TaskReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**ReservationStatus** | Pointer to [**string**](TaskReservationEnumStatus.md) |  |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**TaskSid** | Pointer to **string** | The SID of the reserved Task resource |
**WorkerName** | Pointer to **string** | The friendly_name of the Worker that is reserved |
**WorkerSid** | Pointer to **string** | The SID of the reserved Worker resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that this task is contained within. |
**Url** | Pointer to **string** | The absolute URL of the TaskReservation reservation |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


