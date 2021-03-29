/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrunkingV1TrunkPhoneNumber struct for TrunkingV1TrunkPhoneNumber
type TrunkingV1TrunkPhoneNumber struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Whether the phone number requires an Address registered with Twilio
	AddressRequirements *string `json:"address_requirements,omitempty"`
	// The API version used to start a new TwiML session
	ApiVersion *string `json:"api_version,omitempty"`
	// Whether the phone number is new to the Twilio platform
	Beta *bool `json:"beta,omitempty"`
	// Indicate if a phone can receive calls or messages
	Capabilities *map[string]interface{} `json:"capabilities,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the application that handles SMS messages sent to the phone number
	SmsApplicationSid *string `json:"sms_application_sid,omitempty"`
	// The HTTP method used with sms_fallback_url
	SmsFallbackMethod *string `json:"sms_fallback_method,omitempty"`
	// The URL that we call when an error occurs while retrieving or executing the TwiML
	SmsFallbackUrl *string `json:"sms_fallback_url,omitempty"`
	// The HTTP method to use with sms_url
	SmsMethod *string `json:"sms_method,omitempty"`
	// The URL we call when the phone number receives an incoming SMS message
	SmsUrl *string `json:"sms_url,omitempty"`
	// The URL to send status information to your application
	StatusCallback *string `json:"status_callback,omitempty"`
	// The HTTP method we use to call status_callback
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// The SID of the Trunk that handles calls to the phone number
	TrunkSid *string `json:"trunk_sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// The SID of the application that handles calls to the phone number
	VoiceApplicationSid *string `json:"voice_application_sid,omitempty"`
	// Whether to lookup the caller's name
	VoiceCallerIdLookup *bool `json:"voice_caller_id_lookup,omitempty"`
	// The HTTP method that we use to call voice_fallback_url
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
	// The URL we call when an error occurs in TwiML
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
	// The HTTP method used with the voice_url
	VoiceMethod *string `json:"voice_method,omitempty"`
	// The URL we call when the phone number receives a call
	VoiceUrl *string `json:"voice_url,omitempty"`
}
