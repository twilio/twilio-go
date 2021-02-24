/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListPoliciesResponse struct for ListPoliciesResponse
type ListPoliciesResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"Meta,omitempty"`
	Results []TrusthubV1Policies            `json:"Results,omitempty"`
}
