/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSupportingDocumentResponse struct for ListSupportingDocumentResponse
type ListSupportingDocumentResponse struct {
	Meta    ListBundleResponseMeta                            `json:"meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceSupportingDocument `json:"results,omitempty"`
}
