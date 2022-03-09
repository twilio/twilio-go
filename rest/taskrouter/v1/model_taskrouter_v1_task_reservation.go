/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1TaskReservation struct for TaskrouterV1TaskReservation
type TaskrouterV1TaskReservation struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The current status of the reservation
	ReservationStatus *string `json:"reservation_status,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the reserved Task resource
	TaskSid *string `json:"task_sid,omitempty"`
	// The absolute URL of the TaskReservation reservation
	Url *string `json:"url,omitempty"`
	// The friendly_name of the Worker that is reserved
	WorkerName *string `json:"worker_name,omitempty"`
	// The SID of the reserved Worker resource
	WorkerSid *string `json:"worker_sid,omitempty"`
	// The SID of the Workspace that this task is contained within.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
