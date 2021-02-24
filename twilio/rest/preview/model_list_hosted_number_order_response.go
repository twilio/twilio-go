/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListHostedNumberOrderResponse struct for ListHostedNumberOrderResponse
type ListHostedNumberOrderResponse struct {
	Items []PreviewHostedNumbersHostedNumberOrder `json:"Items,omitempty"`
	Meta  ListDayResponseMeta                     `json:"Meta,omitempty"`
}
