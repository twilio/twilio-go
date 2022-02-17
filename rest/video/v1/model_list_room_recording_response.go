/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRoomRecordingResponse struct for ListRoomRecordingResponse
type ListRoomRecordingResponse struct {
	Meta       ListCompositionHookResponseMeta `json:"meta,omitempty"`
	Recordings []VideoV1RoomRecording          `json:"recordings,omitempty"`
}
