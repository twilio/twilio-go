/*
 * Twilio - Conversations
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
// ConversationsV1Role struct for ConversationsV1Role
type ConversationsV1Role struct {
	AccountSid string `json:"account_sid,omitempty"`
	ChatServiceSid string `json:"chat_service_sid,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Sid string `json:"sid,omitempty"`
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
}
