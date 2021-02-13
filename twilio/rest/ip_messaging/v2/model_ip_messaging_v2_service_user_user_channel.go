/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// IpMessagingV2ServiceUserUserChannel struct for IpMessagingV2ServiceUserUserChannel
type IpMessagingV2ServiceUserUserChannel struct {
	AccountSid               string                 `json:"AccountSid,omitempty"`
	ChannelSid               string                 `json:"ChannelSid,omitempty"`
	LastConsumedMessageIndex *int32                 `json:"LastConsumedMessageIndex,omitempty"`
	Links                    map[string]interface{} `json:"Links,omitempty"`
	MemberSid                string                 `json:"MemberSid,omitempty"`
	NotificationLevel        NotificationLevel      `json:"NotificationLevel,omitempty"`
	ServiceSid               string                 `json:"ServiceSid,omitempty"`
	Status                   ChannelStatus          `json:"Status,omitempty"`
	UnreadMessagesCount      *int32                 `json:"UnreadMessagesCount,omitempty"`
	Url                      string                 `json:"Url,omitempty"`
	UserSid                  string                 `json:"UserSid,omitempty"`
}
