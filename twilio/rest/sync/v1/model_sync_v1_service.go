/*
    * Twilio - Sync
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// SyncV1Service struct for SyncV1Service
type SyncV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Whether token identities in the Service must be granted access to Sync objects by using the Permissions resource
	AclEnabled *bool `json:"AclEnabled,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// Whether every endpoint_disconnected event occurs after a configurable delay
	ReachabilityDebouncingEnabled *bool `json:"ReachabilityDebouncingEnabled,omitempty"`
	// The reachability event delay in milliseconds
	ReachabilityDebouncingWindow *int32 `json:"ReachabilityDebouncingWindow,omitempty"`
	// Whether the service instance calls webhook_url when client endpoints connect to Sync
	ReachabilityWebhooksEnabled *bool `json:"ReachabilityWebhooksEnabled,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"Url,omitempty"`
	// The URL we call when Sync objects are manipulated
	WebhookUrl *string `json:"WebhookUrl,omitempty"`
	// Whether the Service instance should call webhook_url when the REST API is used to update Sync objects
	WebhooksFromRestEnabled *bool `json:"WebhooksFromRestEnabled,omitempty"`
}
