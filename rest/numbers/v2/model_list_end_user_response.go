/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEndUserResponse struct for ListEndUserResponse
type ListEndUserResponse struct {
	Meta    ListBundleResponseMeta                 `json:"meta,omitempty"`
	Results []NumbersV2RegulatoryComplianceEndUser `json:"results,omitempty"`
}
