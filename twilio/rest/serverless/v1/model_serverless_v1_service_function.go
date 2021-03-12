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
// ServerlessV1ServiceFunction struct for ServerlessV1ServiceFunction
type ServerlessV1ServiceFunction struct {
	// The SID of the Account that created the Function resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Function resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the Function resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The string that you assigned to describe the Function resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URLs of nested resources of the Function resource
	Links *map[string]interface{} `json:"Links,omitempty"`
	// The SID of the Service that the Function resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the Function resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Function resource
	Url *string `json:"Url,omitempty"`
}
