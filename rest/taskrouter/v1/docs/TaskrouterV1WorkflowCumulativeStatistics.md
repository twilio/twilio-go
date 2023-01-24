# TaskrouterV1WorkflowCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource. |
**AvgTaskAcceptanceTime** | Pointer to **int** | The average time in seconds between Task creation and acceptance. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations that were created for Workers. |
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations accepted by Workers. |
**ReservationsRejected** | Pointer to **int** | The total number of Reservations that were rejected. |
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that were timed out. |
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations that were canceled. |
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations that were rescinded. |
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the number of Tasks canceled and reservations accepted above and below the thresholds specified in seconds. |
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were accepted. |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were canceled. |
**TasksCanceled** | Pointer to **int** | The total number of Tasks that were canceled. |
**TasksCompleted** | Pointer to **int** | The total number of Tasks that were completed. |
**TasksEntered** | Pointer to **int** | The total number of Tasks that entered the Workflow. |
**TasksDeleted** | Pointer to **int** | The total number of Tasks that were deleted. |
**TasksMoved** | Pointer to **int** | The total number of Tasks that were moved from one queue to another. |
**TasksTimedOutInWorkflow** | Pointer to **int** | The total number of Tasks that were timed out of their Workflows (and deleted). |
**WorkflowSid** | Pointer to **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |
**Url** | Pointer to **string** | The absolute URL of the Workflow statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


