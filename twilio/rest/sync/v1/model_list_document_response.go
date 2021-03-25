/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListDocumentResponse struct for ListDocumentResponse
type ListDocumentResponse struct {
	Documents []SyncV1ServiceDocument `json:"Documents,omitempty"`
	Meta      ListServiceResponseMeta `json:"Meta,omitempty"`
}
