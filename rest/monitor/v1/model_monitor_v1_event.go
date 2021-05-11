/*
 * Twilio - Monitor
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

// MonitorV1Event struct for MonitorV1Event
type MonitorV1Event struct {
	AccountSid      *string                 `json:"account_sid,omitempty"`
	ActorSid        *string                 `json:"actor_sid,omitempty"`
	ActorType       *string                 `json:"actor_type,omitempty"`
	Description     *string                 `json:"description,omitempty"`
	EventData       *map[string]interface{} `json:"event_data,omitempty"`
	EventDate       *time.Time              `json:"event_date,omitempty"`
	EventType       *string                 `json:"event_type,omitempty"`
	Links           *map[string]interface{} `json:"links,omitempty"`
	ResourceSid     *string                 `json:"resource_sid,omitempty"`
	ResourceType    *string                 `json:"resource_type,omitempty"`
	Sid             *string                 `json:"sid,omitempty"`
	Source          *string                 `json:"source,omitempty"`
	SourceIpAddress *string                 `json:"source_ip_address,omitempty"`
	Url             *string                 `json:"url,omitempty"`
}
