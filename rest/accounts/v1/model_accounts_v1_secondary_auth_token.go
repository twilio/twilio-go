/*
 * Twilio - Accounts
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

// AccountsV1SecondaryAuthToken struct for AccountsV1SecondaryAuthToken
type AccountsV1SecondaryAuthToken struct {
	AccountSid         *string    `json:"account_sid,omitempty"`
	DateCreated        *time.Time `json:"date_created,omitempty"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	SecondaryAuthToken *string    `json:"secondary_auth_token,omitempty"`
	Url                *string    `json:"url,omitempty"`
}
