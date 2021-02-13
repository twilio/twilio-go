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

import (
	"time"
)

// IpMessagingV1ServiceUser struct for IpMessagingV1ServiceUser
type IpMessagingV1ServiceUser struct {
	AccountSid          string                 `json:"AccountSid,omitempty"`
	Attributes          string                 `json:"Attributes,omitempty"`
	DateCreated         time.Time              `json:"DateCreated,omitempty"`
	DateUpdated         time.Time              `json:"DateUpdated,omitempty"`
	FriendlyName        string                 `json:"FriendlyName,omitempty"`
	Identity            string                 `json:"Identity,omitempty"`
	IsNotifiable        bool                   `json:"IsNotifiable,omitempty"`
	IsOnline            bool                   `json:"IsOnline,omitempty"`
	JoinedChannelsCount *int32                 `json:"JoinedChannelsCount,omitempty"`
	Links               map[string]interface{} `json:"Links,omitempty"`
	RoleSid             string                 `json:"RoleSid,omitempty"`
	ServiceSid          string                 `json:"ServiceSid,omitempty"`
	Sid                 string                 `json:"Sid,omitempty"`
	Url                 string                 `json:"Url,omitempty"`
}
