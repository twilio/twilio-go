/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SupersimV1Network struct for SupersimV1Network
type SupersimV1Network struct {
	// A human readable identifier of this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The MCC/MNCs included in the Network resource
	Identifiers *[]map[string]interface{} `json:"identifiers,omitempty"`
	// The ISO country code of the Network resource
	IsoCountry *string `json:"iso_country,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Network resource
	Url *string `json:"url,omitempty"`
}
