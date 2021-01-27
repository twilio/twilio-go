/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateRoomParticipantSubscribeRuleRequest struct for UpdateRoomParticipantSubscribeRuleRequest
type UpdateRoomParticipantSubscribeRuleRequest struct {
	// A JSON-encoded array of subscribe rules. See the [Specifying Subscribe Rules](https://www.twilio.com/docs/video/api/track-subscriptions#specifying-sr) section for further information.
	Rules map[string]interface{} `json:"Rules,omitempty"`
}
