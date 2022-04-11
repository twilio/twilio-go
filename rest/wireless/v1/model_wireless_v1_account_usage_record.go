/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// WirelessV1AccountUsageRecord struct for WirelessV1AccountUsageRecord
type WirelessV1AccountUsageRecord struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that describes the aggregated Commands usage for all SIMs during the specified period
	Commands *interface{} `json:"commands,omitempty"`
	// An object that describes the aggregated Data usage for all SIMs over the period
	Data *interface{} `json:"data,omitempty"`
	// The time period for which usage is reported
	Period *interface{} `json:"period,omitempty"`
}
