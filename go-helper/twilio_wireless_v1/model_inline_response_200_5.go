/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Meta InlineResponse200Meta `json:"meta,omitempty"`
	UsageRecords []WirelessV1AccountUsageRecord `json:"usage_records,omitempty"`
}
