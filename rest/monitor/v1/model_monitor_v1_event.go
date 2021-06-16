/*
 * Twilio - Monitor
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

// MonitorV1Event struct for MonitorV1Event
type MonitorV1Event struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the actor that caused the event, if available
	ActorSid *string `json:"actor_sid,omitempty"`
	// The type of actor that caused the event
	ActorType *string `json:"actor_type,omitempty"`
	// A description of the event
	Description *string `json:"description,omitempty"`
	// A JSON string that represents an object with additional data about the event
	EventData *map[string]interface{} `json:"event_data,omitempty"`
	// The ISO 8601 date and time in GMT when the event was recorded
	EventDate *time.Time `json:"event_date,omitempty"`
	// The event's type
	EventType *string `json:"event_type,omitempty"`
	// The absolute URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the resource that was affected
	ResourceSid *string `json:"resource_sid,omitempty"`
	// The type of resource that was affected
	ResourceType *string `json:"resource_type,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The originating system or interface that caused the event
	Source *string `json:"source,omitempty"`
	// The IP address of the source
	SourceIpAddress *string `json:"source_ip_address,omitempty"`
	// The absolute URL of the resource that was affected
	Url *string `json:"url,omitempty"`
}
