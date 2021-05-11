/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountNotification struct for ApiV2010AccountNotification
type ApiV2010AccountNotification struct {
	AccountSid    *string     `json:"account_sid,omitempty"`
	ApiVersion    *string     `json:"api_version,omitempty"`
	CallSid       *string     `json:"call_sid,omitempty"`
	DateCreated   *string     `json:"date_created,omitempty"`
	DateUpdated   *string     `json:"date_updated,omitempty"`
	ErrorCode     *string     `json:"error_code,omitempty"`
	Log           *string     `json:"log,omitempty"`
	MessageDate   *string     `json:"message_date,omitempty"`
	MessageText   *string     `json:"message_text,omitempty"`
	MoreInfo      *string     `json:"more_info,omitempty"`
	RequestMethod *HttpMethod `json:"request_method,omitempty"`
	RequestUrl    *string     `json:"request_url,omitempty"`
	Sid           *string     `json:"sid,omitempty"`
	Uri           *string     `json:"uri,omitempty"`
}
