/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InsightsV1VideoRoomSummaryReadResponse struct for InsightsV1VideoRoomSummaryReadResponse
type InsightsV1VideoRoomSummaryReadResponse struct {
	Meta InsightsV1VideoRoomSummaryReadResponseMeta `json:"Meta,omitempty"`
	Rooms []InsightsV1VideoRoomSummary `json:"Rooms,omitempty"`
}
