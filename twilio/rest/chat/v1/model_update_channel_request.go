/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateChannelRequest struct for UpdateChannelRequest
type UpdateChannelRequest struct {
	// A valid JSON string that contains application-specific data.
	Attributes string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource's `sid` in the URL. This value must be 64 characters or less in length and be unique within the Service.
	UniqueName string `json:"UniqueName,omitempty"`
}
