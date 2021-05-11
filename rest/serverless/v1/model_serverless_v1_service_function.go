/*
 * Twilio - Serverless
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

// ServerlessV1ServiceFunction struct for ServerlessV1ServiceFunction
type ServerlessV1ServiceFunction struct {
	AccountSid   *string                 `json:"account_sid,omitempty"`
	DateCreated  *time.Time              `json:"date_created,omitempty"`
	DateUpdated  *time.Time              `json:"date_updated,omitempty"`
	FriendlyName *string                 `json:"friendly_name,omitempty"`
	Links        *map[string]interface{} `json:"links,omitempty"`
	ServiceSid   *string                 `json:"service_sid,omitempty"`
	Sid          *string                 `json:"sid,omitempty"`
	Url          *string                 `json:"url,omitempty"`
}
