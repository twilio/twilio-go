/*
 * Twilio - Api
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
// ApiV2010AccountUsageUsageRecordUsageRecordLastMonth struct for ApiV2010AccountUsageUsageRecordUsageRecordLastMonth
type ApiV2010AccountUsageUsageRecordUsageRecordLastMonth struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	AsOf string `json:"AsOf,omitempty"`
	Category string `json:"Category,omitempty"`
	Count string `json:"Count,omitempty"`
	CountUnit string `json:"CountUnit,omitempty"`
	Description string `json:"Description,omitempty"`
	EndDate time.Time `json:"EndDate,omitempty"`
	Price float32 `json:"Price,omitempty"`
	PriceUnit string `json:"PriceUnit,omitempty"`
	StartDate time.Time `json:"StartDate,omitempty"`
	SubresourceUris map[string]interface{} `json:"SubresourceUris,omitempty"`
	Uri string `json:"Uri,omitempty"`
	Usage string `json:"Usage,omitempty"`
	UsageUnit string `json:"UsageUnit,omitempty"`
}
