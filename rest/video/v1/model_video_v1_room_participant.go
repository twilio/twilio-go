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

import (
	"time"
)

// VideoV1RoomParticipant struct for VideoV1RoomParticipant
type VideoV1RoomParticipant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Duration of time in seconds the participant was connected
	Duration *int `json:"duration,omitempty"`
	// The time when the participant disconnected from the room in ISO 8601 format
	EndTime *time.Time `json:"end_time,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// The URLs of related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the participant's room
	RoomSid *string `json:"room_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The time of participant connected to the room in ISO 8601 format
	StartTime *time.Time `json:"start_time,omitempty"`
	// The status of the Participant
	Status *string `json:"status,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
