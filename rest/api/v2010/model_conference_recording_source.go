/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ConferenceRecordingSource the model 'ConferenceRecordingSource'
type ConferenceRecordingSource string

// List of conference_recording_source
const (
	CONFERENCERECORDINGSOURCE_DIAL_VERB                      ConferenceRecordingSource = "DialVerb"
	CONFERENCERECORDINGSOURCE_CONFERENCE                     ConferenceRecordingSource = "Conference"
	CONFERENCERECORDINGSOURCE_OUTBOUND_API                   ConferenceRecordingSource = "OutboundAPI"
	CONFERENCERECORDINGSOURCE_TRUNKING                       ConferenceRecordingSource = "Trunking"
	CONFERENCERECORDINGSOURCE_RECORD_VERB                    ConferenceRecordingSource = "RecordVerb"
	CONFERENCERECORDINGSOURCE_START_CALL_RECORDING_API       ConferenceRecordingSource = "StartCallRecordingAPI"
	CONFERENCERECORDINGSOURCE_START_CONFERENCE_RECORDING_API ConferenceRecordingSource = "StartConferenceRecordingAPI"
)
