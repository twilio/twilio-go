/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventsV1SubscriptionSubscribedEvent struct for EventsV1SubscriptionSubscribedEvent
type EventsV1SubscriptionSubscribedEvent struct {
	AccountSid      *string `json:"AccountSid,omitempty"`
	SubscriptionSid *string `json:"SubscriptionSid,omitempty"`
	Type            *string `json:"Type,omitempty"`
	Url             *string `json:"Url,omitempty"`
	Version         *int32  `json:"Version,omitempty"`
}
