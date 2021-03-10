/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BundleStatus the model 'BundleStatus'
type BundleStatus string

// List of bundle_status
const (
	BUNDLESTATUS_DRAFT                  BundleStatus = "draft"
	BUNDLESTATUS_PENDING_REVIEW         BundleStatus = "pending-review"
	BUNDLESTATUS_IN_REVIEW              BundleStatus = "in-review"
	BUNDLESTATUS_TWILIO_REJECTED        BundleStatus = "twilio-rejected"
	BUNDLESTATUS_TWILIO_APPROVED        BundleStatus = "twilio-approved"
	BUNDLESTATUS_PROVISIONALLY_APPROVED BundleStatus = "provisionally-approved"
)
