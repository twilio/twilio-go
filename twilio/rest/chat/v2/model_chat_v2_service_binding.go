/*
    * Twilio - Chat
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
// ChatV2ServiceBinding struct for ChatV2ServiceBinding
type ChatV2ServiceBinding struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The push technology to use for the binding
	BindingType *string `json:"BindingType,omitempty"`
	// The SID of the Credential for the binding
	CredentialSid *string `json:"CredentialSid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The unique endpoint identifier for the Binding
	Endpoint *string `json:"Endpoint,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"Identity,omitempty"`
	// The absolute URLs of the Binding's User
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The Programmable Chat message types the binding is subscribed to
	MessageTypes *[]string `json:"MessageTypes,omitempty"`
	// The SID of the Service that the Binding resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Binding resource
	Url *string `json:"Url,omitempty"`
}
