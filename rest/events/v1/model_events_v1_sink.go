/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1Sink struct for EventsV1Sink
type EventsV1Sink struct {
	// The date this Sink was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Sink was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Sink Description
	Description *string `json:"description,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// A string that uniquely identifies this Sink.
	Sid *string `json:"sid,omitempty"`
	// JSON Sink configuration.
	SinkConfiguration *map[string]interface{} `json:"sink_configuration,omitempty"`
	// Sink type.
	SinkType *string `json:"sink_type,omitempty"`
	// The Status of this Sink
	Status *string `json:"status,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
