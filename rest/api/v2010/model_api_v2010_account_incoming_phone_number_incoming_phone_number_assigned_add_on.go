/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn struct for ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn
type ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOn struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A JSON string that represents the current configuration
	Configuration *map[string]interface{} `json:"configuration,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A short description of the Add-on functionality
	Description *string `json:"description,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The SID of the Phone Number that installed this Add-on
	ResourceSid *string `json:"resource_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
