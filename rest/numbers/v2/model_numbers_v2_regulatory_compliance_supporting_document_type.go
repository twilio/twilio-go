/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// NumbersV2RegulatoryComplianceSupportingDocumentType struct for NumbersV2RegulatoryComplianceSupportingDocumentType
type NumbersV2RegulatoryComplianceSupportingDocumentType struct {
	Fields       *[]map[string]interface{} `json:"fields,omitempty"`
	FriendlyName *string                   `json:"friendly_name,omitempty"`
	MachineName  *string                   `json:"machine_name,omitempty"`
	Sid          *string                   `json:"sid,omitempty"`
	Url          *string                   `json:"url,omitempty"`
}
