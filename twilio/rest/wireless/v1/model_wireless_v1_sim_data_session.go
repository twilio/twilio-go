/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// WirelessV1SimDataSession struct for WirelessV1SimDataSession
type WirelessV1SimDataSession struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated
	CellId *string `json:"CellId,omitempty"`
	// An object with the estimated location where the device's Data Session took place
	CellLocationEstimate *map[string]interface{} `json:"CellLocationEstimate,omitempty"`
	// The date that the record ended, given as GMT in ISO 8601 format
	End *time.Time `json:"End,omitempty"`
	// The unique ID of the device using the SIM to connect
	Imei *string `json:"Imei,omitempty"`
	// The date that the resource was last updated, given as GMT in ISO 8601 format
	LastUpdated *time.Time `json:"LastUpdated,omitempty"`
	// The three letter country code representing where the device's Data Session took place
	OperatorCountry *string `json:"OperatorCountry,omitempty"`
	// The 'mobile country code' is the unique ID of the home country where the Data Session took place
	OperatorMcc *string `json:"OperatorMcc,omitempty"`
	// The 'mobile network code' is the unique ID specific to the mobile operator network where the Data Session took place
	OperatorMnc *string `json:"OperatorMnc,omitempty"`
	// The friendly name of the mobile operator network that the SIM-connected device is attached to
	OperatorName *string `json:"OperatorName,omitempty"`
	// The number of packets downloaded by the device between the start time and when the Data Session was last updated
	PacketsDownloaded *int32 `json:"PacketsDownloaded,omitempty"`
	// The number of packets uploaded by the device between the start time and when the Data Session was last updated
	PacketsUploaded *int32 `json:"PacketsUploaded,omitempty"`
	// The generation of wireless technology that the device was using
	RadioLink *string `json:"RadioLink,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The SID of the Sim resource that the Data Session is for
	SimSid *string `json:"SimSid,omitempty"`
	// The date that the Data Session started, given as GMT in ISO 8601 format
	Start *time.Time `json:"Start,omitempty"`
}
