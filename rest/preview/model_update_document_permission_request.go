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
// UpdateDocumentPermissionRequest struct for UpdateDocumentPermissionRequest
type UpdateDocumentPermissionRequest struct {
	// Boolean flag specifying whether the identity can delete the Sync Document.
	Manage bool `json:"Manage"`
	// Boolean flag specifying whether the identity can read the Sync Document.
	Read bool `json:"Read"`
	// Boolean flag specifying whether the identity can update the Sync Document.
	Write bool `json:"Write"`
}
