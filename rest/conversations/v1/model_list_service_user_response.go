/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListServiceUserResponse struct for ListServiceUserResponse
type ListServiceUserResponse struct {
	Meta ListConversationResponseMeta `json:"Meta,omitempty"`
	Users []ConversationsV1ServiceServiceUser `json:"Users,omitempty"`
}