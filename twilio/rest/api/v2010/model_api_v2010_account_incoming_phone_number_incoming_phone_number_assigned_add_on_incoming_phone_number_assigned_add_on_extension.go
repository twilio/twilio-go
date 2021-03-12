/*
    * Twilio - Api
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
// ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension struct for ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension
type ApiV2010AccountIncomingPhoneNumberIncomingPhoneNumberAssignedAddOnIncomingPhoneNumberAssignedAddOnExtension struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The SID that uniquely identifies the assigned Add-on installation
	AssignedAddOnSid *string `json:"AssignedAddOnSid,omitempty"`
	// Whether the Extension will be invoked
	Enabled *bool `json:"Enabled,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// A string that you assigned to describe the Product this Extension is used within
	ProductName *string `json:"ProductName,omitempty"`
	// The SID of the Phone Number to which the Add-on is assigned
	ResourceSid *string `json:"ResourceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"Uri,omitempty"`
}
