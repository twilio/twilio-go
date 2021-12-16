/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// TrusthubV1EndUserType struct for TrusthubV1EndUserType
type TrusthubV1EndUserType struct {
	// The required information for creating an End-User.
	Fields *[]map[string]interface{} `json:"fields,omitempty"`
	// A human-readable description of the End-User Type resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// A machine-readable description of the End-User Type resource
	MachineName *string `json:"machine_name,omitempty"`
	// The unique string that identifies the End-User Type resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the End-User Type resource
	Url *string `json:"url,omitempty"`
}
