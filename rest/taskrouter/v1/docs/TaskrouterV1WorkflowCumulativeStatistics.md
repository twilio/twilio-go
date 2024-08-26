# TaskrouterV1WorkflowCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource. |
**AvgTaskAcceptanceTime** | Pointer to **int** | The average time in seconds between Task creation and acceptance. |[default to 0]
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ReservationsCreated** | Pointer to **int** | The total number of Reservations that were created for Workers. |[default to 0]
**ReservationsAccepted** | Pointer to **int** | The total number of Reservations accepted by Workers. |[default to 0]
**ReservationsRejected** | Pointer to **int** | The total number of Reservations that were rejected. |[default to 0]
**ReservationsTimedOut** | Pointer to **int** | The total number of Reservations that were timed out. |[default to 0]
**ReservationsCanceled** | Pointer to **int** | The total number of Reservations that were canceled. |[default to 0]
**ReservationsRescinded** | Pointer to **int** | The total number of Reservations that were rescinded. |[default to 0]
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the number of Tasks canceled and reservations accepted above and below the thresholds specified in seconds. |
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were accepted. |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were canceled. |
**TasksCanceled** | Pointer to **int** | The total number of Tasks that were canceled. |[default to 0]
**TasksCompleted** | Pointer to **int** | The total number of Tasks that were completed. |[default to 0]
**TasksEntered** | Pointer to **int** | The total number of Tasks that entered the Workflow. |[default to 0]
**TasksDeleted** | Pointer to **int** | The total number of Tasks that were deleted. |[default to 0]
**TasksMoved** | Pointer to **int** | The total number of Tasks that were moved from one queue to another. |[default to 0]
**TasksTimedOutInWorkflow** | Pointer to **int** | The total number of Tasks that were timed out of their Workflows (and deleted). |[default to 0]
**WorkflowSid** | Pointer to **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |
**Url** | Pointer to **string** | The absolute URL of the Workflow statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


