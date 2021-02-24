/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// TaskrouterV1WorkspaceActivity struct for TaskrouterV1WorkspaceActivity
type TaskrouterV1WorkspaceActivity struct {
	AccountSid   *string    `json:"AccountSid,omitempty"`
	Available    *bool      `json:"Available,omitempty"`
	DateCreated  *time.Time `json:"DateCreated,omitempty"`
	DateUpdated  *time.Time `json:"DateUpdated,omitempty"`
	FriendlyName *string    `json:"FriendlyName,omitempty"`
	Sid          *string    `json:"Sid,omitempty"`
	Url          *string    `json:"Url,omitempty"`
	WorkspaceSid *string    `json:"WorkspaceSid,omitempty"`
}
