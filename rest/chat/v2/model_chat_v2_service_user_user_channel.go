/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ChatV2ServiceUserUserChannel struct for ChatV2ServiceUserUserChannel
type ChatV2ServiceUserUserChannel struct {
	AccountSid               *string                       `json:"account_sid,omitempty"`
	ChannelSid               *string                       `json:"channel_sid,omitempty"`
	LastConsumedMessageIndex *int32                        `json:"last_consumed_message_index,omitempty"`
	Links                    *map[string]interface{}       `json:"links,omitempty"`
	MemberSid                *string                       `json:"member_sid,omitempty"`
	NotificationLevel        *UserChannelNotificationLevel `json:"notification_level,omitempty"`
	ServiceSid               *string                       `json:"service_sid,omitempty"`
	Status                   *UserChannelChannelStatus     `json:"status,omitempty"`
	UnreadMessagesCount      *int32                        `json:"unread_messages_count,omitempty"`
	Url                      *string                       `json:"url,omitempty"`
	UserSid                  *string                       `json:"user_sid,omitempty"`
}
