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

// PreviewSyncServiceSyncList struct for PreviewSyncServiceSyncList
type PreviewSyncServiceSyncList struct {
	AccountSid  string                 `json:"AccountSid,omitempty"`
	CreatedBy   string                 `json:"CreatedBy,omitempty"`
	DateCreated time.Time              `json:"DateCreated,omitempty"`
	DateUpdated time.Time              `json:"DateUpdated,omitempty"`
	Links       map[string]interface{} `json:"Links,omitempty"`
	Revision    string                 `json:"Revision,omitempty"`
	ServiceSid  string                 `json:"ServiceSid,omitempty"`
	Sid         string                 `json:"Sid,omitempty"`
	UniqueName  string                 `json:"UniqueName,omitempty"`
	Url         string                 `json:"Url,omitempty"`
}
