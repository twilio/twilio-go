/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.12.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// TrunkingV1Trunk struct for TrunkingV1Trunk
type TrunkingV1Trunk struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The types of authentication mapped to the domain
	AuthType *string `json:"AuthType,omitempty"`
	// Reserved
	AuthTypeSet *[]string `json:"AuthTypeSet,omitempty"`
	// Whether Caller ID Name (CNAM) lookup is enabled for the trunk
	CnamLookupEnabled *bool `json:"CnamLookupEnabled,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The HTTP method we use to call the disaster_recovery_url
	DisasterRecoveryMethod *string `json:"DisasterRecoveryMethod,omitempty"`
	// The HTTP URL that we call if an error occurs while sending SIP traffic towards your configured Origination URL
	DisasterRecoveryUrl *string `json:"DisasterRecoveryUrl,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic
	DomainName *string `json:"DomainName,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The recording settings for the trunk
	Recording *map[string]interface{} `json:"Recording,omitempty"`
	// Whether Secure Trunking is enabled for the trunk
	Secure *bool `json:"Secure,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The call transfer settings for the trunk
	TransferMode *string `json:"TransferMode,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"Url,omitempty"`
}
