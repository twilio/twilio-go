/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RecordingSource the model 'RecordingSource'
type RecordingSource string

// List of recording_source
const (
	RECORDINGSOURCE_DIAL_VERB                      RecordingSource = "DialVerb"
	RECORDINGSOURCE_CONFERENCE                     RecordingSource = "Conference"
	RECORDINGSOURCE_OUTBOUND_API                   RecordingSource = "OutboundAPI"
	RECORDINGSOURCE_TRUNKING                       RecordingSource = "Trunking"
	RECORDINGSOURCE_RECORD_VERB                    RecordingSource = "RecordVerb"
	RECORDINGSOURCE_START_CALL_RECORDING_API       RecordingSource = "StartCallRecordingAPI"
	RECORDINGSOURCE_START_CONFERENCE_RECORDING_API RecordingSource = "StartConferenceRecordingAPI"
)
