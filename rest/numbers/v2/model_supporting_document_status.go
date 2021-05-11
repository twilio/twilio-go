/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SupportingDocumentStatus the model 'SupportingDocumentStatus'
type SupportingDocumentStatus string

// List of supporting_document_status
const (
	SUPPORTINGDOCUMENTSTATUS_DRAFT                  SupportingDocumentStatus = "draft"
	SUPPORTINGDOCUMENTSTATUS_PENDING_REVIEW         SupportingDocumentStatus = "pending-review"
	SUPPORTINGDOCUMENTSTATUS_REJECTED               SupportingDocumentStatus = "rejected"
	SUPPORTINGDOCUMENTSTATUS_APPROVED               SupportingDocumentStatus = "approved"
	SUPPORTINGDOCUMENTSTATUS_EXPIRED                SupportingDocumentStatus = "expired"
	SUPPORTINGDOCUMENTSTATUS_PROVISIONALLY_APPROVED SupportingDocumentStatus = "provisionally-approved"
)
