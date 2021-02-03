/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// VerifyV2ServiceRateLimitBucket struct for VerifyV2ServiceRateLimitBucket
type VerifyV2ServiceRateLimitBucket struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Interval int32 `json:"Interval,omitempty"`
	Max int32 `json:"Max,omitempty"`
	RateLimitSid string `json:"RateLimitSid,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
