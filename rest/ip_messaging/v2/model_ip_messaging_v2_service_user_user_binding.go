/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// IpMessagingV2ServiceUserUserBinding struct for IpMessagingV2ServiceUserUserBinding
type IpMessagingV2ServiceUserUserBinding struct {
	AccountSid    *string                 `json:"account_sid,omitempty"`
	BindingType   *UserBindingBindingType `json:"binding_type,omitempty"`
	CredentialSid *string                 `json:"credential_sid,omitempty"`
	DateCreated   *time.Time              `json:"date_created,omitempty"`
	DateUpdated   *time.Time              `json:"date_updated,omitempty"`
	Endpoint      *string                 `json:"endpoint,omitempty"`
	Identity      *string                 `json:"identity,omitempty"`
	MessageTypes  *[]string               `json:"message_types,omitempty"`
	ServiceSid    *string                 `json:"service_sid,omitempty"`
	Sid           *string                 `json:"sid,omitempty"`
	Url           *string                 `json:"url,omitempty"`
	UserSid       *string                 `json:"user_sid,omitempty"`
}
