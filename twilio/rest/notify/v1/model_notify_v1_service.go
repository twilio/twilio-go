/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// NotifyV1Service struct for NotifyV1Service
type NotifyV1Service struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Deprecated
	AlexaSkillId *string `json:"AlexaSkillId,omitempty"`
	// The SID of the Credential to use for APN Bindings
	ApnCredentialSid *string `json:"ApnCredentialSid,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// Deprecated
	DefaultAlexaNotificationProtocolVersion *string `json:"DefaultAlexaNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending APNS notifications
	DefaultApnNotificationProtocolVersion *string `json:"DefaultApnNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending FCM notifications
	DefaultFcmNotificationProtocolVersion *string `json:"DefaultFcmNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending GCM notifications
	DefaultGcmNotificationProtocolVersion *string `json:"DefaultGcmNotificationProtocolVersion,omitempty"`
	// Enable delivery callbacks
	DeliveryCallbackEnabled *bool `json:"DeliveryCallbackEnabled,omitempty"`
	// Webhook URL
	DeliveryCallbackUrl *string `json:"DeliveryCallbackUrl,omitempty"`
	// Deprecated
	FacebookMessengerPageId *string `json:"FacebookMessengerPageId,omitempty"`
	// The SID of the Credential to use for FCM Bindings
	FcmCredentialSid *string `json:"FcmCredentialSid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the Credential to use for GCM Bindings
	GcmCredentialSid *string `json:"GcmCredentialSid,omitempty"`
	// The URLs of the resources related to the service
	Links *map[string]interface{} `json:"Links,omitempty"`
	// Whether to log notifications
	LogEnabled *bool `json:"LogEnabled,omitempty"`
	// The SID of the Messaging Service to use for SMS Bindings
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"Url,omitempty"`
}
