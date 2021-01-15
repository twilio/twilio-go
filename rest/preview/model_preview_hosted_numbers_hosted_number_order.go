/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// PreviewHostedNumbersHostedNumberOrder struct for PreviewHostedNumbersHostedNumberOrder
type PreviewHostedNumbersHostedNumberOrder struct {
	AccountSid string `json:"account_sid,omitempty"`
	AddressSid string `json:"address_sid,omitempty"`
	CallDelay int32 `json:"call_delay,omitempty"`
	Capabilities map[string]interface{} `json:"capabilities,omitempty"`
	CcEmails []string `json:"cc_emails,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Email string `json:"email,omitempty"`
	Extension string `json:"extension,omitempty"`
	FailureReason string `json:"failure_reason,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	IncomingPhoneNumberSid string `json:"incoming_phone_number_sid,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Sid string `json:"sid,omitempty"`
	SigningDocumentSid string `json:"signing_document_sid,omitempty"`
	Status string `json:"status,omitempty"`
	UniqueName string `json:"unique_name,omitempty"`
	Url string `json:"url,omitempty"`
	VerificationAttempts int32 `json:"verification_attempts,omitempty"`
	VerificationCallSids []string `json:"verification_call_sids,omitempty"`
	VerificationCode string `json:"verification_code,omitempty"`
	VerificationDocumentSid string `json:"verification_document_sid,omitempty"`
	VerificationType string `json:"verification_type,omitempty"`
}
