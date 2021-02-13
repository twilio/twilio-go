/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// PreviewSyncServiceSyncMapSyncMapItem struct for PreviewSyncServiceSyncMapSyncMapItem
type PreviewSyncServiceSyncMapSyncMapItem struct {
	AccountSid  string                 `json:"AccountSid,omitempty"`
	CreatedBy   string                 `json:"CreatedBy,omitempty"`
	Data        map[string]interface{} `json:"Data,omitempty"`
	DateCreated time.Time              `json:"DateCreated,omitempty"`
	DateUpdated time.Time              `json:"DateUpdated,omitempty"`
	Key         string                 `json:"Key,omitempty"`
	MapSid      string                 `json:"MapSid,omitempty"`
	Revision    string                 `json:"Revision,omitempty"`
	ServiceSid  string                 `json:"ServiceSid,omitempty"`
	Url         string                 `json:"Url,omitempty"`
}
