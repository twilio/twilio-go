/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// IpMessagingV1Service struct for IpMessagingV1Service
type IpMessagingV1Service struct {
	Sid                          *string                 `json:"sid,omitempty"`
	AccountSid                   *string                 `json:"account_sid,omitempty"`
	FriendlyName                 *string                 `json:"friendly_name,omitempty"`
	DateCreated                  *time.Time              `json:"date_created,omitempty"`
	DateUpdated                  *time.Time              `json:"date_updated,omitempty"`
	DefaultServiceRoleSid        *string                 `json:"default_service_role_sid,omitempty"`
	DefaultChannelRoleSid        *string                 `json:"default_channel_role_sid,omitempty"`
	DefaultChannelCreatorRoleSid *string                 `json:"default_channel_creator_role_sid,omitempty"`
	ReadStatusEnabled            *bool                   `json:"read_status_enabled,omitempty"`
	ReachabilityEnabled          *bool                   `json:"reachability_enabled,omitempty"`
	TypingIndicatorTimeout       int                     `json:"typing_indicator_timeout,omitempty"`
	ConsumptionReportInterval    int                     `json:"consumption_report_interval,omitempty"`
	Limits                       *map[string]interface{} `json:"limits,omitempty"`
	Webhooks                     *map[string]interface{} `json:"webhooks,omitempty"`
	PreWebhookUrl                *string                 `json:"pre_webhook_url,omitempty"`
	PostWebhookUrl               *string                 `json:"post_webhook_url,omitempty"`
	WebhookMethod                *string                 `json:"webhook_method,omitempty"`
	WebhookFilters               *[]string               `json:"webhook_filters,omitempty"`
	Notifications                *map[string]interface{} `json:"notifications,omitempty"`
	Url                          *string                 `json:"url,omitempty"`
	Links                        *map[string]interface{} `json:"links,omitempty"`
}
