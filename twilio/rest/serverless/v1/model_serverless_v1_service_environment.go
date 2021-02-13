/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.9.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// ServerlessV1ServiceEnvironment struct for ServerlessV1ServiceEnvironment
type ServerlessV1ServiceEnvironment struct {
	AccountSid   string                 `json:"AccountSid,omitempty"`
	BuildSid     string                 `json:"BuildSid,omitempty"`
	DateCreated  time.Time              `json:"DateCreated,omitempty"`
	DateUpdated  time.Time              `json:"DateUpdated,omitempty"`
	DomainName   string                 `json:"DomainName,omitempty"`
	DomainSuffix string                 `json:"DomainSuffix,omitempty"`
	Links        map[string]interface{} `json:"Links,omitempty"`
	ServiceSid   string                 `json:"ServiceSid,omitempty"`
	Sid          string                 `json:"Sid,omitempty"`
	UniqueName   string                 `json:"UniqueName,omitempty"`
	Url          string                 `json:"Url,omitempty"`
}
