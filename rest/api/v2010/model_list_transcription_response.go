/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTranscriptionResponse struct for ListTranscriptionResponse
type ListTranscriptionResponse struct {
	End             int32                          `json:"end,omitempty"`
	FirstPageUri    string                         `json:"first_page_uri,omitempty"`
	NextPageUri     string                         `json:"next_page_uri,omitempty"`
	Page            int32                          `json:"page,omitempty"`
	PageSize        int32                          `json:"page_size,omitempty"`
	PreviousPageUri string                         `json:"previous_page_uri,omitempty"`
	Start           int32                          `json:"start,omitempty"`
	Transcriptions  []ApiV2010AccountTranscription `json:"transcriptions,omitempty"`
	Uri             string                         `json:"uri,omitempty"`
}
