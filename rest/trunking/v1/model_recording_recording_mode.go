/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// RecordingRecordingMode the model 'RecordingRecordingMode'
type RecordingRecordingMode string

// List of recording_recording_mode
const (
	RECORDINGRECORDINGMODE_DO_NOT_RECORD            RecordingRecordingMode = "do-not-record"
	RECORDINGRECORDINGMODE_RECORD_FROM_RINGING      RecordingRecordingMode = "record-from-ringing"
	RECORDINGRECORDINGMODE_RECORD_FROM_ANSWER       RecordingRecordingMode = "record-from-answer"
	RECORDINGRECORDINGMODE_RECORD_FROM_RINGING_DUAL RecordingRecordingMode = "record-from-ringing-dual"
	RECORDINGRECORDINGMODE_RECORD_FROM_ANSWER_DUAL  RecordingRecordingMode = "record-from-answer-dual"
)
