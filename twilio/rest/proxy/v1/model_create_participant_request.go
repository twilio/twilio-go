/*
 * Twilio - Proxy
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateParticipantRequest struct for CreateParticipantRequest
type CreateParticipantRequest struct {
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Participant create request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict bool `json:"FailOnParticipantConflict,omitempty"`
	// The string that you assigned to describe the participant. This value must be 255 characters or fewer. **This value should not have PII.**
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The phone number of the Participant.
	Identifier string `json:"Identifier"`
	// The proxy phone number to use for the Participant. If not specified, Proxy will select a number from the pool.
	ProxyIdentifier string `json:"ProxyIdentifier,omitempty"`
	// The SID of the Proxy Identifier to assign to the Participant.
	ProxyIdentifierSid string `json:"ProxyIdentifierSid,omitempty"`
}
