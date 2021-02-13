/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ChatV1Credential struct for ChatV1Credential
type ChatV1Credential struct {
	AccountSid   string      `json:"AccountSid,omitempty"`
	DateCreated  time.Time   `json:"DateCreated,omitempty"`
	DateUpdated  time.Time   `json:"DateUpdated,omitempty"`
	FriendlyName string      `json:"FriendlyName,omitempty"`
	Sandbox      string      `json:"Sandbox,omitempty"`
	Sid          string      `json:"Sid,omitempty"`
	Type         PushService `json:"Type,omitempty"`
	Url          string      `json:"Url,omitempty"`
}
