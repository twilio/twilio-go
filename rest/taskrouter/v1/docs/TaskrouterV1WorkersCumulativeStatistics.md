# TaskrouterV1WorkersCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ActivityDurations** | Pointer to **[]map[string]interface{}** | The minimum, average, maximum, and total time that Workers spent in each Activity |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations that were accepted |
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations that were canceled |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations that were created |
**ReservationsRejected** | Pointer to **int** | The total number of Reservations that were rejected |
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations that were rescinded |
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that were timed out |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |
**Url** | Pointer to **string** | The absolute URL of the Workers statistics resource |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workers |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


