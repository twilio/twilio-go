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
	"encoding/json"
	"github.com/twilio/twilio-go/client"
)

// StudioV2ExecutionContext struct for StudioV2ExecutionContext
type StudioV2ExecutionContext struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ExecutionContext resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The current state of the Flow's Execution. As a flow executes, we save its state in this context. We save data that your widgets can access as variables in configuration fields or in text areas as variable substitution.
	Context *map[string]interface{} `json:"context,omitempty"`
	// The SID of the Flow.
	FlowSid *string `json:"flow_sid,omitempty"`
	// The SID of the context's Execution resource.
	ExecutionSid *string `json:"execution_sid,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
