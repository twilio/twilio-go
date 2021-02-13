/*
 * Twilio - Conversations
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

// UpdateConversationRequest struct for UpdateConversationRequest
type UpdateConversationRequest struct {
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	// The human-readable name of this conversation, limited to 256 characters. Optional.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The unique ID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) this conversation belongs to.
	MessagingServiceSid string `json:"MessagingServiceSid,omitempty"`
	// Current state of this conversation. Can be either `active`, `inactive` or `closed` and defaults to `active`
	State string `json:"State,omitempty"`
	// ISO8601 duration when conversation will be switched to `closed` state. Minimum value for this timer is 10 minutes.
	TimersClosed string `json:"TimersClosed,omitempty"`
	// ISO8601 duration when conversation will be switched to `inactive` state. Minimum value for this timer is 1 minute.
	TimersInactive string `json:"TimersInactive,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL.
	UniqueName string `json:"UniqueName,omitempty"`
}
