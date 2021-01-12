/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject13 struct for InlineObject13
type InlineObject13 struct {
	// A valid JSON string that contains application-specific data.
	Attributes string `json:"Attributes,omitempty"`
	// A descriptive string that you create to describe the resource. It is often used for display purposes.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID of the [Role](https://www.twilio.com/docs/api/chat/rest/roles) assigned to this user.
	RoleSid string `json:"RoleSid,omitempty"`
}
