/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// WirelessV1SimDataSession struct for WirelessV1SimDataSession
type WirelessV1SimDataSession struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CellId string `json:"CellId,omitempty"`
	CellLocationEstimate map[string]interface{} `json:"CellLocationEstimate,omitempty"`
	End time.Time `json:"End,omitempty"`
	Imei string `json:"Imei,omitempty"`
	LastUpdated time.Time `json:"LastUpdated,omitempty"`
	OperatorCountry string `json:"OperatorCountry,omitempty"`
	OperatorMcc string `json:"OperatorMcc,omitempty"`
	OperatorMnc string `json:"OperatorMnc,omitempty"`
	OperatorName string `json:"OperatorName,omitempty"`
	PacketsDownloaded int32 `json:"PacketsDownloaded,omitempty"`
	PacketsUploaded int32 `json:"PacketsUploaded,omitempty"`
	RadioLink string `json:"RadioLink,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SimSid string `json:"SimSid,omitempty"`
	Start time.Time `json:"Start,omitempty"`
}
