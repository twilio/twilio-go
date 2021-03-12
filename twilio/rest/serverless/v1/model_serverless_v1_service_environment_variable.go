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
// ServerlessV1ServiceEnvironmentVariable struct for ServerlessV1ServiceEnvironmentVariable
type ServerlessV1ServiceEnvironmentVariable struct {
	// The SID of the Account that created the Variable resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// The ISO 8601 date and time in GMT when the Variable resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The ISO 8601 date and time in GMT when the Variable resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// The SID of the Environment in which the Variable exists
	EnvironmentSid *string `json:"EnvironmentSid,omitempty"`
	// A string by which the Variable resource can be referenced
	Key *string `json:"Key,omitempty"`
	// The SID of the Service that the Variable resource is associated with
	ServiceSid *string `json:"ServiceSid,omitempty"`
	// The unique string that identifies the Variable resource
	Sid *string `json:"Sid,omitempty"`
	// The absolute URL of the Variable resource
	Url *string `json:"Url,omitempty"`
	// A string that contains the actual value of the Variable
	Value *string `json:"Value,omitempty"`
}
