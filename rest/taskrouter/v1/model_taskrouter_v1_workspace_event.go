/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceEvent struct for TaskrouterV1WorkspaceEvent
type TaskrouterV1WorkspaceEvent struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the resource that triggered the event
	ActorSid *string `json:"actor_sid,omitempty"`
	// The type of resource that triggered the event
	ActorType *string `json:"actor_type,omitempty"`
	// The absolute URL of the resource that triggered the event
	ActorUrl *string `json:"actor_url,omitempty"`
	// A description of the event
	Description *string `json:"description,omitempty"`
	// Data about the event
	EventData *map[string]interface{} `json:"event_data,omitempty"`
	// The time the event was sent
	EventDate *time.Time `json:"event_date,omitempty"`
	// The time the event was sent in milliseconds
	EventDateMs *int32 `json:"event_date_ms,omitempty"`
	// The identifier for the event
	EventType *string `json:"event_type,omitempty"`
	// The SID of the object the event is most relevant to
	ResourceSid *string `json:"resource_sid,omitempty"`
	// The type of object the event is most relevant to
	ResourceType *string `json:"resource_type,omitempty"`
	// The URL of the resource the event is most relevant to
	ResourceUrl *string `json:"resource_url,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Where the Event originated
	Source *string `json:"source,omitempty"`
	// The IP from which the Event originated
	SourceIpAddress *string `json:"source_ip_address,omitempty"`
	// The absolute URL of the Event resource
	Url *string `json:"url,omitempty"`
	// The SID of the Workspace that contains the Event
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
}
