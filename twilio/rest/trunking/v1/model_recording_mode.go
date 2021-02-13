/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RecordingMode the model 'RecordingMode'
type RecordingMode string

// List of recording_mode
const (
	RECORDINGMODE_DO_NOT_RECORD            RecordingMode = "do-not-record"
	RECORDINGMODE_RECORD_FROM_RINGING      RecordingMode = "record-from-ringing"
	RECORDINGMODE_RECORD_FROM_ANSWER       RecordingMode = "record-from-answer"
	RECORDINGMODE_RECORD_FROM_RINGING_DUAL RecordingMode = "record-from-ringing-dual"
	RECORDINGMODE_RECORD_FROM_ANSWER_DUAL  RecordingMode = "record-from-answer-dual"
)
