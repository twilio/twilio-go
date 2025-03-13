/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Assistants
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
	"time"
)

// AssistantsV1ToolWithPolicies struct for AssistantsV1ToolWithPolicies
type AssistantsV1ToolWithPolicies struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Tool resource.
	AccountSid string `json:"account_sid,omitempty"`
	// The description of the tool.
	Description string `json:"description"`
	// True if the tool is enabled.
	Enabled bool `json:"enabled"`
	// The tool ID.
	Id string `json:"id"`
	// The metadata related to method, url, input_schema to used with the Tool.
	Meta map[string]interface{} `json:"meta"`
	// The name of the tool.
	Name string `json:"name"`
	// The authentication requirement for the tool.
	RequiresAuth bool `json:"requires_auth"`
	// The type of the tool. ('WEBHOOK')
	Type string `json:"type"`
	// The url of the tool resource.
	Url string `json:"url,omitempty"`
	// The date and time in GMT when the Tool was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated time.Time `json:"date_created"`
	// The date and time in GMT when the Tool was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated time.Time `json:"date_updated"`
	// The Policies associated with the tool.
	Policies []AssistantsV1Policy `json:"policies,omitempty"`
}
