/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiV2010AccountConferenceParticipant struct for ApiV2010AccountConferenceParticipant
type ApiV2010AccountConferenceParticipant struct {
	AccountSid             string `json:"AccountSid,omitempty"`
	CallSid                string `json:"CallSid,omitempty"`
	CallSidToCoach         string `json:"CallSidToCoach,omitempty"`
	Coaching               bool   `json:"Coaching,omitempty"`
	ConferenceSid          string `json:"ConferenceSid,omitempty"`
	DateCreated            string `json:"DateCreated,omitempty"`
	DateUpdated            string `json:"DateUpdated,omitempty"`
	EndConferenceOnExit    bool   `json:"EndConferenceOnExit,omitempty"`
	Hold                   bool   `json:"Hold,omitempty"`
	Label                  string `json:"Label,omitempty"`
	Muted                  bool   `json:"Muted,omitempty"`
	StartConferenceOnEnter bool   `json:"StartConferenceOnEnter,omitempty"`
	Status                 Status `json:"Status,omitempty"`
	Uri                    string `json:"Uri,omitempty"`
}
