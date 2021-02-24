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

// ListAvailableAddOnExtensionResponse struct for ListAvailableAddOnExtensionResponse
type ListAvailableAddOnExtensionResponse struct {
	Extensions []PreviewMarketplaceAvailableAddOnAvailableAddOnExtension `json:"Extensions,omitempty"`
	Meta       ListDayResponseMeta                                       `json:"Meta,omitempty"`
}
