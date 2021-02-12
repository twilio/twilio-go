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
// UpdateSessionRequest struct for UpdateSessionRequest
type UpdateSessionRequest struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
	DateExpiry time.Time `json:"DateExpiry,omitempty"`
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict bool `json:"FailOnParticipantConflict,omitempty"`
	// The new status of the resource. Can be: `in-progress` to re-open a session or `closed` to close a session.
	Status string `json:"Status,omitempty"`
	// The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
	Ttl int32 `json:"Ttl,omitempty"`
}
