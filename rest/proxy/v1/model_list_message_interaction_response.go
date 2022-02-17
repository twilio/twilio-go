/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListMessageInteractionResponse struct for ListMessageInteractionResponse
type ListMessageInteractionResponse struct {
	Interactions []ProxyV1MessageInteraction `json:"interactions,omitempty"`
	Meta         ListServiceResponseMeta     `json:"meta,omitempty"`
}
