/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// EventsV1SubscriptionReadResponse struct for EventsV1SubscriptionReadResponse
type EventsV1SubscriptionReadResponse struct {
	Meta EventsV1SchemaVersionReadResponseMeta `json:"Meta,omitempty"`
	Subscriptions []EventsV1Subscription `json:"Subscriptions,omitempty"`
}
