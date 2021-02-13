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

// UpdateSupportingDocumentRequest struct for UpdateSupportingDocumentRequest
type UpdateSupportingDocumentRequest struct {
	// The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
	Attributes map[string]interface{} `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
}
