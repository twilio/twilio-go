/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListUsageRecordResponse struct for ListUsageRecordResponse
type ListUsageRecordResponse struct {
	Meta         ListCommandResponseMeta `json:"Meta,omitempty"`
	UsageRecords []SupersimV1UsageRecord `json:"UsageRecords,omitempty"`
}
