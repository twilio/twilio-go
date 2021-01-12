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
// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	Meta InlineResponse200Meta `json:"meta,omitempty"`
	Webhooks []ConversationsV1ConversationConversationScopedWebhook `json:"webhooks,omitempty"`
}
