/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1SchemaVersion struct for EventsV1SchemaVersion
type EventsV1SchemaVersion struct {
	// The date the schema version was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The unique identifier of the schema.
	Id  *string `json:"id,omitempty"`
	Raw *string `json:"raw,omitempty"`
	// The version of this schema.
	SchemaVersion *int32 `json:"schema_version,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
