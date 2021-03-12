/*
    * Twilio - Sync
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// SyncV1ServiceSyncListSyncListItem struct for SyncV1ServiceSyncListSyncListItem
type SyncV1ServiceSyncListSyncListItem struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The identity of the List Item's creator
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// An arbitrary, schema-less object that the List Item stores
	Data *map[string]interface{} `json:"Data,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the List Item expires
	DateExpires *time.Time `json:"DateExpires,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The automatically generated index of the List Item
	Index *int32 `json:"Index,omitempty"`
	// The SID of the Sync List that contains the List Item
	ListSid *string `json:"ListSid,omitempty"`
	// The current revision of the item, represented as a string
	Revision *string `json:"Revision,omitempty"`
	// The SID of the Sync Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The absolute URL of the List Item resource
	Url *string `json:"Url,omitempty"`
}
