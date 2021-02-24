/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListRatePlanResponse struct for ListRatePlanResponse
type ListRatePlanResponse struct {
	Meta      ListCommandResponseMeta `json:"Meta,omitempty"`
	RatePlans []WirelessV1RatePlan    `json:"RatePlans,omitempty"`
}
