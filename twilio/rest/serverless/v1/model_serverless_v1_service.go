/*
    * Twilio - Serverless
    *
    * This is the public Twilio REST API.
    *
    * API version: 1.10.0
    * Contact: support@twilio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi
import (
	"time"
)
// ServerlessV1Service struct for ServerlessV1Service
type ServerlessV1Service struct {
	// The SID of the Account that created the Service resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Service resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the Service resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the Service resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to inject Account credentials into a function invocation context
	IncludeCredentials *bool `json:"IncludeCredentials,omitempty"`
	// The URLs of the Service's nested resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The unique string that identifies the Service resource
	Sid *string `json:"Sid,omitempty"`
	// Whether the Service resource's properties and subresources can be edited via the UI
	UiEditable *bool `json:"UiEditable,omitempty"`
	// A user-defined string that uniquely identifies the Service resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The absolute URL of the Service resource
	Url *string `json:"Url,omitempty"`
}
