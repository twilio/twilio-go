/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FaxV1Fax struct for FaxV1Fax
type FaxV1Fax struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to transmit the fax
	ApiVersion *string `json:"api_version,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The direction of the fax
	Direction *string `json:"direction,omitempty"`
	// The time it took to transmit the fax
	Duration *int `json:"duration,omitempty"`
	// The number the fax was sent from
	From *string `json:"from,omitempty"`
	// The URLs of the fax's related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the FaxMedia resource that is associated with the Fax
	MediaSid *string `json:"media_sid,omitempty"`
	// The Twilio-hosted URL that can be used to download fax media
	MediaUrl *string `json:"media_url,omitempty"`
	// The number of pages contained in the fax document
	NumPages *int `json:"num_pages,omitempty"`
	// The fax transmission price
	Price *float32 `json:"price,omitempty"`
	// The ISO 4217 currency used for billing
	PriceUnit *string `json:"price_unit,omitempty"`
	// The quality of the fax
	Quality *string `json:"quality,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the fax
	Status *string `json:"status,omitempty"`
	// The phone number that received the fax
	To *string `json:"to,omitempty"`
	// The absolute URL of the fax resource
	Url *string `json:"url,omitempty"`
}
