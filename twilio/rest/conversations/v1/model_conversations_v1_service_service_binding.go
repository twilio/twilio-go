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
// ConversationsV1ServiceServiceBinding struct for ConversationsV1ServiceServiceBinding
type ConversationsV1ServiceServiceBinding struct {
	AccountSid string `json:"AccountSid,omitempty"`
	BindingType BindingType `json:"BindingType,omitempty"`
	ChatServiceSid string `json:"ChatServiceSid,omitempty"`
	CredentialSid string `json:"CredentialSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Endpoint string `json:"Endpoint,omitempty"`
	Identity string `json:"Identity,omitempty"`
	MessageTypes []string `json:"MessageTypes,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
