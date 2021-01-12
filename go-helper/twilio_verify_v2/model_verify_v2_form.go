/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// VerifyV2Form struct for VerifyV2Form
type VerifyV2Form struct {
	FormMeta map[string]interface{} `json:"form_meta,omitempty"`
	FormType string `json:"form_type,omitempty"`
	Forms map[string]interface{} `json:"forms,omitempty"`
	Url string `json:"url,omitempty"`
}
