/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices struct for PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices
type PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices struct {
	BasePrice           float32  `json:"base_price,omitempty"`
	CurrentPrice        float32  `json:"current_price,omitempty"`
	DestinationPrefixes []string `json:"destination_prefixes,omitempty"`
	FriendlyName        string   `json:"friendly_name,omitempty"`
	OriginationPrefixes []string `json:"origination_prefixes,omitempty"`
}
