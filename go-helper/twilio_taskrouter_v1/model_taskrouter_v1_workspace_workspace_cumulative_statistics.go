/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct for TaskrouterV1WorkspaceWorkspaceCumulativeStatistics
type TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct {
	AccountSid string `json:"account_sid,omitempty"`
	AvgTaskAcceptanceTime int32 `json:"avg_task_acceptance_time,omitempty"`
	EndTime time.Time `json:"end_time,omitempty"`
	ReservationsAccepted int32 `json:"reservations_accepted,omitempty"`
	ReservationsCanceled int32 `json:"reservations_canceled,omitempty"`
	ReservationsCreated int32 `json:"reservations_created,omitempty"`
	ReservationsRejected int32 `json:"reservations_rejected,omitempty"`
	ReservationsRescinded int32 `json:"reservations_rescinded,omitempty"`
	ReservationsTimedOut int32 `json:"reservations_timed_out,omitempty"`
	SplitByWaitTime map[string]interface{} `json:"split_by_wait_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	TasksCanceled int32 `json:"tasks_canceled,omitempty"`
	TasksCompleted int32 `json:"tasks_completed,omitempty"`
	TasksCreated int32 `json:"tasks_created,omitempty"`
	TasksDeleted int32 `json:"tasks_deleted,omitempty"`
	TasksMoved int32 `json:"tasks_moved,omitempty"`
	TasksTimedOutInWorkflow int32 `json:"tasks_timed_out_in_workflow,omitempty"`
	Url string `json:"url,omitempty"`
	WaitDurationUntilAccepted map[string]interface{} `json:"wait_duration_until_accepted,omitempty"`
	WaitDurationUntilCanceled map[string]interface{} `json:"wait_duration_until_canceled,omitempty"`
	WorkspaceSid string `json:"workspace_sid,omitempty"`
}
