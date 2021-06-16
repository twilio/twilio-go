/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack struct for VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack
type VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack struct {
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the track is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The track type
	Kind *string `json:"kind,omitempty"`
	// The track name
	Name *string `json:"name,omitempty"`
	// The SID of the participant that subscribes to the track
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the participant that publishes the track
	PublisherSid *string `json:"publisher_sid,omitempty"`
	// The SID of the room where the track is published
	RoomSid *string `json:"room_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
