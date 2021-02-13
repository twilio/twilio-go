/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// UpdateServiceRequest struct for UpdateServiceRequest
type UpdateServiceRequest struct {
	// The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
	CodeLength int32 `json:"CodeLength,omitempty"`
	// Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled bool `json:"CustomCodeEnabled,omitempty"`
	// Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.**
	DoNotShareWarningEnabled bool `json:"DoNotShareWarningEnabled,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call.
	DtmfInputRequired bool `json:"DtmfInputRequired,omitempty"`
	// A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled bool `json:"LookupEnabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled bool `json:"Psd2Enabled,omitempty"`
	// Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushApnCredentialSid string `json:"PushApnCredentialSid,omitempty"`
	// Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
	PushFcmCredentialSid string `json:"PushFcmCredentialSid,omitempty"`
	// Optional configuration for the Push factors. If true, include the date in the Challenge's reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
	PushIncludeDate bool `json:"PushIncludeDate,omitempty"`
	// Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
	SkipSmsToLandlines bool `json:"SkipSmsToLandlines,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
	TtsName string `json:"TtsName,omitempty"`
}
