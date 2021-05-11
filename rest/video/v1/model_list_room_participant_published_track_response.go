/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRoomParticipantPublishedTrackResponse struct for ListRoomParticipantPublishedTrackResponse
type ListRoomParticipantPublishedTrackResponse struct {
	Meta            ListCompositionHookResponseMeta                           `json:"meta,omitempty"`
	PublishedTracks []VideoV1RoomRoomParticipantRoomParticipantPublishedTrack `json:"published_tracks,omitempty"`
}
