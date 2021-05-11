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

// RoomRecordingStatus the model 'RoomRecordingStatus'
type RoomRecordingStatus string

// List of room_recording_status
const (
	ROOMRECORDINGSTATUS_PROCESSING RoomRecordingStatus = "processing"
	ROOMRECORDINGSTATUS_COMPLETED  RoomRecordingStatus = "completed"
	ROOMRECORDINGSTATUS_DELETED    RoomRecordingStatus = "deleted"
	ROOMRECORDINGSTATUS_FAILED     RoomRecordingStatus = "failed"
)
