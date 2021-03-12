/*
    * Twilio - Pricing
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// PricingV2VoiceVoiceCountryInstance struct for PricingV2VoiceVoiceCountryInstance
type PricingV2VoiceVoiceCountryInstance struct {
	// The name of the country
	Country *string `json:"Country,omitempty"`
	// The list of InboundCallPrice records
	InboundCallPrices *[]PricingV2VoiceVoiceCountryInstanceInboundCallPrices `json:"InboundCallPrices,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The list of OutboundPrefixPriceWithOrigin records
	OutboundPrefixPrices *[]PricingV2VoiceVoiceCountryInstanceOutboundPrefixPrices `json:"OutboundPrefixPrices,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
