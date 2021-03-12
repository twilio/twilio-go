/*
    * Twilio - Api
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
// ApiV2010AccountUsageUsageRecord struct for ApiV2010AccountUsageUsageRecord
type ApiV2010AccountUsageUsageRecord struct {
	// The SID of the Account accrued the usage
	AccountSid *string `json:"AccountSid,omitempty"`
	// The API version used to create the resource
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// Usage records up to date as of this timestamp
	AsOf *string `json:"AsOf,omitempty"`
	// The category of usage
	Category *string `json:"Category,omitempty"`
	// The number of usage events
	Count *string `json:"Count,omitempty"`
	// The units in which count is measured
	CountUnit *string `json:"CountUnit,omitempty"`
	// A plain-language description of the usage category
	Description *string `json:"Description,omitempty"`
	// The last date for which usage is included in the UsageRecord
	EndDate *time.Time `json:"EndDate,omitempty"`
	// The total price of the usage
	Price *float32 `json:"Price,omitempty"`
	// The currency in which `price` is measured
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The first date for which usage is included in this UsageRecord
	StartDate *time.Time `json:"StartDate,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"SubresourceUris,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
	// The amount of usage
	Usage *string `json:"Usage,omitempty"`
	// The units in which usage is measured
	UsageUnit *string `json:"UsageUnit,omitempty"`
}
