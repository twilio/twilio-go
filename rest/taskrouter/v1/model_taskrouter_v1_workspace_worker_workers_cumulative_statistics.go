/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct for TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics
type TaskrouterV1WorkspaceWorkerWorkersCumulativeStatistics struct {
	AccountSid            *string                   `json:"account_sid,omitempty"`
	ActivityDurations     *[]map[string]interface{} `json:"activity_durations,omitempty"`
	EndTime               *time.Time                `json:"end_time,omitempty"`
	ReservationsAccepted  *int32                    `json:"reservations_accepted,omitempty"`
	ReservationsCanceled  *int32                    `json:"reservations_canceled,omitempty"`
	ReservationsCreated   *int32                    `json:"reservations_created,omitempty"`
	ReservationsRejected  *int32                    `json:"reservations_rejected,omitempty"`
	ReservationsRescinded *int32                    `json:"reservations_rescinded,omitempty"`
	ReservationsTimedOut  *int32                    `json:"reservations_timed_out,omitempty"`
	StartTime             *time.Time                `json:"start_time,omitempty"`
	Url                   *string                   `json:"url,omitempty"`
	WorkspaceSid          *string                   `json:"workspace_sid,omitempty"`
}
