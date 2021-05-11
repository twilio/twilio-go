/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1Workspace struct for TaskrouterV1Workspace
type TaskrouterV1Workspace struct {
	AccountSid           *string                 `json:"account_sid,omitempty"`
	DateCreated          *time.Time              `json:"date_created,omitempty"`
	DateUpdated          *time.Time              `json:"date_updated,omitempty"`
	DefaultActivityName  *string                 `json:"default_activity_name,omitempty"`
	DefaultActivitySid   *string                 `json:"default_activity_sid,omitempty"`
	EventCallbackUrl     *string                 `json:"event_callback_url,omitempty"`
	EventsFilter         *string                 `json:"events_filter,omitempty"`
	FriendlyName         *string                 `json:"friendly_name,omitempty"`
	Links                *map[string]interface{} `json:"links,omitempty"`
	MultiTaskEnabled     *bool                   `json:"multi_task_enabled,omitempty"`
	PrioritizeQueueOrder *WorkspaceQueueOrder    `json:"prioritize_queue_order,omitempty"`
	Sid                  *string                 `json:"sid,omitempty"`
	TimeoutActivityName  *string                 `json:"timeout_activity_name,omitempty"`
	TimeoutActivitySid   *string                 `json:"timeout_activity_sid,omitempty"`
	Url                  *string                 `json:"url,omitempty"`
}
