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

// TaskrouterV1WorkspaceTaskQueue struct for TaskrouterV1WorkspaceTaskQueue
type TaskrouterV1WorkspaceTaskQueue struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The name of the Activity to assign Workers when a task is assigned for them
	AssignmentActivityName *string `json:"AssignmentActivityName,omitempty"`
	// The SID of the Activity to assign Workers when a task is assigned for them
	AssignmentActivitySid *string `json:"AssignmentActivitySid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The maximum number of Workers to reserve
	MaxReservedWorkers *int32 `json:"MaxReservedWorkers,omitempty"`
	// The name of the Activity to assign Workers once a task is reserved for them
	ReservationActivityName *string `json:"ReservationActivityName,omitempty"`
	// The SID of the Activity to assign Workers once a task is reserved for them
	ReservationActivitySid *string `json:"ReservationActivitySid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// A string describing the Worker selection criteria for any Tasks that enter the TaskQueue
	TargetWorkers *string `json:"TargetWorkers,omitempty"`
	// How Tasks will be assigned to Workers
	TaskOrder *string `json:"TaskOrder,omitempty"`
	// The absolute URL of the TaskQueue resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Workspace that contains the TaskQueue
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
