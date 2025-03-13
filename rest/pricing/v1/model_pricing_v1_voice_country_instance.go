/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
)

// PricingV1VoiceCountryInstance struct for PricingV1VoiceCountryInstance
type PricingV1VoiceCountryInstance struct {
	// The name of the country.
	Country *string `json:"country,omitempty"`
	// The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	IsoCountry *string `json:"iso_country,omitempty"`
	// The list of OutboundPrefixPrice records, which include a list of the `prefixes`, the `friendly_name`, `base_price`, and the   `current_price` for those prefixes.
	OutboundPrefixPrices *[]PricingV1VoiceVoiceCountryInstanceOutboundPrefixPrices `json:"outbound_prefix_prices,omitempty"`
	// The list of [InboundCallPrice](https://www.twilio.com/docs/voice/pricing#inbound-call-price) records.
	InboundCallPrices *[]PricingV1VoiceVoiceCountryInstanceInboundCallPrices `json:"inbound_call_prices,omitempty"`
	// The currency in which prices are measured, specified in [ISO 4127](http://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`).
	PriceUnit *string `json:"price_unit,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
