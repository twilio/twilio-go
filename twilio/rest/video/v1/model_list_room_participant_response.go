/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListRoomParticipantResponse struct for ListRoomParticipantResponse
type ListRoomParticipantResponse struct {
	Meta ListCompositionHookResponseMeta `json:"Meta,omitempty"`
	Participants []VideoV1RoomRoomParticipant `json:"Participants,omitempty"`
}
