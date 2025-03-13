/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)

// ChatV2ChannelWebhook struct for ChatV2ChannelWebhook
type ChatV2ChannelWebhook struct {
	// The unique string that we created to identify the Channel Webhook resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Channel Webhook resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Channel Webhook resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Channel](https://www.twilio.com/docs/chat/channels) the Channel Webhook resource belongs to.
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The type of webhook. Can be: `webhook`, `studio`, or `trigger`.
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Channel Webhook resource.
	Url *string `json:"url,omitempty"`
	// The JSON string that describes how the channel webhook is configured. The configuration object contains the `url`, `method`, `filters`, and `retry_count` values that are configured by the create and update actions.
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}
