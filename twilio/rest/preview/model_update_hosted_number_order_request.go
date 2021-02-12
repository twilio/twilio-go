/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateHostedNumberOrderRequest struct for UpdateHostedNumberOrderRequest
type UpdateHostedNumberOrderRequest struct {
	// The number of seconds, between 0 and 60, to delay before initiating the verification call. Defaults to 0.
	CallDelay int32 `json:"CallDelay,omitempty"`
	// Optional. A list of emails that LOA document for this HostedNumberOrder will be carbon copied to.
	CcEmails []string `json:"CcEmails,omitempty"`
	// Email of the owner of this phone number that is being hosted.
	Email string `json:"Email,omitempty"`
	// Digits to dial after connecting the verification call.
	Extension string `json:"Extension,omitempty"`
	// A 64 character string that is a human readable text that describes this resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// User can only post to `pending-verification` status to transition the HostedNumberOrder to initiate a verification call or verification of ownership with a copy of a phone bill.
	Status string `json:"Status,omitempty"`
	// Provides a unique and addressable name to be assigned to this HostedNumberOrder, assigned by the developer, to be optionally used in addition to SID.
	UniqueName string `json:"UniqueName,omitempty"`
	// A verification code that is given to the user via a phone call to the phone number that is being hosted.
	VerificationCode string `json:"VerificationCode,omitempty"`
	// Optional. The unique sid identifier of the Identity Document that represents the document for verifying ownership of the number to be hosted. Required when VerificationType is phone-bill.
	VerificationDocumentSid string `json:"VerificationDocumentSid,omitempty"`
	// Optional. The method used for verifying ownership of the number to be hosted. One of phone-call (default) or phone-bill.
	VerificationType string `json:"VerificationType,omitempty"`
}
