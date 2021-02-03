/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Status the model 'Status'
type Status string

// List of status
const (
	DRAFT Status = "draft"
	PENDING_REVIEW Status = "pending-review"
	REJECTED Status = "rejected"
	APPROVED Status = "approved"
	EXPIRED Status = "expired"
	PROVISIONALLY_APPROVED Status = "provisionally-approved"
)
