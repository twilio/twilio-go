/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// PreviewWirelessCommand struct for PreviewWirelessCommand
type PreviewWirelessCommand struct {
	AccountSid  string    `json:"AccountSid,omitempty"`
	Command     string    `json:"Command,omitempty"`
	CommandMode string    `json:"CommandMode,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DeviceSid   string    `json:"DeviceSid,omitempty"`
	Direction   string    `json:"Direction,omitempty"`
	Sid         string    `json:"Sid,omitempty"`
	SimSid      string    `json:"SimSid,omitempty"`
	Status      string    `json:"Status,omitempty"`
	Url         string    `json:"Url,omitempty"`
}
