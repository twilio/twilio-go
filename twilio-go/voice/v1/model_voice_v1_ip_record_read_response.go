/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VoiceV1IpRecordReadResponse struct for VoiceV1IpRecordReadResponse
type VoiceV1IpRecordReadResponse struct {
	IpRecords []VoiceV1IpRecord `json:"ip_records,omitempty"`
	Meta VoiceV1ByocTrunkReadResponseMeta `json:"meta,omitempty"`
}
