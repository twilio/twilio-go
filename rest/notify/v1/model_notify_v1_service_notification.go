/*
 * Twilio - Notify
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
// NotifyV1ServiceNotification struct for NotifyV1ServiceNotification
type NotifyV1ServiceNotification struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Action string `json:"Action,omitempty"`
	Alexa map[string]interface{} `json:"Alexa,omitempty"`
	Apn map[string]interface{} `json:"Apn,omitempty"`
	Body string `json:"Body,omitempty"`
	Data map[string]interface{} `json:"Data,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	FacebookMessenger map[string]interface{} `json:"FacebookMessenger,omitempty"`
	Fcm map[string]interface{} `json:"Fcm,omitempty"`
	Gcm map[string]interface{} `json:"Gcm,omitempty"`
	Identities []string `json:"Identities,omitempty"`
	Priority string `json:"Priority,omitempty"`
	Segments []string `json:"Segments,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Sms map[string]interface{} `json:"Sms,omitempty"`
	Sound string `json:"Sound,omitempty"`
	Tags []string `json:"Tags,omitempty"`
	Title string `json:"Title,omitempty"`
	Ttl int32 `json:"Ttl,omitempty"`
}
