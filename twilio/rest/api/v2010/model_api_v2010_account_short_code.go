/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010AccountShortCode struct for ApiV2010AccountShortCode
type ApiV2010AccountShortCode struct {
	AccountSid        *string     `json:"AccountSid,omitempty"`
	ApiVersion        *string     `json:"ApiVersion,omitempty"`
	DateCreated       *string     `json:"DateCreated,omitempty"`
	DateUpdated       *string     `json:"DateUpdated,omitempty"`
	FriendlyName      *string     `json:"FriendlyName,omitempty"`
	ShortCode         *string     `json:"ShortCode,omitempty"`
	Sid               *string     `json:"Sid,omitempty"`
	SmsFallbackMethod *HttpMethod `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl    *string     `json:"SmsFallbackUrl,omitempty"`
	SmsMethod         *HttpMethod `json:"SmsMethod,omitempty"`
	SmsUrl            *string     `json:"SmsUrl,omitempty"`
	Uri               *string     `json:"Uri,omitempty"`
}
