/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListIpRecordResponse struct for ListIpRecordResponse
type ListIpRecordResponse struct {
	IpRecords []VoiceV1IpRecord         `json:"ip_records,omitempty"`
	Meta      ListByocTrunkResponseMeta `json:"meta,omitempty"`
}
