/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListCustomerProfileResponse struct for ListCustomerProfileResponse
type ListCustomerProfileResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"meta,omitempty"`
	Results []TrusthubV1CustomerProfile     `json:"results,omitempty"`
}
