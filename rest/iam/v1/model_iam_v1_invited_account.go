/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Iam
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

// IamV1InvitedAccount struct for IamV1InvitedAccount
type IamV1InvitedAccount struct {
	// Unique Twilio notification sid
	NotificationSid string `json:"notificationSid,omitempty"`
	// Unique Twilio account sid
	AccountSid string `json:"accountSid,omitempty"`
	// First name of the sender who invited the account
	SenderFirstName string `json:"senderFirstName,omitempty"`
	// Last name of the sender who invited the account
	SenderLastName string `json:"senderLastName,omitempty"`
	// Email of the sender who invited the account
	SenderEmail string `json:"senderEmail,omitempty"`
	// The date and time of the invitation creation
	DateCreated time.Time `json:"dateCreated,omitempty"`
	// Total number of invitations sent
	TotalSends int `json:"totalSends,omitempty"`
	// The date and time of the last invitation sent
	DateLastSent time.Time `json:"dateLastSent,omitempty"`
	// Whether the invitation token has expired
	TokenExpired bool `json:"tokenExpired,omitempty"`
	// Whether the invitation can be resent
	CanResend bool `json:"canResend,omitempty"`
	// Reason why the invitation cannot be resent, if applicable
	CanResendReason string `json:"canResendReason,omitempty"`
}
