/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SupersimV1NetworkAccessProfileNetwork struct for SupersimV1NetworkAccessProfileNetwork
type SupersimV1NetworkAccessProfileNetwork struct {
	// A human readable identifier of this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The MCC/MNCs included in the resource
	Identifiers *[]interface{} `json:"identifiers,omitempty"`
	// The ISO country code of the Network resource
	IsoCountry *string `json:"iso_country,omitempty"`
	// The unique string that identifies the Network Access Profile resource
	NetworkAccessProfileSid *string `json:"network_access_profile_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
