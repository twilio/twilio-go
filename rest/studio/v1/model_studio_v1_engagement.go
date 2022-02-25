/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// StudioV1Engagement struct for StudioV1Engagement
type StudioV1Engagement struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The phone number, SIP address or Client identifier that triggered this Engagement
	ContactChannelAddress *string `json:"contact_channel_address,omitempty"`
	// The SID of the Contact
	ContactSid *string `json:"contact_sid,omitempty"`
	// The current state of the execution flow
	Context *map[string]interface{} `json:"context,omitempty"`
	// The ISO 8601 date and time in GMT when the Engagement was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Engagement was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The SID of the Flow
	FlowSid *string `json:"flow_sid,omitempty"`
	// The URLs of the Engagement's nested resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the Engagement
	Status *string `json:"status,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
