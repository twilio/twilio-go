/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// NotifyV1ServiceNotification struct for NotifyV1ServiceNotification
type NotifyV1ServiceNotification struct {
	AccountSid        *string                 `json:"account_sid,omitempty"`
	Action            *string                 `json:"action,omitempty"`
	Alexa             *map[string]interface{} `json:"alexa,omitempty"`
	Apn               *map[string]interface{} `json:"apn,omitempty"`
	Body              *string                 `json:"body,omitempty"`
	Data              *map[string]interface{} `json:"data,omitempty"`
	DateCreated       *time.Time              `json:"date_created,omitempty"`
	FacebookMessenger *map[string]interface{} `json:"facebook_messenger,omitempty"`
	Fcm               *map[string]interface{} `json:"fcm,omitempty"`
	Gcm               *map[string]interface{} `json:"gcm,omitempty"`
	Identities        *[]string               `json:"identities,omitempty"`
	Priority          *NotificationPriority   `json:"priority,omitempty"`
	Segments          *[]string               `json:"segments,omitempty"`
	ServiceSid        *string                 `json:"service_sid,omitempty"`
	Sid               *string                 `json:"sid,omitempty"`
	Sms               *map[string]interface{} `json:"sms,omitempty"`
	Sound             *string                 `json:"sound,omitempty"`
	Tags              *[]string               `json:"tags,omitempty"`
	Title             *string                 `json:"title,omitempty"`
	Ttl               *int32                  `json:"ttl,omitempty"`
}
