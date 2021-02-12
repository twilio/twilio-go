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
import (
	"time"
)
// SyncV1ServiceSyncMap struct for SyncV1ServiceSyncMap
type SyncV1ServiceSyncMap struct {
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
