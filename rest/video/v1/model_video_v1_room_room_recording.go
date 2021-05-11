/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VideoV1RoomRoomRecording struct for VideoV1RoomRoomRecording
type VideoV1RoomRoomRecording struct {
	AccountSid      *string                 `json:"account_sid,omitempty"`
	Codec           *RoomRecordingCodec     `json:"codec,omitempty"`
	ContainerFormat *RoomRecordingFormat    `json:"container_format,omitempty"`
	DateCreated     *time.Time              `json:"date_created,omitempty"`
	Duration        *int32                  `json:"duration,omitempty"`
	GroupingSids    *map[string]interface{} `json:"grouping_sids,omitempty"`
	Links           *map[string]interface{} `json:"links,omitempty"`
	Offset          *int32                  `json:"offset,omitempty"`
	RoomSid         *string                 `json:"room_sid,omitempty"`
	Sid             *string                 `json:"sid,omitempty"`
	Size            *int32                  `json:"size,omitempty"`
	SourceSid       *string                 `json:"source_sid,omitempty"`
	Status          *RoomRecordingStatus    `json:"status,omitempty"`
	TrackName       *string                 `json:"track_name,omitempty"`
	Type            *RoomRecordingType      `json:"type,omitempty"`
	Url             *string                 `json:"url,omitempty"`
}
