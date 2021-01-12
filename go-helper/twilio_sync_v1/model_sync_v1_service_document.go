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
// SyncV1ServiceDocument struct for SyncV1ServiceDocument
type SyncV1ServiceDocument struct {
	AccountSid string `json:"account_sid,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateExpires time.Time `json:"date_expires,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	Revision string `json:"revision,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	UniqueName string `json:"unique_name,omitempty"`
	Url string `json:"url,omitempty"`
}
