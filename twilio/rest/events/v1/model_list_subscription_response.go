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

// ListSubscriptionResponse struct for ListSubscriptionResponse
type ListSubscriptionResponse struct {
	Meta          ListVersionResponseMeta `json:"Meta,omitempty"`
	Subscriptions []EventsV1Subscription  `json:"Subscriptions,omitempty"`
}
