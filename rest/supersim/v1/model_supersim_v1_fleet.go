/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// SupersimV1Fleet struct for SupersimV1Fleet
type SupersimV1Fleet struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Defines whether SIMs in the Fleet are capable of using data connectivity
	DataEnabled *bool `json:"data_enabled,omitempty"`
	// The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume
	DataLimit *int `json:"data_limit,omitempty"`
	// The model by which a SIM is metered and billed
	DataMetering *string `json:"data_metering,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// A string representing the HTTP method to use when making a request to `ip_commands_url`
	IpCommandsMethod *string `json:"ip_commands_method,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device
	IpCommandsUrl *string `json:"ip_commands_url,omitempty"`
	// The SID of the Network Access Profile of the Fleet
	NetworkAccessProfileSid *string `json:"network_access_profile_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands
	SmsCommandsEnabled *bool `json:"sms_commands_enabled,omitempty"`
	// A string representing the HTTP method to use when making a request to `sms_commands_url`
	SmsCommandsMethod *string `json:"sms_commands_method,omitempty"`
	// The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number
	SmsCommandsUrl *string `json:"sms_commands_url,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the Fleet resource
	Url *string `json:"url,omitempty"`
}
