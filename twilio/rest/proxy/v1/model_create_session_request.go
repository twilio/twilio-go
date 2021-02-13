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

import (
	"time"
)

// CreateSessionRequest struct for CreateSessionRequest
type CreateSessionRequest struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
	DateExpiry time.Time `json:"DateExpiry,omitempty"`
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict bool `json:"FailOnParticipantConflict,omitempty"`
	// The Mode of the Session. Can be: `message-only`, `voice-only`, or `voice-and-message` and the default value is `voice-and-message`.
	Mode string `json:"Mode,omitempty"`
	// The Participant objects to include in the new session.
	Participants []map[string]interface{} `json:"Participants,omitempty"`
	// The initial status of the Session. Can be: `open`, `in-progress`, `closed`, `failed`, or `unknown`. The default is `open` on create.
	Status string `json:"Status,omitempty"`
	// The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
	Ttl int32 `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
	UniqueName string `json:"UniqueName,omitempty"`
}
