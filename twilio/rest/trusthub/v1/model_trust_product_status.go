/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TrustProductStatus the model 'TrustProductStatus'
type TrustProductStatus string

// List of trust_product_status
const (
	TRUSTPRODUCTSTATUS_DRAFT           TrustProductStatus = "draft"
	TRUSTPRODUCTSTATUS_PENDING_REVIEW  TrustProductStatus = "pending-review"
	TRUSTPRODUCTSTATUS_IN_REVIEW       TrustProductStatus = "in-review"
	TRUSTPRODUCTSTATUS_TWILIO_REJECTED TrustProductStatus = "twilio-rejected"
	TRUSTPRODUCTSTATUS_TWILIO_APPROVED TrustProductStatus = "twilio-approved"
)