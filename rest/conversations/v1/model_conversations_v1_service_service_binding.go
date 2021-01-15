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
// ConversationsV1ServiceServiceBinding struct for ConversationsV1ServiceServiceBinding
type ConversationsV1ServiceServiceBinding struct {
	AccountSid string `json:"account_sid,omitempty"`
	BindingType string `json:"binding_type,omitempty"`
	ChatServiceSid string `json:"chat_service_sid,omitempty"`
	CredentialSid string `json:"credential_sid,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
	Identity string `json:"identity,omitempty"`
	MessageTypes []string `json:"message_types,omitempty"`
	Sid string `json:"sid,omitempty"`
	Url string `json:"url,omitempty"`
}
