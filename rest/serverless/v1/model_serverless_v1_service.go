/*
 * Twilio - Serverless
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

// ServerlessV1Service struct for ServerlessV1Service
type ServerlessV1Service struct {
	AccountSid         *string                 `json:"account_sid,omitempty"`
	DateCreated        *time.Time              `json:"date_created,omitempty"`
	DateUpdated        *time.Time              `json:"date_updated,omitempty"`
	FriendlyName       *string                 `json:"friendly_name,omitempty"`
	IncludeCredentials *bool                   `json:"include_credentials,omitempty"`
	Links              *map[string]interface{} `json:"links,omitempty"`
	Sid                *string                 `json:"sid,omitempty"`
	UiEditable         *bool                   `json:"ui_editable,omitempty"`
	UniqueName         *string                 `json:"unique_name,omitempty"`
	Url                *string                 `json:"url,omitempty"`
}
