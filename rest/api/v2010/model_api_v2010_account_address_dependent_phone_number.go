/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAddressDependentPhoneNumber struct for ApiV2010AccountAddressDependentPhoneNumber
type ApiV2010AccountAddressDependentPhoneNumber struct {
	AccountSid           *string                                 `json:"account_sid,omitempty"`
	AddressRequirements  *DependentPhoneNumberAddressRequirement `json:"address_requirements,omitempty"`
	ApiVersion           *string                                 `json:"api_version,omitempty"`
	Capabilities         *map[string]interface{}                 `json:"capabilities,omitempty"`
	DateCreated          *string                                 `json:"date_created,omitempty"`
	DateUpdated          *string                                 `json:"date_updated,omitempty"`
	EmergencyAddressSid  *string                                 `json:"emergency_address_sid,omitempty"`
	EmergencyStatus      *DependentPhoneNumberEmergencyStatus    `json:"emergency_status,omitempty"`
	FriendlyName         *string                                 `json:"friendly_name,omitempty"`
	PhoneNumber          *string                                 `json:"phone_number,omitempty"`
	Sid                  *string                                 `json:"sid,omitempty"`
	SmsApplicationSid    *string                                 `json:"sms_application_sid,omitempty"`
	SmsFallbackMethod    *HttpMethod                             `json:"sms_fallback_method,omitempty"`
	SmsFallbackUrl       *string                                 `json:"sms_fallback_url,omitempty"`
	SmsMethod            *HttpMethod                             `json:"sms_method,omitempty"`
	SmsUrl               *string                                 `json:"sms_url,omitempty"`
	StatusCallback       *string                                 `json:"status_callback,omitempty"`
	StatusCallbackMethod *HttpMethod                             `json:"status_callback_method,omitempty"`
	TrunkSid             *string                                 `json:"trunk_sid,omitempty"`
	Uri                  *string                                 `json:"uri,omitempty"`
	VoiceApplicationSid  *string                                 `json:"voice_application_sid,omitempty"`
	VoiceCallerIdLookup  *bool                                   `json:"voice_caller_id_lookup,omitempty"`
	VoiceFallbackMethod  *HttpMethod                             `json:"voice_fallback_method,omitempty"`
	VoiceFallbackUrl     *string                                 `json:"voice_fallback_url,omitempty"`
	VoiceMethod          *HttpMethod                             `json:"voice_method,omitempty"`
	VoiceUrl             *string                                 `json:"voice_url,omitempty"`
}
