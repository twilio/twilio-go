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
// PreviewDeployedDevicesFleetKey struct for PreviewDeployedDevicesFleetKey
type PreviewDeployedDevicesFleetKey struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DeviceSid string `json:"DeviceSid,omitempty"`
	FleetSid string `json:"FleetSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Secret string `json:"Secret,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
