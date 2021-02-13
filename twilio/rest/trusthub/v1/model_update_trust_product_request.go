/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateTrustProductRequest struct for UpdateTrustProductRequest
type UpdateTrustProductRequest struct {
	// The email address that will receive updates when the Customer-Profile resource changes status.
	Email string `json:"Email,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The verification status of the Customer-Profile resource.
	Status string `json:"Status,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback string `json:"StatusCallback,omitempty"`
}
