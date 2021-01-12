/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject13 struct for InlineObject13
type InlineObject13 struct {
	// The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Amount string `json:"Amount,omitempty"`
	// The 4-10 character string being verified.
	Code string `json:"Code"`
	// The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
	Payee string `json:"Payee,omitempty"`
	// The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the `verification_sid` must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
	To string `json:"To,omitempty"`
	// A SID that uniquely identifies the Verification Check. Either this parameter or the `to` phone number/[email](https://www.twilio.com/docs/verify/email) must be specified.
	VerificationSid string `json:"VerificationSid,omitempty"`
}
