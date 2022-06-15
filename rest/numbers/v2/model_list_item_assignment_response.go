/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListItemAssignmentResponse struct for ListItemAssignmentResponse
type ListItemAssignmentResponse struct {
	Meta    ListBundleResponseMeta    `json:"meta,omitempty"`
	Results []NumbersV2ItemAssignment `json:"results,omitempty"`
}
