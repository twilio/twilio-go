/*
    * Twilio - Verify
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
// VerifyV2VerificationAttempt struct for VerifyV2VerificationAttempt
type VerifyV2VerificationAttempt struct {
	// Account Sid
	AccountSid *string `json:"AccountSid,omitempty"`
	// Channel used for the attempt
	Channel *string `json:"Channel,omitempty"`
	// Object with the channel information for an attempt
	ChannelData *map[string]interface{} `json:"ChannelData,omitempty"`
	// Status of a conversion
	ConversionStatus *string `json:"ConversionStatus,omitempty"`
	// The date this Attempt was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The date this Attempt was updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// A string that uniquely identifies this Verification Attempt
	Sid *string `json:"Sid,omitempty"`
	Url *string `json:"Url,omitempty"`
}
