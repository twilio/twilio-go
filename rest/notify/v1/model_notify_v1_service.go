/*
 * Twilio - Notify
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
// NotifyV1Service struct for NotifyV1Service
type NotifyV1Service struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AlexaSkillId string `json:"AlexaSkillId,omitempty"`
	ApnCredentialSid string `json:"ApnCredentialSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	DefaultAlexaNotificationProtocolVersion string `json:"DefaultAlexaNotificationProtocolVersion,omitempty"`
	DefaultApnNotificationProtocolVersion string `json:"DefaultApnNotificationProtocolVersion,omitempty"`
	DefaultFcmNotificationProtocolVersion string `json:"DefaultFcmNotificationProtocolVersion,omitempty"`
	DefaultGcmNotificationProtocolVersion string `json:"DefaultGcmNotificationProtocolVersion,omitempty"`
	DeliveryCallbackEnabled bool `json:"DeliveryCallbackEnabled,omitempty"`
	DeliveryCallbackUrl string `json:"DeliveryCallbackUrl,omitempty"`
	FacebookMessengerPageId string `json:"FacebookMessengerPageId,omitempty"`
	FcmCredentialSid string `json:"FcmCredentialSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	GcmCredentialSid string `json:"GcmCredentialSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	LogEnabled bool `json:"LogEnabled,omitempty"`
	MessagingServiceSid string `json:"MessagingServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
