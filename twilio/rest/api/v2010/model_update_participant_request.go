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
// UpdateParticipantRequest struct for UpdateParticipantRequest
type UpdateParticipantRequest struct {
	// The HTTP method we should use to call `announce_url`. Can be: `GET` or `POST` and defaults to `POST`.
	AnnounceMethod string `json:"AnnounceMethod,omitempty"`
	// The URL we call using the `announce_method` for an announcement to the participant. The URL must return an MP3 file, a WAV file, or a TwiML document that contains `<Play>` or `<Say>` commands.
	AnnounceUrl string `json:"AnnounceUrl,omitempty"`
	// Whether to play a notification beep to the conference when the participant exits. Can be: `true` or `false`.
	BeepOnExit bool `json:"BeepOnExit,omitempty"`
	// The SID of the participant who is being `coached`. The participant being coached is the only participant who can hear the participant who is `coaching`.
	CallSidToCoach string `json:"CallSidToCoach,omitempty"`
	// Whether the participant is coaching another call. Can be: `true` or `false`. If not present, defaults to `false` unless `call_sid_to_coach` is defined. If `true`, `call_sid_to_coach` must be defined.
	Coaching bool `json:"Coaching,omitempty"`
	// Whether to end the conference when the participant leaves. Can be: `true` or `false` and defaults to `false`.
	EndConferenceOnExit bool `json:"EndConferenceOnExit,omitempty"`
	// Whether the participant should be on hold. Can be: `true` or `false`. `true` puts the participant on hold, and `false` lets them rejoin the conference.
	Hold bool `json:"Hold,omitempty"`
	// The HTTP method we should use to call `hold_url`. Can be: `GET` or `POST` and the default is `GET`.
	HoldMethod string `json:"HoldMethod,omitempty"`
	// The URL we call using the `hold_method` for  music that plays when the participant is on hold. The URL may return an MP3 file, a WAV file, or a TwiML document that contains the `<Play>`, `<Say>` or `<Redirect>` commands.
	HoldUrl string `json:"HoldUrl,omitempty"`
	// Whether the participant should be muted. Can be `true` or `false`. `true` will mute the participant, and `false` will un-mute them. Anything value other than `true` or `false` is interpreted as `false`.
	Muted bool `json:"Muted,omitempty"`
	// The HTTP method we should use to call `wait_url`. Can be `GET` or `POST` and the default is `POST`. When using a static audio file, this should be `GET` so that we can cache the file.
	WaitMethod string `json:"WaitMethod,omitempty"`
	// The URL we should call using the `wait_method` for the music to play while participants are waiting for the conference to start. The default value is the URL of our standard hold music. [Learn more about hold music](https://www.twilio.com/labs/twimlets/holdmusic).
	WaitUrl string `json:"WaitUrl,omitempty"`
}
