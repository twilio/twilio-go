/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// VerifyV2Service struct for VerifyV2Service
type VerifyV2Service struct {
	AccountSid string `json:"account_sid,omitempty"`
	CodeLength int32 `json:"code_length,omitempty"`
	CustomCodeEnabled bool `json:"custom_code_enabled,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	DoNotShareWarningEnabled bool `json:"do_not_share_warning_enabled,omitempty"`
	DtmfInputRequired bool `json:"dtmf_input_required,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
	LookupEnabled bool `json:"lookup_enabled,omitempty"`
	Psd2Enabled bool `json:"psd2_enabled,omitempty"`
	Push map[string]interface{} `json:"push,omitempty"`
	Sid string `json:"sid,omitempty"`
	SkipSmsToLandlines bool `json:"skip_sms_to_landlines,omitempty"`
	TtsName string `json:"tts_name,omitempty"`
	Url string `json:"url,omitempty"`
}
