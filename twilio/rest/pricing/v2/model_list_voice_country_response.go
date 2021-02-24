/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListVoiceCountryResponse struct for ListVoiceCountryResponse
type ListVoiceCountryResponse struct {
	Countries []PricingV2VoiceVoiceCountry `json:"Countries,omitempty"`
	Meta      ListVoiceCountryResponseMeta `json:"Meta,omitempty"`
}
