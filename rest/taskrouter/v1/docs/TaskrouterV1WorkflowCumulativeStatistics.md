# TaskrouterV1WorkflowCumulativeStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Workflow resource. |
**AvgTaskAcceptanceTime** | **int** | The average time in seconds between Task creation and acceptance. |[optional] [default to 0]
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The beginning of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The end of the interval during which these statistics were calculated, in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**ReservationsCreated** | **int** | The total number of Reservations that were created for Workers. |[optional] [default to 0]
**ReservationsAccepted** | **int** | The total number of Reservations accepted by Workers. |[optional] [default to 0]
**ReservationsRejected** | **int** | The total number of Reservations that were rejected. |[optional] [default to 0]
**ReservationsTimedOut** | **int** | The total number of Reservations that were timed out. |[optional] [default to 0]
**ReservationsCanceled** | **int** | The total number of Reservations that were canceled. |[optional] [default to 0]
**ReservationsRescinded** | **int** | The total number of Reservations that were rescinded. |[optional] [default to 0]
**SplitByWaitTime** | Pointer to **interface{}** | A list of objects that describe the number of Tasks canceled and reservations accepted above and below the thresholds specified in seconds. |
**WaitDurationUntilAccepted** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were accepted. |
**WaitDurationUntilCanceled** | Pointer to **interface{}** | The wait duration statistics (`avg`, `min`, `max`, `total`) for Tasks that were canceled. |
**TasksCanceled** | **int** | The total number of Tasks that were canceled. |[optional] [default to 0]
**TasksCompleted** | **int** | The total number of Tasks that were completed. |[optional] [default to 0]
**TasksEntered** | **int** | The total number of Tasks that entered the Workflow. |[optional] [default to 0]
**TasksDeleted** | **int** | The total number of Tasks that were deleted. |[optional] [default to 0]
**TasksMoved** | **int** | The total number of Tasks that were moved from one queue to another. |[optional] [default to 0]
**TasksTimedOutInWorkflow** | **int** | The total number of Tasks that were timed out of their Workflows (and deleted). |[optional] [default to 0]
**WorkflowSid** | Pointer to **string** | Returns the list of Tasks that are being controlled by the Workflow with the specified Sid value. |
**WorkspaceSid** | Pointer to **string** | The SID of the Workspace that contains the Workflow. |
**Url** | Pointer to **string** | The absolute URL of the Workflow statistics resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


