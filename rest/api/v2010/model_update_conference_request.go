/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateConferenceRequest struct for UpdateConferenceRequest
type UpdateConferenceRequest struct {
	// The HTTP method used to call `announce_url`. Can be: `GET` or `POST` and the default is `POST`
	AnnounceMethod string `json:"AnnounceMethod,omitempty"`
	// The URL we should call to announce something into the conference. The URL can return an MP3, a WAV, or a TwiML document with `<Play>` or `<Say>`.
	AnnounceUrl string `json:"AnnounceUrl,omitempty"`
	// The new status of the resource. Can be:  Can be: `init`, `in-progress`, or `completed`. Specifying `completed` will end the conference and hang up all participants
	Status string `json:"Status,omitempty"`
}
