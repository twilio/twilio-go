/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// IpMessagingV1Credential struct for IpMessagingV1Credential
type IpMessagingV1Credential struct {
	AccountSid   *string                `json:"AccountSid,omitempty"`
	DateCreated  *time.Time             `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time             `json:"DateUpdated,omitempty"`
	FriendlyName *string                `json:"FriendlyName,omitempty"`
	Sandbox      *string                `json:"Sandbox,omitempty"`
	Sid          *string                `json:"Sid,omitempty"`
	Type         *CredentialPushService `json:"Type,omitempty"`
	Url          *string                `json:"Url,omitempty"`
}
