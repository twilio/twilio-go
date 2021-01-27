/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ProxyV1Service struct for ProxyV1Service
type ProxyV1Service struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	ChatInstanceSid string `json:"ChatInstanceSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DefaultTtl int32 `json:"DefaultTtl,omitempty"`
	GeoMatchLevel string `json:"GeoMatchLevel,omitempty"`
	InterceptCallbackUrl string `json:"InterceptCallbackUrl,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	NumberSelectionBehavior string `json:"NumberSelectionBehavior,omitempty"`
	OutOfSessionCallbackUrl string `json:"OutOfSessionCallbackUrl,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
