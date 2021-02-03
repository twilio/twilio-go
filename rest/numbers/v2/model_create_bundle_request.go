/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateBundleRequest struct for CreateBundleRequest
type CreateBundleRequest struct {
	// The email address that will receive updates when the Bundle resource changes status.
	Email string `json:"Email"`
	// The type of End User of the Bundle resource.
	EndUserType string `json:"EndUserType,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName"`
	// The ISO country code of the Bundle's phone number country ownership request.
	IsoCountry string `json:"IsoCountry,omitempty"`
	// The type of phone number of the Bundle's ownership request.
	NumberType string `json:"NumberType,omitempty"`
	// The unique string of a regulation that is associated to the Bundle resource.
	RegulationSid string `json:"RegulationSid,omitempty"`
	// The URL we call to inform your application of status changes.
	StatusCallback string `json:"StatusCallback,omitempty"`
}
