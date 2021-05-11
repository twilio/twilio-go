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

// VerifyV2ServiceWebhook struct for VerifyV2ServiceWebhook
type VerifyV2ServiceWebhook struct {
	AccountSid    *string         `json:"account_sid,omitempty"`
	DateCreated   *time.Time      `json:"date_created,omitempty"`
	DateUpdated   *time.Time      `json:"date_updated,omitempty"`
	EventTypes    *[]string       `json:"event_types,omitempty"`
	FriendlyName  *string         `json:"friendly_name,omitempty"`
	ServiceSid    *string         `json:"service_sid,omitempty"`
	Sid           *string         `json:"sid,omitempty"`
	Status        *WebhookStatus  `json:"status,omitempty"`
	Url           *string         `json:"url,omitempty"`
	WebhookMethod *WebhookMethods `json:"webhook_method,omitempty"`
	WebhookUrl    *string         `json:"webhook_url,omitempty"`
}
