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
import (
	"time"
)
// VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack struct for VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack
type VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack struct {
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Enabled bool `json:"Enabled,omitempty"`
	Kind string `json:"Kind,omitempty"`
	Name string `json:"Name,omitempty"`
	ParticipantSid string `json:"ParticipantSid,omitempty"`
	PublisherSid string `json:"PublisherSid,omitempty"`
	RoomSid string `json:"RoomSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
