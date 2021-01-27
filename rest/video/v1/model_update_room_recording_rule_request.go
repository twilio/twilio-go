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
// UpdateRoomRecordingRuleRequest struct for UpdateRoomRecordingRuleRequest
type UpdateRoomRecordingRuleRequest struct {
	// A JSON-encoded array of recording rules.
	Rules map[string]interface{} `json:"Rules,omitempty"`
}
