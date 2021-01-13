/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountRecordingRecordingTranscriptionReadResponse struct for ApiV2010AccountRecordingRecordingTranscriptionReadResponse
type ApiV2010AccountRecordingRecordingTranscriptionReadResponse struct {
	End int32 `json:"end,omitempty"`
	FirstPageUri string `json:"first_page_uri,omitempty"`
	NextPageUri string `json:"next_page_uri,omitempty"`
	Page int32 `json:"page,omitempty"`
	PageSize int32 `json:"page_size,omitempty"`
	PreviousPageUri string `json:"previous_page_uri,omitempty"`
	Start int32 `json:"start,omitempty"`
	Transcriptions []ApiV2010AccountRecordingRecordingTranscription `json:"transcriptions,omitempty"`
	Uri string `json:"uri,omitempty"`
}
