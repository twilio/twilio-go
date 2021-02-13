/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PricingV1VoiceVoiceCountryInstance struct for PricingV1VoiceVoiceCountryInstance
type PricingV1VoiceVoiceCountryInstance struct {
	Country              string                   `json:"Country,omitempty"`
	InboundCallPrices    []map[string]interface{} `json:"InboundCallPrices,omitempty"`
	IsoCountry           string                   `json:"IsoCountry,omitempty"`
	OutboundPrefixPrices []map[string]interface{} `json:"OutboundPrefixPrices,omitempty"`
	PriceUnit            string                   `json:"PriceUnit,omitempty"`
	Url                  string                   `json:"Url,omitempty"`
}
