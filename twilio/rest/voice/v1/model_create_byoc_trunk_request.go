/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateByocTrunkRequest struct for CreateByocTrunkRequest
type CreateByocTrunkRequest struct {
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
	CnamLookupEnabled bool `json:"CnamLookupEnabled,omitempty"`
	// The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
	ConnectionPolicySid string `json:"ConnectionPolicySid,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \"call back\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \"sip.twilio.com\".
	FromDomainSid string `json:"FromDomainSid,omitempty"`
	// The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	StatusCallbackUrl string `json:"StatusCallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the BYOC Trunk receives a call.
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
