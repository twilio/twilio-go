/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// StudioV1FlowExecutionExecutionStep struct for StudioV1FlowExecutionExecutionStep
type StudioV1FlowExecutionExecutionStep struct {
	AccountSid       *string                 `json:"account_sid,omitempty"`
	Context          *map[string]interface{} `json:"context,omitempty"`
	DateCreated      *time.Time              `json:"date_created,omitempty"`
	DateUpdated      *time.Time              `json:"date_updated,omitempty"`
	ExecutionSid     *string                 `json:"execution_sid,omitempty"`
	FlowSid          *string                 `json:"flow_sid,omitempty"`
	Links            *map[string]interface{} `json:"links,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	Sid              *string                 `json:"sid,omitempty"`
	TransitionedFrom *string                 `json:"transitioned_from,omitempty"`
	TransitionedTo   *string                 `json:"transitioned_to,omitempty"`
	Url              *string                 `json:"url,omitempty"`
}
