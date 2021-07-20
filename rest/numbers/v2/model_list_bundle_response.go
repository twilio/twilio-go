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

// ListBundleResponse struct for ListBundleResponse
type ListBundleResponse struct {
	Meta    ListBundleResponseMeta                `json:"meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceBundle `json:"results,omitempty"`
}
