/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateServiceRequest struct for UpdateServiceRequest
type UpdateServiceRequest struct {
	ConsumptionReportInterval int32 `json:"ConsumptionReportInterval,omitempty"`
	DefaultChannelCreatorRoleSid string `json:"DefaultChannelCreatorRoleSid,omitempty"`
	DefaultChannelRoleSid string `json:"DefaultChannelRoleSid,omitempty"`
	DefaultServiceRoleSid string `json:"DefaultServiceRoleSid,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	LimitsChannelMembers int32 `json:"LimitsChannelMembers,omitempty"`
	LimitsUserChannels int32 `json:"LimitsUserChannels,omitempty"`
	MediaCompatibilityMessage string `json:"MediaCompatibilityMessage,omitempty"`
	NotificationsAddedToChannelEnabled bool `json:"NotificationsAddedToChannelEnabled,omitempty"`
	NotificationsAddedToChannelSound string `json:"NotificationsAddedToChannelSound,omitempty"`
	NotificationsAddedToChannelTemplate string `json:"NotificationsAddedToChannelTemplate,omitempty"`
	NotificationsInvitedToChannelEnabled bool `json:"NotificationsInvitedToChannelEnabled,omitempty"`
	NotificationsInvitedToChannelSound string `json:"NotificationsInvitedToChannelSound,omitempty"`
	NotificationsInvitedToChannelTemplate string `json:"NotificationsInvitedToChannelTemplate,omitempty"`
	NotificationsLogEnabled bool `json:"NotificationsLogEnabled,omitempty"`
	NotificationsNewMessageBadgeCountEnabled bool `json:"NotificationsNewMessageBadgeCountEnabled,omitempty"`
	NotificationsNewMessageEnabled bool `json:"NotificationsNewMessageEnabled,omitempty"`
	NotificationsNewMessageSound string `json:"NotificationsNewMessageSound,omitempty"`
	NotificationsNewMessageTemplate string `json:"NotificationsNewMessageTemplate,omitempty"`
	NotificationsRemovedFromChannelEnabled bool `json:"NotificationsRemovedFromChannelEnabled,omitempty"`
	NotificationsRemovedFromChannelSound string `json:"NotificationsRemovedFromChannelSound,omitempty"`
	NotificationsRemovedFromChannelTemplate string `json:"NotificationsRemovedFromChannelTemplate,omitempty"`
	PostWebhookRetryCount int32 `json:"PostWebhookRetryCount,omitempty"`
	PostWebhookUrl string `json:"PostWebhookUrl,omitempty"`
	PreWebhookRetryCount int32 `json:"PreWebhookRetryCount,omitempty"`
	PreWebhookUrl string `json:"PreWebhookUrl,omitempty"`
	ReachabilityEnabled bool `json:"ReachabilityEnabled,omitempty"`
	ReadStatusEnabled bool `json:"ReadStatusEnabled,omitempty"`
	TypingIndicatorTimeout int32 `json:"TypingIndicatorTimeout,omitempty"`
	WebhookFilters []string `json:"WebhookFilters,omitempty"`
	WebhookMethod string `json:"WebhookMethod,omitempty"`
}
