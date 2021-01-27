/*
 * Twilio - Preview
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
// PreviewDeployedDevicesFleetDeployment struct for PreviewDeployedDevicesFleetDeployment
type PreviewDeployedDevicesFleetDeployment struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FleetSid string `json:"FleetSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SyncServiceSid string `json:"SyncServiceSid,omitempty"`
	Url string `json:"Url,omitempty"`
}
