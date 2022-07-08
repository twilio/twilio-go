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

// IpMessagingV1Invite struct for IpMessagingV1Invite
type IpMessagingV1Invite struct {
	AccountSid  *string    `json:"account_sid,omitempty"`
	ChannelSid  *string    `json:"channel_sid,omitempty"`
	CreatedBy   *string    `json:"created_by,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	Identity    *string    `json:"identity,omitempty"`
	RoleSid     *string    `json:"role_sid,omitempty"`
	ServiceSid  *string    `json:"service_sid,omitempty"`
	Sid         *string    `json:"sid,omitempty"`
	Url         *string    `json:"url,omitempty"`
}
