/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListVideoRoomSummaryResponse struct for ListVideoRoomSummaryResponse
type ListVideoRoomSummaryResponse struct {
	Meta  ListVideoRoomSummaryResponseMeta `json:"meta,omitempty"`
	Rooms []InsightsV1VideoRoomSummary     `json:"rooms,omitempty"`
}
