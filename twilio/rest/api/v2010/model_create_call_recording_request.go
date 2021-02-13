/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateCallRecordingRequest struct for CreateCallRecordingRequest
type CreateCallRecordingRequest struct {
	// The number of channels used in the recording. Can be: `mono` or `dual` and the default is `mono`. `mono` records all parties of the call into one channel. `dual` records each party of a 2-party call into separate channels.
	RecordingChannels string `json:"RecordingChannels,omitempty"`
	// The URL we should call using the `recording_status_callback_method` on each recording event specified in  `recording_status_callback_event`. For more information, see [RecordingStatusCallback parameters](https://www.twilio.com/docs/voice/api/recording#recordingstatuscallback).
	RecordingStatusCallback string `json:"RecordingStatusCallback,omitempty"`
	// The recording status events on which we should call the `recording_status_callback` URL. Can be: `in-progress`, `completed` and `absent` and the default is `completed`. Separate multiple event values with a space.
	RecordingStatusCallbackEvent []string `json:"RecordingStatusCallbackEvent,omitempty"`
	// The HTTP method we should use to call `recording_status_callback`. Can be: `GET` or `POST` and the default is `POST`.
	RecordingStatusCallbackMethod string `json:"RecordingStatusCallbackMethod,omitempty"`
	// The audio track to record for the call. Can be: `inbound`, `outbound` or `both`. The default is `both`. `inbound` records the audio that is received by Twilio. `outbound` records the audio that is generated from Twilio. `both` records the audio that is received and generated by Twilio.
	RecordingTrack string `json:"RecordingTrack,omitempty"`
	// Whether to trim any leading and trailing silence in the recording. Can be: `trim-silence` or `do-not-trim` and the default is `do-not-trim`. `trim-silence` trims the silence from the beginning and end of the recording and `do-not-trim` does not.
	Trim string `json:"Trim,omitempty"`
}
