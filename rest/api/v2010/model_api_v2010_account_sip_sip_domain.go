/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountSipSipDomain struct for ApiV2010AccountSipSipDomain
type ApiV2010AccountSipSipDomain struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to process the call
	ApiVersion *string `json:"api_version,omitempty"`
	// The types of authentication mapped to the domain
	AuthType *string `json:"auth_type,omitempty"`
	// The SID of the BYOC Trunk resource.
	ByocTrunkSid *string `json:"byoc_trunk_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The unique address on Twilio to route SIP traffic
	DomainName *string `json:"domain_name,omitempty"`
	// Whether an emergency caller sid is configured for the domain.
	EmergencyCallerSid *string `json:"emergency_caller_sid,omitempty"`
	// Whether emergency calling is enabled for the domain.
	EmergencyCallingEnabled *bool `json:"emergency_calling_enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Whether secure SIP is enabled for the domain
	Secure *bool `json:"secure,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// Whether SIP registration is allowed
	SipRegistration *bool `json:"sip_registration,omitempty"`
	// A list mapping resources associated with the SIP Domain resource
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
	// The HTTP method used with voice_fallback_url
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
	// The URL we call when an error occurs while executing TwiML
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
	// The HTTP method to use with voice_url
	VoiceMethod *string `json:"voice_method,omitempty"`
	// The HTTP method we use to call voice_status_callback_url
	VoiceStatusCallbackMethod *string `json:"voice_status_callback_method,omitempty"`
	// The URL that we call with status updates
	VoiceStatusCallbackUrl *string `json:"voice_status_callback_url,omitempty"`
	// The URL we call when receiving a call
	VoiceUrl *string `json:"voice_url,omitempty"`
}
