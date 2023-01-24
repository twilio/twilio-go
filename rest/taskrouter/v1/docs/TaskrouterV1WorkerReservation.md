# TaskrouterV1WorkerReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the WorkerReservation resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ReservationStatus** | Pointer to [**string**](WorkerReservationEnumStatus.md) |  |
**Sid** | Pointer to **string** | The unique string that we created to identify the WorkerReservation resource. |
**TaskSid** | Pointer to **string** | The SID of the reserved Task resource. |
**WorkerName** | Pointer to **string** | The `friendly_name` of the Worker that is reserved. |
**WorkerSid** | Pointer to **string** | The SID of the reserved Worker resource. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that this worker is contained within. |
**Url** | Pointer to **string** | The absolute URL of the WorkerReservation resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


