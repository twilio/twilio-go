/*
 * Twilio - Taskrouter
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

// TaskrouterV1WorkspaceTaskChannel struct for TaskrouterV1WorkspaceTaskChannel
type TaskrouterV1WorkspaceTaskChannel struct {
	AccountSid              string                 `json:"AccountSid,omitempty"`
	ChannelOptimizedRouting bool                   `json:"ChannelOptimizedRouting,omitempty"`
	DateCreated             time.Time              `json:"DateCreated,omitempty"`
	DateUpdated             time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName            string                 `json:"FriendlyName,omitempty"`
	Links                   map[string]interface{} `json:"Links,omitempty"`
	Sid                     string                 `json:"Sid,omitempty"`
	UniqueName              string                 `json:"UniqueName,omitempty"`
	Url                     string                 `json:"Url,omitempty"`
	WorkspaceSid            string                 `json:"WorkspaceSid,omitempty"`
}
