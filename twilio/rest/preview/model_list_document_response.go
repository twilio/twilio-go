/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListDocumentResponse struct for ListDocumentResponse
type ListDocumentResponse struct {
	Documents []PreviewSyncServiceDocument `json:"Documents,omitempty"`
	Meta      ListDayResponseMeta          `json:"Meta,omitempty"`
}
