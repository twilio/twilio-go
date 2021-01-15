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
// WirelessV1AccountUsageRecord struct for WirelessV1AccountUsageRecord
type WirelessV1AccountUsageRecord struct {
	AccountSid string `json:"account_sid,omitempty"`
	Commands map[string]interface{} `json:"commands,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	Period map[string]interface{} `json:"period,omitempty"`
}
