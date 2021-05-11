/*
 * Twilio - Trusthub
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

// TrusthubV1TrustProduct struct for TrusthubV1TrustProduct
type TrusthubV1TrustProduct struct {
	AccountSid     *string                 `json:"account_sid,omitempty"`
	DateCreated    *time.Time              `json:"date_created,omitempty"`
	DateUpdated    *time.Time              `json:"date_updated,omitempty"`
	Email          *string                 `json:"email,omitempty"`
	FriendlyName   *string                 `json:"friendly_name,omitempty"`
	Links          *map[string]interface{} `json:"links,omitempty"`
	PolicySid      *string                 `json:"policy_sid,omitempty"`
	Sid            *string                 `json:"sid,omitempty"`
	Status         *TrustProductStatus     `json:"status,omitempty"`
	StatusCallback *string                 `json:"status_callback,omitempty"`
	Url            *string                 `json:"url,omitempty"`
	ValidUntil     *time.Time              `json:"valid_until,omitempty"`
}
