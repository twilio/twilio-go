# TaskrouterV1WorkersCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |
**ActivityDurations** | Pointer to **[]interface{}** | The minimum, average, maximum, and total time that Workers spent in each Activity |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations that were created |
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations that were accepted |
**ReservationsRejected** | Pointer to **int** | The total number of Reservations that were rejected |
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that were timed out |
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations that were canceled |
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations that were rescinded |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workers |
**Url** | Pointer to **string** | The absolute URL of the Workers statistics resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


