/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListVideoParticipantSummaryResponse struct for ListVideoParticipantSummaryResponse
type ListVideoParticipantSummaryResponse struct {
	Meta         ListVideoRoomSummaryResponseMeta    `json:"meta,omitempty"`
	Participants []InsightsV1VideoParticipantSummary `json:"participants,omitempty"`
}
