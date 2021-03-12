/*
    * Twilio - Autopilot
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
// AutopilotV1Assistant struct for AutopilotV1Assistant
type AutopilotV1Assistant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"AccountSid,omitempty"`
	// Reserved
	CallbackEvents *string `json:"CallbackEvents,omitempty"`
	// Reserved
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A string describing the state of the assistant.
	DevelopmentStage *string `json:"DevelopmentStage,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Reserved
	LatestModelBuildSid *string `json:"LatestModelBuildSid,omitempty"`
	// A list of the URLs of the Assistant's related resources
	Links *map[string]interface{} `json:"Links,omitempty"`
	// Whether queries should be logged and kept after training
	LogQueries *bool `json:"LogQueries,omitempty"`
	// Whether model needs to be rebuilt
	NeedsModelBuild *bool `json:"NeedsModelBuild,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"Sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"UniqueName,omitempty"`
	// The absolute URL of the Assistant resource
	Url *string `json:"Url,omitempty"`
}
