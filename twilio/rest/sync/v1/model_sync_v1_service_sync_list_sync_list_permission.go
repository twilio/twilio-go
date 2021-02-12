/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SyncV1ServiceSyncListSyncListPermission struct for SyncV1ServiceSyncListSyncListPermission
type SyncV1ServiceSyncListSyncListPermission struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Identity string `json:"Identity,omitempty"`
	ListSid string `json:"ListSid,omitempty"`
	Manage bool `json:"Manage,omitempty"`
	Read bool `json:"Read,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Url string `json:"Url,omitempty"`
	Write bool `json:"Write,omitempty"`
}
