/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// CustomerProfileStatus the model 'CustomerProfileStatus'
type CustomerProfileStatus string

// List of customer_profile_status
const (
	CUSTOMERPROFILESTATUS_DRAFT           CustomerProfileStatus = "draft"
	CUSTOMERPROFILESTATUS_PENDING_REVIEW  CustomerProfileStatus = "pending-review"
	CUSTOMERPROFILESTATUS_IN_REVIEW       CustomerProfileStatus = "in-review"
	CUSTOMERPROFILESTATUS_TWILIO_REJECTED CustomerProfileStatus = "twilio-rejected"
	CUSTOMERPROFILESTATUS_TWILIO_APPROVED CustomerProfileStatus = "twilio-approved"
)
