/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)

// SyncV1Service struct for SyncV1Service
type SyncV1Service struct {
	// The unique string that we created to identify the Service resource.
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. It is a read-only property, it cannot be assigned using REST API.
	UniqueName *string `json:"unique_name,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Service resource.
	Url *string `json:"url,omitempty"`
	// The URL we call when Sync objects are manipulated.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// Whether the Service instance should call `webhook_url` when the REST API is used to update Sync objects. The default is `false`.
	WebhooksFromRestEnabled *bool `json:"webhooks_from_rest_enabled,omitempty"`
	// Whether the service instance calls `webhook_url` when client endpoints connect to Sync. The default is `false`.
	ReachabilityWebhooksEnabled *bool `json:"reachability_webhooks_enabled,omitempty"`
	// Whether token identities in the Service must be granted access to Sync objects by using the [Permissions](https://www.twilio.com/docs/sync/api/sync-permissions) resource. It is disabled (false) by default.
	AclEnabled *bool `json:"acl_enabled,omitempty"`
	// Whether every `endpoint_disconnected` event should occur after a configurable delay. The default is `false`, where the `endpoint_disconnected` event occurs immediately after disconnection. When `true`, intervening reconnections can prevent the `endpoint_disconnected` event.
	ReachabilityDebouncingEnabled *bool `json:"reachability_debouncing_enabled,omitempty"`
	// The reachability event delay in milliseconds if `reachability_debouncing_enabled` = `true`.  Must be between 1,000 and 30,000 and defaults to 5,000. This is the number of milliseconds after the last running client disconnects, and a Sync identity is declared offline, before `webhook_url` is called, if all endpoints remain offline. A reconnection from the same identity by any endpoint during this interval prevents the reachability event from occurring.
	ReachabilityDebouncingWindow int `json:"reachability_debouncing_window,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
