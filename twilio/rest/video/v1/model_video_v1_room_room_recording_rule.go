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
import (
	"time"
)
// VideoV1RoomRoomRecordingRule struct for VideoV1RoomRoomRecordingRule
type VideoV1RoomRoomRecordingRule struct {
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	RoomSid string `json:"RoomSid,omitempty"`
	Rules []map[string]interface{} `json:"Rules,omitempty"`
}
