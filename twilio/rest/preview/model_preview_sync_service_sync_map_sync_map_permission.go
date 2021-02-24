/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PreviewSyncServiceSyncMapSyncMapPermission struct for PreviewSyncServiceSyncMapSyncMapPermission
type PreviewSyncServiceSyncMapSyncMapPermission struct {
	AccountSid *string `json:"AccountSid,omitempty"`
	Identity   *string `json:"Identity,omitempty"`
	Manage     *bool   `json:"Manage,omitempty"`
	MapSid     *string `json:"MapSid,omitempty"`
	Read       *bool   `json:"Read,omitempty"`
	ServiceSid *string `json:"ServiceSid,omitempty"`
	Url        *string `json:"Url,omitempty"`
	Write      *bool   `json:"Write,omitempty"`
}
