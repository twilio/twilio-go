/*
 * Twilio - Studio
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
// StudioV1FlowExecutionExecutionStep struct for StudioV1FlowExecutionExecutionStep
type StudioV1FlowExecutionExecutionStep struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Context map[string]interface{} `json:"Context,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ExecutionSid string `json:"ExecutionSid,omitempty"`
	FlowSid string `json:"FlowSid,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Name string `json:"Name,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TransitionedFrom string `json:"TransitionedFrom,omitempty"`
	TransitionedTo string `json:"TransitionedTo,omitempty"`
	Url string `json:"Url,omitempty"`
}
