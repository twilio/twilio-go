/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateDialingPermissionsSettingsRequest struct for UpdateDialingPermissionsSettingsRequest
type UpdateDialingPermissionsSettingsRequest struct {
	// `true` for the sub-account to inherit voice dialing permissions from the Master Project; otherwise `false`.
	DialingPermissionsInheritance bool `json:"DialingPermissionsInheritance,omitempty"`
}
