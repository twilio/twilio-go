/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010ShortCode struct for ApiV2010ShortCode
type ApiV2010ShortCode struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to start a new TwiML session
	ApiVersion *string `json:"api_version,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A string that you assigned to describe this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The short code. e.g., 894546.
	ShortCode *string `json:"short_code,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// HTTP method we use to call the sms_fallback_url
	SmsFallbackMethod *string `json:"sms_fallback_method,omitempty"`
	// URL Twilio will request if an error occurs in executing TwiML
	SmsFallbackUrl *string `json:"sms_fallback_url,omitempty"`
	// HTTP method to use when requesting the sms url
	SmsMethod *string `json:"sms_method,omitempty"`
	// URL we call when receiving an incoming SMS message to this short code
	SmsUrl *string `json:"sms_url,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
