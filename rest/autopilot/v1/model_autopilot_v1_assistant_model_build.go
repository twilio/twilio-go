/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 5.15.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// AutopilotV1AssistantModelBuild struct for AutopilotV1AssistantModelBuild
type AutopilotV1AssistantModelBuild struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssistantSid string `json:"AssistantSid,omitempty"`
	BuildDuration *int32 `json:"BuildDuration,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ErrorCode *int32 `json:"ErrorCode,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status string `json:"Status,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
