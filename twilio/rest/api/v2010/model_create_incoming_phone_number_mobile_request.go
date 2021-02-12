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
// CreateIncomingPhoneNumberMobileRequest struct for CreateIncomingPhoneNumberMobileRequest
type CreateIncomingPhoneNumberMobileRequest struct {
	// The SID of the Address resource we should associate with the new phone number. Some regions require addresses to meet local regulations.
	AddressSid string `json:"AddressSid,omitempty"`
	// The API version to use for incoming calls made to the new phone number. The default is `2010-04-01`.
	ApiVersion string `json:"ApiVersion,omitempty"`
	// The SID of the Bundle resource that you associate with the phone number. Some regions require a Bundle to meet local Regulations.
	BundleSid string `json:"BundleSid,omitempty"`
	// The SID of the emergency address configuration to use for emergency calling from the new phone number.
	EmergencyAddressSid string `json:"EmergencyAddressSid,omitempty"`
	// The configuration status parameter that determines whether the new phone number is enabled for emergency calling.
	EmergencyStatus string `json:"EmergencyStatus,omitempty"`
	// A descriptive string that you created to describe the new phone number. It can be up to 64 characters long. By default, the is a formatted version of the phone number.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID of the Identity resource that we should associate with the new phone number. Some regions require an identity to meet local regulations.
	IdentitySid string `json:"IdentitySid,omitempty"`
	// The phone number to purchase specified in [E.164](https://www.twilio.com/docs/glossary/what-e164) format.  E.164 phone numbers consist of a + followed by the country code and subscriber number without punctuation characters. For example, +14155551234.
	PhoneNumber string `json:"PhoneNumber"`
	// The SID of the application that should handle SMS messages sent to the new phone number. If an `sms_application_sid` is present, we ignore all of the `sms_*_url` urls and use those of the application.
	SmsApplicationSid string `json:"SmsApplicationSid,omitempty"`
	// The HTTP method that we should use to call `sms_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while requesting or executing the TwiML defined by `sms_url`.
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `sms_url`. Can be: `GET` or `POST` and defaults to `POST`.
	SmsMethod string `json:"SmsMethod,omitempty"`
	// The URL we should call when the new phone number receives an incoming SMS message.
	SmsUrl string `json:"SmsUrl,omitempty"`
	// The URL we should call using the `status_callback_method` to send status information to your application.
	StatusCallback string `json:"StatusCallback,omitempty"`
	// The HTTP method we should use to call `status_callback`. Can be: `GET` or `POST` and defaults to `POST`.
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	// The SID of the Trunk we should use to handle calls to the new phone number. If a `trunk_sid` is present, we ignore all of the voice urls and voice applications and use only those set on the Trunk. Setting a `trunk_sid` will automatically delete your `voice_application_sid` and vice versa.
	TrunkSid string `json:"TrunkSid,omitempty"`
	// The SID of the application we should use to handle calls to the new phone number. If a `voice_application_sid` is present, we ignore all of the voice urls and use only those set on the application. Setting a `voice_application_sid` will automatically delete your `trunk_sid` and vice versa.
	VoiceApplicationSid string `json:"VoiceApplicationSid,omitempty"`
	// Whether to lookup the caller's name from the CNAM database and post it to your app. Can be: `true` or `false` and defaults to `false`.
	VoiceCallerIdLookup bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method that we should use to call `voice_fallback_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs retrieving or executing the TwiML requested by `url`.
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method that we should use to call `voice_url`. Can be: `GET` or `POST` and defaults to `POST`.
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	// The configuration parameter for the new phone number to receive incoming voice calls or faxes. Can be: `fax` or `voice` and defaults to `voice`.
	VoiceReceiveMode string `json:"VoiceReceiveMode,omitempty"`
	// The URL that we should call to answer a call to the new phone number. The `voice_url` will not be called if a `voice_application_sid` or a `trunk_sid` is set.
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
