/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VoiceV1DialingPermissionsDialingPermissionsCountryInstance struct for VoiceV1DialingPermissionsDialingPermissionsCountryInstance
type VoiceV1DialingPermissionsDialingPermissionsCountryInstance struct {
	Continent                       *string                 `json:"continent,omitempty"`
	CountryCodes                    *[]string               `json:"country_codes,omitempty"`
	HighRiskSpecialNumbersEnabled   *bool                   `json:"high_risk_special_numbers_enabled,omitempty"`
	HighRiskTollfraudNumbersEnabled *bool                   `json:"high_risk_tollfraud_numbers_enabled,omitempty"`
	IsoCode                         *string                 `json:"iso_code,omitempty"`
	Links                           *map[string]interface{} `json:"links,omitempty"`
	LowRiskNumbersEnabled           *bool                   `json:"low_risk_numbers_enabled,omitempty"`
	Name                            *string                 `json:"name,omitempty"`
	Url                             *string                 `json:"url,omitempty"`
}
