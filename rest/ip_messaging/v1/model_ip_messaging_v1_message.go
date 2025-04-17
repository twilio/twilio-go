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

// IpMessagingV1Message struct for IpMessagingV1Message
type IpMessagingV1Message struct {
	Sid         *string    `json:"sid,omitempty"`
	AccountSid  *string    `json:"account_sid,omitempty"`
	Attributes  *string    `json:"attributes,omitempty"`
	ServiceSid  *string    `json:"service_sid,omitempty"`
	To          *string    `json:"to,omitempty"`
	ChannelSid  *string    `json:"channel_sid,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	WasEdited   *bool      `json:"was_edited,omitempty"`
	From        *string    `json:"from,omitempty"`
	Body        *string    `json:"body,omitempty"`
	Index       int        `json:"index,omitempty"`
	Url         *string    `json:"url,omitempty"`
}
