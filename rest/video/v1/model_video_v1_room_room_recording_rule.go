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

import (
	"time"
)

// VideoV1RoomRoomRecordingRule struct for VideoV1RoomRoomRecordingRule
type VideoV1RoomRoomRecordingRule struct {
	DateCreated *time.Time                `json:"date_created,omitempty"`
	DateUpdated *time.Time                `json:"date_updated,omitempty"`
	RoomSid     *string                   `json:"room_sid,omitempty"`
	Rules       *[]map[string]interface{} `json:"rules,omitempty"`
}
