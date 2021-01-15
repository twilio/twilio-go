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
// UpdateServiceConfigurationRequest struct for UpdateServiceConfigurationRequest
type UpdateServiceConfigurationRequest struct {
	// The service-level role assigned to users when they are added to the service. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
	DefaultChatServiceRoleSid string `json:"DefaultChatServiceRoleSid,omitempty"`
	// The conversation-level role assigned to a conversation creator when they join a new conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
	DefaultConversationCreatorRoleSid string `json:"DefaultConversationCreatorRoleSid,omitempty"`
	// The conversation-level role assigned to users when they are added to a conversation. See the [Conversation Role](https://www.twilio.com/docs/conversations/api/role-resource) for more info about roles.
	DefaultConversationRoleSid string `json:"DefaultConversationRoleSid,omitempty"`
	// Whether the [Reachability Indicator](https://www.twilio.com/docs/chat/reachability-indicator) is enabled for this Conversations Service. The default is `false`.
	ReachabilityEnabled bool `json:"ReachabilityEnabled,omitempty"`
}
