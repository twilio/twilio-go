/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// SyncV1Service struct for SyncV1Service
type SyncV1Service struct {
	AccountSid                    *string                 `json:"AccountSid,omitempty"`
	AclEnabled                    *bool                   `json:"AclEnabled,omitempty"`
	DateCreated                   *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated                   *time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName                  *string                 `json:"FriendlyName,omitempty"`
	Links                         *map[string]interface{} `json:"Links,omitempty"`
	ReachabilityDebouncingEnabled *bool                   `json:"ReachabilityDebouncingEnabled,omitempty"`
	ReachabilityDebouncingWindow  *int32                  `json:"ReachabilityDebouncingWindow,omitempty"`
	ReachabilityWebhooksEnabled   *bool                   `json:"ReachabilityWebhooksEnabled,omitempty"`
	Sid                           *string                 `json:"Sid,omitempty"`
	UniqueName                    *string                 `json:"UniqueName,omitempty"`
	Url                           *string                 `json:"Url,omitempty"`
	WebhookUrl                    *string                 `json:"WebhookUrl,omitempty"`
	WebhooksFromRestEnabled       *bool                   `json:"WebhooksFromRestEnabled,omitempty"`
}
