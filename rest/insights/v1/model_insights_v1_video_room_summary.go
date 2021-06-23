/*
 * Twilio - Insights
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

// InsightsV1VideoRoomSummary struct for InsightsV1VideoRoomSummary
type InsightsV1VideoRoomSummary struct {
	// Account SID associated with this room.
	AccountSid *string `json:"account_sid,omitempty"`
	// Codecs used by participants in the room.
	Codecs *[]string `json:"codecs,omitempty"`
	// Actual number of concurrent participants.
	ConcurrentParticipants *int `json:"concurrent_participants,omitempty"`
	// Creation time of the room.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// How the room was created.
	CreatedMethod *string `json:"created_method,omitempty"`
	// Total room duration from create time to end time.
	DurationSec *int `json:"duration_sec,omitempty"`
	// Edge location of Twilio media servers for the room.
	EdgeLocation *string `json:"edge_location,omitempty"`
	// Reason the room ended.
	EndReason *string `json:"end_reason,omitempty"`
	// End time for the room.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Room subresources.
	Links *map[string]interface{} `json:"links,omitempty"`
	// Maximum number of participants allowed in the room at the same time allowed by the application settings.
	MaxConcurrentParticipants *int `json:"max_concurrent_participants,omitempty"`
	// Max number of total participants allowed by the application settings.
	MaxParticipants *int `json:"max_participants,omitempty"`
	// Region of Twilio media servers for the room.
	MediaRegion *string `json:"media_region,omitempty"`
	// Video Log Analyzer resource state. Will be either `in-progress` or `complete`.
	ProcessingState *string `json:"processing_state,omitempty"`
	// Boolean indicating if recording is enabled for the room.
	RecordingEnabled *bool `json:"recording_enabled,omitempty"`
	// room friendly name.
	RoomName *string `json:"room_name,omitempty"`
	// Unique identifier for the room.
	RoomSid *string `json:"room_sid,omitempty"`
	// Status of the room.
	RoomStatus *string `json:"room_status,omitempty"`
	// Type of room.
	RoomType *string `json:"room_type,omitempty"`
	// Webhook provided for status callbacks.
	StatusCallback *string `json:"status_callback,omitempty"`
	// HTTP method provided for status callback URL.
	StatusCallbackMethod *string `json:"status_callback_method,omitempty"`
	// Combined amount of participant time in the room.
	TotalParticipantDurationSec *int `json:"total_participant_duration_sec,omitempty"`
	// Combined amount of recorded seconds for participants in the room.
	TotalRecordingDurationSec *int `json:"total_recording_duration_sec,omitempty"`
	// Unique number of participant identities.
	UniqueParticipantIdentities *int `json:"unique_participant_identities,omitempty"`
	// Number of participants. May include duplicate identities for participants who left and rejoined.
	UniqueParticipants *int `json:"unique_participants,omitempty"`
	// URL for the room resource.
	Url *string `json:"url,omitempty"`
}
