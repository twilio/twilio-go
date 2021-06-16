/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.17.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ServerlessV1ServiceEnvironmentLog struct for ServerlessV1ServiceEnvironmentLog
type ServerlessV1ServiceEnvironmentLog struct {
	// The SID of the Account that created the Log resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the build that corresponds to the log
	BuildSid *string `json:"build_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the Log resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The SID of the deployment that corresponds to the log
	DeploymentSid *string `json:"deployment_sid,omitempty"`
	// The SID of the environment in which the log occurred
	EnvironmentSid *string `json:"environment_sid,omitempty"`
	// The SID of the function whose invocation produced the log
	FunctionSid *string `json:"function_sid,omitempty"`
	// The log level
	Level *string `json:"level,omitempty"`
	// The log message
	Message *string `json:"message,omitempty"`
	// The SID of the request associated with the log
	RequestSid *string `json:"request_sid,omitempty"`
	// The SID of the Service that the Log resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the Log resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Log resource
	Url *string `json:"url,omitempty"`
}
