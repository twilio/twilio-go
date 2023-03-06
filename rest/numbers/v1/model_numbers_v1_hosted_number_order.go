/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
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

// NumbersV1HostedNumberOrder struct for NumbersV1HostedNumberOrder
type NumbersV1HostedNumberOrder struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the HostedNumberOrder resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the AuthorizationDocument resource associated with the hosted number order.
	AuthorizationDocumentSid *string                                                   `json:"authorization_document_sid,omitempty"`
	Capabilities             *NumbersV1AuthorizationDocumentDependentOrderCapabilities `json:"capabilities,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The message that explains why the hosted number order `status` became `action-required`.
	FailureReason *string `json:"failure_reason,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The SID of the IncomingPhoneNumber resource created by this hosted number order.
	IncomingPhoneNumberSid *string `json:"incoming_phone_number_sid,omitempty"`
	// The [E.164](https://www.twilio.com/docs/glossary/what-e164) formatted phone number hosted by the hosted number order.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The unique string that we created to identify the HostedNumberOrder resource.
	Sid    *string `json:"sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// The absolute URL of the HostedNumberOrder resource.
	Url *string `json:"url,omitempty"`
	// The number of attempts made to verify ownership via a call for the hosted phone number.
	VerificationAttempts *int `json:"verification_attempts,omitempty"`
	// The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
	VerificationCallDelay *int `json:"verification_call_delay,omitempty"`
	// The numerical extension to dial when making the ownership verification call.
	VerificationCallExtension *string `json:"verification_call_extension,omitempty"`
	// The Call SIDs that identify the calls placed to verify ownership.
	VerificationCallSids *[]string `json:"verification_call_sids,omitempty"`
	// The digits the user must pass in the ownership verification call.
	VerificationCode *string `json:"verification_code,omitempty"`
	// The SID of the identity document resource that represents the document used to verify ownership of the number to be hosted.
	VerificationDocumentSid *string `json:"verification_document_sid,omitempty"`
	VerificationType        *string `json:"verification_type,omitempty"`
}
