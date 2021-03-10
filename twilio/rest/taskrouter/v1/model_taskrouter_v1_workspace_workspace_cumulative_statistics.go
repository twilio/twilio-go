/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct for TaskrouterV1WorkspaceWorkspaceCumulativeStatistics
type TaskrouterV1WorkspaceWorkspaceCumulativeStatistics struct {
	AccountSid                *string                 `json:"AccountSid,omitempty"`
	AvgTaskAcceptanceTime     *int32                  `json:"AvgTaskAcceptanceTime,omitempty"`
	EndTime                   *time.Time              `json:"EndTime,omitempty"`
	ReservationsAccepted      *int32                  `json:"ReservationsAccepted,omitempty"`
	ReservationsCanceled      *int32                  `json:"ReservationsCanceled,omitempty"`
	ReservationsCreated       *int32                  `json:"ReservationsCreated,omitempty"`
	ReservationsRejected      *int32                  `json:"ReservationsRejected,omitempty"`
	ReservationsRescinded     *int32                  `json:"ReservationsRescinded,omitempty"`
	ReservationsTimedOut      *int32                  `json:"ReservationsTimedOut,omitempty"`
	SplitByWaitTime           *map[string]interface{} `json:"SplitByWaitTime,omitempty"`
	StartTime                 *time.Time              `json:"StartTime,omitempty"`
	TasksCanceled             *int32                  `json:"TasksCanceled,omitempty"`
	TasksCompleted            *int32                  `json:"TasksCompleted,omitempty"`
	TasksCreated              *int32                  `json:"TasksCreated,omitempty"`
	TasksDeleted              *int32                  `json:"TasksDeleted,omitempty"`
	TasksMoved                *int32                  `json:"TasksMoved,omitempty"`
	TasksTimedOutInWorkflow   *int32                  `json:"TasksTimedOutInWorkflow,omitempty"`
	Url                       *string                 `json:"Url,omitempty"`
	WaitDurationUntilAccepted *map[string]interface{} `json:"WaitDurationUntilAccepted,omitempty"`
	WaitDurationUntilCanceled *map[string]interface{} `json:"WaitDurationUntilCanceled,omitempty"`
	WorkspaceSid              *string                 `json:"WorkspaceSid,omitempty"`
}
