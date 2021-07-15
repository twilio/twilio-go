/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SyncV1ServiceSyncListSyncListPermission struct for SyncV1ServiceSyncListSyncListPermission
type SyncV1ServiceSyncListSyncListPermission struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The identity of the user to whom the Sync List Permission applies
	Identity *string `json:"identity,omitempty"`
	// The SID of the Sync List to which the Permission applies
	ListSid *string `json:"list_sid,omitempty"`
	// Manage access
	Manage *bool `json:"manage,omitempty"`
	// Read access
	Read *bool `json:"read,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The absolute URL of the Sync List Permission resource
	Url *string `json:"url,omitempty"`
	// Write access
	Write *bool `json:"write,omitempty"`
}
