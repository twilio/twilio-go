/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SupersimV1Network struct for SupersimV1Network
type SupersimV1Network struct {
	FriendlyName string                   `json:"FriendlyName,omitempty"`
	Identifiers  []map[string]interface{} `json:"Identifiers,omitempty"`
	IsoCountry   string                   `json:"IsoCountry,omitempty"`
	Sid          string                   `json:"Sid,omitempty"`
	Url          string                   `json:"Url,omitempty"`
}
