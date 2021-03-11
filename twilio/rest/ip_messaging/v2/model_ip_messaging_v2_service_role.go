/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// IpMessagingV2ServiceRole struct for IpMessagingV2ServiceRole
type IpMessagingV2ServiceRole struct {
	AccountSid   *string    `json:"AccountSid,omitempty"`
	DateCreated  *time.Time `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time `json:"DateUpdated,omitempty"`
	FriendlyName *string    `json:"FriendlyName,omitempty"`
	Permissions  *[]string  `json:"Permissions,omitempty"`
	ServiceSid   *string    `json:"ServiceSid,omitempty"`
	Sid          *string    `json:"Sid,omitempty"`
	Type         *string    `json:"Type,omitempty"`
	Url          *string    `json:"Url,omitempty"`
}
