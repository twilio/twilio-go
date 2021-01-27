/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListDependentHostedNumberOrderResponse struct for ListDependentHostedNumberOrderResponse
type ListDependentHostedNumberOrderResponse struct {
	Items []PreviewHostedNumbersAuthorizationDocumentDependentHostedNumberOrder `json:"Items,omitempty"`
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
}
