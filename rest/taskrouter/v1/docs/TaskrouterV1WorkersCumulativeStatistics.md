# TaskrouterV1WorkersCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Worker resource. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ActivityDurations** | Pointer to **[]interface{}** | The minimum, average, maximum, and total time (in seconds) that Workers spent in each Activity. |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations that were created. |[default to 0]
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations that were accepted. |[default to 0]
**ReservationsRejected** | Pointer to **int** | The total number of Reservations that were rejected. |[default to 0]
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that were timed out. |[default to 0]
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations that were canceled. |[default to 0]
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations that were rescinded. |[default to 0]
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workers. |
**Url** | Pointer to **string** | The absolute URL of the Workers statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


