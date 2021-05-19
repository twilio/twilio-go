/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRoomParticipantResponse struct for ListRoomParticipantResponse
type ListRoomParticipantResponse struct {
	Meta         ListCompositionHookResponseMeta `json:"meta,omitempty"`
	Participants []VideoV1RoomRoomParticipant    `json:"participants,omitempty"`
}
