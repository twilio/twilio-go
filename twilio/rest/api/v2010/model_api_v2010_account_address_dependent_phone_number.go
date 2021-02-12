/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountAddressDependentPhoneNumber struct for ApiV2010AccountAddressDependentPhoneNumber
type ApiV2010AccountAddressDependentPhoneNumber struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AddressRequirements AddressRequirement `json:"AddressRequirements,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	Capabilities map[string]interface{} `json:"Capabilities,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	EmergencyAddressSid string `json:"EmergencyAddressSid,omitempty"`
	EmergencyStatus EmergencyStatus `json:"EmergencyStatus,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SmsApplicationSid string `json:"SmsApplicationSid,omitempty"`
	SmsFallbackMethod HttpMethod `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	SmsMethod HttpMethod `json:"SmsMethod,omitempty"`
	SmsUrl string `json:"SmsUrl,omitempty"`
	StatusCallback string `json:"StatusCallback,omitempty"`
	StatusCallbackMethod HttpMethod `json:"StatusCallbackMethod,omitempty"`
	TrunkSid string `json:"TrunkSid,omitempty"`
	Uri string `json:"Uri,omitempty"`
	VoiceApplicationSid string `json:"VoiceApplicationSid,omitempty"`
	VoiceCallerIdLookup bool `json:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod HttpMethod `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod HttpMethod `json:"VoiceMethod,omitempty"`
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
