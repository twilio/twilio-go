/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateApplicationRequest struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	// The API version to use to start a new TwiML session. Can be: `2010-04-01` or `2008-08-01`. The default value is your account's default API version.
	ApiVersion string `json:"ApiVersion,omitempty"`
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The URL we should call using a POST method to send message status information to your application.
	MessageStatusCallback string `json:"MessageStatusCallback,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`.
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `sms_url`.
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`.
	SmsMethod string `json:"SmsMethod,omitempty"`
	// Same as message_status_callback: The URL we should call using a POST method to send status information about SMS messages sent by the application. Deprecated, included for backwards compatibility.
	SmsStatusCallback string `json:"SmsStatusCallback,omitempty"`
	// The URL we should call when the phone number receives an incoming SMS message.
	SmsUrl string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST`.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// Whether we should look up the caller's caller-ID name from the CNAM database (additional charges apply). Can be: `true` or `false`.
	VoiceCallerIdLookup bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	// The URL we should call when the phone number assigned to this application receives a call.
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
