/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct for TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics
type TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The minimum, average, maximum, and total time that Workers spent in each Activity
	ActivityDurations *[]map[string]interface{} `json:"activity_durations,omitempty"`
	// The end of the interval during which these statistics were calculated
	EndTime *time.Time `json:"end_time,omitempty"`
	// The total number of Reservations that were accepted
	ReservationsAccepted *int `json:"reservations_accepted,omitempty"`
	// The total number of Reservations that were canceled
	ReservationsCanceled *int `json:"reservations_canceled,omitempty"`
	// The total number of Reservations that were created
	ReservationsCreated *int `json:"reservations_created,omitempty"`
	// The total number of Reservations that were rejected
	ReservationsRejected *int `json:"reservations_rejected,omitempty"`
	// The total number of Reservations that were rescinded
	ReservationsRescinded *int `json:"reservations_rescinded,omitempty"`
	// The total number of Reservations that were timed out
	ReservationsTimedOut *int `json:"reservations_timed_out,omitempty"`
	// The beginning of the interval during which these statistics were calculated
	StartTime *time.Time `json:"start_time,omitempty"`
	// The absolute URL of the Workers statistics resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Workers
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
