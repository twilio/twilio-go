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
// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	// The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
	ByocTrunkSid string `json:"ByocTrunkSid,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \"-\".
	DomainName string `json:"DomainName,omitempty"`
	// Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
	EmergencyCallerSid string `json:"EmergencyCallerSid,omitempty"`
	// Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
	EmergencyCallingEnabled bool `json:"EmergencyCallingEnabled,omitempty"`
	// A descriptive string that you created to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
	Secure bool `json:"Secure,omitempty"`
	// Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not.
	SipRegistration bool `json:"SipRegistration,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`.
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	// The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`.
	VoiceStatusCallbackMethod string `json:"VoiceStatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	VoiceStatusCallbackUrl string `json:"VoiceStatusCallbackUrl,omitempty"`
	// The URL we should call when the domain receives a call.
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
