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
// CreateServiceConversationScopedWebhookRequest struct for CreateServiceConversationScopedWebhookRequest
type CreateServiceConversationScopedWebhookRequest struct {
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters []string `json:"ConfigurationFilters,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid string `json:"ConfigurationFlowSid,omitempty"`
	// The HTTP method to be used when sending a webhook request.
	ConfigurationMethod string `json:"ConfigurationMethod,omitempty"`
	// The message index for which and it's successors the webhook will be replayed. Not set by default
	ConfigurationReplayAfter int32 `json:"ConfigurationReplayAfter,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers []string `json:"ConfigurationTriggers,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl string `json:"ConfigurationUrl,omitempty"`
	// The target of this webhook: `webhook`, `studio`, `trigger`
	Target string `json:"Target"`
}
