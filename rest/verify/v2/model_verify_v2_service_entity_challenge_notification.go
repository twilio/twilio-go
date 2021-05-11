/*
 * Twilio - Verify
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

// VerifyV2ServiceEntityChallengeNotification struct for VerifyV2ServiceEntityChallengeNotification
type VerifyV2ServiceEntityChallengeNotification struct {
	AccountSid   *string    `json:"account_sid,omitempty"`
	ChallengeSid *string    `json:"challenge_sid,omitempty"`
	DateCreated  *time.Time `json:"date_created,omitempty"`
	EntitySid    *string    `json:"entity_sid,omitempty"`
	Identity     *string    `json:"identity,omitempty"`
	Priority     *string    `json:"priority,omitempty"`
	ServiceSid   *string    `json:"service_sid,omitempty"`
	Sid          *string    `json:"sid,omitempty"`
	Ttl          *int32     `json:"ttl,omitempty"`
}
