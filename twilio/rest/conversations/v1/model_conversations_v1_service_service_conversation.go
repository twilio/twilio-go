/*
 * Twilio - Conversations
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

// ConversationsV1ServiceServiceConversation struct for ConversationsV1ServiceServiceConversation
type ConversationsV1ServiceServiceConversation struct {
	AccountSid          string                 `json:"AccountSid,omitempty"`
	Attributes          string                 `json:"Attributes,omitempty"`
	ChatServiceSid      string                 `json:"ChatServiceSid,omitempty"`
	DateCreated         time.Time              `json:"DateCreated,omitempty"`
	DateUpdated         time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName        string                 `json:"FriendlyName,omitempty"`
	Links               map[string]interface{} `json:"Links,omitempty"`
	MessagingServiceSid string                 `json:"MessagingServiceSid,omitempty"`
	Sid                 string                 `json:"Sid,omitempty"`
	State               State                  `json:"State,omitempty"`
	Timers              map[string]interface{} `json:"Timers,omitempty"`
	UniqueName          string                 `json:"UniqueName,omitempty"`
	Url                 string                 `json:"Url,omitempty"`
}
