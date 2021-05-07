# TaskrouterV1WorkspaceTaskQueueTaskQueueCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AvgTaskAcceptanceTime** | Pointer to **int32** | The average time in seconds between Task creation and acceptance |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |
**ReservationsAccepted** | Pointer to **int32** | The total number of Reservations accepted for Tasks in the TaskQueue |
**ReservationsCanceled** | Pointer to **int32** | The total number of Reservations canceled for Tasks in the TaskQueue |
**ReservationsCreated** | Pointer to **int32** | The total number of Reservations created for Tasks in the TaskQueue |
**ReservationsRejected** | Pointer to **int32** | The total number of Reservations rejected for Tasks in the TaskQueue |
**ReservationsRescinded** | Pointer to **int32** | The total number of Reservations rescinded |
**ReservationsTimedOut** | Pointer to **int32** | The total number of Reservations that timed out for Tasks in the TaskQueue |
**SplitByWaitTime** | Pointer to **map[string]interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |
**TaskQueueSid** | Pointer to **string** | The SID of the TaskQueue from which these statistics were calculated |
**TasksCanceled** | Pointer to **int32** | The total number of Tasks canceled in the TaskQueue |
**TasksCompleted** | Pointer to **int32** | The total number of Tasks completed in the TaskQueue |
**TasksDeleted** | Pointer to **int32** | The total number of Tasks deleted in the TaskQueue |
**TasksEntered** | Pointer to **int32** | The total number of Tasks entered into the TaskQueue |
**TasksMoved** | Pointer to **int32** | The total number of Tasks that were moved from one queue to another |
**Url** | Pointer to **string** | The absolute URL of the TaskQueue statistics resource |
**WaitDurationInQueueUntilAccepted** | Pointer to **map[string]interface{}** | The relative wait duration statistics for Tasks accepted while in the TaskQueue |
**WaitDurationUntilAccepted** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks accepted while in the TaskQueue |
**WaitDurationUntilCanceled** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks canceled while in the TaskQueue |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the TaskQueue |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


