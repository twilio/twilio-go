# TaskrouterV1TaskQueueCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AvgTaskAcceptanceTime** | **int** | The average time in seconds between Task creation and acceptance |[optional] 
**StartTime** | [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |[optional] 
**ReservationsCreated** | **int** | The total number of Reservations created for Tasks in the TaskQueue |[optional] 
**ReservationsAccepted** | **int** | The total number of Reservations accepted for Tasks in the TaskQueue |[optional] 
**ReservationsRejected** | **int** | The total number of Reservations rejected for Tasks in the TaskQueue |[optional] 
**ReservationsTimedOut** | **int** | The total number of Reservations that timed out for Tasks in the TaskQueue |[optional] 
**ReservationsCanceled** | **int** | The total number of Reservations canceled for Tasks in the TaskQueue |[optional] 
**ReservationsRescinded** | **int** | The total number of Reservations rescinded |[optional] 
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds |
**TaskQueueSid** | **string** | The SID of the TaskQueue from which these statistics were calculated |[optional] 
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics for Tasks accepted while in the TaskQueue |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics for Tasks canceled while in the TaskQueue |
**WaitDurationInQueueUntilAccepted** | Pointer to **interface{}** | The relative wait duration statistics for Tasks accepted while in the TaskQueue |
**TasksCanceled** | **int** | The total number of Tasks canceled in the TaskQueue |[optional] 
**TasksCompleted** | **int** | The total number of Tasks completed in the TaskQueue |[optional] 
**TasksDeleted** | **int** | The total number of Tasks deleted in the TaskQueue |[optional] 
**TasksEntered** | **int** | The total number of Tasks entered into the TaskQueue |[optional] 
**TasksMoved** | **int** | The total number of Tasks that were moved from one queue to another |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the TaskQueue |[optional] 
**Url** | **string** | The absolute URL of the TaskQueue statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


