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
// ConversationsV1ServiceServiceConversationServiceConversationMessage struct for ConversationsV1ServiceServiceConversationServiceConversationMessage
type ConversationsV1ServiceServiceConversationServiceConversationMessage struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Attributes string `json:"Attributes,omitempty"`
	Author string `json:"Author,omitempty"`
	Body string `json:"Body,omitempty"`
	ChatServiceSid string `json:"ChatServiceSid,omitempty"`
	ConversationSid string `json:"ConversationSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Delivery map[string]interface{} `json:"Delivery,omitempty"`
	Index *int32 `json:"Index,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Media []map[string]interface{} `json:"Media,omitempty"`
	ParticipantSid string `json:"ParticipantSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
