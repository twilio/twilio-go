/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// MonitorV1Event struct for MonitorV1Event
type MonitorV1Event struct {
	AccountSid      string                 `json:"AccountSid,omitempty"`
	ActorSid        string                 `json:"ActorSid,omitempty"`
	ActorType       string                 `json:"ActorType,omitempty"`
	Description     string                 `json:"Description,omitempty"`
	EventData       map[string]interface{} `json:"EventData,omitempty"`
	EventDate       time.Time              `json:"EventDate,omitempty"`
	EventType       string                 `json:"EventType,omitempty"`
	Links           map[string]interface{} `json:"Links,omitempty"`
	ResourceSid     string                 `json:"ResourceSid,omitempty"`
	ResourceType    string                 `json:"ResourceType,omitempty"`
	Sid             string                 `json:"Sid,omitempty"`
	Source          string                 `json:"Source,omitempty"`
	SourceIpAddress string                 `json:"SourceIpAddress,omitempty"`
	Url             string                 `json:"Url,omitempty"`
}
