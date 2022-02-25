/*
 * Twilio - Autopilot
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

// AutopilotV1FieldType struct for AutopilotV1FieldType
type AutopilotV1FieldType struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A list of the URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the FieldType resource
	Url *string `json:"url,omitempty"`
}
