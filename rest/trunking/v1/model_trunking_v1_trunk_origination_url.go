/*
 * Twilio - Trunking
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
// TrunkingV1TrunkOriginationUrl struct for TrunkingV1TrunkOriginationUrl
type TrunkingV1TrunkOriginationUrl struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Enabled bool `json:"Enabled,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Priority int32 `json:"Priority,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SipUrl string `json:"SipUrl,omitempty"`
	TrunkSid string `json:"TrunkSid,omitempty"`
	Url string `json:"Url,omitempty"`
	Weight int32 `json:"Weight,omitempty"`
}