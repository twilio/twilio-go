/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NumbersV2RegulatoryComplianceBundleItemAssignmentReadResponse struct for NumbersV2RegulatoryComplianceBundleItemAssignmentReadResponse
type NumbersV2RegulatoryComplianceBundleItemAssignmentReadResponse struct {
	Meta NumbersV2RegulatoryComplianceBundleReadResponseMeta `json:"Meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceBundleItemAssignment `json:"Results,omitempty"`
}
