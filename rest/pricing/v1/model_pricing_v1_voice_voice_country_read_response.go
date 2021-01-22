/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PricingV1VoiceVoiceCountryReadResponse struct for PricingV1VoiceVoiceCountryReadResponse
type PricingV1VoiceVoiceCountryReadResponse struct {
	Countries []PricingV1VoiceVoiceCountry `json:"Countries,omitempty"`
	Meta PricingV1MessagingMessagingCountryReadResponseMeta `json:"Meta,omitempty"`
}
