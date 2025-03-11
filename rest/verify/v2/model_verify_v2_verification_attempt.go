/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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

// VerifyV2VerificationAttempt struct for VerifyV2VerificationAttempt
type VerifyV2VerificationAttempt struct {
	// The SID that uniquely identifies the verification attempt resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Verification resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/verify/api/service) used to generate the attempt.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Verification](https://www.twilio.com/docs/verify/api/verification) that generated the attempt.
	VerificationSid *string `json:"verification_sid,omitempty"`
	// The date that this Attempt was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this Attempt was updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated      *time.Time `json:"date_updated,omitempty"`
	ConversionStatus *string    `json:"conversion_status,omitempty"`
	Channel          *string    `json:"channel,omitempty"`
	// An object containing the charge for this verification attempt related to the channel costs and the currency used. The costs related to the succeeded verifications are not included. May not be immediately available. More information on pricing is available [here](https://www.twilio.com/en-us/verify/pricing).
	Price *map[string]interface{} `json:"price,omitempty"`
	// An object containing the channel specific information for an attempt.
	ChannelData *map[string]interface{} `json:"channel_data,omitempty"`
	Url         *string                 `json:"url,omitempty"`
}
