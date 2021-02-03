/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// RoomStatus the model 'RoomStatus'
type RoomStatus string

// List of room_status
const (
	IN_PROGRESS RoomStatus = "in-progress"
	COMPLETED RoomStatus = "completed"
	FAILED RoomStatus = "failed"
)
