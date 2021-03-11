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

// TaskrouterV1WorkspaceTask struct for TaskrouterV1WorkspaceTask
type TaskrouterV1WorkspaceTask struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// An object that contains the addon data for all installed addons
	Addons *string `json:"Addons,omitempty"`
	// The number of seconds since the Task was created
	Age *int32 `json:"Age,omitempty"`
	// The current status of the Task's assignment
	AssignmentStatus *string `json:"AssignmentStatus,omitempty"`
	// The JSON string with custom attributes of the work
	Attributes *string `json:"Attributes,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// Retrieve the list of all Tasks in the Workspace with the specified priority
	Priority *int32 `json:"Priority,omitempty"`
	// The reason the Task was canceled or completed
	Reason *string `json:"Reason,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The SID of the TaskChannel
	TaskChannelSid *string `json:"TaskChannelSid,omitempty"`
	// The unique name of the TaskChannel
	TaskChannelUniqueName *string `json:"TaskChannelUniqueName,omitempty"`
	// The ISO 8601 date and time in GMT when the Task entered the TaskQueue.
	TaskQueueEnteredDate *time.Time `json:"TaskQueueEnteredDate,omitempty"`
	// The friendly name of the TaskQueue
	TaskQueueFriendlyName *string `json:"TaskQueueFriendlyName,omitempty"`
	// The SID of the TaskQueue
	TaskQueueSid *string `json:"TaskQueueSid,omitempty"`
	// The amount of time in seconds that the Task can live before being assigned
	Timeout *int32 `json:"Timeout,omitempty"`
	// The absolute URL of the Task resource
	Url *string `json:"Url,omitempty"`
	// The friendly name of the Workflow that is controlling the Task
	WorkflowFriendlyName *string `json:"WorkflowFriendlyName,omitempty"`
	// The SID of the Workflow that is controlling the Task
	WorkflowSid *string `json:"WorkflowSid,omitempty"`
	// The SID of the Workspace that contains the Task
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
