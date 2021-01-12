/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// MessagingV1ServiceShortCode struct for MessagingV1ServiceShortCode
type MessagingV1ServiceShortCode struct {
	AccountSid string `json:"account_sid,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	ShortCode string `json:"short_code,omitempty"`
	Sid string `json:"sid,omitempty"`
	Url string `json:"url,omitempty"`
}
