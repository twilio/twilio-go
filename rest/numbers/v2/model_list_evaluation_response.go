/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEvaluationResponse struct for ListEvaluationResponse
type ListEvaluationResponse struct {
	Meta    ListBundleResponseMeta                          `json:"meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceBundleEvaluation `json:"results,omitempty"`
}
