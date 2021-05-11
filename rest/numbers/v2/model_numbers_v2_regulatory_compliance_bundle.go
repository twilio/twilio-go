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

import (
	"time"
)

// NumbersV2RegulatoryComplianceBundle struct for NumbersV2RegulatoryComplianceBundle
type NumbersV2RegulatoryComplianceBundle struct {
	AccountSid     *string                 `json:"account_sid,omitempty"`
	DateCreated    *time.Time              `json:"date_created,omitempty"`
	DateUpdated    *time.Time              `json:"date_updated,omitempty"`
	Email          *string                 `json:"email,omitempty"`
	FriendlyName   *string                 `json:"friendly_name,omitempty"`
	Links          *map[string]interface{} `json:"links,omitempty"`
	RegulationSid  *string                 `json:"regulation_sid,omitempty"`
	Sid            *string                 `json:"sid,omitempty"`
	Status         *BundleStatus           `json:"status,omitempty"`
	StatusCallback *string                 `json:"status_callback,omitempty"`
	Url            *string                 `json:"url,omitempty"`
	ValidUntil     *time.Time              `json:"valid_until,omitempty"`
}
