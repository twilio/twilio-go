# TaskrouterV1TaskQueueCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the TaskQueue resource. |
**AvgTaskAcceptanceTime** | Pointer to **int** | The average time in seconds between Task creation and acceptance. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations created for Tasks in the TaskQueue. |
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations accepted for Tasks in the TaskQueue. |
**ReservationsRejected** | Pointer to **int** | The total number of Reservations rejected for Tasks in the TaskQueue. |
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that timed out for Tasks in the TaskQueue. |
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations canceled for Tasks in the TaskQueue. |
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations rescinded. |
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the number of Tasks canceled and reservations accepted above and below the thresholds specified in seconds. |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated. |
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks accepted while in the TaskQueue. Calculation is based on the time when the Tasks were created. For transfers, the wait duration is counted from the moment ***the Task was created***, and not from when the transfer was initiated. |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks canceled while in the TaskQueue. |
**WaitDurationInQueueUntilAccepted** | Pointer to **interface{}** | The relative wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks accepted while in the TaskQueue. Calculation is based on the time when the Tasks entered the TaskQueue. |
**TasksCanceled** | Pointer to **int** | The total number of Tasks canceled in the TaskQueue. |
**TasksCompleted** | Pointer to **int** | The total number of Tasks completed in the TaskQueue. |
**TasksDeleted** | Pointer to **int** | The total number of Tasks deleted in the TaskQueue. |
**TasksEntered** | Pointer to **int** | The total number of Tasks entered into the TaskQueue. |
**TasksMoved** | Pointer to **int** | The total number of Tasks that were moved from one queue to another. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue. |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


