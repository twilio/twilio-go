/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// SyncV1ServiceDocumentDocumentPermission struct for SyncV1ServiceDocumentDocumentPermission
type SyncV1ServiceDocumentDocumentPermission struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The Sync Document SID
	DocumentSid *string `json:"DocumentSid,omitempty"`
	// The identity of the user to whom the Sync Document Permission applies
	Identity *string `json:"Identity,omitempty"`
	// Manage access
	Manage *bool `json:"Manage,omitempty"`
	// Read access
	Read *bool `json:"Read,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The absolute URL of the Sync Document Permission resource
	Url *string `json:"Url,omitempty"`
	// Write access
	Write *bool `json:"Write,omitempty"`
}
