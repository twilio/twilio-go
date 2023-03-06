/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// StudioV2ExecutionStep struct for StudioV2ExecutionStep
type StudioV2ExecutionStep struct {
	// The unique string that we created to identify the ExecutionStep resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ExecutionStep resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Flow.
	FlowSid *string `json:"flow_sid,omitempty"`
	// The SID of the Step's Execution resource.
	ExecutionSid  *string `json:"execution_sid,omitempty"`
	ParentStepSid *string `json:"parent_step_sid,omitempty"`
	// The event that caused the Flow to transition to the Step.
	Name *string `json:"name,omitempty"`
	// The current state of the Flow's Execution. As a flow executes, we save its state in this context. We save data that your widgets can access as variables in configuration fields or in text areas as variable substitution.
	Context *interface{} `json:"context,omitempty"`
	// The Widget that preceded the Widget for the Step.
	TransitionedFrom *string `json:"transitioned_from,omitempty"`
	// The Widget that will follow the Widget for the Step.
	TransitionedTo *string `json:"transitioned_to,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
