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

// TaskrouterV1WorkspaceWorkflow struct for TaskrouterV1WorkspaceWorkflow
type TaskrouterV1WorkspaceWorkflow struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The URL that we call when a task managed by the Workflow is assigned to a Worker
	AssignmentCallbackUrl *string `json:"AssignmentCallbackUrl,omitempty"`
	// A JSON string that contains the Workflow's configuration
	Configuration *string `json:"Configuration,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The MIME type of the document
	DocumentContentType *string `json:"DocumentContentType,omitempty"`
	// The URL that we call when a call to the `assignment_callback_url` fails
	FallbackAssignmentCallbackUrl *string `json:"FallbackAssignmentCallbackUrl,omitempty"`
	// The string that you assigned to describe the Workflow resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// How long TaskRouter will wait for a confirmation response from your application after it assigns a Task to a Worker
	TaskReservationTimeout *int32 `json:"TaskReservationTimeout,omitempty"`
	// The absolute URL of the Workflow resource
	Url *string `json:"Url,omitempty"`
	// The SID of the Workspace that contains the Workflow
	WorkspaceSid *string `json:"WorkspaceSid,omitempty"`
}
