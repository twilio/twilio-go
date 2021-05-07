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

// TaskrouterV1Workspace struct for TaskrouterV1Workspace
type TaskrouterV1Workspace struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The name of the default activity
	DefaultActivityName *string `json:"default_activity_name,omitempty"`
	// The SID of the Activity that will be used when new Workers are created in the Workspace
	DefaultActivitySid *string `json:"default_activity_sid,omitempty"`
	// The URL we call when an event occurs
	EventCallbackUrl *string `json:"event_callback_url,omitempty"`
	// The list of Workspace events for which to call event_callback_url
	EventsFilter *string `json:"events_filter,omitempty"`
	// The string that you assigned to describe the Workspace resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// Whether multi-tasking is enabled
	MultiTaskEnabled *bool `json:"multi_task_enabled,omitempty"`
	// The type of TaskQueue to prioritize when Workers are receiving Tasks from both types of TaskQueues
	PrioritizeQueueOrder *string `json:"prioritize_queue_order,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The name of the timeout activity
	TimeoutActivityName *string `json:"timeout_activity_name,omitempty"`
	// The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response
	TimeoutActivitySid *string `json:"timeout_activity_sid,omitempty"`
	// The absolute URL of the Workspace resource
	Url *string `json:"url,omitempty"`
}
