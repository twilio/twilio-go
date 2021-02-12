/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListEndUserResponse struct for ListEndUserResponse
type ListEndUserResponse struct {
	Meta ListBundleResponseMeta `json:"Meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceEndUser `json:"Results,omitempty"`
}
