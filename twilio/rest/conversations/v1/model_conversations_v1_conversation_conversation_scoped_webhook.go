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

// ConversationsV1ConversationConversationScopedWebhook struct for ConversationsV1ConversationConversationScopedWebhook
type ConversationsV1ConversationConversationScopedWebhook struct {
	AccountSid      string                 `json:"AccountSid,omitempty"`
	Configuration   map[string]interface{} `json:"Configuration,omitempty"`
	ConversationSid string                 `json:"ConversationSid,omitempty"`
	DateCreated     time.Time              `json:"DateCreated,omitempty"`
	DateUpdated     time.Time              `json:"DateUpdated,omitempty"`
	Sid             string                 `json:"Sid,omitempty"`
	Target          string                 `json:"Target,omitempty"`
	Url             string                 `json:"Url,omitempty"`
}
