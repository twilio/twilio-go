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
// InsightsV1CallMetricReadResponse struct for InsightsV1CallMetricReadResponse
type InsightsV1CallMetricReadResponse struct {
	Meta InsightsV1VideoRoomSummaryReadResponseMeta `json:"Meta,omitempty"`
	Metrics []InsightsV1CallMetric `json:"Metrics,omitempty"`
}
