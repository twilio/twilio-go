/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateSubscriptionRequest struct for CreateSubscriptionRequest
type CreateSubscriptionRequest struct {
	// A human readable description for the Subscription **This value should not contain PII.**
	Description string `json:"Description"`
	// The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created.
	SinkSid string `json:"SinkSid"`
	// Contains a dictionary of URL links to nested resources of this Subscription.
	Types []map[string]interface{} `json:"Types"`
}
