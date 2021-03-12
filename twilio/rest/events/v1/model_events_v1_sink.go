/*
    * Twilio - Events
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
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
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date this Sink was updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Sink Description
	Description *string `json:"Description,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"Links,omitempty"`
	// A string that uniquely identifies this Sink.
	Sid *string `json:"Sid,omitempty"`
	// JSON Sink configuration.
	SinkConfiguration *map[string]interface{} `json:"SinkConfiguration,omitempty"`
	// Sink type.
	SinkType *string `json:"SinkType,omitempty"`
	// The Status of this Sink
	Status *string `json:"Status,omitempty"`
	// The URL of this resource.
	Url *string `json:"Url,omitempty"`
}
