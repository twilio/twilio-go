/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
// InlineObject18 struct for InlineObject18
type InlineObject18 struct {
	// The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
	CodeLength int32 `json:"CodeLength,omitempty"`
	// Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
	CustomCodeEnabled bool `json:"CustomCodeEnabled,omitempty"`
	// Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.**
	DoNotShareWarningEnabled bool `json:"DoNotShareWarningEnabled,omitempty"`
	// Whether to ask the user to press a number before delivering the verify code in a phone call.
	DtmfInputRequired bool `json:"DtmfInputRequired,omitempty"`
	// A descriptive string that you create to describe the verification service. It can be up to 64 characters long. **This value should not contain PII.**
	FriendlyName string `json:"FriendlyName,omitempty"`
	// Whether to perform a lookup with each verification started and return info about the phone number.
	LookupEnabled bool `json:"LookupEnabled,omitempty"`
	// Whether to pass PSD2 transaction parameters when starting a verification.
	Psd2Enabled bool `json:"Psd2Enabled,omitempty"`
	// Configurations for the Push factors (channel) created under this Service. If present, it must be a json string with the following format: {\"notify_service_sid\": \"ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\", \"include_date\": true}. If `include_date` is set to `true`, which is the default, that means that the push challenge’s response will include the date created value. If `include_date` is set to `false`, then the date created value will not be included. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info
	Push map[string]interface{} `json:"Push,omitempty"`
	// Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`.
	SkipSmsToLandlines bool `json:"SkipSmsToLandlines,omitempty"`
	// The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.
	TtsName string `json:"TtsName,omitempty"`
}