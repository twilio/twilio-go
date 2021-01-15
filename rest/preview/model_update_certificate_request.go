/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateCertificateRequest struct for UpdateCertificateRequest
type UpdateCertificateRequest struct {
	// Provides the unique string identifier of an existing Device to become authenticated with this Certificate credential.
	DeviceSid string `json:"DeviceSid,omitempty"`
	// Provides a human readable descriptive text for this Certificate credential, up to 256 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
}
