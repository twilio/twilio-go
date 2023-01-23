/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// TaskrouterV1Worker struct for TaskrouterV1Worker
type TaskrouterV1Worker struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Worker resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The `friendly_name` of the Worker's current Activity.
	ActivityName *string `json:"activity_name,omitempty"`
	// The SID of the Worker's current Activity.
	ActivitySid *string `json:"activity_sid,omitempty"`
	// The JSON string that describes the Worker. For example: `{ \"email\": \"Bob@example.com\", \"phone\": \"+5095551234\" }`. **Note** If this property has been assigned a value, it will only be displayed in FETCH actions that return a single resource. Otherwise, this property will be null, even if it has a value. This data is passed to the `assignment_callback_url` when TaskRouter assigns a Task to the Worker.
	Attributes *string `json:"attributes,omitempty"`
	// Whether the Worker is available to perform tasks.
	Available *bool `json:"available,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT of the last change to the Worker's activity specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Used to calculate Workflow statistics.
	DateStatusChanged *time.Time `json:"date_status_changed,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource. Friendly names are case insensitive, and unique within the TaskRouter Workspace.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that we created to identify the Worker resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the Workspace that contains the Worker.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// The absolute URL of the Worker resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
