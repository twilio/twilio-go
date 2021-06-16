/*
 * Twilio - Sync
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

// SyncV1ServiceDocument struct for SyncV1ServiceDocument
type SyncV1ServiceDocument struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The identity of the Sync Document's creator
	CreatedBy *string `json:"created_by,omitempty"`
	// An arbitrary, schema-less object that the Sync Document stores
	Data *map[string]interface{} `json:"data,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Sync Document expires
	DateExpires *time.Time `json:"date_expires,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URLs of resources related to the Sync Document
	Links *map[string]interface{} `json:"links,omitempty"`
	// The current revision of the Sync Document, represented by a string identifier
	Revision *string `json:"revision,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Document resource
	Url *string `json:"url,omitempty"`
}
