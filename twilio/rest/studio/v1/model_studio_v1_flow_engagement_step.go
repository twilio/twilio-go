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

// StudioV1FlowEngagementStep struct for StudioV1FlowEngagementStep
type StudioV1FlowEngagementStep struct {
	AccountSid       *string                 `json:"AccountSid,omitempty"`
	Context          *map[string]interface{} `json:"Context,omitempty"`
	DateCreated      *time.Time              `json:"DateCreated,omitempty"`
	DateUpdated      *time.Time              `json:"DateUpdated,omitempty"`
	EngagementSid    *string                 `json:"EngagementSid,omitempty"`
	FlowSid          *string                 `json:"FlowSid,omitempty"`
	Links            *map[string]interface{} `json:"Links,omitempty"`
	Name             *string                 `json:"Name,omitempty"`
	Sid              *string                 `json:"Sid,omitempty"`
	TransitionedFrom *string                 `json:"TransitionedFrom,omitempty"`
	TransitionedTo   *string                 `json:"TransitionedTo,omitempty"`
	Url              *string                 `json:"Url,omitempty"`
}
