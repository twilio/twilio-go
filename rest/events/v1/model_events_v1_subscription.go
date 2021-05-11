/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// EventsV1Subscription struct for EventsV1Subscription
type EventsV1Subscription struct {
	AccountSid  *string                 `json:"account_sid,omitempty"`
	DateCreated *time.Time              `json:"date_created,omitempty"`
	DateUpdated *time.Time              `json:"date_updated,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Links       *map[string]interface{} `json:"links,omitempty"`
	Sid         *string                 `json:"sid,omitempty"`
	SinkSid     *string                 `json:"sink_sid,omitempty"`
	Url         *string                 `json:"url,omitempty"`
}
