/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TrusthubV1Policies struct for TrusthubV1Policies
type TrusthubV1Policies struct {
	FriendlyName *string                 `json:"friendly_name,omitempty"`
	Requirements *map[string]interface{} `json:"requirements,omitempty"`
	Sid          *string                 `json:"sid,omitempty"`
	Url          *string                 `json:"url,omitempty"`
}
