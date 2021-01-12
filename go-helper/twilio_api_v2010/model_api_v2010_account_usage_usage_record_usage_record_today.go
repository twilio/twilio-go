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
// ApiV2010AccountUsageUsageRecordUsageRecordToday struct for ApiV2010AccountUsageUsageRecordUsageRecordToday
type ApiV2010AccountUsageUsageRecordUsageRecordToday struct {
	AccountSid string `json:"account_sid,omitempty"`
	ApiVersion string `json:"api_version,omitempty"`
	AsOf string `json:"as_of,omitempty"`
	Category string `json:"category,omitempty"`
	Count string `json:"count,omitempty"`
	CountUnit string `json:"count_unit,omitempty"`
	Description string `json:"description,omitempty"`
	EndDate time.Time `json:"end_date,omitempty"`
	Price float32 `json:"price,omitempty"`
	PriceUnit string `json:"price_unit,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
	Uri string `json:"uri,omitempty"`
	Usage string `json:"usage,omitempty"`
	UsageUnit string `json:"usage_unit,omitempty"`
}
