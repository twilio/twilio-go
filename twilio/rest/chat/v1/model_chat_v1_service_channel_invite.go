/*
    * Twilio - Chat
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
// ChatV1ServiceChannelInvite struct for ChatV1ServiceChannelInvite
type ChatV1ServiceChannelInvite struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID of the Channel the new resource belongs to
	ChannelSid *string `json:"ChannelSid,omitempty"`
	// The identity of the User that created the invite
	CreatedBy *string `json:"CreatedBy,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"Identity,omitempty"`
	// The SID of the Role assigned to the member
	RoleSid *string `json:"RoleSid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Invite resource
	Url *string `json:"Url,omitempty"`
}
