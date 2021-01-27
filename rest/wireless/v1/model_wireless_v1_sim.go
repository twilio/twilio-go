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
// WirelessV1Sim struct for WirelessV1Sim
type WirelessV1Sim struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CommandsCallbackMethod string `json:"CommandsCallbackMethod,omitempty"`
	CommandsCallbackUrl string `json:"CommandsCallbackUrl,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	EId string `json:"EId,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Iccid string `json:"Iccid,omitempty"`
	IpAddress string `json:"IpAddress,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	RatePlanSid string `json:"RatePlanSid,omitempty"`
	ResetStatus string `json:"ResetStatus,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	SmsMethod string `json:"SmsMethod,omitempty"`
	SmsUrl string `json:"SmsUrl,omitempty"`
	Status string `json:"Status,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
