/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// SyncV1ServiceSyncList struct for SyncV1ServiceSyncList
type SyncV1ServiceSyncList struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CreatedBy string `json:"CreatedBy,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateExpires time.Time `json:"DateExpires,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Revision string `json:"Revision,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
