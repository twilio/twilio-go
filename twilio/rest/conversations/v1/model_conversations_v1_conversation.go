/*
 * Twilio - Conversations
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

// ConversationsV1Conversation struct for ConversationsV1Conversation
type ConversationsV1Conversation struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"AccountSid,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string `json:"Attributes,omitempty"`
	// The unique ID of the Conversation Service this conversation belongs to.
	ChatServiceSid *string `json:"ChatServiceSid,omitempty"`
	// The date that this resource was created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The human-readable name of this conversation.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Absolute URLs to access the participants, messages and webhooks of this conversation.
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The unique ID of the Messaging Service this conversation belongs to.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// A 34 character string that uniquely identifies this resource.
	Sid *string `json:"Sid,omitempty"`
	// Current state of this conversation.
	State *string `json:"State,omitempty"`
	// Timer date values for this conversation.
	Timers *map[string]interface{} `json:"Timers,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// An absolute URL for this conversation.
	Url *string `json:"Url,omitempty"`
}
