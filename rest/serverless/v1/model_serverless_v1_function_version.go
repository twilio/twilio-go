/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ServerlessV1FunctionVersion struct for ServerlessV1FunctionVersion
type ServerlessV1FunctionVersion struct {
	// The SID of the Account that created the Function Version resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Function Version resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The SID of the Function resource that is the parent of the Function Version resource
	FunctionSid *string                 `json:"function_sid,omitempty"`
	Links       *map[string]interface{} `json:"links,omitempty"`
	// The URL-friendly string by which the Function Version resource can be referenced
	Path *string `json:"path,omitempty"`
	// The SID of the Service that the Function Version resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Function Version resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Function Version resource
	Url *string `json:"url,omitempty"`
	// The access control that determines how the Function Version resource can be accessed
	Visibility *string `json:"visibility,omitempty"`
}
