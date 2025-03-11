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

// StudioV2Flow struct for StudioV2Flow
type StudioV2Flow struct {
	// The unique string that we created to identify the Flow resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flow resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the Flow.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// JSON representation of flow definition.
	Definition *map[string]interface{} `json:"definition,omitempty"`
	Status     *string                 `json:"status,omitempty"`
	// The latest revision number of the Flow's definition.
	Revision int `json:"revision,omitempty"`
	// Description of change made in the revision.
	CommitMessage *string `json:"commit_message,omitempty"`
	// Boolean if the flow definition is valid.
	Valid *bool `json:"valid,omitempty"`
	// List of error in the flow definition.
	Errors *[]map[string]interface{} `json:"errors,omitempty"`
	// List of warnings in the flow definition.
	Warnings *[]map[string]interface{} `json:"warnings,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	WebhookUrl  *string    `json:"webhook_url,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of the Flow's nested resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
