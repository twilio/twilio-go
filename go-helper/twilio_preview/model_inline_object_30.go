/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject30 struct for InlineObject30
type InlineObject30 struct {
	// Whether the Terms of Service were accepted.
	AcceptTermsOfService bool `json:"AcceptTermsOfService"`
	// The SID of the AvaliableAddOn to install.
	AvailableAddOnSid string `json:"AvailableAddOnSid"`
	// The JSON object that represents the configuration of the new Add-on being installed.
	Configuration map[string]interface{} `json:"Configuration,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within the Account.
	UniqueName string `json:"UniqueName,omitempty"`
}
