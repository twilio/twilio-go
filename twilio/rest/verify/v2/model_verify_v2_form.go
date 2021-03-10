/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VerifyV2Form struct for VerifyV2Form
type VerifyV2Form struct {
	FormMeta *map[string]interface{} `json:"FormMeta,omitempty"`
	FormType *FormFormTypes          `json:"FormType,omitempty"`
	Forms    *map[string]interface{} `json:"Forms,omitempty"`
	Url      *string                 `json:"Url,omitempty"`
}
