/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateDialingPermissionsCountryBulkUpdateRequest struct for CreateDialingPermissionsCountryBulkUpdateRequest
type CreateDialingPermissionsCountryBulkUpdateRequest struct {
	// URL encoded JSON array of update objects. example : `[ { \"iso_code\": \"GB\", \"low_risk_numbers_enabled\": \"true\", \"high_risk_special_numbers_enabled\":\"true\", \"high_risk_tollfraud_numbers_enabled\": \"false\" } ]`
	UpdateRequest string `json:"UpdateRequest"`
}
