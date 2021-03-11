# TaskrouterV1WorkspaceWorkspaceCumulativeStatistics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AvgTaskAcceptanceTime** | Pointer to **int32** | The average time in seconds between Task creation and acceptance |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |
**ReservationsAccepted** | Pointer to **int32** | The total number of Reservations accepted by Workers |
**ReservationsCanceled** | Pointer to **int32** | The total number of Reservations that were canceled |
**ReservationsCreated** | Pointer to **int32** | The total number of Reservations that were created for Workers |
**ReservationsRejected** | Pointer to **int32** | The total number of Reservations that were rejected |
**ReservationsRescinded** | Pointer to **int32** | The total number of Reservations that were rescinded |
**ReservationsTimedOut** | Pointer to **int32** | The total number of Reservations that were timed out |
**SplitByWaitTime** | Pointer to **map[string]interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |
**TasksCanceled** | Pointer to **int32** | The total number of Tasks that were canceled |
**TasksCompleted** | Pointer to **int32** | The total number of Tasks that were completed |
**TasksCreated** | Pointer to **int32** | The total number of Tasks created |
**TasksDeleted** | Pointer to **int32** | The total number of Tasks that were deleted |
**TasksMoved** | Pointer to **int32** | The total number of Tasks that were moved from one queue to another |
**TasksTimedOutInWorkflow** | Pointer to **int32** | The total number of Tasks that were timed out of their Workflows |
**Url** | Pointer to **string** | The absolute URL of the Workspace statistics resource |
**WaitDurationUntilAccepted** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks that were accepted |
**WaitDurationUntilCanceled** | Pointer to **map[string]interface{}** | The wait duration statistics for Tasks that were canceled |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


