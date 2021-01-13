/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SupersimV1UsageRecordReadResponse struct for SupersimV1UsageRecordReadResponse
type SupersimV1UsageRecordReadResponse struct {
	Meta SupersimV1CommandReadResponseMeta `json:"meta,omitempty"`
	UsageRecords []SupersimV1UsageRecord `json:"usage_records,omitempty"`
}
