/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1Subscription struct for EventsV1Subscription
type EventsV1Subscription struct {
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date this Subscription was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Subscription was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Subscription description
	Description *string `json:"description,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// A string that uniquely identifies this Subscription.
	Sid *string `json:"sid,omitempty"`
	// Sink SID.
	SinkSid *string `json:"sink_sid,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
