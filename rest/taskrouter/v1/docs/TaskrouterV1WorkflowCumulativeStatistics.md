# TaskrouterV1WorkflowCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**AvgTaskAcceptanceTime** | **int** | The average time in seconds between Task creation and acceptance |[optional] 
**StartTime** | [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated |[optional] 
**EndTime** | [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated |[optional] 
**ReservationsCreated** | **int** | The total number of Reservations that were created for Workers |[optional] 
**ReservationsAccepted** | **int** | The total number of Reservations accepted by Workers |[optional] 
**ReservationsRejected** | **int** | The total number of Reservations that were rejected |[optional] 
**ReservationsTimedOut** | **int** | The total number of Reservations that were timed out |[optional] 
**ReservationsCanceled** | **int** | The total number of Reservations that were canceled |[optional] 
**ReservationsRescinded** | **int** | The total number of Reservations that were rescinded |[optional] 
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the Tasks canceled and reservations accepted above and below the specified thresholds |
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics for Tasks that were accepted |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics for Tasks that were canceled |
**TasksCanceled** | **int** | The total number of Tasks that were canceled |[optional] 
**TasksCompleted** | **int** | The total number of Tasks that were completed |[optional] 
**TasksEntered** | **int** | The total number of Tasks that entered the Workflow |[optional] 
**TasksDeleted** | **int** | The total number of Tasks that were deleted |[optional] 
**TasksMoved** | **int** | The total number of Tasks that were moved from one queue to another |[optional] 
**TasksTimedOutInWorkflow** | **int** | The total number of Tasks that were timed out of their Workflows |[optional] 
**WorkflowSid** | **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value |[optional] 
**WorkspaceSid** | **string** | The SID of the Workspace that contains the Workflow. |[optional] 
**Url** | **string** | The absolute URL of the Workflow statistics resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


