/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListAvailableAddOnResponse struct for ListAvailableAddOnResponse
type ListAvailableAddOnResponse struct {
	AvailableAddOns []PreviewMarketplaceAvailableAddOn `json:"AvailableAddOns,omitempty"`
	Meta            ListDayResponseMeta                `json:"Meta,omitempty"`
}
