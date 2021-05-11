/*
 * Twilio - Studio
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

// StudioV1Flow struct for StudioV1Flow
type StudioV1Flow struct {
	AccountSid   *string                 `json:"account_sid,omitempty"`
	DateCreated  *time.Time              `json:"date_created,omitempty"`
	DateUpdated  *time.Time              `json:"date_updated,omitempty"`
	FriendlyName *string                 `json:"friendly_name,omitempty"`
	Links        *map[string]interface{} `json:"links,omitempty"`
	Sid          *string                 `json:"sid,omitempty"`
	Status       *FlowStatus             `json:"status,omitempty"`
	Url          *string                 `json:"url,omitempty"`
	Version      *int32                  `json:"version,omitempty"`
}
