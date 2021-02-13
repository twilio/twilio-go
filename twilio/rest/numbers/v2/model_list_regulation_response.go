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

// ListRegulationResponse struct for ListRegulationResponse
type ListRegulationResponse struct {
	Meta    ListBundleResponseMeta                    `json:"Meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceRegulation `json:"Results,omitempty"`
}
