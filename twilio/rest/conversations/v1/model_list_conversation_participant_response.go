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

// ListConversationParticipantResponse struct for ListConversationParticipantResponse
type ListConversationParticipantResponse struct {
	Meta         ListConversationResponseMeta                         `json:"Meta,omitempty"`
	Participants []ConversationsV1ConversationConversationParticipant `json:"Participants,omitempty"`
}
