/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// ConversationsV1ServiceServiceUser struct for ConversationsV1ServiceServiceUser
type ConversationsV1ServiceServiceUser struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Attributes string `json:"Attributes,omitempty"`
	ChatServiceSid string `json:"ChatServiceSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Identity string `json:"Identity,omitempty"`
	IsNotifiable bool `json:"IsNotifiable,omitempty"`
	IsOnline bool `json:"IsOnline,omitempty"`
	RoleSid string `json:"RoleSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
