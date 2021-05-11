/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// EventsV1SubscriptionSubscribedEvent struct for EventsV1SubscriptionSubscribedEvent
type EventsV1SubscriptionSubscribedEvent struct {
	AccountSid      *string `json:"account_sid,omitempty"`
	SchemaVersion   *int32  `json:"schema_version,omitempty"`
	SubscriptionSid *string `json:"subscription_sid,omitempty"`
	Type            *string `json:"type,omitempty"`
	Url             *string `json:"url,omitempty"`
}
