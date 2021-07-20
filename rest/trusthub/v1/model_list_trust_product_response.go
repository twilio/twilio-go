/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTrustProductResponse struct for ListTrustProductResponse
type ListTrustProductResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"meta,omitempty"`
	Results []TrusthubV1TrustProduct        `json:"results,omitempty"`
}
