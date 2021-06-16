/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// WirelessV1SimUsageRecord struct for WirelessV1SimUsageRecord
type WirelessV1SimUsageRecord struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that describes the SIM's usage of Commands during the specified period
	Commands *map[string]interface{} `json:"commands,omitempty"`
	// An object that describes the SIM's data usage during the specified period
	Data *map[string]interface{} `json:"data,omitempty"`
	// The time period for which the usage is reported
	Period *map[string]interface{} `json:"period,omitempty"`
	// The SID of the Sim resource that this Usage Record is for
	SimSid *string `json:"sim_sid,omitempty"`
}
