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
// InsightsV1CallMetric struct for InsightsV1CallMetric
type InsightsV1CallMetric struct {
	AccountSid string `json:"account_sid,omitempty"`
	CallSid string `json:"call_sid,omitempty"`
	CarrierEdge map[string]interface{} `json:"carrier_edge,omitempty"`
	ClientEdge map[string]interface{} `json:"client_edge,omitempty"`
	Direction string `json:"direction,omitempty"`
	Edge string `json:"edge,omitempty"`
	SdkEdge map[string]interface{} `json:"sdk_edge,omitempty"`
	SipEdge map[string]interface{} `json:"sip_edge,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
