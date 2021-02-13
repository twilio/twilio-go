/*
 * Twilio - Wireless
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

// WirelessV1Command struct for WirelessV1Command
type WirelessV1Command struct {
	AccountSid               string      `json:"AccountSid,omitempty"`
	Command                  string      `json:"Command,omitempty"`
	CommandMode              CommandMode `json:"CommandMode,omitempty"`
	DateCreated              time.Time   `json:"DateCreated,omitempty"`
	DateUpdated              time.Time   `json:"DateUpdated,omitempty"`
	DeliveryReceiptRequested bool        `json:"DeliveryReceiptRequested,omitempty"`
	Direction                Direction   `json:"Direction,omitempty"`
	Sid                      string      `json:"Sid,omitempty"`
	SimSid                   string      `json:"SimSid,omitempty"`
	Status                   Status      `json:"Status,omitempty"`
	Transport                Transport   `json:"Transport,omitempty"`
	Url                      string      `json:"Url,omitempty"`
}
