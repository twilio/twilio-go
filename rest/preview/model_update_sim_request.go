/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateSimRequest struct for UpdateSimRequest
type UpdateSimRequest struct {
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	CommandsCallbackMethod string `json:"CommandsCallbackMethod,omitempty"`
	CommandsCallbackUrl string `json:"CommandsCallbackUrl,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	RatePlan string `json:"RatePlan,omitempty"`
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	SmsMethod string `json:"SmsMethod,omitempty"`
	SmsUrl string `json:"SmsUrl,omitempty"`
	Status string `json:"Status,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
