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

// RoomStatus the model 'RoomStatus'
type RoomStatus string

// List of room_status
const (
	ROOMSTATUS_IN_PROGRESS RoomStatus = "in-progress"
	ROOMSTATUS_COMPLETED   RoomStatus = "completed"
	ROOMSTATUS_FAILED      RoomStatus = "failed"
)
