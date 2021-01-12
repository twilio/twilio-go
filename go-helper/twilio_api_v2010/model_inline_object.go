/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// One or more issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, `unsolicited-call`, or `one-way-audio`.
	Issue []string `json:"Issue,omitempty"`
	// The call quality expressed as an integer from `1` to `5` where `1` represents very poor call quality and `5` represents a perfect call.
	QualityScore int32 `json:"QualityScore"`
}
