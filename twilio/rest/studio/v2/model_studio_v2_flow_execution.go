/*
 * Twilio - Studio
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

// StudioV2FlowExecution struct for StudioV2FlowExecution
type StudioV2FlowExecution struct {
	AccountSid            *string                 `json:"AccountSid,omitempty"`
	ContactChannelAddress *string                 `json:"ContactChannelAddress,omitempty"`
	Context               *map[string]interface{} `json:"Context,omitempty"`
	DateCreated           *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated           *time.Time              `json:"DateUpdated,omitempty"`
	FlowSid               *string                 `json:"FlowSid,omitempty"`
	Links                 *map[string]interface{} `json:"Links,omitempty"`
	Sid                   *string                 `json:"Sid,omitempty"`
	Status                *ExecutionStatus        `json:"Status,omitempty"`
	Url                   *string                 `json:"Url,omitempty"`
}
