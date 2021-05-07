/*
 * Twilio - Sync
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

// SyncV1ServiceSyncStream struct for SyncV1ServiceSyncStream
type SyncV1ServiceSyncStream struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The Identity of the Stream's creator
	CreatedBy *string `json:"created_by,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Message Stream expires
	DateExpires *time.Time `json:"date_expires,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URLs of the Stream's nested resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Message Stream resource
	Url *string `json:"url,omitempty"`
}
