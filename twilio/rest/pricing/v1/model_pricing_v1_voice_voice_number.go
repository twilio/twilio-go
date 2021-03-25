/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV1VoiceVoiceNumber struct for PricingV1VoiceVoiceNumber
type PricingV1VoiceVoiceNumber struct {
	// The name of the country
	Country          *string                                    `json:"Country,omitempty"`
	InboundCallPrice *PricingV1VoiceVoiceNumberInboundCallPrice `json:"InboundCallPrice,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// The phone number
	Number            *string                                     `json:"Number,omitempty"`
	OutboundCallPrice *PricingV1VoiceVoiceNumberOutboundCallPrice `json:"OutboundCallPrice,omitempty"`
	// The currency in which prices are measured, in ISO 4127 format (e.g. usd, eur, jpy)
	PriceUnit *string `json:"PriceUnit,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
