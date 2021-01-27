/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ChatV2ServiceUserUserChannel struct for ChatV2ServiceUserUserChannel
type ChatV2ServiceUserUserChannel struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ChannelSid string `json:"ChannelSid,omitempty"`
	LastConsumedMessageIndex *int32 `json:"LastConsumedMessageIndex,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	MemberSid string `json:"MemberSid,omitempty"`
	NotificationLevel string `json:"NotificationLevel,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Status string `json:"Status,omitempty"`
	UnreadMessagesCount *int32 `json:"UnreadMessagesCount,omitempty"`
	Url string `json:"Url,omitempty"`
	UserSid string `json:"UserSid,omitempty"`
}
