/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// NumbersV2RegulatoryComplianceBundleItemAssignment struct for NumbersV2RegulatoryComplianceBundleItemAssignment
type NumbersV2RegulatoryComplianceBundleItemAssignment struct {
	AccountSid  *string    `json:"account_sid,omitempty"`
	BundleSid   *string    `json:"bundle_sid,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	ObjectSid   *string    `json:"object_sid,omitempty"`
	Sid         *string    `json:"sid,omitempty"`
	Url         *string    `json:"url,omitempty"`
}
