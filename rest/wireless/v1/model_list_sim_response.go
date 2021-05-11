/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSimResponse struct for ListSimResponse
type ListSimResponse struct {
	Meta ListCommandResponseMeta `json:"meta,omitempty"`
	Sims []WirelessV1Sim         `json:"sims,omitempty"`
}
