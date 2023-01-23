# TaskrouterV1WorkerReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**ReservationStatus** | Pointer to [**string**](WorkerReservationEnumStatus.md) |  |
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**TaskSid** | **string** | The SID of the reserved Task resource |[optional] 
**WorkerName** | **string** | The friendly_name of the Worker that is reserved |[optional] 
**WorkerSid** | **string** | The SID of the reserved Worker resource |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that this worker is contained within. |[optional] 
**Url** | **string** | The absolute URL of the WorkerReservation resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of related resources |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


