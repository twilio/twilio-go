/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAddressDependentPhoneNumber struct for ApiV2010AccountAddressDependentPhoneNumber
type ApiV2010AccountAddressDependentPhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Whether the phone number requires an Address registered with Twilio
	AddressRequirements *string `json:"AddressRequirements,omitempty"`
	// The API version used to start a new TwiML session
	ApiVersion *string `json:"ApiVersion,omitempty"`
	// Indicate if a phone can receive calls or messages
	Capabilities *map[string]interface{} `json:"Capabilities,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"DateUpdated,omitempty"`
	// The emergency address configuration to use for emergency calling
	EmergencyAddressSid *string `json:"EmergencyAddressSid,omitempty"`
	// Whether the phone number is enabled for emergency calling
	EmergencyStatus *string `json:"EmergencyStatus,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"PhoneNumber,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The SID of the application that handles SMS messages sent to the phone number
	SmsApplicationSid *string `json:"SmsApplicationSid,omitempty"`
	// The HTTP method used with sms_fallback_url
	SmsFallbackMethod *string `json:"SmsFallbackMethod,omitempty"`
	// The URL that we call when an error occurs while retrieving or executing the TwiML
	SmsFallbackUrl *string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method to use with sms_url
	SmsMethod *string `json:"SmsMethod,omitempty"`
	// The URL we call when the phone number receives an incoming SMS message
	SmsUrl *string `json:"SmsUrl,omitempty"`
	// The URL to send status information to your application
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method we use to call status_callback
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
	// The SID of the Trunk that handles calls to the phone number
	TrunkSid *string `json:"TrunkSid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
	// The SID of the application that handles calls to the phone number
	VoiceApplicationSid *string `json:"VoiceApplicationSid,omitempty"`
	// Whether to lookup the caller's name
	VoiceCallerIdLookup *bool `json:"VoiceCallerIdLookup,omitempty"`
	// The HTTP method used with voice_fallback_url
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL we call when an error occurs in TwiML
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method used with the voice_url
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The URL we call when the phone number receives a call
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}
