/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateCallRecordingRequest struct for UpdateCallRecordingRequest
type UpdateCallRecordingRequest struct {
	// Whether to record during a pause. Can be: `skip` or `silence` and the default is `silence`. `skip` does not record during the pause period, while `silence` will replace the actual audio of the call with silence during the pause period. This parameter only applies when setting `status` is set to `paused`.
	PauseBehavior string `json:"PauseBehavior,omitempty"`
	// The new status of the recording. Can be: `stopped`, `paused`, `in-progress`.
	Status string `json:"Status"`
}
