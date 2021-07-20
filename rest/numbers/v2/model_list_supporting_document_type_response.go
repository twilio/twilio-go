/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSupportingDocumentTypeResponse struct for ListSupportingDocumentTypeResponse
type ListSupportingDocumentTypeResponse struct {
	Meta                    ListBundleResponseMeta                                `json:"meta,omitempty"`
	SupportingDocumentTypes []NumbersV2RegulatoryComplianceSupportingDocumentType `json:"supporting_document_types,omitempty"`
}
