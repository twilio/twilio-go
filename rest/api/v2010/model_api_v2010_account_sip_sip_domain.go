/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountSipSipDomain struct for ApiV2010AccountSipSipDomain
type ApiV2010AccountSipSipDomain struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	AuthType string `json:"AuthType,omitempty"`
	ByocTrunkSid string `json:"ByocTrunkSid,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	DomainName string `json:"DomainName,omitempty"`
	EmergencyCallerSid string `json:"EmergencyCallerSid,omitempty"`
	EmergencyCallingEnabled bool `json:"EmergencyCallingEnabled,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Secure bool `json:"Secure,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SipRegistration bool `json:"SipRegistration,omitempty"`
	SubresourceUris map[string]interface{} `json:"SubresourceUris,omitempty"`
	Uri string `json:"Uri,omitempty"`
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	VoiceStatusCallbackMethod string `json:"VoiceStatusCallbackMethod,omitempty"`
	VoiceStatusCallbackUrl string `json:"VoiceStatusCallbackUrl,omitempty"`
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
